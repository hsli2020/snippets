package main

import (
	"flag"
	"fmt"
	"nft/internal/combiner"
	"nft/internal/domain"
	"nft/internal/generator"
	"nft/internal/helpers"
	"nft/internal/inmemory"
	"nft/internal/pinata"
	"nft/internal/trait"
	"os"
)

func main() {
	// Flags.
	appContext := domain.NewAppContext()

	// Generator flags.
	flag.IntVar(&appContext.GeneratorParams.Length, "generated-image-length", 2048, "generated image length")
	flag.IntVar(&appContext.GeneratorParams.Width, "generated-image-width", 2048, "generated image width")
	flag.StringVar(&appContext.GeneratorParams.InputDirectory, "generated-image-input", "input-dir", "generated image input directory")
	flag.StringVar(&appContext.GeneratorParams.OutputDirectory, "generated-image-output", "output-dir", "generated image output directory")
	flag.IntVar(&appContext.GeneratorParams.Number, "generated-image-number", 100, "generated image number")
	generate := flag.Bool("generate", false, "Generate images")

	// Ipfs flags.
	flag.StringVar(&appContext.IpfsParams.InputDirectory, "ipfs-input", "output-dir", "ipfs input directory")
	flag.StringVar(&appContext.IpfsParams.OutputDirectory, "ipfs-output", "ipfs-metadata", "ipfs input directory")
	flag.StringVar(&appContext.IpfsParams.APIKey, "ipfs-api-key", "", "ipfs api key")
	flag.StringVar(&appContext.IpfsParams.SecretKey, "ipfs-secret-key", "", "ipfs secret key")
	ipfsUpload := flag.Bool("ipfs-upload", false, "Upload files to ipfs")

	// Other.
	printInfo := flag.Bool("info", false, "Print info")
	flag.CommandLine.SetOutput(os.Stdout)
	flag.Parse()

	if *generate {
		fmt.Println("Generating images ...")
		traitService := trait.NewBasicTraitService(
			inmemory.NewGroupRepository(),
			inmemory.NewTraitRepository(),
		)
		_, err := traitService.Import(appContext.GeneratorParams.InputDirectory)
		if err != nil {
			panic(err)
		}
		generatorService := generator.NewBasicImageGenerator(
			appContext.GeneratorParams,
			traitService,
			combiner.NewBasicImageCombiner(),
		)
		err = generatorService.GenerateImages()
		if err != nil {
			panic(err)
		}
		return
	}
	if *ipfsUpload {
		fmt.Println("Uploading files to ipfs ...")
		ipfsService := pinata.NewIpfsService(appContext.IpfsParams)
		if err := ipfsService.UploadImages(); err != nil {
			panic(err)
		}
	}
	if *printInfo {
		helpers.PrintInfoV2(
			appContext.IpfsParams.OutputDirectory,
			appContext.GeneratorParams.OutputDirectory,
		)
		helpers.PrintInfo(appContext.GeneratorParams.OutputDirectory)
	}
}

// internal/trait/service.go
package trait

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"nft/internal/domain"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type service struct {
	groupRepository domain.GroupRepository
	traitRepository domain.TraitRepository
}

// NewBasicTraitService - returns a naïve, stateless implementation of a service.
func NewBasicTraitService(
	groupRepository domain.GroupRepository,
	traitRepository domain.TraitRepository,
) domain.TraitService {
	return &service{
		groupRepository: groupRepository,
		traitRepository: traitRepository,
	}
}

