package misc

import (
	"wadverifier/wad"
	"wadverifier/wad/games"
	"wadverifier/wad/status"
)

func BuildREKKRInfo() []wad.Entry {
	return []wad.Entry{
		// REKKR
		{
			MD5Hash:      "d666daa88cca9ff59816ab2d32aeb4c6",
			Name:         "REKKR (PWAD Version)",
			Version:      "1.16",
			Patchinfo:    games.REKKR,
			Status:       status.FINAL,
			PWADRequires: "DOOM.WAD or Freedoom - Phase 1",
			Additional:   "May require the DEHacked file if playing on a source port based on Chocolate Doom.",
		},
		{
			MD5Hash:    "b6f4bb3a80f096b6045cfaeb57d4cf29",
			Name:       "REKKR (Standalone version)",
			Version:    "1.16",
			Patchinfo:  games.REKKR,
			Status:     status.FINAL,
			Additional: "Requires the .deh file to properly play it. However, the batch included provides everything to play it immediately.",
		},
	}
}
