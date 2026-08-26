package strife

import (
	"wadverifier/wad"
	"wadverifier/wad/flags"
	"wadverifier/wad/games"
	"wadverifier/wad/status"
)

func BuildStrifeInfo() []wad.Entry {
	return []wad.Entry{
		{
			MD5Hash:    "2fed2031a5b03892106e0f117f17901f",
			Name:       "Strafe (Registered)",
			Version:    "1.2 - 1.31",
			Patchinfo:  games.IWAD,
			Status:     status.FINAL,
			Additional: "Your IWAD is up-to-date. However, the latest updates of Strife only modify the executable. Please make sure it is also updated to the latest version.",
		},
		{
			MD5Hash:   "8f2d3a6a289f5d2f2f9c1eec02b47299",
			Name:      "Strife (Registered)",
			Version:   "1.0",
			Patchinfo: games.IWAD,
			Status:    status.NOTFINAL,
		},
		// Additionnal IWADs
		{
			MD5Hash:      "082234d6a3f7086424856478b5aa9e95",
			Name:         "Strife (voice acting samples)",
			Patchinfo:    games.NONE,
			PWADRequires: "Strife (Registered)",
		},
	}
}

func BuildSharewareInfo() []wad.Entry {
	return []wad.Entry{
		{
			MD5Hash:   "bb545b9c4eca0ff92c14d466b3294023",
			Name:      "Strife (Shareware)",
			Version:   "1.1",
			Patchinfo: games.STRIFE_SHAREWARE,
			Status:    status.FINAL,
		},
		{
			MD5Hash:    "de2c8dcad7cca206292294bdab524292",
			Name:       "Strife (Shareware)",
			Version:    "1.0",
			Patchinfo:  games.STRIFE_SHAREWARE,
			Status:     status.NOTFINAL,
			Additional: "Download the latest shareware of Strife at https://www.doomworld.com/idgames/roguestuff/strife11",
		},
	}
}

func BuildConsolePortInfo() []wad.Entry {
	return []wad.Entry{
		{
			MD5Hash:   "47958a4fea8a54116e4b51fc155799c0",
			Name:      "Strife: Veteran Edition",
			Version:   "1.2",
			Patchinfo: games.STRIFE_VETERAN_EDITION,
			Status:    status.FINAL,
			Flags:     flags.RERELEASE,
		},
		{
			MD5Hash:   "2c0a712d3e39b010519c879f734d79ae",
			Name:      "Strife: Veteran Edition",
			Version:   "1.1",
			Patchinfo: games.STRIFE_VETERAN_EDITION,
			Status:    status.NOTFINAL,
			Flags:     flags.RERELEASE,
		},
		{
			MD5Hash:   "06a8f99b9b756ac908917c3868b8e3bc",
			Name:      "Strife: Veteran Edition",
			Version:   "1.0",
			Patchinfo: games.STRIFE_VETERAN_EDITION,
			Status:    status.NOTFINAL,
			Flags:     flags.RERELEASE,
		},
	}
}

func Populate() []wad.Entry {
	list := BuildStrifeInfo()
	list = append(list, BuildSharewareInfo()...)
	list = append(list, BuildConsolePortInfo()...)
	return list
}
