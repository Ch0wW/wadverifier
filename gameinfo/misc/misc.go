package misc

import (
	"wadverifier/wad"
	"wadverifier/wad/flags"
)

func BuildMiscFiles() []wad.Entry {
	return []wad.Entry{
		{
			MD5Hash:    "e16e17f59afe7b3297f53ebe7e9de815",
			Name:       "DOOM 64 (2020 Re-Release)",
			Version:    "1.0",
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