func (s *service) Import(root string) (int, error) {
	priority := 0
	err := filepath.Walk(root, func(path string, info os.FileInfo, errr error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(info.Name()) != ".png" {
			return fmt.Errorf("Bad file extension: %s", filepath.Ext(info.Name()))
		}
		splited := strings.Split(path, "/")
		groupName := splited[1][3:]
		traitName := info.Name()[:len(info.Name())-4]

		traitOptions := strings.Split(traitName, ".")
		traitName = traitOptions[0]
		rarenessKind := domain.RarenessKindCommon
		if len(traitOptions) > 1 {
			switch traitOptions[1] {
			case "silver":
				rarenessKind = domain.RarenessKindSilver
			case "gold":
				rarenessKind = domain.RarenessKindGold
			}
		}

		// Create group if not exist.
		foundGroup, _ := s.groupRepository.GetByName(groupName)
		if foundGroup == nil {
			_, err := s.groupRepository.Create(&domain.GroupWrite{
				Name:     groupName,
				Priotiry: priority,
			})
			priority++

			fmt.Printf("%s - %d\n", groupName, priority)
			if err != nil {
				return err
			}
		}
		foundGroup, _ = s.groupRepository.GetByName(groupName)
		if foundGroup == nil {
			return fmt.Errorf("group not found %s", groupName)
		}

		// Read image.
		imgBytes, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		// Create trait.
		_, err = s.traitRepository.Create(&domain.TraitWrite{
			Name:         traitName,
			Group:        foundGroup,
			Image:        imgBytes,
			RarenessKind: rarenessKind,
		})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return 0, err
	}
	return 0, nil
}

func (s *service) GetRandomTraits() ([]*domain.TraitRead, error) {
	// Get all groups.
	groups, err := s.groupRepository.GetAll()
	if err != nil {
		return nil, err
	}

	// Group traits by their types.
	groupedTraits := make(map[domain.GroupID][]*domain.TraitRead)
	for _, group := range groups {
		traits, err := s.traitRepository.GetByGroupID(group.ID)
		if err != nil {
			return nil, err
		}
		if len(traits) == 0 {
			return nil, fmt.Errorf("traits not found for group %v (%s)", group.ID, group.Name)
		}
		groupedTraits[group.ID] = traits
	}

	// Choose random traits for each available group.
	randomTraits := make([]*domain.TraitRead, 0)
	for _, traits := range groupedTraits {
		randomTrait, err := GetRandomTrait(traits)
		if err != nil {
			return nil, err
		}
		randomTraits = append(randomTraits, randomTrait)
	}
	if len(groups) != len(randomTraits) {
		return nil, fmt.Errorf(
			fmt.Sprintf("expected traits size %d but got %d", len(groups), len(randomTraits)),
		)
	}
	return randomTraits, nil
}

func GetRandomTrait(traits []*domain.TraitRead) (*domain.TraitRead, error) {

	pdf, err := getProbabilityDensityVector(traits)
	if err != nil {
		return nil, err
	}

	// get cdf
	len := len(traits)
	cdf := make([]float32, len)
	cdf[0] = pdf[0]
	for i := 1; i < len; i++ {
		cdf[i] = cdf[i-1] + pdf[i]
	}
	random := sample(cdf)
	if !(len > random) {
		return nil, fmt.Errorf(
			fmt.Sprintf("random generated trait index out of range, max size: %d, generated index: %d", len, random),
		)
	}
	return traits[random], nil
}

func getProbabilityDensityVector(traits []*domain.TraitRead) ([]float32, error) {
	var (
		len               = len(traits)
		probabilityVector = make([]float32, len)
	)
	var (
		baseChance   = float32(100/len) / 100
		silverChance = baseChance / 2
		goldChance   = baseChance / 4
	)
	var (
		chanceOffset  float32 = 1.00
		commonCounter         = 0
	)
	for i, t := range traits {
		switch t.RarenessKind {
		case domain.RarenessKindSilver:
			probabilityVector[i] = silverChance
			chanceOffset -= silverChance
		case domain.RarenessKindGold:
			probabilityVector[i] = goldChance
			chanceOffset -= goldChance
		default:
			commonCounter++
		}
	}
	for i, p := range probabilityVector {
		if p == 0 {
			probabilityVector[i] = chanceOffset / float32(commonCounter)
		}
	}
	if err := checkProbabilityVector(probabilityVector); err != nil {
		return nil, err
	}
	return probabilityVector, nil
}

func checkProbabilityVector(vector []float32) error {
	var (
		sum         float32 = 0
		controllSum float32 = 1
	)
	for _, p := range vector {
		sum += p
	}
	if !(sum >= controllSum-0.1 || sum >= controllSum+0.1) {
		return fmt.Errorf("Expected probability vector controll sum %v but got %v", controllSum, sum)
	}
	return nil
}

func sample(cdf []float32) int {
	rand.Seed(time.Now().UnixNano())
	r := rand.Float32()
	bucket := 0
	for r > cdf[bucket] {
		bucket++
	}
	return bucket
}

// internal/pinata/service.go

package pinata

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"nft/internal/domain"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type service struct {
	params domain.IpfsParams
	client *http.Client
}

const (
	pinFileURL = "https://api.pinata.cloud/pinning/pinFileToIPFS"
)

// NewIpfsService - new pinata service.
func NewIpfsService(params domain.IpfsParams) domain.IpfsService {
	return &service{
		params: params,
		client: http.DefaultClient,
	}
}

func (s *service) UploadImages() error {
	var (
		uploadImageParams = make([]*uploadImageParam, 0)
	)
	counter := 0
	err := filepath.Walk(s.params.InputDirectory, func(path string, info os.FileInfo, _ error) error {
		if info.IsDir() {
			return nil
		}
		var (
			fileName = info.Name()
		)
		if filepath.Ext(fileName) != ".png" {
			return nil
		}
		counter++
		uploadImageParams = append(uploadImageParams, &uploadImageParam{
			number:   counter,
			path:     path,
			fileName: fileName,
		})
		return nil
	})
	if err != nil {
		return err
	}

	fmt.Printf("Files to upload: %d\n", len(uploadImageParams))
	var (
		wg       sync.WaitGroup
		poolSize = 10
		channel  = make(chan *uploadImageParam, poolSize)
	)
	for i := 0; i < poolSize; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for p := range channel {
				err := s.uploadImage(p)
				if err != nil {
					panic(err)
				}
			}
		}()
	}
	for _, p := range uploadImageParams {
		channel <- p
	}
	return nil
}

