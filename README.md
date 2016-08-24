[![Go Report Card](https://goreportcard.com/badge/github.com/drud/yamlshell)](https://goreportcard.com/report/github.com/drud/yamlshell)

# yamlshell

yamlshell is a golang package for running arbitrary commands from a yaml file.

## Binary Usage

To use the yamlshell binary, run `make all` and create a command file:

```
run:
    # Run from current working directory
    - commands: 
      - "echo $FOO"
    # Execute from another location
    - workdir: "/path/to/working/directory"
      commands: 
      - ls -alh
```

You can then run the commands within via `yamlshell example.yaml`

## API Usage

API Usage is dead simple, your best bet is to just refer to the [cobra command file](https://github.com/drud/yamlshell/blob/master/cmd/run.go).
