package misc

import (
	"wadverifier/wad"
	"wadverifier/wad/games"
	"wadverifier/wad/status"
)

func BuildHacxInfo() []wad.Entry {
	return []wad.Entry{
		// HACX
		{
			MD5Hash:   "65ed74d522bdf6649c2831b13b9e02b4",
			Name:      "HacX",
			Version:   "1.2",
			Patchinfo: games.HACX,
			Status:    status.FINAL,
		},
		{
			MD5Hash:   "b7fd2f43f3382cf012dc6b097a3cb182",
			Name:      "HacX",
			Version:   "1.1",
			Patchinfo: games.HACX,
			Status:    status.NOTFINAL,
		},
		{
			MD5Hash:   "1511a7032ebc834a3884cf390d7f186e",
			Name:      "HacX",
			Version:   "1.0",
			Patchinfo: games.HACX,
			Status:    status.NOTFINAL,
		},
	}
}
