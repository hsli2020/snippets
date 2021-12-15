// ===== main.go
package main

import (
	"fmt"
	"os"

	"github.com/wilsontwm/filezy/cmd"
)

func main() {
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

// ===== model/file.go
package model

import (
	"path/filepath"
	"strings"
)

type File struct {
	FullPath string // The full path including the folder
	Folder   string // The folder of the file
	File     string // The file name including extension
	FileName string // The file name excluding extension
	Ext      string // Extension of the file
}

// Construct from string
func ConstructFile(path string) File {
	base := filepath.Base(path)
	ext := filepath.Ext(path)

	return File{
		FullPath: path,
		File:     base,
		Folder:   strings.TrimSuffix(path, base),
		FileName: strings.TrimSuffix(base, ext),
		Ext:      ext,
	}
}

// ===== cmd/cmd.go
package cmd

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"

	"github.com/wilsontwm/filezy/helper"
	"github.com/wilsontwm/filezy/model"
)

var isRecursive bool
var folder string
var prefix string
var suffix string
var regexPattern string
var extension string
var enableLog bool

var RootCmd = &cobra.Command{
	Use:     "filezy",
	Short:   "filezy is a CLI-based file management tool",
	Version: "0.0.1",
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&isRecursive, "recursive", "r", false, "Scan files in sub-directories recursively")
	RootCmd.PersistentFlags().StringVarP(&prefix, "prefix", "p", "", "Return files that have the specified prefix in the file name")
	RootCmd.PersistentFlags().StringVarP(&suffix, "suffix", "s", "", "Return files that have the specified suffix in the file name")
	RootCmd.PersistentFlags().StringVarP(&regexPattern, "regex", "x", "", "Return files that match the regex pattern in the file name")
	RootCmd.PersistentFlags().StringVarP(&extension, "ext", "e", "", "Return files that have the specified extension")
	RootCmd.PersistentFlags().BoolVarP(&enableLog, "log", "l", false, "Print logs")
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

// Get the filtered files
func getFilteredFiles(folder string) ([]model.File, error) {
	files := make([]model.File, 0)
	if filesInFolder, err := helper.GetFiles(folder, isRecursive); err == nil {
		for _, file := range filesInFolder {
			if prefix != "" && !strings.HasPrefix(file.FileName, prefix) {
				continue
			} else if suffix != "" && !strings.HasSuffix(file.FileName, suffix) {
				continue
			} else if regexPattern != "" && !regexp.MustCompile(regexPattern).MatchString(file.FileName) {
				continue
			} else if extension != "" && strings.TrimSuffix(file.Ext, extension) != "." {
				continue
			}

			files = append(files, file)
		}
	} else {
		return files, err
	}

	return files, nil
}

// ===== cmd/compress.go
package cmd

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"

	"github.com/wilsontwm/filezy/model"
)

var compressCmd = &cobra.Command{
	Use:   "compress [filename]",
	Short: "Compress files in folder",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]

		files, err := getFilteredFiles(folder)
		must(err)

		err = compressFiles(filename, files)
		must(err)
	},
}

func init() {
	compressCmd.Flags().StringVarP(&folder, "folder", "f", "./", "Target folder to be scanned")

	RootCmd.AddCommand(compressCmd)
}

func compressFiles(filename string, files []model.File) error {
	zipfile, err := os.Create(filename + ".zip")
	if err != nil {
		return err
	}
	defer zipfile.Close()

	zipWriter := zip.NewWriter(zipfile)
	defer zipWriter.Close()

	// Switch to targeted folder
	if folder != "" {
		os.Chdir(folder)
	}

	currentFolder, err := os.Getwd()
	if err != nil {
		return err
	}

	for _, file := range files {
		relativeFolder, err := filepath.Rel(currentFolder, file.Folder)
		if err != nil {
			return err
		}

		path := filepath.Join(relativeFolder, file.File)
		if err = addFileToArchive(zipWriter, path); err != nil {
			return err
		}

		if enableLog {
			fmt.Printf("%v: Compress %v\n", time.Now().Format(time.RFC3339), path)
		}
	}

	if enableLog {
		fmt.Printf("Output zip file: %v\n", filename+".zip")
	}

	return nil
}

func addFileToArchive(zipWriter *zip.Writer, filename string) error {
	fileToZip, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	// Get the file information
	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	header.Name = filename
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)

	return err
}

// ===== cmd/copy.go
package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"

	"github.com/wilsontwm/filezy/helper"
	"github.com/wilsontwm/filezy/model"
)

