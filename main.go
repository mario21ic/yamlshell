package main

import (
	"fmt"
	"os"

	"github.com/drud/yamlshell/cmd"
)

func main() {
	if err := cmd.YamlSH.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
