package cmd

import (
	"io/ioutil"
	"log"
	"os"

	yamlshell "github.com/drud/yamlshell/api"
	"github.com/spf13/cobra"
)

// YamlSH runs commands from a yaml file.
var YamlSH = &cobra.Command{
	Use:   "yamlshell [filename]",
	Short: "yamlshell is a way to run shell commands from a yaml file.",
	Long:  `yamlshell is a way to run shell commands from a yaml file.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("You must provide a file containing yaml commands.")
		}

		if _, err := os.Stat(args[0]); os.IsNotExist(err) {
			log.Fatalf("Cannot read from file: %s", args[0])
		}

		buf, err := ioutil.ReadFile(args[0])

		if err != nil {
			log.Fatalf("Could not read file: %s", err)
		}

		commands := yamlshell.GetCommands(buf)
		out, err := yamlshell.RunCommands(commands)
		for _, v := range out {
			log.Println(string(v))
		}

		if err != nil {
			log.Fatalf("Error running commands: %s", err)
		}
	},
}
