package heretic

import (
	"wadverifier/wad"
	"wadverifier/wad/flags"
	"wadverifier/wad/games"
	"wadverifier/wad/status"
)

const (
	AddendumKexHeretic = `This WAD is incompatible with sourceports that supports either Heretic or Hexen due to major differences with its original files.
• You will have to use the original WADs instead, found in the following sub-directories from:
	- STEAM: "<yoursteamfolder>\steamapps\common\Heretic + Hexen\dos\base\",
	- GOG: "<installfolder>\dos\base\" `
)

func BuildReleaseInfo() []wad.Entry {
	return []wad.Entry{

		{
			MD5Hash:   "66d686b1ed6d35ff103f15dbd30e0341",
			Name:      "Heretic: Shadow of the Serpent Riders",
			Version:   "1.3",
			Patchinfo: games.IWAD,
			Status:    status.FINAL,
		},
		{
			MD5Hash:   "1e4cb4ef075ad344dd63971637307e04",
			Name:      "Heretic",
			Version:   "1.2",
			Patchinfo: games.IWAD,
			Status:    status.NOTFINAL,
		},
		{
			MD5Hash:   "3117e399cdb4298eaa3941625f4b2923",
			Name:      "Heretic",
			Version:   "1.0",
			Patchinfo: games.IWAD,
			Status:    status.NOTFINAL,
		},
	}
}

func BuildPrototypeInfo() []wad.Entry {
	return []wad.Entry{
		{
			MD5Hash:    "fc7eab659f6ee522bb57acc1a946912f",
			Name:       "Heretic",
			Version:    "Wide-Area Beta",
			Patchinfo:  games.NONE,
			Status:     status.UNKNOWN,
			Flags:      flags.PRERELEASE,
			Additional: "This is the latest Beta version of Heretic.",
		},
	}
}

func BuildShareWareInfo() []wad.Entry {
	return []wad.Entry{
		{
			MD5Hash:    "ae779722390ec32fa37b0d361f7d82f8",
			Name:       "Heretic (Shareware)",
			Version:    "1.2",
			Patchinfo:  games.HERETIC_SHAREWARE,
			Status:     status.FINAL,
			Additional: "This is the latest Shareware version of Heretic.",
		},
		{
			MD5Hash:   "023b52175d2f260c3bdc5528df5d0a8c",
			Name:      "Heretic (Shareware)",
			Version:   "1.0",
			Patchinfo: games.HERETIC_SHAREWARE,
			Status:    status.NOTFINAL,
		},
	}
}

func BuildConsolePortInfo() []wad.Entry {
	return []wad.Entry{
		{
			MD5Hash:    "a5de95a3162e71b5ffc568ba2343cd46",
			Name:       "Heretic (Heretic + Hexen)",
			Version:    "1.0.4575 - 46e45cbd (Jul 17, 2025)",
			Patchinfo:  games.KEX_HERETIC_HEXEN2025,
			Flags:      flags.RERELEASE,
			Additional: AddendumKexHeretic,
		},
		{
			MD5Hash:   "aa50f3cf21d1bcf15241b8310d67e316",
			Name:      "Heretic Extras (Heretic + Hexen)",
			Version:   "1.0.4575 - 46e45cbd (Jul 17, 2025)",
			Patchinfo: games.KEX_HERETIC_HEXEN2025,
			Flags:     flags.RERELEASE,
		},
		{
			MD5Hash:   "683588a48557a317537fc74a5dcd26c6",
			Name:      "Heretic: Faith Renewed (Heretic + Hexen)",
			Version:   "1.0.4575 - 46e45cbd (Jul 17, 2025)",
			Patchinfo: games.KEX_HERETIC_HEXEN2025,
			Flags:     flags.RERELEASE,
		},
		{
			MD5Hash:   "ad7e85b4b8fdb74b80ba9598ea333f76",
			Name:      "Heretic Original Soundtrack (Heretic + Hexen)",
			Version:   "1.0.4575 - 46e45cbd (Jul 17, 2025)",
			Patchinfo: games.KEX_HERETIC_HEXEN2025,
			Flags:     flags.RERELEASE | flags.HIDDEN,
		},
		{
			MD5Hash:    "4a354d6575382dca17f580bdf5c67f66",
			Name:       "Heretic Remade Soundtrack (Heretic + Hexen)",
			Version:    "1.0.4575 - 46e45cbd (Jul 17, 2025)",
			Patchinfo:  games.KEX_HERETIC_HEXEN2025,
			Flags:      flags.RERELEASE | flags.HIDDEN,
			Additional: "OST remade by Andrew Hulshault.",
		},
	}
}

func Populate() []wad.Entry {
	list := BuildReleaseInfo()
	list = append(list, BuildPrototypeInfo()...)
	list = append(list, BuildShareWareInfo()...)
	list = append(list, BuildConsolePortInfo()...)
	return list
}
