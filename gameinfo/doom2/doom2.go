package doom2

import (
	"wadverifier/wad"
	"wadverifier/wad/flags"
	"wadverifier/wad/games"
	"wadverifier/wad/status"
)

const (
	PatchInfo_CanBeDowngraded = `If you want to use it on multiplayer source ports, you'll need to patch it using Peter Vaskovics's tool, available below:
	• Windows binaries: http://downloads.zdaemon.org/iwadpatcher-1.2-bin.zip
	• Source code: https://github.com/petervas/iwadpatcher`

	AddendumDoomMP_Rerelease = `This WAD is incompatible with sourceports due to major differences with its original files.
	• You need to use the original WAD instead, found in the following directory:
		- "<yoursteamfolder>\steamapps\common\DOOM 2\base\DOOM2.WAD" for the Steam release,
		- "<bethesdafolder>\games\Doom 2\base\DOOM2.WAD" for the Bethesda Launcher release.`

	AddendumDoomIIKexDoom = `This WAD is incompatible with sourceports due to major differences with its original files.
	• You need to use the original WAD instead, found in the following directory:
		- "<yoursteamfolder>\steamapps\common\Ultimate DOOM\base\doom2\DOOM2.WAD",
		- "<installfolder>\base\doom2\DOOM2.WAD" for the GOG release.`
)

func BuildReleaseInfo() []wad.Entry {
	return []wad.Entry{
		{
			MD5Hash:   "30e3c2d0350b67bfbf47271970b74b2f",
			Name:      "DOOM II",
			Version:   "1.666",
			Patchinfo: games.IWAD,
			Status:    status.NOTFINAL,
		},
		{
			MD5Hash:   "ea74a47a791fdef2e9f2ea8b8a9da13b",
			Name:      "DOOM II",
			Version:   "1.7",
			Patchinfo: games.IWAD,
			Status:    status.NOTFINAL,
		},
		{
			MD5Hash:   "d7a07e5d3f4625074312bc299d7ed33f",
			Name:      "DOOM II",
			Version:   "1.7a",
			Patchinfo: games.IWAD,
			Status:    status.NOTFINAL,
		},
		{
			MD5Hash:   "c236745bb01d89bbb866c8fed81b6f8c",
			Name:      "DOOM II",
			Version:   "1.8",
			Patchinfo: games.IWAD,
			Status:    status.NOTFINAL,
		},
		{
			MD5Hash:   "25e1459ca71d321525f84628f45ca8cd",
			Name:      "DOOM II",
			Version:   "1.9",
			Patchinfo: games.IWAD,
			Status:    status.FINAL,
		},
	}
}

func BuildLocalizedInfo() []wad.Entry {
	return []wad.Entry{
		{
			MD5Hash:    "d9153ced9fd5b898b36cc5844e35b520",
			Name:       "DOOM II (German Release)",
			Version:    "1.666",
			Patchinfo:  games.NONE,
			Status:     status.FINAL,
			Flags:      flags.HIDDEN,
			Additional: "This is the latest release officially released in Germany, but it is not the latest version of Doom II. It is strongly recommended to update it to v1.9 to restore censored contents and play online.",
		},
		{
			MD5Hash:    "3cb02349b3df649c86290907eed64e7b",
			Name:       "DOOM II (French Release)",
			Version:    "1.8",
			Patchinfo:  games.NONE,
			Status:     status.FINAL,
			Flags:      flags.HIDDEN,
			Additional: "This is the latest release officially translated in french, but it is not the latest release of Doom II. It is strongly recommended to update it to v1.9 in order to play online.",
		},
	}
}

