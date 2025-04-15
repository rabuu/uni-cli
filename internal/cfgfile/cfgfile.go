package cfgfile

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/rabuu/uni-cli/internal/exit"
)

type (
	Config struct {
		ExportDir string `toml:"export_directory"`
		Semester string `toml:"semester"`
		Courses map[string]Course `toml:"courses"`
	}

	Course struct {
		FullName string `toml:"full-name"`
		Prefix string `toml:"prefix,omitempty"`
		Export []Export `toml:"export"`
		Members []GroupMember `toml:"members"`
		Tutor string `toml:"tutor,omitempty"`
		Link string `toml:"link,omitempty"`
	}

	Export struct {
		Filename string `toml:"filename"`
		Output string `toml:"output"`
	}

	GroupMember struct {
		Firstname string `toml:"firstname"`
		Lastname string `toml:"lastname"`
		Id string `toml:"id"`
	}
)

func ParseConfig(path string, uniDirectory string) Config {
	var config Config
	_, err := toml.DecodeFile(path, &config)
	exit.ExitWithErr(err)

	// validate export directory
	if config.ExportDir == "" {
		exit.ExitWithMsg("No export directory is specified.")
	}
	exportDirPath := filepath.Join(uniDirectory, config.ExportDir)
	exportDirInfo, err := os.Stat(exportDirPath)
	if os.IsNotExist(err) {
		err := os.Mkdir(exportDirPath, 0755)
		exit.ExitWithErr(err)
		exportDirInfo, err = os.Stat(exportDirPath)
		exit.ExitWithErr(err)
	} else {
		exit.ExitWithErr(err)
	}
	if !exportDirInfo.IsDir() {
		exit.ExitWithMsg("Specified export directory is no directory.")
	}

	if config.Courses == nil {
		config.Courses = make(map[string]Course)
	}

	return config
}

func (config *Config) WriteToFile(path string) {
	var buf bytes.Buffer
	err := toml.NewEncoder(&buf).Encode(config)
	exit.ExitWithErr(err)

	os.WriteFile(path, buf.Bytes(), 0644)
}

func (config *Config) PrintToStdout() {
	err := toml.NewEncoder(os.Stdout).Encode(config)
	exit.ExitWithErr(err)
}

func (config *Config) ContainsCourse(name string) bool {
	_, ok := config.Courses[name]
	return ok
}

func (config *Config) PrintCoursesHumanReadable() {
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

func (config *Config) PrintCoursesFishCompletion() {
	for name, course := range config.Courses {
		fmt.Print(name)
		if course.FullName != "" {
			fmt.Printf("\t%s", course.FullName)
		}
		fmt.Println()
	}
}
