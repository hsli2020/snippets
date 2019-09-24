package u

import (
	//"crypto/sha1"
	//"fmt"
	"io"
	"os"
    "path/filepath"
	"strings"
	"time"
)

// FileExists returns true if a given path exists and is a file
func FileExists(path string) bool {
	st, err := os.Stat(path)
	return err == nil && !st.IsDir() && st.Mode().IsRegular()
}

func ExistFile(name string) bool {
    _, err := os.Stat(name)
    return !os.IsNotExist(err)
}

// DirExists returns true if a given path exists and is a directory
func DirExists(path string) bool {
	st, err := os.Stat(path)
	return err == nil && st.IsDir()
}

// PathIsDir returns true if a path exists and is a directory
// Returns false, nil if a path exists and is not a directory (e.g. a file)
// Returns undefined, error if there was an error e.g. because a path doesn't exists
func IsDir(path string) (isDir bool, err error) {
	fi, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fi.IsDir(), nil
}

// GetFileSize returns size of the file
func GetFileSize(path string) (int64, error) {
	fi, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}
	return fi.Size(), nil
}

// CopyFile copies a file
func CopyFile(dst, src string) error {
	fsrc, err := os.Open(src)
	if err != nil {
		return err
	}
	defer fsrc.Close()
	fdst, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer fdst.Close()
	if _, err = io.Copy(fdst, fsrc); err != nil {
		return err
	}
	return nil
}

// Save user uploaded file to target dir
func SaveFile(filename, dir string) (string, error) {
    if dir == "" {
        dir = time.Now().Format("2006/01/02/15/04")
        os.MkdirAll(dir, 0755)
    }

    fin, err := os.OpenFile(filename, os.O_RDONLY, 0755)
	if err != nil {
        return "", err
	}
    defer fin.Close()

    ext := filepath.Ext(filename)

    // create sha for file name
    //h := sha1.New()
    //io.Copy(h, fin)
    //fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext

    fname := time.Now().Format("20060102-150405.000000") + ext
    path := filepath.Join(dir, fname)

    // create new file
    fout, err := os.Create(path)
    if err != nil {
        return "", err
    }
    defer fout.Close()

    // copy
    fin.Seek(0, 0)
    io.Copy(fout, fin)

    return path, nil
}

func BackupFile(filename, dir string) (string, error) {
    file, err := os.Stat(filename)
	if err != nil {
        return "", err
	}
    modTime := file.ModTime().Format("20060102-150405")

    ext := filepath.Ext(filename)
    name := strings.TrimSuffix(filename, ext)

    newName := name + "-" + modTime + ext
    if dir != "" {
        newName = filepath.Join(dir, newName)
    }

    err = os.Rename(filename, newName)
    if err != nil {
        return "", err
    }

    return newName, nil
}

func ArchiveFile(filename string) (string, error) {
    return BackupFile(filename, "archive")
}
