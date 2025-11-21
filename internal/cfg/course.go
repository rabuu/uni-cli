package cfg

type (
	Course struct {
		Name string
		Prefix string `toml:",omitempty"`
		RetrieveFile []FileMap
		RetrieveZip []FileMap
		ExportFile []FileMap
		GroupDescription string `toml:",omitempty"`
		Team []TeamMember
		Tutor string `toml:",omitempty"`
		Link string `toml:",omitempty"`
		Web map[string]string
	}

	FileMap struct {
		From string
		To string `toml:",omitempty"`
		Move bool `toml:",omitempty"`
	}

	TeamMember struct {
		First string
		Last string
		ID string `toml:",omitempty"`
		Email string `toml:",omitempty"`
	}
)

func (course Course) ListTeamNames(sep string) (names string) {
	length := len(course.Team)
	for i, member := range course.Team {
		names += member.First + " " + member.Last
		if i < length - 1 {
			names += sep
		}
	}
	return
}

func (course Course) ListTeamLastnames(sep string) (names string) {
	length := len(course.Team)
	for i, member := range course.Team {
		names += member.Last
		if i < length - 1 {
			names += sep
		}
	}
	return
}
