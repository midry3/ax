package index

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const (
	INDEX_FILE = ".axconf"
)

var (
	index map[string]string
)

func Lookup(name string) string {
	v, ok := index[name]
	if ok {
		return v
	}
	return ""
}

func load() {
	dir, err := os.UserHomeDir()
	if err != nil {
		dir = "."
	}
	path := filepath.Join(dir, INDEX_FILE)
	b, err := os.ReadFile(path)
	if err != nil {
		os.WriteFile(path, []byte{}, os.ModePerm)
		index = map[string]string{}
		return
	}
	err = json.Unmarshal(b, &index)
	if err != nil {
		panic(err)
	}
}

func init() {
	load()
}
