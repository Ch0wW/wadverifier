package doom

import (
	"wadverifier/wad"
	"wadverifier/wad/flags"
	"wadverifier/wad/games"
	"wadverifier/wad/status"
)

const (
	PatchInfo_Shareware = "You can download the most updated shareware release of DooM at https://www.doomworld.com/idgames/idstuff/doom/doom19s"

	PatchInfo_CanBeDowngraded = `If you want to use it on multiplayer source ports, you'll need to patch it using Peter Vaskovics's tool, available below:
	• Windows binaries: http://downloads.zdaemon.org/iwadpatcher-1.2-bin.zip
	• Source code: https://github.com/petervas/iwadpatcher`

	AddendumDoomMP_Rerelease = `This WAD is incompatible with sourceports due to major differences with its original files.
	• You will have to use the original WAD instead, found in the following directory:
		- "<yoursteamfolder>\steamapps\common\Ultimate DOOM\base\DOOM.WAD" for the Steam release,
		- "<bethesdafolder>\games\Ultimate Doom\base\DOOM.WAD" for the Bethesda Launcher release.`

	AddendumDoomMP_KexDoom = `This WAD is incompatible with sourceports due to major differences with its original files.
	• You will have to use the original WAD instead, found in the following directory:
		- "<yoursteamfolder>\steamapps\common\Ultimate DOOM\base\DOOM.WAD" for the Steam release,
		- "<installfolder>\base\DOOM.WAD" for the GOG release.`
)

func BuildReleaseInfo() []wad.Entry {

	// DOOM/UDOOM population
	list := []wad.Entry{

		// ToDo: MISSING DOOM v1.0 !!!

		{
			MD5Hash:   "981b03e6d1dc033301aa3095acc437ce",
			Name:      "DOOM (Registered)",
			Version:   "1.1",
			Patchinfo: games.IWAD,
			Status:    status.NOTFINAL,
		},
		{
			MD5Hash:   "792fd1fea023d61210857089a7c1e351",
			Name:      "DOOM (Registered)",
			Version:   "1.2",
			Patchinfo: games.IWAD,
			Status:    status.NOTFINAL,
		},
		{
			MD5Hash:   "54978d12de87f162b9bcc011676cb3c0",
			Name:      "DOOM (Registered)",
			Version:   "1.666",
			Patchinfo: games.IWAD,
			Status:    status.NOTFINAL,
		},
		{
			MD5Hash:   "11e1cd216801ea2657723abc86ecb01f",
			Name:      "DOOM (Registered)",
			Version:   "1.8",
			Patchinfo: games.IWAD,
			Status:    status.NOTFINAL,
		},
		{
			MD5Hash:    "1cd63c5ddff1bf8ce844237f580e9cf3",
			Name:       "DOOM (Registered)",
			Version:    "1.9",
			Patchinfo:  games.IWAD,
			Status:     status.NOTFINAL,
			Additional: "This IWAD is the latest version of non-Ultimate DOOM, but it is highly recommended to upgrade to The Ultimate Doom version.",
		},
		{
			MD5Hash:   "c4fe9fd920207691a9f493668e0a2083",
			Name:      "The Ultimate DOOM",
			Version:   "1.9",
			Patchinfo: games.IWAD,
			Status:    status.FINAL,
		},
	}

	return list
}

func BuildPrototypeInfo() []wad.Entry {

	prototype := []wad.Entry{
		{
			MD5Hash:   "740901119ba2953e3c7f3764eca6e128",
			Name:      "DOOM (Press Release)",
			Version:   "0.2",
			Patchinfo: games.NONE,
			Status:    status.UNKNOWN,
			Flags:     flags.PRERELEASE,
		},
		{
			MD5Hash:   "dae9b1eea1a8e090fdfa5707187f4a43",
			Name:      "DOOM (Press Release)",
			Version:   "0.3",
			Patchinfo: games.NONE,
			Status:    status.UNKNOWN,
			Flags:     flags.PRERELEASE,
		},
		{
			MD5Hash:   "b6afa12a8b22e2726a8ff5bd249223de",
			Name:      "DOOM (Press Release)",
			Version:   "0.4",
			Patchinfo: games.NONE,
			Status:    status.UNKNOWN,
			Flags:     flags.PRERELEASE,
		},
		{
			MD5Hash:   "9c877480b8ef33b7074f1f0c07ed6487",
			Name:      "DOOM (Press Release)",
			Version:   "0.5",
			Patchinfo: games.NONE,
			Status:    status.UNKNOWN,
			Flags:     flags.PRERELEASE,
		},
		{
			MD5Hash:   "049e32f18d9c9529630366cfc72726ea",
			Name:      "DOOM (Press Release - DOOMPRES.WAD)",
			Version:   "October 4th, 1993",
			Patchinfo: games.NONE,
			Status:    status.UNKNOWN,
			Flags:     flags.PRERELEASE,
		},
	}

	return prototype
}

