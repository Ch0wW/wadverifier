package main

const (
	AddentumMPDowngrade = `
If you want to use it on multiplayer source ports, you'll need to patch it using Peter Vaskovics's tool, available below:
	• Windows binaries: http://downloads.zdaemon.org/iwadpatcher-1.2-bin.zip
	• Source code: https://github.com/petervas/iwadpatcher`

	AddentumDoomBethesda = `This WAD is incompatible with sourceports due to major differences with its original files.
• You need to use the original WAD instead, found in the following directory:
	- "<yoursteamfolder>\steamapps\common\Ultimate DOOM\base\DOOM.WAD" for the Steam version,
	- "<bethesdafolder>\games\Ultimate Doom\base\DOOM.WAD" for the Bethesda Launcher version.`

	AddentumDoomIIBethesda = `This WAD is incompatible with sourceports due to major differences with its original files.
• You need to use the original WAD instead, found in the following directory:
	- "<yoursteamfolder>\steamapps\common\DOOM 2\base\DOOM2.WAD" for the Steam version,
	- "<bethesdafolder>\games\Doom 2\base\DOOM2.WAD" for the Bethesda Launcher version.`

	AddentumDoomKexDoom = `This WAD is incompatible with sourceports due to major differences with its original files.
• You need to use the original WAD instead, found in the following directory:
	- "<yoursteamfolder>\steamapps\common\Ultimate DOOM\base\DOOM.WAD" for the Steam version,
	- "<installfolder>\base\DOOM.WAD" for the GOG version.`

	AddentumDoomIIKexDoom = `This WAD is incompatible with sourceports due to major differences with its original files.
• You need to use the original WAD instead, found in the following directory:
	- "<yoursteamfolder>\steamapps\common\Ultimate DOOM\base\doom2\DOOM2.WAD",
	- "<installfolder>\base\doom2\DOOM2.WAD" for the GOG version`

	AddentumTNTKexDoom = `This WAD is incompatible with sourceports due to major differences with its original files.
• You need to use the original WAD instead, found in the following directory:
	- "<yoursteamfolder>\steamapps\common\Ultimate DOOM\base\tnt\TNT.WAD",
	- "<installfolder>\base\tnt\TNT.WAD" for the GOG version`

	AddentumPLUTONIAKexDoom = `This WAD is incompatible with sourceports due to major differences with its original files.
• You need to use the original WAD instead, found in the following directory:
	- "<yoursteamfolder>\steamapps\common\Ultimate DOOM\base\plutonia\PLUTONIA.WAD",
	- "<installfolder>\base\plutonia\PLUTONIA.WAD" for the GOG version`

	AddentumKexHeretic = `This WAD is incompatible with sourceports that supports either Heretic or Hexen due to major differences with its original files.
• You will have to use the original WADs instead, found in the following sub-directories from:
	- STEAM: "<yoursteamfolder>\steamapps\common\Heretic + Hexen\dos\base\",
	- GOG: "<installfolder>\dos\base\" `
)

// =======================
// PopulateIWADInfos
// Populate the IWAD lists with their MD5 and versions
// =======================
func PopulateIWADInfos() {
	Populate_Doom()
	Populate_DoomII()
	Populate_FinalDOOM()
	Populate_HereticHexen()
	Populate_Strife()
	Populate_FreeDoom()
	Populate_Misc()

	if len(customdata.Data) > 0 {
		PWADInfo_Custom = customdata.Data
	}

}

