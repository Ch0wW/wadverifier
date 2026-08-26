package misc

import (
	"wadverifier/wad"
	"wadverifier/wad/flags"
	"wadverifier/wad/status"
)

func BuildMiscFiles() []wad.Entry {
	return []wad.Entry{
		{
			MD5Hash:    "0aaba212339c72250f8a53a0a2b6189e",
			Name:       "DOOM 64 (2020 Re-Release)",
			Version:    "1.1",
			Status:     status.FINAL,
			Flags:      flags.RERELEASE | flags.HIDDEN,
			Additional: "This WAD can only be used with the NightDive Studios port of DOOM 64!",
		},
		{
			MD5Hash:    "e16e17f59afe7b3297f53ebe7e9de815",
			Name:       "DOOM 64 (2020 Re-Release)",
			Version:    "1.0",
			Status:     status.NOTFINAL,
			Flags:      flags.RERELEASE | flags.HIDDEN,
			Additional: "This WAD can only be used with the NightDive Studios port of DOOM 64!",
		},
	}
}

func Populate() []wad.Entry {
	list := BuildMiscFiles()
	list = append(list, BuildChexQuestInfo()...)
	list = append(list, BuildHacxInfo()...)
	list = append(list, BuildMasterLevelsInfo()...)
	list = append(list, BuildNerveInfo()...)
	list = append(list, BuildREKKRInfo()...)
	list = append(list, BuildSIGILInfo()...)
	list = append(list, BuildSIGIL_2_Info()...)
	list = append(list, BuildKexDoomExtraInfo()...)

	return list
}
