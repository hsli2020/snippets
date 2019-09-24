package main

import (
	"archive/zip"
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"unsafe"
)

//#cgo LDFLAGS: libjbig2enc.a -lpng -llept
//#include <stdlib.h>
//#include <stdio.h>
//
//struct Pix;
//typedef struct Pix PIX;
//void pixDestroy ( PIX **ppix );
//PIX * pixReadMemPng ( const uint8_t*filedata, size_t filesize );
//PIX * pixRemoveColormap ( PIX *pixs, int type );
//
//struct jbig2ctx;
//typedef struct jbig2ctx JB2;
//JB2 *jb2Init(float thresh, float weight, int xres, int yres, int full_headers, int refine_level);
//void jb2Destroy(JB2 *ctx);
//void jb2AddPage(JB2 *ctx, PIX *bw);
//uint8_t *jb2ProducePage(JB2 *ctx, int page_no, int xres, int yres, int *const length);
//uint8_t *jb2PagesComplete(JB2 *ctx, int *const length);
import "C"

func ref(x int) string {
	return fmt.Sprintf("%d 0 R", x)
}

type Dict map[string]string

func (d Dict) Set(key string, val string) {
	d[key] = val
}

func (d Dict) String() string {
	var s []string
	s = append(s, "<< ")
	for k, v := range d {
		s = append(s, fmt.Sprintf("/%s ", k), v, "\n")
	}
	s = append(s, ">>\n")

	return strings.Join(s, "")
}

type Obj struct {
	id     int
	dict   Dict
	stream *bytes.Buffer
}

func NewObj(id int, dict Dict, stream *bytes.Buffer) *Obj {
	obj := &Obj{id: id, dict: dict, stream: stream}
	if obj.stream != nil {
		obj.Set("Length", fmt.Sprintf("%d", stream.Len()))
	}

	return obj
}

func (o *Obj) String() string {
	var s []string
	s = append(s, o.dict.String())
	if o.stream != nil {
		s = append(s, "stream\n", o.stream.String(), "\nendstream\n")
	}
	s = append(s, "endobj\n")

	return strings.Join(s, "")
}

func (o *Obj) Set(key string, val string) {
	o.dict.Set(key, val)
}

type Doc struct {
	next_id   int
	objs      []*Obj
	pages     []*Obj
	pages_obj *Obj
}

func NewDoc() *Doc {
	return &Doc{next_id: 1}
}

func (d *Doc) NextID() int {
	id := d.next_id
	d.next_id += 1
	return id
}

func (d *Doc) AddObject(obj *Obj) *Obj {
	d.objs = append(d.objs, obj)
	return obj
}

func (d *Doc) AddNewObject(m Dict, stream *bytes.Buffer) *Obj {
	return d.AddObject(NewObj(d.NextID(), m, stream))
}

func (d *Doc) String() string {
	var a []string
	var offsets []int
	var j0 int

	add := func(x string) {
		a = append(a, x)
		j0 += len(x) + 1
	}

	add("%PDF-1.4")
	for _, o := range d.objs {
		offsets = append(offsets, j0)
		add(fmt.Sprintf("%d 0 obj", o.id))
		add(o.String())
	}

	xrefstart := j0
	a = append(a, "xref")
	a = append(a, fmt.Sprintf("0 %d", len(offsets)+1))
	a = append(a, "0000000000 65535 f ")
	for _, o := range offsets {
		a = append(a, fmt.Sprintf("%010d 00000 n ", o))
	}
	a = append(a, "")
	a = append(a, "trailer")
	a = append(a, fmt.Sprintf("<< /Size %d\n/Root 1 0 R >>", len(offsets)+1))
	a = append(a, "startxref")
	a = append(a, fmt.Sprintf("%d", xrefstart))
	a = append(a, "%%EOF")

	return strings.Join(a, "\n")
}

type Image struct {
	name, format       string
	width, height      int
	color_space        string
	bits_per_component int
	filter             string
	decode_params      string
	data               *bytes.Buffer

	jbig2_page_no int
}

func NewJBIG2Image(name string, width, height int, data *bytes.Buffer) *Image {
	img := &Image{
		name:               name,
		format:             "jbig2",
		width:              width,
		height:             height,
		color_space:        "DeviceGray",
		bits_per_component: 1,
		filter:             "JBIG2Decode",
		data:               data}

	return img
}

func NewJPEGImage(name string, width, height int, cs string, data *bytes.Buffer) *Image {
	return &Image{
		name:               name,
		format:             "jpeg",
		width:              width,
		height:             height,
		color_space:        cs,
		bits_per_component: 8,
		filter:             "DCTDecode",
		data:               data,
	}
}

