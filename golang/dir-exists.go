import "os"

func DirExists(path string) bool {
	if fi, err := os.Stat(path); err == nil {
		return fi.Mode().IsDir()
	}
	return false
}