type uploadImageParam struct {
	number   int
	path     string
	fileName string
}

func (s *service) uploadImage(p *uploadImageParam) error {
	var (
		key = strings.TrimSuffix(p.fileName, filepath.Ext(p.fileName))
	)
	// Upload image to ipfs.
	imgBytes, err := ioutil.ReadFile(p.path)
	if err != nil {
		return err
	}
	ipfsImageHash, err := s.pinFile(p.fileName, imgBytes, false)
	if err != nil {
		return err
	}

	// Read image traits.
	traitsBytes, err := ioutil.ReadFile(strings.ReplaceAll(p.path, ".png", ".json"))
	if err != nil {
		return err
	}
	traits := []*domain.ERC721Trait{}
	if err := json.Unmarshal(traitsBytes, &traits); err != nil {
		return err
	}

	// Create image metadata file.
	erc721Metadata := &domain.ERC721Metadata{
		Image:      fmt.Sprintf("ipfs://%s", ipfsImageHash),
		Attributes: traits,
	}
	erc721MetadataBytes, err := json.Marshal(erc721Metadata)
	if err != nil {
		return err
	}
	metadataFile, err := os.Create(fmt.Sprintf("%s/%d.%s.json", s.params.OutputDirectory, p.number, key))
	if err != nil {
		return err
	}
	defer metadataFile.Close()
	_, err = metadataFile.Write(erc721MetadataBytes)
	if err != nil {
		return err
	}
	fmt.Printf("Image successfully uploaded to ipfs: %d, %s \n", p.number, p.fileName)
	return nil
}

