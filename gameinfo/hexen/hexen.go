package hexen

import (
	"wadverifier/wad"
	"wadverifier/wad/flags"
	"wadverifier/wad/games"
	"wadverifier/wad/status"
)

func BuildReleaseInfo() []wad.Entry {
	return []wad.Entry{
		{
			MD5Hash:   "abb033caf81e26f12a2103e1fa25453f",
			Name:      "Hexen",
			Version:   "1.1",
			Patchinfo: games.IWAD,
			Status:    status.FINAL,
		},
		{
			MD5Hash:   "b2543a03521365261d0a0f74d5dd90f0",
			Name:      "Hexen",
			Version:   "1.0",
			Patchinfo: games.IWAD,
			Status:    status.NOTFINAL,
		},

		// MAC versions
		{
			MD5Hash:    "b68140a796f6fd7f3a5d3226a32b93be",
			Name:       "Hexen (Macintosh version)",
			Version:    "1.1",
			Patchinfo:  games.IWAD,
			Status:     status.FINAL,
			Additional: "Despite being the Macintosh release, it contains differences that makes it incompatible with the Windows release of the game!",
		},

		// hexdd.wad
		// Steam only distributes version 1.0, btw!
		{
			MD5Hash:    "1077432e2690d390c256ac908b5f4efa",
			Name:       "Hexen: Deathkings of the Dark Citadel",
			Version:    "1.0",
			Patchinfo:  games.IWAD,
			Status:     status.NOTFINAL,
			Additional: "Requires Hexen 1.1 and needs to be run as a PWAD. Steam's release uses this version!",
		},
		{
			MD5Hash:    "78d5898e99e220e4de64edaa0e479593",
			Name:       "Hexen: Deathkings of the Dark Citadel",
			Version:    "1.1",
			Patchinfo:  games.IWAD,
			Status:     status.FINAL,
			Additional: "Requires Hexen 1.1 and needs to be run as a PWAD.",
		},
	}
}

func BuildShareWareInfo() []wad.Entry {
	return []wad.Entry{
		{
			MD5Hash:    "b68140a796f6fd7f3a5d3226a32b93be",
			Name:       "Hexen Demo (Macintosh version)",
			Patchinfo:  games.NONE,
			Status:     status.FINAL,
			Additional: "Despite being the Macintosh release, it contains differences that makes it incompatible with the Windows release of the game!",
		},
		// OTHER IRREVELENT THINGS
		{
			MD5Hash:   "876a5a44c7b68f04b3bb9bc7a5bd69d6",
			Name:      "Hexen Demo",
			Version:   "1.0",
			Patchinfo: games.NONE,
			Status:    status.FINAL,
		},
	}
}

func BuildPrototypeInfo() []wad.Entry {
	return []wad.Entry{
		{
			MD5Hash:   "9178a32a496ff5befebfe6c47dac106c",
			Name:      "Hexen Demo Beta",
			Patchinfo: games.NONE,
			Flags:     flags.PRERELEASE,
		},
		{
			MD5Hash:   "c88a2bb3d783e2ad7b599a8e301e099e",
			Version:   "Hexen Beta",
			Patchinfo: games.NONE,
			Flags:     flags.PRERELEASE,
		},
	}
}

