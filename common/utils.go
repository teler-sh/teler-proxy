package common

import (
	"fmt"
	"os"
)

func PrintBanner() {
	fmt.Fprintf(os.Stderr, "%s\n\n", Banner)
}

func PrintUsage() {
	fmt.Fprint(os.Stderr, Usage)
}

func PrintVersion() {
	version := Version
	if version == "" {
		version = "unknown (go-get)"
	}

	fmt.Println(App, "version", version)
}
