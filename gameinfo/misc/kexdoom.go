package misc

import (
	"wadverifier/wad"
	"wadverifier/wad/flags"
	"wadverifier/wad/games"
	"wadverifier/wad/status"
)

func BuildKexDoomExtraInfo() []wad.Entry {

	// DOOM/UDOOM population
	return []wad.Entry{
		// DOOM 64

		// Additionnal Doom + Doom II wads
		{
			MD5Hash:    "2e76d93d52ef64fb9db3cee2437c686b",
			Name:       "extras.wad (Doom + Doom II release)",
			Version:    "Update 1",
			Patchinfo:  games.KEX_DOOM2024,
			Status:     status.FINAL,
			Flags:      flags.RERELEASE | flags.HIDDEN,
			Additional: "Used for shared graphics and contains Andrew Hulshult's IDKFA soundtrack in OGG format.",
		},
		{
			MD5Hash:    "95f21547be5e0bff38d412017440f656",
			Name:       "id1.wad (Doom + Doom II release)",
			Version:    "Update 1",
			Patchinfo:  games.KEX_DOOM2024,
			Status:     status.NOTFINAL,
			Flags:      flags.RERELEASE,
			Additional: "Used for 'Legacy of Rust' addon.",
		},
		{
			MD5Hash:    "713c5a3c1734b1d55b2813a3dd0136d9",
			Name:       "id1.wad (Doom + Doom II release)",
			Version:    "Update 2+",
			Patchinfo:  games.KEX_DOOM2024,
			Status:     status.FINAL,
			Flags:      flags.RERELEASE,
			Additional: "Used for 'Legacy of Rust' addon.",
		},
		{
			MD5Hash:    "187bfe543f8328b379e46957976e800d",
			Name:       "id1-tex.wad (Doom + Doom II)",
			Patchinfo:  games.NONE,
			Flags:      flags.RERELEASE | flags.HIDDEN,
			Additional: "Used for the 'Legacy of Rust' addon.",
		},
		{
			MD5Hash:    "f8fbab472230bfa090d6a9234d65fae6",
			Name:       "id1-res.wad (Doom + Doom II)",
			Patchinfo:  games.NONE,
			Flags:      flags.RERELEASE | flags.HIDDEN,
			Additional: "Used for the 'Legacy of Rust' addon.",
		},
		{
			MD5Hash:    "85d25c8c3d06a05a1283ae4afe749c9f",
			Name:       "id1-weap.wad (Doom + Doom II)",
			Patchinfo:  games.NONE,
			Flags:      flags.RERELEASE | flags.HIDDEN,
			Additional: "Used for the 'Legacy of Rust' addon.",
		},
		{
			MD5Hash:    "436c83dd83a47f8dd251ba15108e9459",
			Name:       "id1-mus.wad (Doom + Doom II)",
			Patchinfo:  games.NONE,
			Flags:      flags.RERELEASE | flags.HIDDEN,
			Additional: "Used for the 'Legacy of Rust' addon.",
		},
		{
			MD5Hash:    "4f0651accebc007b853943ac12aa95b8",
			Name:       "id24res.wad (Doom + Doom II)",
			Patchinfo:  games.NONE,
			Flags:      flags.RERELEASE | flags.HIDDEN,
			Additional: "Resources from *that* proprietary ID24 standard.",
		},
		{
			MD5Hash:    "5670fd8fe8eb6910ec28f9e27969d84f",
			Name:       "iddm1.wad (Doom + Doom II)",
			Version:    "Update 1",
			Patchinfo:  games.KEX_DOOM2024,
			Status:     status.NOTFINAL,
			Flags:      flags.RERELEASE,
			Additional: "Used as the exclusive Deathmatch WAD for Doom + Doom II.",
		},
		{
			MD5Hash:    "cb92010b8ec05f8924ac966a8ed95b74",
			Name:       "iddm1.wad (Doom + Doom II)",
			Version:    "Update 2+",
			Patchinfo:  games.KEX_DOOM2024,
			Status:     status.FINAL,
			Flags:      flags.RERELEASE,
			Additional: "Used as the exclusive Deathmatch WAD for Doom + Doom II.",
		},
	}
}
