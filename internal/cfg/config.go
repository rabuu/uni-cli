package cfg

import (
	"bytes"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type (
	Config struct {
		Courses map[string]Course `toml:"courses"`
	}

	Course struct {
		FullName string `toml:"full-name"`
		Prefix string `toml:"prefix"`
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

	if config.Courses == nil {
		config.Courses = make(map[string]Course)
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
	_, ok := config.Courses[name]
	return ok
}

func (config *Config) PrintCourses() {
	if len(config.Courses) == 0 {
		fmt.Println("There are no registered courses.")
		return
	}

	fmt.Println("Courses:")
	for name, course := range config.Courses {
		fmt.Printf("  - %s", name)
		if course.FullName != "" {
			fmt.Printf(" (%s)", course.FullName)
		}
		fmt.Println()
	}
}