func (s *service) pinFile(fileName string, data []byte, wrapWithDirectory bool) (string, error) {
	type pinataResponse struct {
		IPFSHash  string `json:"IpfsHash"`
		PinSize   int    `json:"PinSize"`
		Timestamp string `json:"Timestamp"`
	}

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// this step is very important
	fileWriter, err := bodyWriter.CreateFormFile("file", fileName)
	if err != nil {
		return "", err
	}
	if _, err := fileWriter.Write(data); err != nil {
		return "", err
	}

	// wrap with directory.
	if wrapWithDirectory {
		fileWriter, err = bodyWriter.CreateFormField("pinataOptions")
		if err != nil {
			return "", err
		}
		if _, err := fileWriter.Write([]byte(`{"wrapWithDirectory": true}`)); err != nil {
			return "", err
		}
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	req, err := http.NewRequest("POST", pinFileURL, bodyBuf)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Set("pinata_api_key", s.params.APIKey)
	req.Header.Set("pinata_secret_api_key", s.params.SecretKey)

	// Do request.
	var (
		retries = 3
		resp    *http.Response
	)
	for retries > 0 {
		resp, err = s.client.Do(req)
		if err != nil {
			retries -= 1
		} else {
			break
		}
	}
	if resp == nil {
		return "", fmt.Errorf("Failed to upload files to ipfs, err: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		errMsg := make([]byte, resp.ContentLength)
		_, _ = resp.Body.Read(errMsg)
		return "", fmt.Errorf("Failed to upload file, response code %d, msg: %s", resp.StatusCode, string(errMsg))
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	pinataResp := pinataResponse{}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&pinataResp)
	if err != nil {
		return "", fmt.Errorf("Failed to decode json, err: %v", err)
	}
	if len(pinataResp.IPFSHash) == 0 {
		return "", errors.New("Ipfs hash not found in the response body")
	}
	return pinataResp.IPFSHash, nil
}

// internal/inmemory/group_repository.go

package inmemory

import (
	"errors"
	"fmt"
	"nft/internal/domain"

	"github.com/google/uuid"
)

type groupRepo struct {
	m map[domain.GroupID]*domain.GroupRead
	n map[string]*domain.GroupRead
}

func NewGroupRepository() domain.GroupRepository {
	return &groupRepo{
		m: make(map[domain.GroupID]*domain.GroupRead),
		n: make(map[string]*domain.GroupRead),
	}
}

func (r *groupRepo) Create(group *domain.GroupWrite) (domain.GroupID, error) {
	if group == nil {
		return "", errors.New("cannot store nil group")
	}
	if len(group.Name) == 0 {
		return "", errors.New("cannot store group with empty name")
	}
	groupID := domain.GroupID(uuid.New().String())
	readGroup := &domain.GroupRead{
		ID:         groupID,
		GroupWrite: *group,
	}
	r.m[groupID] = readGroup
	r.n[group.Name] = readGroup
	return groupID, nil
}

func (r *groupRepo) GetByID(groupID domain.GroupID) (*domain.GroupRead, error) {
	foundGroup, ok := r.m[groupID]
	if !ok {
		return nil, fmt.Errorf("group not found by id %v", groupID)
	}
	return foundGroup, nil
}

func (r *groupRepo) GetByName(name string) (*domain.GroupRead, error) {
	foundGroup, ok := r.n[name]
	if !ok {
		return nil, fmt.Errorf("group not found by name %s", name)
	}
	return foundGroup, nil
}

func (r *groupRepo) GetAll() ([]*domain.GroupRead, error) {
	list := make([]*domain.GroupRead, 0)
	for _, group := range r.m {
		list = append(list, group)
	}
	return list, nil
}

// internal/inmemory/trait_repository.go

package inmemory

import (
	"errors"
	"fmt"
	"nft/internal/domain"

	"github.com/google/uuid"
)

type traitRepo struct {
	m map[domain.TraitID]*domain.TraitRead
	n map[string]*domain.TraitRead
	g map[domain.GroupID][]*domain.TraitRead
}

func NewTraitRepository() domain.TraitRepository {
	return &traitRepo{
		m: make(map[domain.TraitID]*domain.TraitRead),
		n: make(map[string]*domain.TraitRead),
		g: make(map[domain.GroupID][]*domain.TraitRead),
	}
}

func (r *traitRepo) Create(trait *domain.TraitWrite) (domain.TraitID, error) {
	if trait == nil {
		return "", errors.New("cannot store nil trait")
	}
	if len(trait.Name) == 0 {
		return "", errors.New("cannot store trait with empty name")
	}
	traitID := domain.TraitID(uuid.New().String())
	traitRead := &domain.TraitRead{
		ID:         traitID,
		TraitWrite: *trait,
	}
	r.m[traitID] = traitRead
	r.n[trait.Name] = traitRead

	_, ok := r.g[trait.Group.ID]
	if !ok {
		r.g[trait.Group.ID] = make([]*domain.TraitRead, 0)
	}
	r.g[trait.Group.ID] = append(r.g[trait.Group.ID], traitRead)
	return traitID, nil
}

func (r *traitRepo) GetByID(traitID domain.TraitID) (*domain.TraitRead, error) {
	foundTrait, ok := r.m[traitID]
	if !ok {
		return nil, fmt.Errorf("trait not found by id %v", traitID)
	}
	return foundTrait, nil
}

func (r *traitRepo) IsExistByName(name string) (bool, error) {
	_, ok := r.n[name]
	return ok, nil
}

func (r *traitRepo) GetAll() ([]*domain.TraitRead, error) {
	list := make([]*domain.TraitRead, 0)
	for _, trait := range r.m {
		list = append(list, trait)
	}
	return list, nil
}

func (r *traitRepo) GetByGroupID(groupID domain.GroupID) ([]*domain.TraitRead, error) {
	return r.g[groupID], nil
}

// internal/helpers/utils.go

package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"nft/internal/domain"
	"os"
	"path/filepath"
	"strings"
)

func PrintInfo(root string) {
	m := map[string]int{}
	types := map[string]bool{}
	err := filepath.Walk(root, func(path string, info os.FileInfo, errr error) error {
		if filepath.Ext(info.Name()) != ".json" {
			return nil
		}
		jsonBytes, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		traits := []*domain.ERC721Trait{}
		if err := json.Unmarshal(jsonBytes, &traits); err != nil {
			return err
		}
		if len(traits) == 0 {
			return fmt.Errorf("Traits not found for %s", path)
		}
		for _, trait := range traits {
			key := fmt.Sprintf("%s-%s", trait.TraitType, trait.Value)
			m[key] = m[key] + 1
			types[trait.TraitType] = true
		}
		return nil
	})

	for t := range types {
		fmt.Println(t)
		for k, v := range m {
			vals := strings.Split(k, "-")
			if t == vals[0] {
				fmt.Printf("	%s = %d \n", vals[1], v)
			}
		}
	}
	if err != nil {
		panic(err)
	}
}

func PrintInfoV2(dir1, dir2 string) {
	err := filepath.Walk(dir1, func(path string, info os.FileInfo, _ error) error {
		ext := filepath.Ext(info.Name())
		if ext != ".json" {
			return nil
		}
		key := strings.Split(info.Name(), ".")[1]

		// metadata traits
		jsonBytes, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		metadata := domain.ERC721Metadata{}
		if err := json.Unmarshal(jsonBytes, &metadata); err != nil {
			return err
		}
		if len(metadata.Attributes) == 0 {
			return fmt.Errorf("Traits not found for %s", path)
		}

		// traits
		path2 := fmt.Sprintf("%s/%s.json", dir2, key)
		jsonBytes, err = ioutil.ReadFile(path2)
		if err != nil {
			return err
		}
		traits := []*domain.ERC721Trait{}
		if err := json.Unmarshal(jsonBytes, &traits); err != nil {
			return err
		}
		if len(traits) == 0 {
			return fmt.Errorf("Traits not found for %s", path2)
		}
		if err := compareTraits(key, traits, metadata.Attributes); err != nil {
			return err
		}
		return compareTraits(key, metadata.Attributes, traits)
	})
	if err != nil {
		panic(err)
	}
}

func compareTraits(key string, traits1, traits2 []*domain.ERC721Trait) error {
	if len(traits1) == 0 || len(traits2) == 0 {
		return fmt.Errorf("Traits are empty for key %s", key)
	}
	if len(traits1) != len(traits2) {
		return fmt.Errorf("Traits len does not equal, %d vs %d for key %s", len(traits1), len(traits2), key)
	}
	for _, t1 := range traits1 {
		found := false
		for _, t2 := range traits2 {
			if t1 != nil && t2 != nil && *t1 == *t2 {
				found = true
			}
		}
		if !found {
			return fmt.Errorf("Trait %v not found for key %s", t1, key)
		}
	}
	return nil
}

// internal/generator/service.go

package generator

import (
	"encoding/json"
	"errors"
	"fmt"
	"hash/fnv"
	"image/color"
	"nft/internal/domain"
	"os"
	"sort"
	"sync"
)

type service struct {
	params        domain.GeneratorParams
	traitService  domain.TraitService
	imageCombiner domain.ImageCombiner
}

// NewBasicRarityService returns a naïve, stateless implementation of a service.
func NewBasicImageGenerator(
	params domain.GeneratorParams,
	traitService domain.TraitService,
	imageCombiner domain.ImageCombiner,
) domain.ImageGenerator {
	return &service{
		params:        params,
		traitService:  traitService,
		imageCombiner: imageCombiner,
	}
}

func (s *service) GenerateImages() error {
	if !(s.params.Number > 0) {
		return fmt.Errorf("specify at least one element to generate")
	}
	var (
		wg       sync.WaitGroup
		poolSize = 20
		channel  = make(chan int, poolSize)
	)
	for i := 0; i < poolSize; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for _ = range channel {
				err := s.generateImageInternal()
				if err != nil {
					panic(err)
				}
			}
		}()
	}
	for i := 0; i < s.params.Number; i++ {
		channel <- i
	}
	return nil
}