var copyCmd = &cobra.Command{
	Use:   "copy [source folder] [target folder]",
	Short: "Copy files from source folder to target folder",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		sourceFolder := args[0]
		targetFolder := args[1]

		source, err := filepath.Abs(sourceFolder)
		must(err)
		target, err := filepath.Abs(targetFolder)
		must(err)

		// Check if source folder and target folder exists
		if _, err := os.Stat(source); os.IsNotExist(err) {
			must(fmt.Errorf("Source folder %v does not exists.", sourceFolder))
		}
		if _, err := os.Stat(target); os.IsNotExist(err) {
			must(fmt.Errorf("Target folder %v does not exists.", targetFolder))
		}

		files, err := getFilteredFiles(source)
		must(err)

		for _, file := range files {
			newFileName, err := helper.GetNewFilePath(file, source, target)
			must(err)

			newFile := model.ConstructFile(newFileName)
			err = copy(file, newFile)
			must(err)

			if enableLog {
				fmt.Printf("%v: Copy %v --> %v\n", time.Now().Format(time.RFC3339), file.FullPath, newFileName)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(copyCmd)
}

func copy(src, dest model.File) error {
	source, err := os.Open(src.FullPath)
	if err != nil {
		return err
	}
	defer source.Close()

	// Create the target folder at destination if not exists (for subfolder only)
	// Root target folder must exists, else error will be thrown. Validation is performed earlier on
	if _, err := os.Stat(dest.Folder); os.IsNotExist(err) {
		if err = os.MkdirAll(dest.Folder, os.ModePerm); err != nil {
			return err
		}
	}

	destination, err := os.Create(dest.FullPath)

	if err != nil {
		return err
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)

	return err
}

// ===== cmd/move.go
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"

	"github.com/wilsontwm/filezy/helper"
	"github.com/wilsontwm/filezy/model"
)

var moveCmd = &cobra.Command{
	Use:   "move [source folder] [target folder]",
	Short: "Move files from source folder to target folder",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		sourceFolder := args[0]
		targetFolder := args[1]

		source, err := filepath.Abs(sourceFolder)
		must(err)
		target, err := filepath.Abs(targetFolder)
		must(err)

		// Check if source folder and target folder exists
		if _, err := os.Stat(source); os.IsNotExist(err) {
			must(fmt.Errorf("Source folder %v does not exists.", sourceFolder))
		}
		if _, err := os.Stat(target); os.IsNotExist(err) {
			must(fmt.Errorf("Target folder %v does not exists.", targetFolder))
		}

		files, err := getFilteredFiles(source)
		must(err)

		for _, file := range files {
			newFileName, err := helper.GetNewFilePath(file, source, target)
			must(err)

			newFile := model.ConstructFile(newFileName)
			err = move(file, newFile)
			must(err)

			if enableLog {
				fmt.Printf("%v: Move %v --> %v\n", time.Now().Format(time.RFC3339), file.FullPath, newFileName)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(moveCmd)
}

func move(src, dest model.File) error {
	// Create the target folder at destination if not exists (for subfolder only)
	// Root target folder must exists, else error will be thrown. Validation is performed earlier on
	if _, err := os.Stat(dest.Folder); os.IsNotExist(err) {
		if err = os.MkdirAll(dest.Folder, os.ModePerm); err != nil {
			return err
		}
	}

	return os.Rename(src.FullPath, dest.FullPath)
}

// ===== cmd/rename.go
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/wilsontwm/filezy/helper"
)

var renameCmd = &cobra.Command{
	Use:   "rename [filename]",
	Short: "Rename files in batch, auto-increment number will be added as suffix",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]

		files, err := getFilteredFiles(folder)
		must(err)

		total := len(files)
		noOfDigits := helper.NumberOfDigits(total)
		for i, file := range files {
			newFileName := file.Folder + filename + file.Ext
			numberStr := helper.ToString(i+1, noOfDigits)
			if total > 1 {
				newFileName = file.Folder + filename + "-" + numberStr + file.Ext
			}

			err := os.Rename(file.FullPath, newFileName)
			must(err)

			if enableLog {
				fmt.Printf("%v: Rename %v --> %v\n", time.Now().Format(time.RFC3339), file.FullPath, newFileName)
			}
		}
	},
}

func init() {
	renameCmd.Flags().StringVarP(&folder, "folder", "f", "./", "Target folder to be scanned")

	RootCmd.AddCommand(renameCmd)
}

// ===== helper/file.go
package helper

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/wilsontwm/filezy/model"
)

// HasFile : Check if file exists in the current directory
func HasFile(filename string) bool {
	if info, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	} else {
		return !info.IsDir()
	}
}

// GetFiles :
func GetFiles(folder string, isRecursive bool) ([]model.File, error) {
	var files []model.File

	if isRecursive {
		err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
			file, err := filepath.Abs(path)
			if err != nil {
				return nil
			}

			if !info.IsDir() {
				files = append(files, model.ConstructFile(file))
			}
			return nil
		})

		return files, err
	} else {
		f, err := os.Open(folder)
		defer f.Close()

		if err != nil {
			return files, err
		}

		if fileinfo, err := f.Readdir(-1); err == nil {
			for _, file := range fileinfo {
				if !file.IsDir() {
					folder, err := filepath.Abs(folder)
					if err != nil {
						return files, err
					}
					files = append(files, model.ConstructFile(folder+"\\"+file.Name()))
				}
			}
		} else {
			return files, err
		}
	}

	return files, nil
}

// GetNewFilePath :
func GetNewFilePath(file model.File, sourceFolder, targetFolder string) (string, error) {
	if file.FullPath == "" {
		return "", fmt.Errorf("Original file is not valid.")
	}

	subfolder := GetSubfolder(file, sourceFolder)

	return targetFolder + subfolder + file.File, nil
}

func GetSubfolder(file model.File, sourceFolder string) string {
	return strings.TrimPrefix(file.Folder, sourceFolder)
}

// ===== helper/number.go
package helper

import (
	"strconv"
)

func NumberOfDigits(number int) int {
	if number < 10 {
		return 1
	}

	return 1 + NumberOfDigits(number/10)
}

func ToString(number int, digits ...int) string {
	str := strconv.Itoa(number)
	if len(digits) == 0 || digits[0] <= 0 {
		return str
	}

	digit := digits[0]
	for i := len(str); i < digit; i++ {
		str = "0" + str
	}

	return str
}