func BuildConsolePortInfo() []wad.Entry {
	return []wad.Entry{
		{
			MD5Hash:   "a66ae0448436a990b3aecd018bc2708a",
			Name:      "Hexen (Heretic + Hexen)",
			Version:   "1.0.4575 - 46e45cbd (Jul 17, 2025)",
			Patchinfo: games.KEX_HERETIC_HEXEN2025,
			Flags:     flags.RERELEASE,
			//Additional: AddendumKexHeretic,
		},
		{
			MD5Hash:   "c078b329f53378044b8fc28d60db8e51",
			Name:      "Hexen: Deathkings of the Dark Citadel (Heretic + Hexen)",
			Version:   "Update 1 (Sep 26, 2025)",
			Patchinfo: games.KEX_HERETIC_HEXEN2025,
			Status:    status.FINAL,
			Flags:     flags.RERELEASE,
			//Additional: AddendumKexHeretic,
		},
		{
			MD5Hash:   "2bb0f56ab1f98000990524c1b67e8759",
			Name:      "Hexen: Deathkings of the Dark Citadel (Heretic + Hexen)",
			Version:   "1.0.4575 - 46e45cbd (Jul 17, 2025)",
			Patchinfo: games.KEX_HERETIC_HEXEN2025,
			Status:    status.NOTFINAL,
			Flags:     flags.RERELEASE,
			//Additional: AddendumKexHeretic,
		},
		{
			MD5Hash:   "aee983213be00b2e5e02c09675dc608f",
			Name:      "Hexen: Vestiges of Grandeur (Heretic + Hexen)",
			Version:   "1.0.4575 - 46e45cbd (Jul 17, 2025)",
			Patchinfo: games.KEX_HERETIC_HEXEN2025,
			Flags:     flags.RERELEASE,
			//Additional: AddendumKexHeretic,
		},
	}
}

func BuildExtraWADInfo() []wad.Entry {
	return []wad.Entry{
		{
			MD5Hash:   "aee441fadfaa8bf9c06bd39d6abd6775",
			Name:      "Hexen: Test WAD (Heretic + Hexen)",
			Version:   "1.0.4575 - 46e45cbd (Jul 17, 2025)",
			Patchinfo: games.KEX_HERETIC_HEXEN2025,
			Flags:     flags.EXTRADATA | flags.HIDDEN,
			//Additional: AddendumKexHeretic,
		},
		{
			MD5Hash:   "2fca43a18a535511efa8c7c27e202c12",
			Name:      "Hexen Original Soundtrack (Heretic + Hexen)",
			Version:   "1.0.4575 - 46e45cbd (Jul 17, 2025)",
			Patchinfo: games.KEX_HERETIC_HEXEN2025,
			Flags:     flags.RERELEASE | flags.HIDDEN,
		},
		{
			MD5Hash:    "4a4cbb09f2ecb1b6f7bf3477ca54c18c",
			Name:       "Hexen Remade Soundtrack (Heretic + Hexen)",
			Version:    "1.0.4575 - 46e45cbd (Jul 17, 2025)",
			Patchinfo:  games.KEX_HERETIC_HEXEN2025,
			Flags:      flags.RERELEASE | flags.HIDDEN,
			Additional: "OST remade by Andrew Hulshault.",
		},
		{
			MD5Hash:   "e0ba5039f3baf750714d845c992a5331",
			Name:      "Hexen Extras (Heretic + Hexen)",
			Version:   "1.0.4575 - 46e45cbd (Jul 17, 2025)",
			Patchinfo: games.KEX_HERETIC_HEXEN2025,
			Flags:     flags.RERELEASE,
		},
		{
			MD5Hash:   "b7d98538d46dd96e36ee60aa41e8fb4b",
			Name:      "Hexen: Deathkings of the Dark Citadel Extras (Heretic + Hexen)",
			Version:   "Update 1 (Sep 26, 2025)",
			Patchinfo: games.KEX_HERETIC_HEXEN2025,
			Status:    status.FINAL,
			Flags:     flags.RERELEASE,
		},
		{
			MD5Hash:   "335336b45cb3e99f6f35b46418d8e31f",
			Name:      "Hexen: Deathkings of the Dark Citadel Extras (Heretic + Hexen)",
			Version:   "1.0.4575 - 46e45cbd (Jul 17, 2025)",
			Patchinfo: games.KEX_HERETIC_HEXEN2025,
			Status:    status.NOTFINAL,
			Flags:     flags.RERELEASE,
		},
	}
}

func Populate() []wad.Entry {
	list := BuildReleaseInfo()
	list = append(list, BuildPrototypeInfo()...)
	list = append(list, BuildShareWareInfo()...)
	list = append(list, BuildConsolePortInfo()...)
	list = append(list, BuildExtraWADInfo()...)
	return list
}