func (s *service) generateImageInternal() error {
	img, traits, err := s.generateInternal()
	if err != nil {
		return err
	}
	key := ""
	for _, flat := range traits {
		key = key + " " + fmt.Sprintf("%s-%s", flat.TraitType, flat.Value)
	}
	key2 := hash(key)
	// Create image file.
	imgFile, err := os.Create(fmt.Sprintf("%s/%d.png", s.params.OutputDirectory, key2))
	if err != nil {
		return err
	}
	defer imgFile.Close()
	_, err = imgFile.Write(img)
	if err != nil {
		return err
	}
	// Create traits file.
	traitsFile, err := os.Create(fmt.Sprintf("%s/%d.json", s.params.OutputDirectory, key2))
	if err != nil {
		return err
	}
	defer traitsFile.Close()
	traitsBytes, err := json.Marshal(traits)
	if err != nil {
		return err
	}
	_, err = traitsFile.Write(traitsBytes)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) generateInternal() ([]byte, []*domain.ERC721Trait, error) {
	// Get random traits.
	traits, err := s.traitService.GetRandomTraits()
	if err != nil {
		return nil, nil, err
	}
	if len(traits) == 0 {
		return nil, nil, errors.New("cannot generate with empty random traits")
	}

	// Convert traints to ERC721 format.
	erc721Traits := make([]*domain.ERC721Trait, len(traits))
	for i, trait := range traits {
		erc721Trait, err := trait.ToERC721()
		if err != nil {
			return nil, nil, err
		}
		erc721Traits[i] = erc721Trait
	}

	// Convert traits to image layers.
	layers := make([]*domain.ImageLayer, len(traits))
	for i, trait := range traits {
		layer, err := trait.ToImageLayer()
		if err != nil {
			return nil, nil, err
		}
		layers[i] = layer
	}

	// Combine image layers together.
	img, err := s.imageCombiner.CombineLayers(layers, &domain.BgProperty{
		Width:   s.params.Width,
		Length:  s.params.Length,
		BgColor: color.Transparent,
	})
	if err != nil {
		return nil, nil, err
	}
	return img, sortTraits(erc721Traits), nil
}