func Populate_Doom() {

	// DOOM/UDOOM population
	IWADInfo_Doom = []WadInfo{
		{
			MD5Hash: "740901119ba2953e3c7f3764eca6e128",
			Name:    "DOOM (Press Release)",
			Version: "0.2",
			Game:    GAME_NONE,
			Status:  IS_UNKNOWN,
			Flags:   FL_PRERELEASE,
		},
		{
			MD5Hash: "dae9b1eea1a8e090fdfa5707187f4a43",
			Name:    "DOOM (Press Release)",
			Version: "0.3",
			Game:    GAME_NONE,
			Status:  IS_UNKNOWN,
			Flags:   FL_PRERELEASE,
		},
		{
			MD5Hash: "b6afa12a8b22e2726a8ff5bd249223de",
			Name:    "DOOM (Press Release)",
			Version: "0.4",
			Game:    GAME_NONE,
			Status:  IS_UNKNOWN,
			Flags:   FL_PRERELEASE,
		},
		{
			MD5Hash: "9c877480b8ef33b7074f1f0c07ed6487",
			Name:    "DOOM (Press Release)",
			Version: "0.5",
			Game:    GAME_NONE,
			Status:  IS_UNKNOWN,
			Flags:   FL_PRERELEASE,
		},
		{
			MD5Hash: "049e32f18d9c9529630366cfc72726ea",
			Name:    "DOOM (Press Release - DOOMPRES.WAD)",
			Version: "October 4th, 1993",
			Game:    GAME_NONE,
			Status:  IS_UNKNOWN,
			Flags:   FL_PRERELEASE,
		},
		{
			MD5Hash: "90facab21eede7981be10790e3f82da2",
			Name:    "DOOM (Shareware)",
			Version: "1.0",
			Game:    GAME_SHAREWARE,
			Status:  IS_NOTFINAL,
		},
		{
			MD5Hash: "52cbc8882f445573ce421fa5453513c1",
			Name:    "DOOM (Shareware)",
			Version: "1.1",
			Game:    GAME_SHAREWARE,
			Status:  IS_NOTFINAL,
		},
		{
			MD5Hash: "981b03e6d1dc033301aa3095acc437ce",
			Name:    "DOOM (Registered)",
			Version: "1.1",
			Game:    GAME_IWAD,
			Status:  IS_NOTFINAL,
		},
		{
			MD5Hash: "30aa5beb9e5ebfbbe1e1765561c08f38",
			Name:    "DOOM (Shareware)",
			Version: "1.2",
			Game:    GAME_SHAREWARE,
			Status:  IS_NOTFINAL,
		},
		{
			MD5Hash: "792fd1fea023d61210857089a7c1e351",
			Name:    "DOOM (Registered)",
			Version: "1.2",
			Game:    GAME_IWAD,
			Status:  IS_NOTFINAL,
		},
		{
			MD5Hash: "a21ae40c388cb6f2c3cc1b95589ee693",
			Name:    "DOOM (Shareware Beta)",
			Version: "1.4",
			Game:    GAME_SHAREWARE,
			Status:  IS_NOTFINAL,
			Flags:   FL_PRERELEASE,
		},
		{
			MD5Hash: "e280233d533dcc28c1acd6ccdc7742d4",
			Name:    "DOOM (Shareware Beta)",
			Version: "1.5",
			Game:    GAME_SHAREWARE,
			Status:  IS_NOTFINAL,
			Flags:   FL_PRERELEASE,
		},
		{
			MD5Hash: "762fd6d4b960d4b759730f01387a50a1",
			Name:    "DOOM (Shareware Beta)",
			Version: "1.6",
			Game:    GAME_SHAREWARE,
			Status:  IS_NOTFINAL,
			Flags:   FL_PRERELEASE,
		},
		{
			MD5Hash:    "c428ea394dc52835f2580d5bfd50d76f",
			Name:       "DOOM (Shareware)",
			Version:    "1.666",
			Game:       GAME_SHAREWARE,
			Status:     IS_NOTFINAL,
			Additional: "Get the latest shareware version of DooM at https://www.doomworld.com/idgames/idstuff/doom/doom19s",
		},
		{
			MD5Hash: "54978d12de87f162b9bcc011676cb3c0",
			Name:    "DOOM (Registered)",
			Version: "1.666",
			Game:    GAME_IWAD,
			Status:  IS_NOTFINAL,
		},
		{
			MD5Hash:    "5f4eb849b1af12887dec04a2a12e5e62",
			Name:       "DOOM (Shareware)",
			Version:    "1.8",
			Game:       GAME_SHAREWARE,
			Status:     IS_NOTFINAL,
			Additional: "Get the latest shareware version of DooM at https://www.doomworld.com/idgames/idstuff/doom/doom19s",
		},
		{
			MD5Hash: "11e1cd216801ea2657723abc86ecb01f",
			Name:    "DOOM (Registered)",
			Version: "1.8",
			Game:    GAME_IWAD,
			Status:  IS_NOTFINAL,
		},
		{
			MD5Hash: "f0cefca49926d00903cf57551d901abe",
			Name:    "DOOM (Shareware)",
			Version: "1.9",
			Game:    GAME_IWAD,
			Status:  IS_FINAL,
		},
		{
			MD5Hash:    "1cd63c5ddff1bf8ce844237f580e9cf3",
			Name:       "DOOM (Registered)",
			Version:    "1.9",
			Game:       GAME_IWAD,
			Status:     IS_NOTFINAL,
			Additional: "Technically speaking, this IWAD is still the latest version of DOOM, but it is highly recommended to upgrade to The Ultimate Doom version.",
		},
		{
			MD5Hash: "dae77aff77a0491e3b7254c9c8401aa8",
			Name:    "DOOM for Pocket PC",
			Game:    GAME_IWAD,
			Status:  IS_FINAL,
		},
		{
			MD5Hash: "c4fe9fd920207691a9f493668e0a2083",
			Name:    "The Ultimate DOOM",
			Version: "1.9",
			Game:    GAME_IWAD,
			Status:  IS_FINAL,
		},
		{
			MD5Hash:    "0c8758f102ccafe26a3040bee8ba5021",
			Name:       "The Ultimate DOOM (XBox Version)",
			Game:       GAME_IWAD,
			Flags:      FL_RERELEASE | FL_HIDDEN,
			Additional: AddentumMPDowngrade,
		},
		{
			MD5Hash:    "72286ddc680d47b9138053dd944b2a3d",
			Version:    "The Ultimate DOOM (XBox Live Arcade version)",
			Game:       GAME_IWAD,
			Flags:      FL_RERELEASE | FL_HIDDEN,
			Additional: AddentumMPDowngrade,
		},
		{
			MD5Hash:    "fb35c4a5a9fd49ec29ab6e900572c524",
			Version:    "The Ultimate DOOM (Doom 3 - BFG Edition)",
			Game:       GAME_IWAD,
			Flags:      FL_RERELEASE | FL_HIDDEN,
			Additional: AddentumMPDowngrade,
		},
		{
			MD5Hash:    "8517c4e8f0eef90b82852667d345eb86",
			Name:       "DOOM (Bethesda Version / DOOM Unity)",
			Version:    "2020_08_21 Build #13735 doom",
			Game:       GAME_DOOMUNITY,
			Status:     IS_FINAL,
			Flags:      FL_RERELEASE,
			Additional: AddentumDoomBethesda,
		},
		{
			MD5Hash:    "3b37188f6337f15718b617c16e6e7a9c",
			Name:       "DOOM (Doom + Doom II)",
			Version:    "Update 1",
			Game:       GAME_KEXDOOM2024,
			Status:     IS_FINAL,
			Flags:      FL_RERELEASE,
			Additional: AddentumDoomKexDoom,
		},
	}
}

func Populate_DoomII() {

	// DOOM/UDOOM population
	IWADInfo_Doom2 = []WadInfo{
		{
			MD5Hash: "30e3c2d0350b67bfbf47271970b74b2f",
			Name:    "DOOM II",
			Version: "1.666",
			Game:    GAME_IWAD,
			Status:  IS_NOTFINAL,
		},
		{
			MD5Hash:    "d9153ced9fd5b898b36cc5844e35b520",
			Name:       "DOOM II (German Release)",
			Version:    "1.666",
			Game:       GAME_IWAD,
			Status:     IS_FINAL,
			Flags:      FL_HIDDEN,
			Additional: "This is the latest release officially released in Germany, but it is not the latest version of Doom II. It is strongly recommended to update it to v1.9 to restore censored contents and play online.",
		},
		{
			MD5Hash: "ea74a47a791fdef2e9f2ea8b8a9da13b",
			Name:    "DOOM II",
			Version: "1.7",
			Game:    GAME_IWAD,
			Status:  IS_NOTFINAL,
		},
		{
			MD5Hash: "d7a07e5d3f4625074312bc299d7ed33f",
			Name:    "DOOM II",
			Version: "1.7a",
			Game:    GAME_IWAD,
			Status:  IS_NOTFINAL,
		},
		{
			MD5Hash: "c236745bb01d89bbb866c8fed81b6f8c",
			Name:    "DOOM II",
			Version: "1.8",
			Game:    GAME_IWAD,
			Status:  IS_NOTFINAL,
		},
		{
			MD5Hash:    "3cb02349b3df649c86290907eed64e7b",
			Name:       "DOOM II (French Release)",
			Version:    "1.8",
			Game:       GAME_IWAD,
			Status:     IS_FINAL,
			Flags:      FL_HIDDEN,
			Additional: "This is the latest release officially translated in french, but it is not the latest release of Doom II. It is strongly recommended to update it to v1.9 in order to play online.",
		},
		{
			MD5Hash: "25e1459ca71d321525f84628f45ca8cd",
			Name:    "DOOM II",
			Version: "1.9",
			Game:    GAME_IWAD,
			Status:  IS_FINAL,
		},
		{
			MD5Hash:    "a793ebcdd790afad4a1f39cc39a893bd",
			Name:       "DOOM II (XBox version)",
			Game:       GAME_IWAD,
			Status:     IS_FINAL,
			Flags:      FL_RERELEASE | FL_HIDDEN,
			Additional: AddentumMPDowngrade,
		},
		{
			MD5Hash:    "43c2df32dc6c740cb11f34dc5ab693fa",
			Name:       "Doom II (XBox Live Arcade version)",
			Version:    "1.9",
			Game:       GAME_IWAD,
			Status:     IS_FINAL,
			Flags:      FL_RERELEASE | FL_HIDDEN,
			Additional: AddentumMPDowngrade,
		},
		{
			MD5Hash:    "c3bea40570c23e511a7ed3ebcd9865f7",
			Name:       "DOOM II (Doom 3 BFG Edition)",
			Game:       GAME_IWAD,
			Status:     IS_FINAL,
			Flags:      FL_RERELEASE | FL_HIDDEN,
			Additional: AddentumMPDowngrade,
		},
		{
			MD5Hash: "f617591a6c5d07037eb716dc4863e26b",
			Name:    "DOOM II (Doom 3 BFG Edition - XBox 360 release)",
			Game:    GAME_IWAD,
			Flags:   FL_RERELEASE | FL_HIDDEN,
		},
		{
			MD5Hash: "9640fc4b2c8447bbd28f2080725d5c51",
			Name:    "DOOM II (Tapwave Zodiac Version)",
			Game:    GAME_IWAD,
			Status:  IS_FINAL,
			Flags:   FL_RERELEASE | FL_HIDDEN,
		},
		{
			MD5Hash:    "8ab6d0527a29efdc1ef200e5687b5cae",
			Name:       "DOOM II (Bethesda Version / DOOM Unity)",
			Version:    "2020_08_21 Build #13736 doom2",
			Game:       GAME_IWAD,
			Status:     IS_FINAL,
			Flags:      FL_RERELEASE | FL_HIDDEN,
			Additional: AddentumDoomIIBethesda,
		},
		{
			MD5Hash:    "64a4c88a871da67492aaa2020a068cd8",
			Name:       "DOOM II (Doom + Doom II)",
			Version:    "Update 1",
			Game:       GAME_KEXDOOM2024,
			Flags:      FL_RERELEASE,
			Additional: AddentumDoomKexDoom,
		},
	}
}

