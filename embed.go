package app

import (
	"embed"
	"github.com/gin-contrib/static"
	"io/fs"
	"log"
	"net/http"
)

//go:embed _ui/build
var UI embed.FS

type embedFileSystem struct {
	http.FileSystem
	indexes bool
}

func (e embedFileSystem) Exists(prefix string, path string) bool {
	f, err := e.Open(path)
	if err != nil {
		return false
	}

	s, _ := f.Stat()
	if s.IsDir() && !e.indexes {
		// Indexing not allowed
		return false
	}

	return true
}

func EmbedFolder() static.ServeFileSystem {
	uiFS, err := fs.Sub(UI, "_ui/build")
	if err != nil {
		log.Fatal("failed to get ui fs", err)
	}
	return embedFileSystem{
		FileSystem: http.FS(uiFS),
		indexes:    true,
	}
}
