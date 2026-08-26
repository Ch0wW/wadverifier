package misc

import (
	"wadverifier/wad"
	"wadverifier/wad/flags"
	"wadverifier/wad/games"
	"wadverifier/wad/status"
)

func BuildSIGILInfo() []wad.Entry {
	return []wad.Entry{
		// SIGIL 1 & 2 by John Romero
		{
			MD5Hash:      "f53ffc4fb89e966839bb8d20c632819a",
			Name:         "SIGIL",
			Version:      "1.0",
			Patchinfo:    games.SIGIL,
			Status:       status.NOTFINAL,
			PWADRequires: "The Ultimate Doom v1.9",
		},
		{
			MD5Hash:      "a775262ca0e423468196803b71a57a43",
			Name:         "SIGIL (Compatibility WAD)",
			Version:      "1.0",
			Patchinfo:    games.SIGIL,
			Status:       status.NOTFINAL,
			PWADRequires: "The Ultimate Doom v1.9",
		},

		{
			MD5Hash:      "1fe9daa0e837c7452eb2f91aac2cc983",
			Name:         "SIGIL",
			Version:      "1.1",
			Patchinfo:    games.SIGIL,
			Status:       status.NOTFINAL,
			PWADRequires: "The Ultimate Doom v1.9",
		},
		{
			MD5Hash:      "c04912beab6aa82c114a19c976ec8c0d",
			Name:         "SIGIL (Compatibility WAD)",
			Version:      "1.1",
			Patchinfo:    games.SIGIL,
			Status:       status.NOTFINAL,
			PWADRequires: "The Ultimate Doom v1.9",
		},

		{
			MD5Hash:      "427ca995600970abcd2efcc684a64c88",
			Name:         "SIGIL",
			Version:      "1.2",
			Patchinfo:    games.SIGIL,
			Status:       status.NOTFINAL,
			PWADRequires: "The Ultimate Doom v1.9",
		},
		{
			MD5Hash:      "9285e9cc2dbd87d238baab37d700c644",
			Name:         "SIGIL (Compatibility WAD)",
			Version:      "1.2",
			Patchinfo:    games.SIGIL,
			Status:       status.NOTFINAL,
			PWADRequires: "The Ultimate Doom v1.9",
		},

		{
			MD5Hash:      "743d6323cb2b9be24c258ff0fc350883",
			Name:         "SIGIL",
			Version:      "1.21",
			Patchinfo:    games.SIGIL,
			Status:       status.NOTFINAL,
			PWADRequires: "The Ultimate Doom v1.9",
		},
		{
			MD5Hash:      "573f3f178c76709f512089ed15484391",
			Name:         "SIGIL (Compatibility WAD)",
			Version:      "1.21",
			Patchinfo:    games.SIGIL,
			Status:       status.NOTFINAL,
			PWADRequires: "The Ultimate Doom v1.9",
		},

		{
			MD5Hash:      "edd5c3dfd3fb1c981cf7390c5c14454e",
			Name:         "SIGIL",
			Version:      "1.23",
			Patchinfo:    games.SIGIL,
			Status:       status.FINAL,
			PWADRequires: "The Ultimate Doom v1.9",
		},
		{
			MD5Hash:      "e4c5ab58e226bfcc8761f35204aeb3fc",
			Name:         "SIGIL (Compatibility WAD)",
			Version:      "1.23",
			Patchinfo:    games.SIGIL,
			Status:       status.FINAL,
			PWADRequires: "The Ultimate Doom v1.9",
		},

		{
			MD5Hash:      "b424dcf46ae55a496c34ac37cce32646",
			Name:         "SIGIL - BucketHead soundtrack",
			Patchinfo:    games.NONE,
			PWADRequires: "SIGIL & The Ultimate DOOM v1.9",
		},
		{
			MD5Hash:      "343faa815928c58faa08939a4502d5d2",
			Name:         "SIGIL - BucketHead soundtrack (Compatibility WAD)",
			Patchinfo:    games.NONE,
			PWADRequires: "SIGIL & The Ultimate DOOM v1.9",
		},
		{
			MD5Hash:    "08ee05388c137db5f5d7996e89425b95",
			Name:       "SIGIL (Doom + Doom II)",
			Version:    "Update 1",
			Patchinfo:  games.KEX_DOOM2024,
			Status:     status.FINAL,
			Flags:      flags.RERELEASE,
			Additional: "File is not identical to the original releases of SIGIL and won't be compatible with multiplayer sourceports.",
		},
	}
}

func BuildSIGIL_2_Info() []wad.Entry {
	return []wad.Entry{
		// SIGIL II
		{
			MD5Hash:      "d0442f5a75f2faef3405c09a0c3acc58",
			Name:         "SIGIL II",
			Version:      "1.0",
			Patchinfo:    games.SIGIL_2,
			PWADRequires: "The Ultimate Doom v1.9",
			Additional:   "You will need a limit-removing source port to be able to run this.",
		},
		{
			MD5Hash:    "953f65cf079d0ba9a25be2c407da7ec1",
			Name:       "SIGIL II (Doom + Doom II)",
			Version:    "Update 3",
			Patchinfo:  games.KEX_DOOM2024,
			Status:     status.FINAL,
			Flags:      flags.RERELEASE,
			Additional: "File is not identical to the original release of SIGIL II and won't be compatible with multiplayer sourceports.",
		},
	}
}