func Populate_FinalDOOM() {
	// DOOM/UDOOM population
	IWADInfo_FinalDoom = []WadInfo{
		{
			MD5Hash: "75c8cf89566741fa9d22447604053bd7",
			Name:    "Final DOOM: The Plutonia Experiment",
			Game:    GAME_IWAD,
			Status:  IS_FINAL,
		},
		{
			MD5Hash:    "3493be7e1e2588bc9c8b31eab2587a04",
			Name:       "Final DOOM: The Plutonia Experiment",
			Version:    "id Anthology release",
			Game:       GAME_IWAD,
			Flags:      FL_RERELEASE,
			Status:     IS_UNKNOWN,
			Additional: "It is unknown if this version is fully compatible with online servers or demos",
		},
		{
			MD5Hash: "4e158d9953c79ccf97bd0663244cc6b6",
			Name:    "Final DOOM: TNT: Evilution",
			Game:    GAME_IWAD,
			Status:  IS_FINAL,
		},
		{
			MD5Hash:    "1d39e405bf6ee3df69a8d2646c8d5c49",
			Name:       "Final DOOM: TNT: Evilution",
			Version:    "id Anthology release",
			Game:       GAME_IWAD,
			Flags:      FL_RERELEASE,
			Status:     IS_UNKNOWN,
			Additional: "It is unknown if this version is fully compatible with online servers or demos",
		},
		{
			MD5Hash:    "7a77ee148fd9ee5bc599356218f6f6b5",
			Name:       "Final DOOM: TNT: Evilution - Fix for MAP31",
			Game:       GAME_IWAD,
			Flags:      FL_HIDDEN,
			Additional: "This PWAD fixes the yellow keycard not appearing on TNT MAP31.",
		},
		{
			MD5Hash:    "ad7885c17a6b9b79b09d7a7634dd7e2c",
			Name:       "Final DOOM: TNT: Evilution (Doom + Doom II)",
			Version:    "Update 1",
			Game:       GAME_IWAD,
			Flags:      FL_RERELEASE,
			Additional: AddentumTNTKexDoom,
		},
		{
			MD5Hash:    "e47cf6d82a0ccedf8c1c16a284bb5937",
			Name:       "Final DOOM: The Plutonia Experiment (Doom + Doom II)",
			Version:    "Update 1",
			Game:       GAME_KEXDOOM2024,
			Flags:      FL_RERELEASE,
			Additional: AddentumPLUTONIAKexDoom,
		},
	}
}