func BuildShareWareInfo() []wad.Entry {

	shareware := []wad.Entry{
		{
			MD5Hash:   "90facab21eede7981be10790e3f82da2",
			Name:      "DOOM (Shareware)",
			Version:   "1.0",
			Patchinfo: games.DOOM_SHAREWARE,
			Status:    status.NOTFINAL,
		},
		{
			MD5Hash:    "cea4989df52b65f4d481b706234a3dca",
			Name:       "DOOM (Shareware)",
			Version:    "1.1 (December 15th, 1993)",
			Patchinfo:  games.DOOM_SHAREWARE,
			Status:     status.NOTFINAL,
			Additional: "This version was superceded by an update released one day later.",
		},
		{
			MD5Hash:   "52cbc8882f445573ce421fa5453513c1",
			Name:      "DOOM (Shareware)",
			Version:   "1.1",
			Patchinfo: games.DOOM_SHAREWARE,
			Status:    status.NOTFINAL,
		},
		{
			MD5Hash:    "2a380f28e813fb0989cae5e4762ebb4c",
			Name:       "DOOM (Shareware)",
			Version:    "1.2 (February 4th, 1994)",
			Patchinfo:  games.DOOM_SHAREWARE,
			Status:     status.NOTFINAL,
			Additional: "This version was originally released for NeXTSTEP and potentially MacOS.",
		},
		{
			MD5Hash:    "30aa5beb9e5ebfbbe1e1765561c08f38",
			Name:       "DOOM (Shareware)",
			Version:    "1.2",
			Patchinfo:  games.DOOM_SHAREWARE,
			Status:     status.NOTFINAL,
			Additional: PatchInfo_Shareware,
		},
		{
			MD5Hash:    "17aebd6b5f2ed8ce07aa526a32af8d99",
			Name:       "DOOM (Shareware)",
			Version:    "1.25",
			Patchinfo:  games.DOOM_SHAREWARE,
			Status:     status.NOTFINAL,
			Additional: "Also known as 'Sybex special edition'",
		},
		{
			MD5Hash:    "a21ae40c388cb6f2c3cc1b95589ee693",
			Name:       "DOOM (Shareware Beta)",
			Version:    "1.4",
			Patchinfo:  games.DOOM_SHAREWARE,
			Status:     status.NOTFINAL,
			Flags:      flags.PRERELEASE,
			Additional: PatchInfo_Shareware,
		},
		{
			MD5Hash:    "e280233d533dcc28c1acd6ccdc7742d4",
			Name:       "DOOM (Shareware Beta)",
			Version:    "1.5",
			Patchinfo:  games.DOOM_SHAREWARE,
			Status:     status.NOTFINAL,
			Flags:      flags.PRERELEASE,
			Additional: PatchInfo_Shareware,
		},
		{
			MD5Hash:    "762fd6d4b960d4b759730f01387a50a1",
			Name:       "DOOM (Shareware Beta)",
			Version:    "1.6",
			Patchinfo:  games.DOOM_SHAREWARE,
			Status:     status.NOTFINAL,
			Flags:      flags.PRERELEASE,
			Additional: PatchInfo_Shareware,
		},
		{
			MD5Hash:    "c428ea394dc52835f2580d5bfd50d76f",
			Name:       "DOOM (Shareware)",
			Version:    "1.666",
			Patchinfo:  games.DOOM_SHAREWARE,
			Status:     status.NOTFINAL,
			Additional: PatchInfo_Shareware,
		},
		{
			MD5Hash:    "5f4eb849b1af12887dec04a2a12e5e62",
			Name:       "DOOM (Shareware)",
			Version:    "1.8",
			Patchinfo:  games.DOOM_SHAREWARE,
			Status:     status.NOTFINAL,
			Additional: PatchInfo_Shareware,
		},
		{
			MD5Hash:   "f0cefca49926d00903cf57551d901abe",
			Name:      "DOOM (Shareware)",
			Version:   "1.9",
			Patchinfo: games.DOOM_SHAREWARE,
			Status:    status.FINAL,
		},
	}

	return shareware
}