func (d *Doc) AddImagePage(img *Image) *Obj {
	var xres, yres float32
	xres = 72.0
	yres = 72.0

	width := float32(img.width*72) / xres
	height := float32(img.height*72) / yres

	xobj := d.AddNewObject(Dict{"Type": "/XObject", "Subtype": "/Image",
		"Width":            fmt.Sprintf("%d", img.width),
		"Height":           fmt.Sprintf("%d", img.height),
		"ColorSpace":       "/" + img.color_space,
		"BitsPerComponent": fmt.Sprintf("%d", img.bits_per_component),
		"Filter":           "/" + img.filter,
	}, img.data)
	if img.decode_params != "" {
		xobj.Set("DecodeParms", img.decode_params)
	}

	s := fmt.Sprintf("q %f 0 0 %f 0 0 cm /Im1 Do Q", width, height)
	contents := d.AddNewObject(Dict{}, bytes.NewBufferString(s))
	resources := d.AddNewObject(Dict{"ProcSet": "[/PDF /ImageB]",
		"XObject": fmt.Sprintf("<< /Im1 %d 0 R >>", xobj.id)}, nil)

	page := d.AddNewObject(Dict{"Type": "/Page",
		"Parent":    "3 0 R",
		"MediaBox":  fmt.Sprintf("[ 0 0 %f %f ]", width, height),
		"Contents":  ref(contents.id),
		"Resources": ref(resources.id)}, nil)
	d.pages = append(d.pages, page)
	return page
}

func (d *Doc) Start() {
	d.AddNewObject(Dict{"Type": "/Catalog", "Outlines": ref(2), "Pages": ref(3)}, nil)
	d.AddNewObject(Dict{"Type": "/Outlines", "Count": "0"}, nil)

	d.pages_obj = d.AddNewObject(Dict{"Type": "/Pages"}, nil)
}

func (d *Doc) Finish() {
	d.pages_obj.Set("Count", fmt.Sprintf("%d", len(d.pages)))

	var pids []string
	for _, page := range d.pages {
		pids = append(pids, ref(page.id))
	}

	d.pages_obj.Set("Kids", fmt.Sprintf("[%s]", strings.Join(pids, " ")))
}

// convert a list of image files to pdf
func NewPDF(files []string, sortNames bool) (*Doc, error) {
	var images []*Image
	for _, file := range files {
		var data bytes.Buffer
		r, err := os.Open(file)
		if err != nil {
			log.Printf("Failed to open %s, %v\n", file, err)
			return nil, err
		}

		if _, err = data.ReadFrom(r); err != nil {
			log.Printf("Failed to read from %s, %v\n", file, err)
			return nil, err
		}
		r.Close()

		bs := data.Bytes()
		config, format, err := image.DecodeConfig(bytes.NewReader(bs))
		if err != nil {
			log.Printf("Failed to decode header for %s, %v\n", file, err)
			continue
		}

		name := path.Base(file)
		var image *Image
		if format == "jpeg" {
			var cs string
			switch config.ColorModel {
			case color.GrayModel:
				cs = "DeviceGray"
			case color.YCbCrModel:
				cs = "DeviceRGB"
			case color.CMYKModel:
				cs = "DeviceCMYK"
			default:
			}
			//			log.Printf("new JPEG page %s: %d\n", name, data.Len())
			image = NewJPEGImage(name, config.Width, config.Height, cs, &data)
		} else if format == "png" {
			if pal, ok := config.ColorModel.(color.Palette); ok && len(pal) == 2 {
				image = NewJBIG2Image(name, config.Width, config.Height, &data)
			} else {
				log.Printf("Unsupported PNG: %s\n", file)
				return nil, fmt.Errorf("Unsupported PNG %s", file)
			}
		} else {
			log.Printf("Unsupported format: %s, %s\n", file, format)
			return nil, fmt.Errorf("Unsupported format %s (%s)", file, format)
		}

		if image == nil {
			log.Printf("No image found for %s\n", file)
			continue
		}

		images = append(images, image)
	}

	if sortNames {
		sort.Sort(ByImageName(images))
	}

	var numPages int
	ctx := C.jb2Init(0.85, 0.5, 0, 0, 0, -1)

	log.Printf("Processing %d images...\n", len(images))
	// adding all pages to jbig2 encoder
	for _, img := range images {
		if img.format != "jbig2" {
			continue
		}
		bs := img.data.Bytes()

		pix := C.pixReadMemPng((*C.uint8_t)(unsafe.Pointer(&bs[0])), C.size_t(len(bs)))
		if pix == nil {
			log.Printf("Failed to decode png: %s\n", img.name)
			return nil, fmt.Errorf("Failed to decode png: %s\n", img.name)
		}

		pixl := C.pixRemoveColormap(pix, 4)
		if pixl == nil {
			log.Printf("Failed to remove colormap: %s\n", img.name)
			return nil, fmt.Errorf("Failed to remove colormap: %s\n", img.name)
		}
		C.pixDestroy(&pix)

		C.jb2AddPage(ctx, pixl)
		C.pixDestroy(&pixl)

		log.Printf("adding page %d %s\n", numPages, img.name)
		img.jbig2_page_no = numPages + 1
		numPages += 1
	}

	log.Printf("Added %d images...\n", numPages)

	var sym []byte
	if numPages > 0 {
		var symlen C.int = 0
		csym := C.jb2PagesComplete(ctx, &symlen)
		if csym == nil {
			log.Printf("No sym found")
			return nil, fmt.Errorf("No sym")
		}

		log.Printf("symlen %d\n", symlen)
		sym = C.GoBytes(unsafe.Pointer(csym), symlen)
		C.free(unsafe.Pointer(csym))

		for _, img := range images {
			if img.jbig2_page_no == 0 {
				continue
			}
			var datalen C.int = 0
			data := C.jb2ProducePage(ctx, C.int(img.jbig2_page_no-1), -1, -2, &datalen)
			if data == nil {
				log.Printf("Failed to produce page for %s\n", img.name)
				return nil, fmt.Errorf("Failed to produce page for %s\n", img.name)
			}

			img.data = bytes.NewBuffer(C.GoBytes(unsafe.Pointer(data), datalen))
			//		log.Printf("page %s len %d\n", img.name, datalen)
			C.free(unsafe.Pointer(data))
		}
	}

	doc := NewDoc()
	doc.Start()

	var symid int
	if len(sym) > 0 {
		symd := doc.AddNewObject(Dict{}, bytes.NewBuffer(sym))
		symid = symd.id
	}

	// output pages
	for _, img := range images {
		if img.format == "jbig2" {
			img.decode_params = fmt.Sprintf(" << /JBIG2Globals %d 0 R >>", symid)
		}

		doc.AddImagePage(img)
	}

	doc.Finish()
	return doc, nil
}