func Populate_HereticHexen() {
	// Heretic population
	IWADInfo_Heretic = []WadInfo{
		{
			MD5Hash:    "fc7eab659f6ee522bb57acc1a946912f",
			Name:       "Heretic",
			Version:    "Wide-Area Beta",
			Status:     IS_UNKNOWN,
			Flags:      FL_PRERELEASE,
			Additional: "This is the latest Beta version of Heretic.",
		},
		{
			MD5Hash: "023b52175d2f260c3bdc5528df5d0a8c",
			Name:    "Heretic (Shareware)",
			Version: "1.0",
			Game:    GAME_IWAD,
			Status:  IS_NOTFINAL,
		},
		{
			MD5Hash:    "ae779722390ec32fa37b0d361f7d82f8",
			Name:       "Heretic (Shareware)",
			Version:    "1.2",
			Game:       GAME_IWAD,
			Status:     IS_FINAL,
			Additional: "This is the latest Shareware version of Heretic.",
		},

		{
			MD5Hash: "66d686b1ed6d35ff103f15dbd30e0341",
			Name:    "Heretic: Shadow of the Serpent Riders",
			Version: "1.3",
			Game:    GAME_IWAD,
			Status:  IS_FINAL,
		},
		{
			MD5Hash: "3117e399cdb4298eaa3941625f4b2923",
			Name:    "Heretic",
			Version: "1.0",
			Game:    GAME_IWAD,
			Status:  IS_NOTFINAL,
		},
		{
			MD5Hash: "1e4cb4ef075ad344dd63971637307e04",
			Name:    "Heretic",
			Version: "1.2",
			Game:    GAME_IWAD,
			Status:  IS_NOTFINAL,
		},
	}

	IWADInfo_Hexen = []WadInfo{
		{
			MD5Hash: "abb033caf81e26f12a2103e1fa25453f",
			Name:    "Hexen",
			Version: "1.1",
			Game:    GAME_IWAD,
			Status:  IS_FINAL,
		},
		{
			MD5Hash: "b2543a03521365261d0a0f74d5dd90f0",
			Name:    "Hexen",
			Version: "1.0",
			Game:    GAME_IWAD,
			Status:  IS_NOTFINAL,
		},

		// MAC versions
		{
			MD5Hash:    "b68140a796f6fd7f3a5d3226a32b93be",
			Name:       "Hexen (Macintosh version)",
			Version:    "1.1",
			Game:       GAME_IWAD,
			Status:     IS_FINAL,
			Additional: "Despite being the Macintosh release, it contains differences that makes it incompatible with the Windows release of the game!",
		},
		{
			MD5Hash:    "b68140a796f6fd7f3a5d3226a32b93be",
			Name:       "Hexen Demo (Macintosh version)",
			Game:       GAME_IWAD,
			Status:     IS_FINAL,
			Additional: "Despite being the Macintosh release, it contains differences that makes it incompatible with the Windows release of the game!",
		},

		// hexdd.wad
		// Steam only distributes version 1.0
		{
			MD5Hash:    "1077432e2690d390c256ac908b5f4efa",
			Name:       "Hexen: Deathkings of the Dark Citadel",
			Version:    "1.0",
			Game:       GAME_IWAD,
			Status:     IS_NOTFINAL,
			Additional: "Requires Hexen 1.1",
		},
		{
			MD5Hash:    "78d5898e99e220e4de64edaa0e479593",
			Name:       "Hexen: Deathkings of the Dark Citadel",
			Version:    "1.1",
			Game:       GAME_IWAD,
			Status:     IS_FINAL,
			Additional: "Requires Hexen 1.1 and needs to be run as a PWAD.",
		},

		// OTHER IRREVELENT THINGS
		{
			MD5Hash: "876a5a44c7b68f04b3bb9bc7a5bd69d6",
			Name:    "Hexen Demo",
			Version: "1.0",
			Game:    GAME_IWAD,
			Status:  IS_FINAL,
		},
		{
			MD5Hash: "9178a32a496ff5befebfe6c47dac106c",
			Name:    "Hexen Demo Beta",
			Game:    GAME_IWAD,
			Flags:   FL_PRERELEASE,
		},
		{
			MD5Hash: "c88a2bb3d783e2ad7b599a8e301e099e",
			Version: "Hexen Beta",
			Game:    GAME_IWAD,
			Flags:   FL_PRERELEASE,
		},

		// Heretic + HeXen
		{
			MD5Hash:    "a5de95a3162e71b5ffc568ba2343cd46",
			Name:       "Heretic (Heretic + Hexen)",
			Version:    "1.0.4575 - 46e45cbd (Jul 17, 2025)",
			Game:       GAME_KEXHEREXEN2025,
			Flags:      FL_RERELEASE,
			Additional: AddentumKexHeretic,
		},
		{
			MD5Hash: "aa50f3cf21d1bcf15241b8310d67e316",
			Name:    "Heretic Extras (Heretic + Hexen)",
			Version: "1.0.4575 - 46e45cbd (Jul 17, 2025)",
			Game:    GAME_KEXHEREXEN2025,
			Flags:   FL_RERELEASE,
		},
		{
			MD5Hash: "683588a48557a317537fc74a5dcd26c6",
			Name:    "Heretic: Faith Renewed (Heretic + Hexen)",
			Version: "1.0.4575 - 46e45cbd (Jul 17, 2025)",
			Game:    GAME_KEXHEREXEN2025,
			Flags:   FL_RERELEASE,
		},
		{
			MD5Hash:    "ad7e85b4b8fdb74b80ba9598ea333f76",
			Name:       "Heretic Original Soundtrack (Heretic + Hexen)",
			Version:    "1.0.4575 - 46e45cbd (Jul 17, 2025)",
			Game:       GAME_KEXHEREXEN2025,
			Flags:      FL_RERELEASE | FL_HIDDEN,
			Additional: AddentumKexHeretic,
		},
		{
			MD5Hash:    "4a354d6575382dca17f580bdf5c67f66",
			Name:       "Heretic Remade Soundtrack (Heretic + Hexen)",
			Version:    "1.0.4575 - 46e45cbd (Jul 17, 2025)",
			Game:       GAME_KEXHEREXEN2025,
			Flags:      FL_RERELEASE | FL_HIDDEN,
			Additional: "OST remade by Andrew Hulshault.",
		},
		{
			MD5Hash:    "a66ae0448436a990b3aecd018bc2708a",
			Name:       "Hexen (Heretic + Hexen)",
			Version:    "1.0.4575 - 46e45cbd (Jul 17, 2025)",
			Game:       GAME_KEXHEREXEN2025,
			Flags:      FL_RERELEASE,
			Additional: AddentumKexHeretic,
		},
		{
			MD5Hash: "e0ba5039f3baf750714d845c992a5331",
			Name:    "Hexen Extras (Heretic + Hexen)",
			Version: "1.0.4575 - 46e45cbd (Jul 17, 2025)",
			Game:    GAME_KEXHEREXEN2025,
			Flags:   FL_RERELEASE,
		},
		{
			MD5Hash:    "2bb0f56ab1f98000990524c1b67e8759",
			Name:       "Hexen: Deathkings of the Dark Citadel (Heretic + Hexen)",
			Version:    "1.0.4575 - 46e45cbd (Jul 17, 2025)",
			Game:       GAME_KEXHEREXEN2025,
			Flags:      FL_RERELEASE,
			Additional: AddentumKexHeretic,
		}, {
			MD5Hash: "335336b45cb3e99f6f35b46418d8e31f",
			Name:    "Hexen: Deathkings of the Dark Citadel Extras (Heretic + Hexen)",
			Version: "1.0.4575 - 46e45cbd (Jul 17, 2025)",
			Game:    GAME_KEXHEREXEN2025,
			Flags:   FL_RERELEASE,
		},
		{
			MD5Hash:    "aee983213be00b2e5e02c09675dc608f",
			Name:       "Hexen: Vestiges of Grandeur (Heretic + Hexen)",
			Version:    "1.0.4575 - 46e45cbd (Jul 17, 2025)",
			Game:       GAME_KEXHEREXEN2025,
			Flags:      FL_RERELEASE,
			Additional: AddentumKexHeretic,
		},
		{
			MD5Hash:    "aee441fadfaa8bf9c06bd39d6abd6775",
			Name:       "Hexen: Test WAD (Heretic + Hexen)",
			Version:    "1.0.4575 - 46e45cbd (Jul 17, 2025)",
			Game:       GAME_KEXHEREXEN2025,
			Flags:      FL_RERELEASE | FL_HIDDEN,
			Additional: AddentumKexHeretic,
		},
		{
			MD5Hash: "2fca43a18a535511efa8c7c27e202c12",
			Name:    "Hexen Original Soundtrack (Heretic + Hexen)",
			Version: "1.0.4575 - 46e45cbd (Jul 17, 2025)",
			Game:    GAME_KEXHEREXEN2025,
			Flags:   FL_RERELEASE | FL_HIDDEN,
		},
		{
			MD5Hash:    "4a4cbb09f2ecb1b6f7bf3477ca54c18c",
			Name:       "Hexen Remade Soundtrack (Heretic + Hexen)",
			Version:    "1.0.4575 - 46e45cbd (Jul 17, 2025)",
			Game:       GAME_KEXHEREXEN2025,
			Flags:      FL_RERELEASE | FL_HIDDEN,
			Additional: "OST remade by Andrew Hulshault.",
		},
	}
}

