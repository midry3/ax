package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/midry3/ax/internal/index"
)

func show_bash_init_func() {
	fmt.Printf(`ax() {
	dir=$(ax-lookup "$1") && cd "$dir"
}`)
}

func show_help() {
	fmt.Println(`[description]
Change directory tool with named path.

[usage]
ax [path or name]`)
}

func isDir(path string) bool {
	i, err := os.Stat(path)
	if err != nil {
		return false
	}
	return i.IsDir()
}

func main() {
	if slices.Contains(os.Args, "-h") || slices.Contains(os.Args, "--help") {
		show_help()
		return
	}
	switch len(os.Args) {
	case 1:
		show_bash_init_func()
	case 2:
		name := os.Args[1]
		if isDir(name) || name == "-" {
			fmt.Printf(name)
		} else {
			name = index.Lookup(strings.TrimLeft(name, "@"))
			if name == "" {
				os.Exit(1)
			}
			fmt.Printf(name)
		}
	default:
		fmt.Fprintln(os.Stderr, "too many args!\n")
		os.Exit(1)
	}
}
