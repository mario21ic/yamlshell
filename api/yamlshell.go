package api

import (
	"os"
	"os/exec"

	"github.com/mattn/go-shellwords"
	"gopkg.in/yaml.v2"
)

// CommandList represents a list of commands from a yaml file.
type CommandList struct {
	Run []struct {
		WorkDir  string   `yaml:"workdir,omitempty"`
		Commands []string `yaml:"commands"`
	} `yaml:"run"`
}

// GetCommands will unmarshal the commands from a buffer.
func GetCommands(buf []byte) CommandList {
	var commandList CommandList

	yaml.Unmarshal(buf, &commandList)
	return commandList
}

// RunCommands will run the commands from a CommandList
func RunCommands(commandList CommandList) ([][]byte, error) {
	var out [][]byte
	p := shellwords.NewParser()
	p.ParseEnv = true
	p.ParseBacktick = true
	for _, commandSet := range commandList.Run {
		if commandSet.WorkDir != "" {
			os.Chdir(commandSet.WorkDir)
		}

		for _, v := range commandSet.Commands {

			args, err := p.Parse(v)
			if err != nil {
				return out, err
			}

			command, args := args[0], args[1:]
			cmdOut, err := exec.Command(command, args...).CombinedOutput()
			out = append(out, cmdOut)

			if err != nil {
				return out, err
			}
		}

	}

	return out, nil
}