const NameLen = len("000001.pdg")

var ss = flag.Bool("ss", false, "enable ssreader mode")

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		log.Printf("Usage: %s dir/zipfile/imagefile\n", os.Args[0])
		return
	}

	target := args[0]
	info, err := os.Stat(target)
	if err != nil {
		log.Fatal(err)
	}

	var files []string
	if info.IsDir() {
		fis, err := ioutil.ReadDir(target)
		if err != nil {
			log.Fatal(err)
		}
		for _, fi := range fis {
			if fi.Mode().IsRegular() {
				files = append(files, filepath.Join(target, fi.Name()))
			}
		}
	} else if info.Mode().IsRegular() {
		ext := strings.ToLower(filepath.Ext(target))
		if ext == "zip" {
			log.Printf("Extract files from %s\n", target)
			dir := path.Base(target) + "_tmp"
			err = os.Mkdir(dir, 0755)
			if err != nil {
				log.Fatal(err)
			}

			files, err = extractZipFiles(target, dir)
			if err != nil {
				log.Fatal(err)
			}
			defer os.RemoveAll(dir)

		} else {
			files = append(files, target)
		}
	}

	var images []string
	for _, file := range files {
		if *ss && len(path.Base(file)) != NameLen {
			log.Printf("Skipping file %s\n", file)
			continue
		}

		images = append(images, file)
	}

	log.Printf("Processing %d files from %s\n", len(images), target)

	sortNames := true
	doc, err := NewPDF(images, sortNames)
	if err != nil {
		log.Fatal(err)
	}

	outfile := strings.TrimSuffix(path.Base(target), filepath.Ext(target)) + ".pdf"
	log.Printf("Outputing file to %s\n", outfile)
	ioutil.WriteFile(outfile, []byte(doc.String()), 0644)
}

//pdg
func tag(name string) int {
	if len(name) < 2 {
		return 10
	}
	switch name[:2] {
	case "bo":
		return 1 //bok
	case "co":
		return 2 //cov
	case "fo":
		return 3 //fow
	case "!0":
		return 4 //toc
	case "00":
		fallthrough
	case "01":
		fallthrough
	case "02":
		return 5
	case "le":
		return 7 //leg
	}

	return 10
}

type ByImageName []*Image

func (a ByImageName) Len() int      { return len(a) }
func (a ByImageName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByImageName) Less(i, j int) bool {
	ti, tj := tag(a[i].name), tag(a[j].name)
	if ti != tj {
		return ti < tj
	}
	return a[i].name < a[j].name
}

//zip
func extractZipFiles(zipfile, dir string) ([]string, error) {
	var files []string
	r, err := zip.OpenReader(zipfile)
	if err != nil {
		return files, err
	}
	defer r.Close()

	for _, f := range r.File {
		fpath := filepath.Join(dir, path.Base(f.Name))

		if f.FileInfo().IsDir() {
			continue
		}

		files = append(files, fpath)

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return files, err
		}

		rc, err := f.Open()
		if err != nil {
			return files, err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()

		if err != nil {
			return files, err
		}
	}
	return files, nil
}