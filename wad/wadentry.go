package wad

import (
	"wadverifier/wad/flags"
	"wadverifier/wad/games"
	"wadverifier/wad/status"
)

// WadInfo : all WAD data returned from the program
type Entry struct {
	MD5Hash      string          `json:"md5"`
	Name         string          `json:"name"`
	Version      string          `json:"version"`
	Status       status.Status   `json:"status"` // Is it the final release of this?
	Flags        flags.Flags     `json:"flags"`  // Special Flags
	Patchinfo    games.PatchType `json:"patchinfo"`
	PWADRequires string          `json:"requires"`   // If the official PWAD requires an IWAD to run
	Additional   string          `json:"additional"` // If I need to display an additionnal message for this IWAD.
}

/*func LoadDefinitions() [][]Info {


	return [][]
}*/

func (s *Entry) IsFinal() bool {
	return s.Status == status.FINAL
}