func BuildConsolePortInfo() []wad.Entry {

	list := []wad.Entry{
		{
			MD5Hash:   "dae77aff77a0491e3b7254c9c8401aa8",
			Name:      "DOOM for Pocket PC",
			Patchinfo: games.NONE,
			Status:    status.FINAL,
		},
		{
			MD5Hash:    "0c8758f102ccafe26a3040bee8ba5021",
			Name:       "The Ultimate DOOM (XBox Version)",
			Patchinfo:  games.NONE,
			Flags:      flags.RERELEASE | flags.HIDDEN,
			Additional: PatchInfo_CanBeDowngraded,
		},
		{
			MD5Hash:    "72286ddc680d47b9138053dd944b2a3d",
			Version:    "The Ultimate DOOM (XBox Live Arcade version)",
			Patchinfo:  games.NONE,
			Flags:      flags.RERELEASE | flags.HIDDEN,
			Additional: PatchInfo_CanBeDowngraded,
		},
		{
			MD5Hash:    "fb35c4a5a9fd49ec29ab6e900572c524",
			Version:    "The Ultimate DOOM (Doom 3 - BFG Edition)",
			Patchinfo:  games.NONE,
			Flags:      flags.RERELEASE | flags.HIDDEN,
			Additional: PatchInfo_CanBeDowngraded,
		},
		{
			MD5Hash:    "232a79f7121b22d7401905ee0ee1e487",
			Name:       "DOOM (Unity port)",
			Version:    "<!! Missing DATE Entry !!> - Version 1.0",
			Patchinfo:  games.DOOM_UNITY,
			Status:     status.NOTFINAL,
			Flags:      flags.RERELEASE,
			Additional: AddendumDoomMP_Rerelease,
		},
		{
			MD5Hash:    "21b200688d0fa7c1b6f63703d2bdd455",
			Name:       "DOOM (Unity port)",
			Version:    "<!! Missing DATE Entry !!> - Version 1.1",
			Patchinfo:  games.DOOM_UNITY,
			Status:     status.NOTFINAL,
			Flags:      flags.RERELEASE,
			Additional: AddendumDoomMP_Rerelease,
		},
		{
			MD5Hash:    "8517c4e8f0eef90b82852667d345eb86",
			Name:       "DOOM (Unity port)",
			Version:    "2020_08_21 Build #13735 doom",
			Patchinfo:  games.DOOM_UNITY,
			Status:     status.FINAL,
			Flags:      flags.RERELEASE,
			Additional: AddendumDoomMP_Rerelease,
		},
		{
			MD5Hash:    "4461d4511386518e784c647e3128e7bc",
			Name:       "DOOM (Doom + Doom II)",
			Version:    "Original Release",
			Patchinfo:  games.KEX_DOOM2024,
			Status:     status.NOTFINAL,
			Flags:      flags.RERELEASE,
			Additional: AddendumDoomMP_KexDoom,
		},
		{
			MD5Hash:    "3b37188f6337f15718b617c16e6e7a9c",
			Name:       "DOOM (Doom + Doom II)",
			Version:    "Update 1",
			Patchinfo:  games.KEX_DOOM2024,
			Status:     status.FINAL,
			Flags:      flags.RERELEASE,
			Additional: AddendumDoomMP_KexDoom,
		},
	}

	return list
}

func Populate() []wad.Entry {
	list := BuildReleaseInfo()
	list = append(list, BuildShareWareInfo()...)
	list = append(list, BuildConsolePortInfo()...)
	list = append(list, BuildPrototypeInfo()...)
	return list
}