func sortTraits(list []*domain.ERC721Trait) []*domain.ERC721Trait {
	sort.Slice(list, func(i, j int) bool {
		t1 := hash(list[i].TraitType)
		t2 := hash(list[j].TraitType)
		return t1 < t2
	})
	return list
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

// internal/domain/app.go
package domain

// AppContext - application main context.
type AppContext struct {
	GeneratorParams GeneratorParams
	IpfsParams      IpfsParams
}

// NewAppContext - constructs a new application context.
func NewAppContext() *AppContext {
	return &AppContext{
		GeneratorParams: GeneratorParams{},
		IpfsParams:      IpfsParams{},
	}
}

// internal/domain/combiner.go
package domain

import (
	"image"
	"image/color"
)

// ImageLayer struct.
type ImageLayer struct {
	Image    image.Image
	Priotiry int
	XPos     int
	YPos     int
}

//BgProperty is background property struct.
type BgProperty struct {
	Width   int
	Length  int
	BgColor color.Color
}

// ImageCombiner interface.
type ImageCombiner interface {
	CombineLayers(layers []*ImageLayer, bgProperty *BgProperty) ([]byte, error)
}

// internal/domain/erc721.go
package domain

// ERC721Trait - ERC721 trait format.
type ERC721Trait struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
}

// ERC721Metadata - metadata schema.
type ERC721Metadata struct {
	Image      string         `json:"image"`
	Attributes []*ERC721Trait `json:"attributes"`
}

// internal/domain/generator.go
package domain

// GeneratorParams struct.
type GeneratorParams struct {
	Width           int
	Length          int
	InputDirectory  string
	OutputDirectory string
	Number          int
}

// ImageGenerator interface.
type ImageGenerator interface {
	GenerateImages() error
}

// internal/domain/group.go
package domain

// GroupID - group id.
type GroupID string

// GroupWrite struct.
type GroupWrite struct {
	Name     string `json:"name"`
	Priotiry int    `json:"priotiry"`
	XPos     int    `json:"xpos"`
	YPos     int    `json:"ypos"`
}

