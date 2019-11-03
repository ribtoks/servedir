package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// flags
var (
	dirPathFlag = flag.String("d", "./", "Path to directory to serve")
	portFlag    = flag.Int("p", 8080, "Port to use")
	rootFlag    = flag.String("r", "/", "Http path to serve")
)

type FileSystem struct {
	fs http.FileSystem
}

func GetFileSystemHandler(directory, path string) http.HandlerFunc {
	fileServer := http.FileServer(FileSystem{http.Dir(directory)})
	return http.StripPrefix(strings.TrimRight(path, "/"), fileServer).ServeHTTP
}

func (fs FileSystem) Open(path string) (http.File, error) {
	f, err := fs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := strings.TrimSuffix(path, "/") + "/index.html"
		if _, err := fs.fs.Open(index); err != nil {
			return nil, err
		}
	}

	return f, nil
}

func main() {
	flag.Parse()
	dirInfo, err := os.Stat(*dirPathFlag)
	if os.IsNotExist(err) {
		os.Stderr.WriteString("Specified path cannot be found")
		os.Exit(1)
	}
	if !dirInfo.IsDir() {
		os.Stderr.WriteString("Specified path is not a directory")
		os.Exit(1)
	}

	fmt.Printf("Listening at http://localhost:%d", *portFlag)
	http.Handle("/", GetFileSystemHandler(*dirPathFlag, *rootFlag))
	http.ListenAndServe(fmt.Sprintf(":%d", *portFlag), nil)
}
