package misc

import (
	"wadverifier/wad"
	"wadverifier/wad/flags"
	"wadverifier/wad/games"
	"wadverifier/wad/status"
)

func BuildMasterLevelsInfo() []wad.Entry {
	return []wad.Entry{
		{
			MD5Hash:   "cb03fd0cd84b10579c2b2b313199d4c1",
			Name:      "Attack (Master Levels for DOOM II)",
			Patchinfo: games.NONE,
			Status:    status.FINAL,
		},
		{
			MD5Hash:   "a421ca18cea00a2696162f8d2a2beeca",
			Name:      "Black Tower (Master Levels for DOOM II)",
			Patchinfo: games.NONE,
			Status:    status.FINAL,
		},
		{
			MD5Hash:   "18eb4ffb3094ddb690e62211dc6169a1",
			Name:      "Bloodsea Keep (Master Levels for DOOM II)",
			Patchinfo: games.NONE,
			Status:    status.FINAL,
		},
		{
			MD5Hash:   "33493942592d764e7787fb0ad7d03044",
			Name:      "Canyon (Master Levels for DOOM II)",
			Patchinfo: games.NONE,
			Status:    status.FINAL,
		},
		{
			MD5Hash:   "e7c273033376824edf95e1328261e7de",
			Name:      "The Catwalk (Master Levels for DOOM II)",
			Patchinfo: games.NONE,
			Status:    status.FINAL,
		},
		{
			MD5Hash:   "77c179948df47a7a613bd1181c959892",
			Name:      "The Combine (Master Levels for DOOM II)",
			Patchinfo: games.NONE,
			Status:    status.FINAL,
		},
		{
			MD5Hash:   "cbf714b499ebdef2682990eaf93fdb5f",
			Name:      "The Fistula (Master Levels for DOOM II)",
			Patchinfo: games.NONE,
			Status:    status.FINAL,
		},
		{
			MD5Hash:   "f000701a3ed1f49249ee08550c03dfa5",
			Name:      "The Garrison (Master Levels for DOOM II)",
			Patchinfo: games.NONE,
			Status:    status.FINAL,
		},
		{
			MD5Hash:   "a1efff02df6d873762ebac6b12358bbc",
			Name:      "Geryon: 6th Canto of Inferno (Master Levels for DOOM II)",
			Patchinfo: games.NONE,
			Status:    status.FINAL,
		},
		{
			MD5Hash:   "787fa80fe9665c322f853b74838e77cc",
			Name:      "Titan Manor (Master Levels for DOOM II)",
			Patchinfo: games.NONE,
			Status:    status.FINAL,
		},
		{
			MD5Hash:   "b4eaf844b135cc2a0058c6e0149b4408",
			Name:      "Mephisto's Maosoleum (Master Levels for DOOM II)",
			Patchinfo: games.NONE,
			Status:    status.FINAL,
		},
		{
			MD5Hash:   "aea597159dee96bcc58f3f9e3e586182",
			Name:      "Minos' Judgement: 4th Canto of Inferno (Master Levels for DOOM II)",
			Patchinfo: games.NONE,
			Status:    status.FINAL,
		},
		{
			MD5Hash:   "46f58580e7792f486c747cf1117c4ca1",
			Name:      "Nessus: 5th Canto of Inferno (Master Levels for DOOM II)",
			Patchinfo: games.NONE,
			Status:    status.FINAL,
		},
		{
			MD5Hash:   "d560abb6d5719d46ebb47b27d7813a4b",
			Name:      "Paradox (Master Levels for DOOM II)",
			Patchinfo: games.NONE,
			Status:    status.FINAL,
		},
		{
			MD5Hash:   "b572d518d564c7d7b6b259a726538cbb",
			Name:      "Subspace (Master Levels for DOOM II)",
			Patchinfo: games.NONE,
			Status:    status.FINAL,
		},
		{
			MD5Hash:   "bb417f07804373415a6ed8e533242c3c",
			Name:      "Subterra (Master Levels for DOOM II)",
			Patchinfo: games.NONE,
			Status:    status.FINAL,
		},
		{
			MD5Hash:   "65b4abcb74e7a386d5c024cf250d6336",
			Name:      "“The Express Elevator to Hell” and “Bad Dream” (Master Levels for DOOM II)",
			Patchinfo: games.NONE,
			Status:    status.FINAL,
		},
		{
			MD5Hash:   "8474f6d663f04630de05ecac36b574d1",
			Name:      "Trapped on Titan (Master Levels for DOOM II)",
			Patchinfo: games.NONE,
			Status:    status.FINAL,
		},
		{
			MD5Hash:   "a49dccebb5f32307246b7f32da121cf7",
			Name:      "Vesperas: 7th Canto of Inferno (Master Levels for DOOM II)",
			Patchinfo: games.NONE,
			Status:    status.FINAL,
		},
		{
			MD5Hash:   "3c0874f2df3c06a002ee2a18aba0f0e8",
			Name:      "Virgil's Lead: 3rd Canto of Inferno (Master Levels for DOOM II)",
			Patchinfo: games.NONE,
			Status:    status.FINAL,
		},

		// Aggregate WAD from the Xbox 360 and PlayStation 3 BFG Edition
		{
			MD5Hash:    "84cb8640f599c4a17c8eb526f90d2b7a",
			Name:       "Master Levels for DOOM II - Xbox 360/PlayStation 3 BFG Edition",
			Patchinfo:  games.NONE,
			Flags:      flags.RERELEASE | flags.HIDDEN,
			Additional: "File is incompatible with demo files or multiplayer sourceports",
		},
		{
			MD5Hash:   "ab3ce78e085e50a61f6dff46aabbfaeb",
			Name:      "Master Levels for Doom II (Doom + Doom II release)",
			Version:   "Update 1",
			Patchinfo: games.KEX_DOOM2024,
			Flags:     flags.RERELEASE,
			//			Status:     status.NOTFINAL,
			Additional: "File is incompatible with demo files or multiplayer sourceports, all masterlevel files are available in <installfolder>/base/master/wads",
		},
	}
}
