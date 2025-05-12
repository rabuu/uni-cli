package cfg

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
		ExportDirectory string
		Semester string
		Courses map[string]Course
	}

	Course struct {
		Name string
		Prefix string `toml:",omitempty"`
		ExportFile []FileMap
		ExportZip []ZipConfig
		GroupDescription string `toml:",omitempty"`
		Members []GroupMember
		Tutor string `toml:",omitempty"`
		Link string `toml:",omitempty"`
	}

	FileMap struct {
		From string
		To string
	}

	ZipConfig struct {
		ArchiveFile string
		Include []FileMap
	}

	GroupMember struct {
		First string
		Last string
		ID string
	}
)

func ParseConfig(path string, uniDirectory string) Config {
	var config Config
	_, err := toml.DecodeFile(path, &config)
	exit.ExitWithErr(err)

	// validate export directory
	if config.ExportDirectory == "" {
		exit.ExitWithMsg("No export directory is specified.")
	}
	exportDirPath := filepath.Join(uniDirectory, config.ExportDirectory)
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
		if course.Name != "" {
			fmt.Printf(" (%s)", course.Name)
		}
		fmt.Println()
	}
}

func (config *Config) PrintCoursesFishCompletion() {
	for name, course := range config.Courses {
		fmt.Print(name)
		if course.Name != "" {
			fmt.Printf("\t%s", course.Name)
		}
		fmt.Println()
	}
}
