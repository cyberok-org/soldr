package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	directory = "soldr-modules"
)

type (
	Config struct {
		Modules []Module `json:"modules"`
	}

	Module struct {
		Url    string `json:"url"`
		Branch string `json:"branch"`
		Tag    string `json:"tag"`
	}
)

func main() {
	tmpfile, err := os.ReadFile("soldr-modules.json")
	if err != nil {
		panic(err)
	}

	cfg := &Config{}

	err = json.Unmarshal(tmpfile, cfg)
	if err != nil {
		panic(err)
	}

	for _, module := range cfg.Modules {
		s := strings.Split(module.Url, "/")
		fmt.Println("Delete ", s[len(s)-1])
		cmd := exec.Command("rm", "-R", directory+"/"+s[len(s)-1])
		err = cmd.Run()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Cloning ", module.Url, module.Branch)
		args := make([]string, 0)
		args = append(args, "clone")
		if module.Branch != "" {
			args = append(args, "-b")
			args = append(args, module.Branch)
		}

		if module.Tag != "" {
			args = append(args, "-b")
			args = append(args, module.Tag)
		}

		args = append(args, module.Url)

		cmd = exec.Command("git", args...)
		cmd.Dir = directory
		err = cmd.Run()
		if err != nil {
			panic(err)
		}
	}
	println("Download modules finish successful")
}
