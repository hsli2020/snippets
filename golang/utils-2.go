package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

func fileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		// Schrodinger, but return false anyway
		return false
	}
}

func Sha256String(str string) (string, error) {
	hasher := sha256.New()
	_, err := hasher.Write([]byte(str))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func HashOfImage(path string) (string, error) {
	hasher := sha256.New()
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	hasher.Write(file)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func hash(v interface{}) string {
	j, _ := json.Marshal(v)
	hash := sha256.Sum256(j)
	return hex.EncodeToString(hash[:])
}

func subPlaceholders(str, id string) string {
	return strings.ReplaceAll(str, "{id}", id)
}

func Dump(v interface{}) {
	j, _ := json.Marshal(v)
	fmt.Println(string(j))
}

func Hash(v interface{}) [32]byte {
	j, _ := json.Marshal(v)
	return sha256.Sum256(j)
}

func LsDirs(path string) ([]string, error) {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		return nil, err
	}

	var dirs []string

	for _, f := range files {
		if strings.HasPrefix(f.Name(), ".") {
			continue
		}

		if !f.IsDir() {
			return nil, fmt.Errorf("directory should contain only directories, found: %s", f.Name())
		}

		dirs = append(dirs, f.Name())
	}

	return dirs, nil
}

func LsFiles(path string) ([]string, error) {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		return nil, err
	}

	var actualFiles []string

	for _, f := range files {
		if strings.HasPrefix(f.Name(), ".") {
			continue
		}

		if f.IsDir() {
			return nil, fmt.Errorf("directory should contain only files, found: %s", f.Name())
		}

		actualFiles = append(actualFiles, f.Name())
	}

	return actualFiles, nil
}

func WriteFileAsJson(v interface{}, out string) error {
	file, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("error when marshaling json: %w", err)
	}

	err = ioutil.WriteFile(out, file, 0644)
	if err != nil {
		return fmt.Errorf("error when generating json: %w", err)
	}
	return nil
}

func WriteFileAsYaml(v interface{}, out string) error {
	file, err := yaml.Marshal(v)
	if err != nil {
		return fmt.Errorf("error when marshaling json: %w", err)
	}

	err = ioutil.WriteFile(out, file, 0644)
	if err != nil {
		return fmt.Errorf("error when generating json: %w", err)
	}
	return nil
}

func LoadJsonFileIntoStruct(path string, v interface{}) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read file at path '%s': %w", path, err)
	}

	err = json.Unmarshal(file, v)
	if err != nil {
		return fmt.Errorf("failed to unmarshal attrs file at path '%s': %w", path, err)
	}

	return nil
}

func LoadYamlFileIntoStruct(path string, v interface{}) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read file at path '%s': %w", path, err)
	}

	err = yaml.Unmarshal(file, v)
	if err != nil {
		return fmt.Errorf("failed to unmarshal attrs file at path '%s': %w", path, err)
	}

	return nil
}
