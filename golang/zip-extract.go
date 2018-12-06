// https://gist.github.com/svett/424e6784facc0ba907ae

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func unzip(archive, target string) error {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}

	for _, file := range reader.File {
		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}

	return nil
}

The extract function will generate too many open file issue as the defer is only called once the function return, here an altered version:

func Unzip(archive, target string) error {
    reader, err := zip.OpenReader(archive)
    if err != nil {
        return err
    }

    if err := os.MkdirAll(target, 0755); err != nil {
        return err
    }

    for _, file := range reader.File {
        path := filepath.Join(target, file.Name)
        if file.FileInfo().IsDir() {
            os.MkdirAll(path, file.Mode())
            continue
        }

        fileReader, err := file.Open()
        if err != nil {

            if fileReader != nil {
                fileReader.Close()
            }

            return err
        }

        targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
        if err != nil {
            fileReader.Close()

            if targetFile != nil {
                targetFile.Close()
            }

            return err
        }

        if _, err := io.Copy(targetFile, fileReader); err != nil {
            fileReader.Close()
            targetFile.Close()

            return err
        }

        fileReader.Close()
        targetFile.Close()
    }

    return nil
}