func Populate_Strife() {
	// Strife population
	IWADInfo_Strife = []WadInfo{
		{
			MD5Hash:    "2fed2031a5b03892106e0f117f17901f",
			Name:       "Strafe (Registered)",
			Version:    "1.2 - 1.31",
			Status:     IS_FINAL,
			Additional: "Your IWAD is up-to-date. However, the latest updates of Strife only modify the executable. Please make sure it is updated to the latest version.",
		},
		{
			MD5Hash: "bb545b9c4eca0ff92c14d466b3294023",
			Name:    "Strife (Shareware)",
			Version: "1.1",
			Status:  IS_FINAL,
		},
		{
			MD5Hash: "8f2d3a6a289f5d2f2f9c1eec02b47299",
			Name:    "Strife (Registered)",
			Version: "1.0",
			Status:  IS_NOTFINAL,
		},
		{
			MD5Hash:    "de2c8dcad7cca206292294bdab524292",
			Name:       "Strife (Shareware)",
			Version:    "1.0",
			Status:     IS_NOTFINAL,
			Additional: "Download the latest shareware of Strife at https://www.doomworld.com/idgames/roguestuff/strife11",
		},

		// Additionnal IWADs
		{
			MD5Hash:      "082234d6a3f7086424856478b5aa9e95",
			Name:         "Strife (voice acting samples)",
			Status:       IS_FINAL,
			PWADRequires: "Strife Registered",
		},
	}

	// Strife: Veteran Edition
	IWADInfo_SVE = []WadInfo{
		{
			MD5Hash: "47958a4fea8a54116e4b51fc155799c0",
			Name:    "Strife: Veteran Edition",
			Version: "1.2",
			Game:    GAME_STRIFE_VE,
			Status:  IS_FINAL,
			Flags:   FL_RERELEASE,
		},
		{
			MD5Hash: "2c0a712d3e39b010519c879f734d79ae",
			Name:    "Strife: Veteran Edition",
			Version: "1.1",
			Game:    GAME_STRIFE_VE,
			Status:  IS_NOTFINAL,
			Flags:   FL_RERELEASE,
		},
		{
			MD5Hash: "06a8f99b9b756ac908917c3868b8e3bc",
			Name:    "Strife: Veteran Edition",
			Version: "1.0",
			Game:    GAME_STRIFE_VE,
			Status:  IS_NOTFINAL,
			Flags:   FL_RERELEASE,
		},
	}
}

func Populate_FreeDoom() {
	IWADInfo_FreeDoom = []WadInfo{

		//---------------
		// First, the latest IWADs
		{
			MD5Hash: "b93be13d05148dd01614bc205a03648e",
			Name:    "Freedoom: Phase 1",
			Version: "0.13.0",
			Game:    GAME_FREEDOOM,
			Status:  IS_FINAL,
		},
		{
			MD5Hash: "cd666466759b5e5f63af93c5f0ffd0a1",
			Name:    "Freedoom: Phase 2",
			Version: "0.13.0",
			Game:    GAME_FREEDOOM,
			Status:  IS_FINAL,
		},
		{
			MD5Hash: "908dfd77a14cc490c4cea94b62d13449",
			Name:    "FreeDM",
			Version: "0.13.0",
			Game:    GAME_FREEDOOM,
			Status:  IS_FINAL,
		},

		//----------------
		// Now, list the outdated ones.

		// 0.12.1
		{
			MD5Hash: "b36aa44a23045e503c19af4b4c438a78",
			Name:    "Freedoom: Phase 1",
			Version: "0.12.1",
			Game:    GAME_FREEDOOM,
			Status:  IS_NOTFINAL,
		},
		{
			MD5Hash: "ca9a4159a7833544a89144c7f5053412",
			Name:    "Freedoom: Phase 2",
			Version: "0.12.1",
			Game:    GAME_FREEDOOM,
			Status:  IS_NOTFINAL,
		},
		{
			MD5Hash: "d40c932a9183ded919afa89f4a729668",
			Name:    "FreeDM",
			Version: "FreeDM 0.12.1",
			Game:    GAME_FREEDOOM,
			Status:  IS_NOTFINAL,
		},

		// 0.12.0
		{
			MD5Hash: "0c5f8ff45cc3538d368a0f8d8fc11ce3",
			Name:    "Freedoom: Phase 1",
			Version: "0.12.0",
			Game:    GAME_FREEDOOM,
			Status:  IS_NOTFINAL,
		},
		{
			MD5Hash: "83560b2963424fa4a2eb971194428bf8",
			Name:    "Freedoom: Phase 2",
			Version: "0.12.0",
			Game:    GAME_FREEDOOM,
			Status:  IS_NOTFINAL,
		},
		{
			MD5Hash: "3250aad8b1d40fb7b25b7df6573eb29f",
			Name:    "FreeDM",
			Version: "FreeDM 0.12.0",
			Game:    GAME_FREEDOOM,
			Status:  IS_NOTFINAL,
		},

		// 0.11.3
		{
			MD5Hash: "ea471a3d38fcee0fb3a69bcd3221e335",
			Name:    "Freedoom: Phase 1",
			Version: "0.11.3",
			Game:    GAME_FREEDOOM,
			Status:  IS_NOTFINAL,
		},
		{
			MD5Hash: "984f99af08f085e38070f51095ab7c31",
			Name:    "Freedoom: Phase 2",
			Version: "0.11.3",
			Game:    GAME_FREEDOOM,
			Status:  IS_NOTFINAL,
		},
		{
			MD5Hash: "87ee2494d921633420ce9bdb418127c4",
			Name:    "FreeDM",
			Version: "0.11.3",
			Game:    GAME_FREEDOOM,
			Status:  IS_NOTFINAL,
		},

		// 0.11.2
		{
			MD5Hash: "9352b09ae878dc52c6c18aa38acda6eb",
			Name:    "FreeDM",
			Version: "0.11.2",
			Game:    GAME_FREEDOOM,
			Status:  IS_NOTFINAL,
		},
		{
			MD5Hash: "6d00c49520be26f08a6bd001814a32ab",
			Name:    "Freedoom: Phase 1",
			Version: "0.11.2",
			Game:    GAME_FREEDOOM,
			Status:  IS_NOTFINAL,
		},
		{
			MD5Hash: "90832a872b5bb0aca4ca0b20419aad5d",
			Name:    "Freedoom: Phase 2",
			Version: "0.11.2",
			Game:    GAME_FREEDOOM,
			Status:  IS_NOTFINAL,
		},

		// 0.11.1
		{
			MD5Hash: "77ba9c0f75c32e4a729490688bb99241",
			Name:    "FreeDM",
			Version: "0.11.1",
			Game:    GAME_FREEDOOM,
			Status:  IS_NOTFINAL,
		},
		{
			MD5Hash: "35312e99d2473297aabe0602700bee8a",
			Name:    "Freedoom: Phase 1",
			Version: "0.11.1",
			Game:    GAME_FREEDOOM,
			Status:  IS_NOTFINAL,
		},
		{
			MD5Hash: "ec5b38b30ba2b70e278205776af3fbb5",
			Name:    "Freedoom: Phase 2",
			Version: "0.11.1",
			Game:    GAME_FREEDOOM,
			Status:  IS_NOTFINAL,
		},

		// 0.11.0
		{
			MD5Hash: "d76d3973c075b069ecb4e16dc9eacbb4",
			Name:    "FreeDM",
			Version: "0.11",
			Game:    GAME_FREEDOOM,
			Status:  IS_NOTFINAL,
		},
		{
			MD5Hash: "21a4707fc25d29edf4b098bd400c5c42",
			Name:    "Freedoom: Phase 1",
			Version: "0.11",
			Game:    GAME_FREEDOOM,
			Status:  IS_NOTFINAL,
		},
		{
			MD5Hash: "b1018017c61b06e33c11102d8bafaad0",
			Name:    "Freedoom: Phase 2",
			Version: "0.11",
			Game:    GAME_FREEDOOM,
			Status:  IS_NOTFINAL,
		},

		// 0.10.1
		{
			MD5Hash: "bd4f359f1963e388beda014c5548b420",
			Name:    "FreeDM",
			Version: "0.10.1",
			Game:    GAME_FREEDOOM,
			Status:  IS_NOTFINAL,
		},
		{
			MD5Hash: "91de79621a393a08c39a9ab2c034b766",
			Name:    "Freedoom: Phase 1",
			Version: "0.10.1",
			Game:    GAME_FREEDOOM,
			Status:  IS_NOTFINAL,
		},
		{
			MD5Hash: "dd9c9e73f5f50d3778c85573cd08d9a4",
			Name:    "Freedoom: Phase 2",
			Version: "0.10.1",
			Game:    GAME_FREEDOOM,
			Status:  IS_NOTFINAL,
		},

		// 0.10.0
		{
			MD5Hash: "9b8d72b59fd93b2b3e116149baa1b142",
			Name:    "Freedoom: Phase 1",
			Version: "0.10",
			Game:    GAME_FREEDOOM,
			Status:  IS_NOTFINAL,
		},
		{
			MD5Hash: "c5a4f2d38d78b251d8557cb2d93e40ee",
			Name:    "Freedoom: Phase 2",
			Version: "0.10",
			Game:    GAME_FREEDOOM,
			Status:  IS_NOTFINAL,
		},
		{
			MD5Hash: "f37b8b70e1394289a7ec404f67cdec1a",
			Name:    "FreeDM",
			Version: "0.10",
			Game:    GAME_FREEDOOM,
			Status:  IS_NOTFINAL,
		},

		// 0.9
		{
			MD5Hash: "aca90cf5ac36e996edc58bd0329b979a",
			Name:    "Freedoom: Phase 1",
			Version: "0.9",
			Game:    GAME_FREEDOOM,
			Status:  IS_NOTFINAL,
		},
		{
			MD5Hash: "8fa57dbc7687f84528eba39dde3a20e0",
			Name:    "Freedoom: Phase 2",
			Version: "0.9",
			Game:    GAME_FREEDOOM,
			Status:  IS_NOTFINAL,
		},
		{
			MD5Hash: "cbb27c5f3c2c44d34843cf63daa627f6",
			Name:    "FreeDM",
			Version: "0.9",
			Game:    GAME_FREEDOOM,
			Status:  IS_NOTFINAL,
		},

		// 0.8
		{
			MD5Hash: "30095b256dd3a1566bbc30286f72bc47",
			Name:    "Ultimate FreeDoom (ex-Freedoom: Phase 1)",
			Version: "0.8",
			Game:    GAME_FREEDOOM,
			Status:  IS_NOTFINAL,
		},
		{
			MD5Hash: "e3668912fc37c479b2840516c887018b",
			Name:    "Freedoom (ex-Freedoom: Phase 1)",
			Version: "0.8",
			Game:    GAME_FREEDOOM,
			Status:  IS_NOTFINAL,
		},
		{
			MD5Hash: "05859098bf191899903ef343afba369d",
			Name:    "FreeDM",
			Version: "0.8",
			Game:    GAME_FREEDOOM,
			Status:  IS_NOTFINAL,
		},
	}
}

