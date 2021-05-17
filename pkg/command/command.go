package command

import (
	"flag"
	"os/exec"
)

func GetCommandOut() string {
	var command string
	var parse []string

	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		panic("Please give command")
	}

	for i, arg := range args {
		if i == 0 {
			command = arg
		} else {
			parse = append(parse, arg)
		}

	}

	cmd := exec.Command(command, parse...)
	out, err := cmd.Output()
	if err != nil {
		panic("command error")
	}

	return string(out)
}
