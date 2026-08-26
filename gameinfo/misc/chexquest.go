package misc

import (
	"wadverifier/wad"
	"wadverifier/wad/flags"
	"wadverifier/wad/games"
	"wadverifier/wad/status"
)

func BuildChexQuestInfo() []wad.Entry {
	return []wad.Entry{
		// Chex Quest
		{
			MD5Hash:    "25485721882b050afa96a56e5758dd52",
			Name:       "Chex Quest",
			Patchinfo:  games.IWAD,
			Status:     status.FINAL,
			Additional: "May require the DEHacked file if playing on a source port, available at https://www.doomworld.com/idgames/utils/exe_edit/patches/chexdeh",
		},
		{
			MD5Hash:   "f428a9a226f143a01b5782af611a83dd",
			Name:      "Chex Quest (Prototype)",
			Patchinfo: games.NONE,
			Flags:     flags.PRERELEASE,
		},

		{
			MD5Hash:      "fdc4ffa57e1983e30912c006284a3e01",
			Name:         "Chex Quest 2",
			Status:       status.FINAL,
			PWADRequires: "Chex Quest",
			Additional:   "May require the DEHacked file if playing on a source port, available at https://www.doomworld.com/idgames/utils/exe_edit/patches/chexdeh",
		},

		{
			MD5Hash:   "bce163d06521f9d15f9686786e64df13",
			Name:      "Chex Quest 3",
			Version:   "1.4",
			Patchinfo: games.CHEX_QUEST_3,
			Status:    status.FINAL,
		},
		{
			MD5Hash:    "cb001c34e424687191f299cc1dff4d68",
			Name:       "Chex Quest 3 (ModDB)",
			Version:    "Unidentified version",
			Patchinfo:  games.CHEX_QUEST_3,
			Status:     status.NOTFINAL,
			Additional: "This version of Chex Quest 3 has been released on ModDB. Since its status is unknown, it is still preferable to use the latest known build instead.",
		},
		{
			MD5Hash:   "59c985995db55cd2623c1893550d82b3",
			Name:      "Chex Quest 3",
			Version:   "1.0",
			Patchinfo: games.CHEX_QUEST_3,
			Status:    status.NOTFINAL,
		},
	}
}