func Populate_Misc() {

	// Master Levels for DOOM II
	IWADInfo_MasterLevels = []WadInfo{
		// Standard PWADs
		{
			MD5Hash: "cb03fd0cd84b10579c2b2b313199d4c1",
			Name:    "Attack (Master Levels for DOOM II)",
			Status:  IS_FINAL,
		},
		{
			MD5Hash: "a421ca18cea00a2696162f8d2a2beeca",
			Name:    "Black Tower (Master Levels for DOOM II)",
			Status:  IS_FINAL,
		},
		{
			MD5Hash: "18eb4ffb3094ddb690e62211dc6169a1",
			Name:    "Bloodsea Keep (Master Levels for DOOM II)",
			Status:  IS_FINAL,
		},
		{
			MD5Hash: "33493942592d764e7787fb0ad7d03044",
			Name:    "Canyon (Master Levels for DOOM II)",
			Status:  IS_FINAL,
		},
		{
			MD5Hash: "e7c273033376824edf95e1328261e7de",
			Name:    "The Catwalk (Master Levels for DOOM II)",
			Status:  IS_FINAL,
		},
		{
			MD5Hash: "77c179948df47a7a613bd1181c959892",
			Name:    "The Combine (Master Levels for DOOM II)",
			Status:  IS_FINAL,
		},
		{
			MD5Hash: "cbf714b499ebdef2682990eaf93fdb5f",
			Name:    "The Fistula (Master Levels for DOOM II)",
			Status:  IS_FINAL,
		},
		{
			MD5Hash: "f000701a3ed1f49249ee08550c03dfa5",
			Name:    "The Garrison (Master Levels for DOOM II)",
			Status:  IS_FINAL,
		},
		{
			MD5Hash: "a1efff02df6d873762ebac6b12358bbc",
			Name:    "Geryon: 6th Canto of Inferno (Master Levels for DOOM II)",
			Status:  IS_FINAL,
		},
		{
			MD5Hash: "787fa80fe9665c322f853b74838e77cc",
			Name:    "Titan Manor (Master Levels for DOOM II)",
			Status:  IS_FINAL,
		},
		{
			MD5Hash: "b4eaf844b135cc2a0058c6e0149b4408",
			Name:    "Mephisto’s Maosoleum (Master Levels for DOOM II)",
			Status:  IS_FINAL,
		},
		{
			MD5Hash: "aea597159dee96bcc58f3f9e3e586182",
			Name:    "Minos’ Judgement: 4th Canto of Inferno (Master Levels for DOOM II)",
			Status:  IS_FINAL,
		},
		{
			MD5Hash: "46f58580e7792f486c747cf1117c4ca1",
			Name:    "Nessus: 5th Canto of Inferno (Master Levels for DOOM II)",
			Status:  IS_FINAL,
		},
		{
			MD5Hash: "d560abb6d5719d46ebb47b27d7813a4b",
			Name:    "Paradox (Master Levels for DOOM II)",
			Status:  IS_FINAL,
		},
		{
			MD5Hash: "b572d518d564c7d7b6b259a726538cbb",
			Name:    "Subspace (Master Levels for DOOM II)",
			Status:  IS_FINAL,
		},
		{
			MD5Hash: "bb417f07804373415a6ed8e533242c3c",
			Name:    "Subterra (Master Levels for DOOM II)",
			Status:  IS_FINAL,
		},
		{
			MD5Hash: "65b4abcb74e7a386d5c024cf250d6336",
			Name:    "“The Express Elevator to Hell” and “Bad Dream” (Master Levels for DOOM II)",
			Status:  IS_FINAL,
		},
		{
			MD5Hash: "8474f6d663f04630de05ecac36b574d1",
			Name:    "Trapped on Titan (Master Levels for DOOM II)",
			Status:  IS_FINAL,
		},
		{
			MD5Hash: "a49dccebb5f32307246b7f32da121cf7",
			Name:    "Vesperas: 7th Canto of Inferno (Master Levels for DOOM II)",
			Status:  IS_FINAL,
		},
		{
			MD5Hash: "3c0874f2df3c06a002ee2a18aba0f0e8",
			Name:    "Virgil’s Lead: 3rd Canto of Inferno (Master Levels for DOOM II)",
			Status:  IS_FINAL,
		},

		// Aggregate WAD from the Xbox 360 and PlayStation 3 BFG Edition
		{
			MD5Hash:    "84cb8640f599c4a17c8eb526f90d2b7a",
			Name:       "Master Levels for DOOM II - Xbox 360/PlayStation 3 BFG Edition",
			Flags:      FL_RERELEASE | FL_HIDDEN,
			Additional: "File is incompatible with demo files or multiplayer sourceports",
		},
		{
			MD5Hash:    "ab3ce78e085e50a61f6dff46aabbfaeb",
			Name:       "Master Levels for Doom II (Doom + Doom II release)",
			Version:    "Update 1",
			Game:       GAME_KEXDOOM2024,
			Flags:      FL_RERELEASE,
			Additional: "File is incompatible with demo files or multiplayer sourceports, all masterlevel files are available in <installfolder>/base/master/wads",
		},
	}

	// DOOM/UDOOM population
	IWADInfo_Misc = []WadInfo{
		{
			MD5Hash:      "967d5ae23daf45196212ae1b605da3b0",
			Version:      "No Rest for the Living",
			Status:       IS_FINAL,
			PWADRequires: "DOOM II v1.9",
			Additional:   "You will need a limit-removing source port to be able to run this.",
		},
		{
			MD5Hash:    "23422eb42833ac7b0dd59c0c7ae18a6f",
			Name:       "No Rest For The Living (Doom + Doom II release)",
			Version:    "Update 1",
			Game:       GAME_KEXDOOM2024,
			Flags:      FL_RERELEASE,
			Additional: "File is not identical to the original release of No Rest for the Living and won't be compatible with multiplayer sourceports.",
		},
		// HACX
		{
			MD5Hash: "65ed74d522bdf6649c2831b13b9e02b4",
			Name:    "HacX",
			Version: "1.2",
			Game:    GAME_HACX,
			Status:  IS_FINAL,
		},
		{
			MD5Hash: "b7fd2f43f3382cf012dc6b097a3cb182",
			Name:    "HacX",
			Version: "1.1",
			Game:    GAME_HACX,
			Status:  IS_NOTFINAL,
		},
		{
			MD5Hash: "1511a7032ebc834a3884cf390d7f186e",
			Name:    "HacX",
			Version: "1.0",
			Game:    GAME_HACX,
			Status:  IS_NOTFINAL,
		},

		// Chex Quest
		{
			MD5Hash:    "25485721882b050afa96a56e5758dd52",
			Name:       "Chex Quest",
			Game:       GAME_IWAD,
			Status:     IS_FINAL,
			Additional: "May require the DEHacked file if playing on a source port, available at https://www.doomworld.com/idgames/utils/exe_edit/patches/chexdeh",
		},

		{
			MD5Hash:      "fdc4ffa57e1983e30912c006284a3e01",
			Name:         "Chex Quest 2",
			Status:       IS_FINAL,
			PWADRequires: "Chex Quest",
			Additional:   "May require the DEHacked file if playing on a source port, available at https://www.doomworld.com/idgames/utils/exe_edit/patches/chexdeh",
		},

		{
			MD5Hash: "bce163d06521f9d15f9686786e64df13",
			Name:    "Chex Quest 3",
			Version: "1.4",
			Game:    GAME_CHEX_QUEST_3,
			Status:  IS_FINAL,
		},
		{
			MD5Hash:    "cb001c34e424687191f299cc1dff4d68",
			Name:       "Chex Quest 3 (ModDB)",
			Version:    "Unidentified version",
			Game:       GAME_CHEX_QUEST_3,
			Status:     IS_NOTFINAL,
			Additional: "This version of Chex Quest 3 has been released on ModDB, and seems outdated.",
		},
		{
			MD5Hash: "59c985995db55cd2623c1893550d82b3",
			Name:    "Chex Quest 3",
			Version: "1.0",
			Game:    GAME_CHEX_QUEST_3,
			Status:  IS_NOTFINAL,
		},

		// SIGIL 1 & 2 by John Romero
		{
			MD5Hash:      "f53ffc4fb89e966839bb8d20c632819a",
			Name:         "SIGIL",
			Version:      "1.0",
			Game:         GAME_SIGIL,
			PWADRequires: "The Ultimate Doom v1.9",
			Status:       IS_NOTFINAL,
		},
		{
			MD5Hash:      "a775262ca0e423468196803b71a57a43",
			Name:         "SIGIL (Compatibility WAD)",
			Version:      "1.0",
			Game:         GAME_SIGIL,
			PWADRequires: "The Ultimate Doom v1.9",
			Status:       IS_NOTFINAL,
		},

		{
			MD5Hash:      "1fe9daa0e837c7452eb2f91aac2cc983",
			Name:         "SIGIL",
			Version:      "1.1",
			Game:         GAME_SIGIL,
			PWADRequires: "The Ultimate Doom v1.9",
			Status:       IS_NOTFINAL,
		},
		{
			MD5Hash:      "c04912beab6aa82c114a19c976ec8c0d",
			Name:         "SIGIL (Compatibility WAD)",
			Version:      "1.1",
			Game:         GAME_SIGIL,
			PWADRequires: "The Ultimate Doom v1.9",
			Status:       IS_NOTFINAL,
		},

		{
			MD5Hash:      "427ca995600970abcd2efcc684a64c88",
			Name:         "SIGIL",
			Version:      "1.2",
			Game:         GAME_SIGIL,
			PWADRequires: "The Ultimate Doom v1.9",
			Status:       IS_NOTFINAL,
		},
		{
			MD5Hash:      "9285e9cc2dbd87d238baab37d700c644",
			Name:         "SIGIL (Compatibility WAD)",
			Version:      "1.2",
			Game:         GAME_SIGIL,
			PWADRequires: "The Ultimate Doom v1.9",
			Status:       IS_NOTFINAL,
		},

		{
			MD5Hash:      "743d6323cb2b9be24c258ff0fc350883",
			Name:         "SIGIL",
			Version:      "1.21",
			Game:         GAME_SIGIL,
			PWADRequires: "The Ultimate Doom v1.9",
			Status:       IS_FINAL,
		},
		{
			MD5Hash:      "573f3f178c76709f512089ed15484391",
			Name:         "SIGIL (Compatibility WAD)",
			Version:      "1.21",
			Game:         GAME_SIGIL,
			PWADRequires: "The Ultimate Doom v1.9",
			Status:       IS_FINAL,
		},
		{
			MD5Hash:      "b424dcf46ae55a496c34ac37cce32646",
			Name:         "SIGIL - BucketHead soundtrack",
			Game:         GAME_SIGIL,
			PWADRequires: "SIGIL & The Ultimate DOOM v1.9",
			Flags:        FL_HIDDEN,
		},
		{
			MD5Hash:      "343faa815928c58faa08939a4502d5d2",
			Name:         "SIGIL - BucketHead soundtrack (Compatibility WAD)",
			Game:         GAME_SIGIL,
			PWADRequires: "SIGIL & The Ultimate DOOM v1.9",
			Flags:        FL_HIDDEN,
		},
		{
			MD5Hash:    "08ee05388c137db5f5d7996e89425b95",
			Name:       "SIGIL (Doom + Doom II)",
			Version:    "Update 1",
			Game:       GAME_KEXDOOM2024,
			Flags:      FL_RERELEASE,
			Additional: "File is not identical to the original releases of SIGIL and won't be compatible with multiplayer sourceports.",
		},

		// SIGIL II
		{
			MD5Hash:      "d0442f5a75f2faef3405c09a0c3acc58",
			Name:         "SIGIL II",
			Version:      "1.0",
			Game:         GAME_SIGIL_2,
			PWADRequires: "The Ultimate Doom v1.9",
			Additional:   "You will need a limit-removing source port to be able to run this.",
		},
		{
			MD5Hash:    "953f65cf079d0ba9a25be2c407da7ec1",
			Name:       "SIGIL II (Doom + Doom II)",
			Version:    "Update 3",
			Game:       GAME_KEXDOOM2024,
			Additional: "File is not identical to the original release of SIGIL II and won't be compatible with multiplayer sourceports.",
		},

		// DOOM 64
		{
			MD5Hash:    "e16e17f59afe7b3297f53ebe7e9de815",
			Name:       "DOOM 64 (2020 Re-Release)",
			Version:    "1.0",
			Flags:      FL_RERELEASE | FL_HIDDEN,
			Additional: "This WAD can only be used with the NightDive Studios port of DOOM 64!",
		},

		// REKKR
		{
			MD5Hash:      "d666daa88cca9ff59816ab2d32aeb4c6",
			Name:         "REKKR (PWAD Version)",
			Version:      "1.16",
			Game:         GAME_REKKR,
			PWADRequires: "DOOM.WAD or Freedoom - Phase 1",
			Status:       IS_FINAL,
			Additional:   "May require the DEHacked file if playing on a source port based on Chocolate Doom.",
		},
		{
			MD5Hash:    "b6f4bb3a80f096b6045cfaeb57d4cf29",
			Name:       "REKKR (Standalone version)",
			Version:    "1.16",
			Game:       GAME_REKKR,
			Status:     IS_FINAL,
			Additional: "Requires the .deh file to properly play it. However, the batch included provides everything to play it immediately.",
		},

		// Additionnal Doom + Doom II wads
		{
			MD5Hash:    "2e76d93d52ef64fb9db3cee2437c686b",
			Name:       "extras.wad (Doom + Doom II release)",
			Version:    "Update 1",
			Game:       GAME_KEXDOOM2024,
			Status:     IS_FINAL,
			Flags:      FL_RERELEASE | FL_HIDDEN,
			Additional: "Used for shared graphics and contains Andrew Hulshult's IDKFA soundtrack in OGG format.",
		},
		{
			MD5Hash:    "95f21547be5e0bff38d412017440f656",
			Name:       "id1.wad (Doom + Doom II release)",
			Version:    "Update 1",
			Game:       GAME_KEXDOOM2024,
			Status:     IS_NOTFINAL,
			Flags:      FL_RERELEASE,
			Additional: "Used for 'Legacy of Rust' addon.",
		},
		{
			MD5Hash:    "713c5a3c1734b1d55b2813a3dd0136d9",
			Name:       "id1.wad (Doom + Doom II release)",
			Version:    "Update 2+",
			Game:       GAME_KEXDOOM2024,
			Status:     IS_FINAL,
			Flags:      FL_RERELEASE,
			Additional: "Used for 'Legacy of Rust' addon.",
		},
		{
			MD5Hash:    "187bfe543f8328b379e46957976e800d",
			Name:       "id1-tex.wad (Doom + Doom II)",
			Game:       GAME_KEXDOOM2024,
			Flags:      FL_RERELEASE | FL_HIDDEN,
			Additional: "Used for the 'Legacy of Rust' addon.",
		},
		{
			MD5Hash:    "f8fbab472230bfa090d6a9234d65fae6",
			Name:       "id1-res.wad (Doom + Doom II)",
			Game:       GAME_KEXDOOM2024,
			Flags:      FL_RERELEASE | FL_HIDDEN,
			Additional: "Used for the 'Legacy of Rust' addon.",
		},
		{
			MD5Hash:    "85d25c8c3d06a05a1283ae4afe749c9f",
			Name:       "id1-weap.wad (Doom + Doom II)",
			Game:       GAME_KEXDOOM2024,
			Status:     IS_FINAL,
			Flags:      FL_RERELEASE,
			Additional: "Used for the 'Legacy of Rust' addon.",
		},
		{
			MD5Hash:    "436c83dd83a47f8dd251ba15108e9459",
			Name:       "id1-mus.wad (Doom + Doom II)",
			Game:       GAME_KEXDOOM2024,
			Flags:      FL_RERELEASE | FL_HIDDEN,
			Additional: "Used for the 'Legacy of Rust' addon.",
		},
		{
			MD5Hash:    "4f0651accebc007b853943ac12aa95b8",
			Name:       "id24res.wad (Doom + Doom II)",
			Game:       GAME_KEXDOOM2024,
			Status:     IS_FINAL,
			Flags:      FL_RERELEASE,
			Additional: "Resources from *that* proprietary ID24 standard.",
		},
		{
			MD5Hash:    "5670fd8fe8eb6910ec28f9e27969d84f",
			Name:       "iddm1.wad (Doom + Doom II)",
			Version:    "Update 1",
			Game:       GAME_KEXDOOM2024,
			Status:     IS_NOTFINAL,
			Flags:      FL_RERELEASE,
			Additional: "Used as the exclusive Deathmatch WAD for Doom + Doom II.",
		},
		{
			MD5Hash:    "cb92010b8ec05f8924ac966a8ed95b74",
			Name:       "iddm1.wad (Doom + Doom II)",
			Version:    "Update 2+",
			Game:       GAME_KEXDOOM2024,
			Status:     IS_FINAL,
			Flags:      FL_RERELEASE,
			Additional: "Used as the exclusive Deathmatch WAD for Doom + Doom II.",
		},
	}
}
