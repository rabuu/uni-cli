package cfg

import (
	"bytes"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type (
	Config struct {
		Courses []Course `toml:"course"`
	}

	Course struct {
		Name string `toml:"name"`
		LongName string `toml:"long-name"`
		Group []GroupMember `toml:"member"`
	}

	GroupMember struct {
		Firstname string `toml:"firstname"`
		Lastname string `toml:"lastname"`
		Id string `toml:"id"`
	}
)

func ParseConfig(path string) Config {
	var config Config
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	return config
}

func (config *Config) WriteToFile(path string) {
	var buf bytes.Buffer
	err := toml.NewEncoder(&buf).Encode(config)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	os.WriteFile(path, buf.Bytes(), 0644)
}

func (config *Config) ContainsCourse(name string) bool {
	for i := 0; i < len(config.Courses); i++ {
		if config.Courses[i].Name == name {
			return true
		}
	}

	return false
}

func (config *Config) PrintCourses() {
	fmt.Println("Courses:")
	for i := 0; i < len(config.Courses); i++ {
		course := config.Courses[i]
		if course.LongName == "" {
			fmt.Printf("  %d. %s\n", i + 1, course.Name)
		} else {
			fmt.Printf("  %d. %s (%s)\n", i + 1, course.Name, course.LongName)
		}
	}
}
