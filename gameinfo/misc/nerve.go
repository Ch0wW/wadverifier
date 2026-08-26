package misc

import (
	"wadverifier/wad"
	"wadverifier/wad/flags"
	"wadverifier/wad/games"
	"wadverifier/wad/status"
)

func BuildNerveInfo() []wad.Entry {
	return []wad.Entry{
		{
			MD5Hash:      "967d5ae23daf45196212ae1b605da3b0",
			Version:      "No Rest for the Living",
			Patchinfo:    games.NONE,
			Status:       status.FINAL,
			PWADRequires: "DOOM II v1.9",
			Additional:   "You will need a limit-removing source port to be able to run this.",
		},
		{
			MD5Hash:    "23422eb42833ac7b0dd59c0c7ae18a6f",
			Name:       "No Rest For The Living (Doom + Doom II release)",
			Version:    "Update 1",
			Patchinfo:  games.KEX_DOOM2024,
			Flags:      flags.RERELEASE,
			Additional: "File is not identical to the original release of No Rest for the Living and won't be compatible with multiplayer sourceports.",
		},
	}
}