func BuildConsolePortInfo() []wad.Entry {

	return []wad.Entry{
		{
			MD5Hash:   "b96683d113c4f4e9a916e1c7d1d71ffd",
			Name:      "DOOM II (PC-98 Version)",
			Patchinfo: games.NONE,
			Status:    status.FINAL,
			Flags:     flags.HIDDEN,
		},
		{
			MD5Hash:   "9640fc4b2c8447bbd28f2080725d5c51",
			Name:      "DOOM II (Tapwave Zodiac Version)",
			Patchinfo: games.NONE,
			Status:    status.FINAL,
			Flags:     flags.HIDDEN,
		},
		{
			MD5Hash:    "a793ebcdd790afad4a1f39cc39a893bd",
			Name:       "DOOM II (XBox version)",
			Patchinfo:  games.NONE,
			Status:     status.FINAL,
			Flags:      flags.RERELEASE | flags.HIDDEN,
			Additional: AddendumDoomMP_Rerelease,
		},
		{
			MD5Hash:    "43c2df32dc6c740cb11f34dc5ab693fa",
			Name:       "Doom II (XBox Live Arcade version)",
			Version:    "1.9",
			Patchinfo:  games.NONE,
			Status:     status.FINAL,
			Flags:      flags.RERELEASE | flags.HIDDEN,
			Additional: AddendumDoomMP_Rerelease,
		},
		{
			MD5Hash:    "c3bea40570c23e511a7ed3ebcd9865f7",
			Name:       "DOOM II (Doom 3 BFG Edition)",
			Patchinfo:  games.NONE,
			Status:     status.FINAL,
			Flags:      flags.RERELEASE | flags.HIDDEN,
			Additional: AddendumDoomMP_Rerelease,
		},
		{
			MD5Hash:   "4c3db5f23b145fccd24c9b84aba3b7dd",
			Name:      "DOOM II (Playstation - American build)",
			Patchinfo: games.NONE,
			Status:    status.FINAL,
			Flags:     flags.RERELEASE | flags.HIDDEN,
		},
		{
			MD5Hash:    "a793ebcdd790afad4a1f39cc39a893bd",
			Name:       "DOOM II (Playstation - European build / Doom Eternal) ",
			Patchinfo:  games.NONE,
			Status:     status.FINAL,
			Flags:      flags.RERELEASE | flags.HIDDEN,
			Additional: "AddendumDoomMP_Rerelease",
		},
		{
			MD5Hash:    "43c2df32dc6c740cb11f34dc5ab693fa",
			Name:       "Doom II (XBox Live Arcade version)",
			Version:    "1.9",
			Patchinfo:  games.NONE,
			Status:     status.FINAL,
			Flags:      flags.RERELEASE | flags.HIDDEN,
			Additional: AddendumDoomMP_Rerelease,
		},
		{
			MD5Hash:    "8ab6d0527a29efdc1ef200e5687b5cae",
			Name:       "DOOM II (Bethesda Version / DOOM Unity)",
			Version:    "2020_08_21 Build #13736 doom2",
			Patchinfo:  games.DOOM_UNITY,
			Status:     status.FINAL,
			Flags:      flags.RERELEASE | flags.HIDDEN,
			Additional: AddendumDoomMP_Rerelease,
		},
		{
			MD5Hash:    "64a4c88a871da67492aaa2020a068cd8",
			Name:       "DOOM II (Doom + Doom II)",
			Version:    "Update 1",
			Patchinfo:  games.KEX_DOOM2024,
			Flags:      flags.RERELEASE,
			Additional: AddendumDoomIIKexDoom,
		},
	}
}

func BuildMiscFiles() []wad.Entry {
	return []wad.Entry{
		{
			MD5Hash:    "1abb3f69cd793b77a9b623d15470d48f",
			Name:       "DOOM II (Non-Interactive Demo)",
			Version:    "1.666",
			Patchinfo:  games.NONE,
			Status:     status.FINAL,
			Flags:      flags.HIDDEN,
			Additional: "Did you know you can replicate this build by removing MAP01's lump in a WAD editor?",
		},
		{
			MD5Hash:    "3cb02349b3df649c86290907eed64e7b",
			Name:       "DOOM II (French Release)",
			Version:    "1.8",
			Patchinfo:  games.NONE,
			Status:     status.FINAL,
			Flags:      flags.HIDDEN,
			Additional: "This is the latest release officially translated in french, but it is not the latest release of Doom II. It is strongly recommended to update it to v1.9 in order to play online.",
		},
	}
}

func Populate() []wad.Entry {
	list := BuildReleaseInfo()
	list = append(list, BuildLocalizedInfo()...)
	list = append(list, BuildConsolePortInfo()...)
	return list
}