// GroupRead struct.
type GroupRead struct {
	ID GroupID `json:"id"`
	GroupWrite
}

// GroupRepository - provides access to the storage.
type GroupRepository interface {
	Create(group *GroupWrite) (GroupID, error)
	GetByID(groupID GroupID) (*GroupRead, error)
	GetByName(name string) (*GroupRead, error)
	GetAll() ([]*GroupRead, error)
}

// internal/domain/ipfs.go
package domain

// IpfsParams - ipfs parameters.
type IpfsParams struct {
	InputDirectory  string
	OutputDirectory string
	APIKey          string
	SecretKey       string
}

// IpfsService - ipfs service.
type IpfsService interface {
	UploadImages() error
}

// internal/domain/traits.go
package domain

import (
	"bytes"
	"errors"
	"image/png"
	"strings"
)

const (
	RarenessKindCommon = RarenessKind(0)
	RarenessKindSilver = RarenessKind(1)
	RarenessKindGold   = RarenessKind(2)
)

// TraitID - trait id.
type TraitID string

// Rareness - rareness.
type RarenessKind int

// TraitWrite struct.
type TraitWrite struct {
	Name         string       `json:"name"`
	Group        *GroupRead   `json:"group"`
	Image        []byte       `json:"image"`
	RarenessKind RarenessKind `json:"rareness"`
}

// TraitRead struct.
type TraitRead struct {
	ID TraitID `json:"id"`
	TraitWrite
}

// ToImageLayer - returns image layer.
func (r *TraitRead) ToImageLayer() (*ImageLayer, error) {
	img, err := png.Decode(bytes.NewReader(r.Image))
	if err != nil {
		return nil, err
	}
	return &ImageLayer{
		Image:    img,
		Priotiry: r.Group.Priotiry,
		XPos:     r.Group.XPos,
		YPos:     r.Group.YPos,
	}, nil
}

// ToERC721 - returns ERC721 trait format.
func (r *TraitRead) ToERC721() (*ERC721Trait, error) {
	if len(r.Group.Name) == 0 {
		return nil, errors.New("trait type required")
	}
	if len(r.Name) == 0 {
		return nil, errors.New("trait value required")
	}
	return &ERC721Trait{
		TraitType: strings.ToUpper(r.Group.Name),
		Value:     strings.ToUpper(r.Name),
	}, nil
}

// TraitRepository - provides access to the storage.
type TraitRepository interface {
	Create(trait *TraitWrite) (TraitID, error)
	GetByID(traitID TraitID) (*TraitRead, error)
	IsExistByName(name string) (bool, error)
	GetAll() ([]*TraitRead, error)
	GetByGroupID(groupID GroupID) ([]*TraitRead, error)
}

// TraitService - provides access to the business logic.
type TraitService interface {
	Import(root string) (int, error)
	GetRandomTraits() ([]*TraitRead, error)
}

// internal/combiner/service.go

package combiner

import (
	"bytes"
	"image"
	"image/draw"
	"image/png"
	"nft/internal/domain"
	"sort"
)

type service struct {}

func NewBasicImageCombiner() domain.ImageCombiner {
	return &service{}
}

func (s *service) CombineLayers(layers []*domain.ImageLayer, bgProperty *domain.BgProperty) ([]byte, error) {

	// Sort list by position.
	layers = sortByPriotiry(layers)

	// Create image's background.
	bgImg := image.NewRGBA(image.Rect(0, 0, bgProperty.Width, bgProperty.Length))

	// Set the background color.
	draw.Draw(bgImg, bgImg.Bounds(), &image.Uniform{bgProperty.BgColor}, image.Point{}, draw.Src)

	// Looping image layers, higher position -> upper layer.
	for _, img := range layers {

		// Set the image offset.
		offset := image.Pt(img.XPos, img.YPos)

		// Combine the image.
		draw.Draw(bgImg, img.Image.Bounds().Add(offset), img.Image, image.Point{}, draw.Over)
	}

	// Encode image to buffer.
	buff := new(bytes.Buffer)
	if err := png.Encode(buff, bgImg); err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}

func sortByPriotiry(list []*domain.ImageLayer) []*domain.ImageLayer {
	sort.Slice(list, func(i, j int) bool {
		return list[i].Priotiry < list[j].Priotiry
	})
	return list
}
