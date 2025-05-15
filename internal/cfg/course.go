package cfg

type (
	Course struct {
		Name string
		Prefix string `toml:",omitempty"`
		RetrieveFile []FileMap
		ExportFile []FileMap
		GroupDescription string `toml:",omitempty"`
		Members []GroupMember
		Tutor string `toml:",omitempty"`
		Link string `toml:",omitempty"`
	}

	FileMap struct {
		From string
		To string `toml:",omitempty"`
	}

	GroupMember struct {
		First string
		Last string
		ID string
	}
)

func (course Course) ListNames(sep string) (names string) {
	length := len(course.Members)
	for i, member := range course.Members {
		names += member.First + " " + member.Last
		if i < length - 1 {
			names += sep
		}
	}
	return
}

func (course Course) ListLastnames(sep string) (names string) {
	length := len(course.Members)
	for i, member := range course.Members {
		names += member.Last
		if i < length - 1 {
			names += sep
		}
	}
	return
}
