package finaldoom

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

	AddendumTNTKexDoom = `This WAD is incompatible with sourceports due to major differences with its original files.
	• You need to use the original WAD instead, found in the following directory:
		- "<yoursteamfolder>\steamapps\common\Ultimate DOOM\base\tnt\TNT.WAD for the Steam release,",
		- "<installfolder>\base\tnt\TNT.WAD" for the GOG release.`

	AddendumPLUTONIAKexDoom = `This WAD is incompatible with sourceports due to major differences with its original files.
	• You need to use the original WAD instead, found in the following directory:
		- "<yoursteamfolder>\steamapps\common\Ultimate DOOM\base\plutonia\PLUTONIA.WAD for the Steam release,",
		- "<installfolder>\base\plutonia\PLUTONIA.WAD" for the GOG version for the GOG release.`
)

func BuildReleaseInfo() []wad.Entry {
	return []wad.Entry{
		{
			MD5Hash:   "75c8cf89566741fa9d22447604053bd7",
			Name:      "Final DOOM: The Plutonia Experiment",
			Patchinfo: games.IWAD,
			Status:    status.FINAL,
		},
		{
			MD5Hash:    "3493be7e1e2588bc9c8b31eab2587a04",
			Name:       "Final DOOM: The Plutonia Experiment",
			Version:    "id Anthology release",
			Patchinfo:  games.IWAD,
			Status:     status.UNKNOWN,
			Flags:      flags.RERELEASE,
			Additional: PatchInfo_CanBeDowngraded,
		},
		{
			MD5Hash:   "4e158d9953c79ccf97bd0663244cc6b6",
			Name:      "Final DOOM: TNT: Evilution",
			Patchinfo: games.IWAD,
			Status:    status.FINAL,
		},
		{
			MD5Hash:    "1d39e405bf6ee3df69a8d2646c8d5c49",
			Name:       "Final DOOM: TNT: Evilution",
			Version:    "id Anthology release",
			Patchinfo:  games.IWAD,
			Status:     status.UNKNOWN,
			Flags:      flags.RERELEASE,
			Additional: PatchInfo_CanBeDowngraded,
		},
		{
			MD5Hash:    "7a77ee148fd9ee5bc599356218f6f6b5",
			Name:       "Final DOOM: TNT: Evilution - Fix for MAP31",
			Patchinfo:  games.NONE,
			Flags:      flags.HIDDEN,
			Additional: "This PWAD fixes the yellow keycard not appearing on TNT MAP31.",
		},
	}
}

func BuildConsolePortInfo() []wad.Entry {
	return []wad.Entry{
		{
			MD5Hash:    "24037397056e919961005e08611623f4",
			Name:       "Final DOOM: The Plutonia Experiment (Doom + Doom II)",
			Version:    "Original Release",
			Patchinfo:  games.KEX_DOOM2024,
			Status:     status.NOTFINAL,
			Flags:      flags.RERELEASE,
			Additional: AddendumPLUTONIAKexDoom,
		},
		{
			MD5Hash:    "e47cf6d82a0ccedf8c1c16a284bb5937",
			Name:       "Final DOOM: The Plutonia Experiment (Doom + Doom II)",
			Version:    "Update 1",
			Patchinfo:  games.KEX_DOOM2024,
			Flags:      flags.RERELEASE,
			Additional: AddendumPLUTONIAKexDoom,
		},
		{
			MD5Hash:    "ad7885c17a6b9b79b09d7a7634dd7e2c",
			Name:       "Final DOOM: TNT: Evilution (Doom + Doom II)",
			Version:    "Update 1",
			Patchinfo:  games.KEX_DOOM2024,
			Flags:      flags.RERELEASE,
			Additional: AddendumTNTKexDoom,
		},
	}

}

func Populate() []wad.Entry {
	list := BuildReleaseInfo()
	list = append(list, BuildConsolePortInfo()...)
	return list
}
