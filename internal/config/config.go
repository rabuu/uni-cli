package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type (
	Config struct {
		Courses []Course `toml:"course"`
	}

	Course struct {
		Name string
		LongName string `toml:"long-name"`
		Group []GroupMember
	}

	GroupMember struct {
		Firstname string
		Lastname string
		Id string
	}
)

func Parse(path string) Config {
	var config Config
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	return config
}
