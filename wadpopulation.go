package main

import (
	"fmt"

	"github.com/fatih/color"
	ansi "github.com/k0kubun/go-ansi"
)

//
//=======================
// PopulateIWADInfos
// Populate the IWAD lists with their MD5 and versions
//=======================
//
func PopulateIWADInfos() {
	Populate_Doom()
	Populate_DoomII()
	Populate_FinalDOOM()
	Populate_HereticHexen()
	Populate_Strife()
	Populate_FreeDoom()
}

func Populate_Doom() {

	// DOOM/UDOOM population
	IWADInfo_Doom = []WadInfo{
		WadInfo{
			MD5Hash: "90facab21eede7981be10790e3f82da2",
			Version: "DOOM 1.0 Shareware",
			IsFinal: false,
		},
		WadInfo{
			MD5Hash: "52cbc8882f445573ce421fa5453513c1",
			Version: "DOOM 1.1 Shareware",
			IsFinal: false,
		},
		WadInfo{
			MD5Hash: "981b03e6d1dc033301aa3095acc437ce",
			Version: "DOOM 1.1 Registered",
			IsFinal: false,
		},
		WadInfo{
			MD5Hash: "30aa5beb9e5ebfbbe1e1765561c08f38",
			Version: "DOOM 1.2 Shareware",
			IsFinal: false,
		},
		WadInfo{
			MD5Hash: "792fd1fea023d61210857089a7c1e351",
			Version: "DOOM 1.2 Registered",
			IsFinal: false,
		},
		WadInfo{
			MD5Hash: "a21ae40c388cb6f2c3cc1b95589ee693",
			Version: "DOOM 1.4 Shareware Beta",
			IsFinal: false,
		},
		WadInfo{
			MD5Hash: "e280233d533dcc28c1acd6ccdc7742d4",
			Version: "DOOM 1.5 Shareware Beta",
			IsFinal: false,
		},
		WadInfo{
			MD5Hash: "762fd6d4b960d4b759730f01387a50a1",
			Version: "DOOM 1.6 Shareware Beta",
			IsFinal: false,
		},
		WadInfo{
			MD5Hash: "c428ea394dc52835f2580d5bfd50d76f",
			Version: "DOOM 1.666 Shareware",
			IsFinal: false,
		},
		WadInfo{
			MD5Hash: "54978d12de87f162b9bcc011676cb3c0",
			Version: "DOOM 1.666 Registered",
			IsFinal: false,
		},
		WadInfo{
			MD5Hash: "5f4eb849b1af12887dec04a2a12e5e62",
			Version: "DOOM 1.8 Shareware",
			IsFinal: false,
		},
		WadInfo{
			MD5Hash: "11e1cd216801ea2657723abc86ecb01f",
			Version: "DOOM 1.8 Registered",
			IsFinal: false,
		},
		WadInfo{
			MD5Hash: "f0cefca49926d00903cf57551d901abe",
			Version: "DOOM 1.9 Shareware",
			IsFinal: true,
		},
		WadInfo{
			MD5Hash: "1cd63c5ddff1bf8ce844237f580e9cf3",
			Version: "DOOM 1.9 Registered",
			IsFinal: false,
		},
		WadInfo{
			MD5Hash: "c4fe9fd920207691a9f493668e0a2083",
			Version: "The Ultimate DOOM (v1.9)",
			IsFinal: true,
		},
		WadInfo{
			MD5Hash: "0c8758f102ccafe26a3040bee8ba5021",
			Version: "The Ultimate DOOM (1.9) - Xbox Version",
			IsFinal: false,
		},
		WadInfo{
			MD5Hash: "72286ddc680d47b9138053dd944b2a3d",
			Version: "The Ultimate DOOM (1.9) - XBLA Version",
			IsFinal: false,
		},
		WadInfo{
			MD5Hash: "fb35c4a5a9fd49ec29ab6e900572c524",
			Version: "The Ultimate DOOM (1.9) - BFG Edition",
			IsFinal: false,
		},
	}
}

func Populate_DoomII() {

	// DOOM/UDOOM population
	IWADInfo_Doom2 = []WadInfo{
		WadInfo{
			MD5Hash: "30e3c2d0350b67bfbf47271970b74b2f",
			Version: "DOOM II 1.666",
			IsFinal: false,
		},
		WadInfo{
			MD5Hash: "d9153ced9fd5b898b36cc5844e35b520",
			Version: "DOOM II 1.666 (German Edition)",
			IsFinal: false,
		},
		WadInfo{
			MD5Hash: "ea74a47a791fdef2e9f2ea8b8a9da13b",
			Version: "DOOM II 1.7",
			IsFinal: false,
		},
		WadInfo{
			MD5Hash: "d7a07e5d3f4625074312bc299d7ed33f",
			Version: "DOOM II 1.7a",
			IsFinal: false,
		},
		WadInfo{
			MD5Hash: "c236745bb01d89bbb866c8fed81b6f8c",
			Version: "DOOM II 1.8",
			IsFinal: false,
		},
		WadInfo{
			MD5Hash: "25e1459ca71d321525f84628f45ca8cd",
			Version: "DOOM II 1.9",
			IsFinal: true,
		},
		WadInfo{
			MD5Hash: "a793ebcdd790afad4a1f39cc39a893bd",
			Version: "DOOM II 1.9 - Xbox Version",
			IsFinal: false,
		},
		WadInfo{
			MD5Hash: "43c2df32dc6c740cb11f34dc5ab693fa",
			Version: "DOOM II 1.9 - XBLA Version",
			IsFinal: false,
		},
		WadInfo{
			MD5Hash: "c3bea40570c23e511a7ed3ebcd9865f7",
			Version: "DOOM II 1.9 - BFG Edition",
			IsFinal: false,
		},
	}
}

func Populate_FinalDOOM() {
	// DOOM/UDOOM population
	IWADInfo_FinalDoom = []WadInfo{
		WadInfo{
			MD5Hash: "75c8cf89566741fa9d22447604053bd7",
			Version: "Final DOOM: The Plutonia Experiment 1.9",
			IsFinal: true,
		},
		WadInfo{
			MD5Hash: "4e158d9953c79ccf97bd0663244cc6b6",
			Version: "Final DOOM: Evilution 1.9",
			IsFinal: true,
		},
	}
}

func Populate_HereticHexen() {
	// Heretic population
	IWADInfo_Heretic = []WadInfo{
		WadInfo{
			MD5Hash: "3117e399cdb4298eaa3941625f4b2923",
			Version: "Heretic 1.0",
			IsFinal: false,
		},
		WadInfo{
			MD5Hash: "66d686b1ed6d35ff103f15dbd30e0341",
			Version: "Heretic: Shadow of the Serpent Riders (1.3)",
			IsFinal: true,
		},
	}

	IWADInfo_Hexen = []WadInfo{
		WadInfo{
			MD5Hash: "b2543a03521365261d0a0f74d5dd90f0",
			Version: "Hexen 1.0",
			IsFinal: false,
		},
		WadInfo{
			MD5Hash: "abb033caf81e26f12a2103e1fa25453f",
			Version: "Hexen 1.1",
			IsFinal: true,
		},
	}
}

func Populate_Strife() {

	// DOOM/UDOOM population
	IWADInfo_Doom2 = []WadInfo{
		WadInfo{
			MD5Hash: "8f2d3a6a289f5d2f2f9c1eec02b47299",
			Version: "Strife 1.0 Registered",
			IsFinal: false,
		},
		WadInfo{
			MD5Hash: "bb545b9c4eca0ff92c14d466b3294023",
			Version: "Strife 1.1 Shareware",
			IsFinal: false,
		},
		WadInfo{
			MD5Hash:     "2fed2031a5b03892106e0f117f17901f",
			Version:     "Strife 1.2 - 1.31 Registered",
			IsFinal:     true,
			Additionnal: "This IWAD is the latest version of Strife. However, the latest updates of Strife modify the executable. Please make sure it is the latest version.",
		},
	}
}

func Populate_FreeDoom() {
	IWADInfo_FreeDoom = []WadInfo{

		//---------------
		// First, the latest IWADs
		WadInfo{
			MD5Hash:    "ea471a3d38fcee0fb3a69bcd3221e335",
			Version:    "Freedoom: Phase 1 0.11.3",
			IsFinal:    true,
			IsFreedoom: true,
		},
		WadInfo{
			MD5Hash:    "984f99af08f085e38070f51095ab7c31",
			Version:    "Freedoom: Phase 2 0.11.3",
			IsFreedoom: true,
			IsFinal:    true,
		},
		WadInfo{
			MD5Hash:    "87ee2494d921633420ce9bdb418127c4",
			Version:    "FreeDM 0.11.3",
			IsFreedoom: true,
			IsFinal:    true,
		},

		//----------------
		// Now, list the outdated ones.

		// 0.11.2
		WadInfo{
			MD5Hash:    "9352b09ae878dc52c6c18aa38acda6eb",
			Version:    "FreeDM 0.11.2",
			IsFinal:    false,
			IsFreedoom: true,
		},
		WadInfo{
			MD5Hash:    "6d00c49520be26f08a6bd001814a32ab",
			Version:    "Freedoom: Phase 1 0.11.2",
			IsFinal:    false,
			IsFreedoom: true,
		},
		WadInfo{
			MD5Hash:    "90832a872b5bb0aca4ca0b20419aad5d",
			Version:    "Freedoom: Phase 2 0.11.2",
			IsFinal:    false,
			IsFreedoom: true,
		},

		// 0.11.1
		WadInfo{
			MD5Hash:    "77ba9c0f75c32e4a729490688bb99241",
			Version:    "FreeDM 0.11.1",
			IsFinal:    false,
			IsFreedoom: true,
		},
		WadInfo{
			MD5Hash:    "35312e99d2473297aabe0602700bee8a",
			Version:    "Freedoom: Phase 1 0.11.1",
			IsFinal:    false,
			IsFreedoom: true,
		},
		WadInfo{
			MD5Hash:    "ec5b38b30ba2b70e278205776af3fbb5",
			Version:    "Freedoom: Phase 2 0.11.1",
			IsFinal:    false,
			IsFreedoom: true,
		},

		// 0.11
		WadInfo{
			MD5Hash:    "d76d3973c075b069ecb4e16dc9eacbb4",
			Version:    "FreeDM 0.11",
			IsFinal:    false,
			IsFreedoom: true,
		},
		WadInfo{
			MD5Hash:    "21a4707fc25d29edf4b098bd400c5c42",
			Version:    "Freedoom: Phase 1 0.11",
			IsFinal:    false,
			IsFreedoom: true,
		},
		WadInfo{
			MD5Hash:    "b1018017c61b06e33c11102d8bafaad0",
			Version:    "Freedoom: Phase 2 0.11",
			IsFinal:    false,
			IsFreedoom: true,
		},

		// 0.10.1
		WadInfo{
			MD5Hash:    "bd4f359f1963e388beda014c5548b420",
			Version:    "FreeDM 0.10.1",
			IsFinal:    false,
			IsFreedoom: true,
		},
		WadInfo{
			MD5Hash:    "91de79621a393a08c39a9ab2c034b766",
			Version:    "Freedoom: Phase 1 0.10.1",
			IsFinal:    false,
			IsFreedoom: true,
		},
		WadInfo{
			MD5Hash:    "dd9c9e73f5f50d3778c85573cd08d9a4",
			Version:    "Freedoom: Phase 2 0.10.1",
			IsFinal:    false,
			IsFreedoom: true,
		},

		// 0.10.0
		WadInfo{
			MD5Hash:    "9b8d72b59fd93b2b3e116149baa1b142",
			Version:    "Freedoom: Phase 1 0.10",
			IsFinal:    false,
			IsFreedoom: true,
		},
		WadInfo{
			MD5Hash:    "c5a4f2d38d78b251d8557cb2d93e40ee",
			Version:    "Freedoom: Phase 2 0.10",
			IsFinal:    false,
			IsFreedoom: true,
		},
		WadInfo{
			MD5Hash:    "f37b8b70e1394289a7ec404f67cdec1a",
			Version:    "FreeDM 0.10",
			IsFinal:    false,
			IsFreedoom: true,
		},

		// 0.9
		WadInfo{
			MD5Hash:    "aca90cf5ac36e996edc58bd0329b979a",
			Version:    "Freedoom: Phase 1 0.9",
			IsFinal:    false,
			IsFreedoom: true,
		},
		WadInfo{
			MD5Hash:    "8fa57dbc7687f84528eba39dde3a20e0",
			Version:    "Freedoom: Phase 2 0.9",
			IsFinal:    false,
			IsFreedoom: true,
		},
		WadInfo{
			MD5Hash:    "cbb27c5f3c2c44d34843cf63daa627f6",
			Version:    "FreeDM 0.9",
			IsFinal:    false,
			IsFreedoom: true,
		},

		// 0.8
		WadInfo{
			MD5Hash:    "30095b256dd3a1566bbc30286f72bc47",
			Version:    "Ultimate Freedoom 0.8",
			IsFinal:    false,
			IsFreedoom: true,
		},
		WadInfo{
			MD5Hash:    "e3668912fc37c479b2840516c887018b",
			Version:    "Freedoom 0.8",
			IsFinal:    false,
			IsFreedoom: true,
		},
		WadInfo{
			MD5Hash:    "05859098bf191899903ef343afba369d",
			Version:    "FreeDM 0.8",
			IsFinal:    false,
			IsFreedoom: true,
		},
	}
}

func CompIWADData(data []WadInfo, hash string) (WadInfo, bool) {
	for i := range data {
		if hash == data[i].MD5Hash {
			return data[i], true
		}
	}
	return WadInfo{}, false
}

func YesorNo(b bool) string {
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	if b == true {
		return red("Yes")
	} else {
		return green("No")
	}
}

func CheckIWAD(filename string, hash string) {

	yellow := color.New(color.FgYellow).SprintFunc()

	bFound := false
	fmt.Println("Checking file :", filename)

	// Check DOOM
	IWAD, bFound := CompIWADData(IWADInfo_Doom, hash)

	// Check DOOM2
	if bFound == false {
		IWAD, bFound = CompIWADData(IWADInfo_Doom2, hash)
	}

	// Check TNT/Plutonia
	if bFound == false {
		IWAD, bFound = CompIWADData(IWADInfo_FinalDoom, hash)
	}

	// Check Heretic
	if bFound == false {
		IWAD, bFound = CompIWADData(IWADInfo_Heretic, hash)
	}

	// Check Hexen
	if bFound == false {
		IWAD, bFound = CompIWADData(IWADInfo_Hexen, hash)
	}

	// Check Strife
	if bFound == false {
		IWAD, bFound = CompIWADData(IWADInfo_Strife, hash)
	}

	// Check FreeDoom/FreeDM
	if bFound == false {
		IWAD, bFound = CompIWADData(IWADInfo_FreeDoom, hash)
	}

	if bFound {
		fmt.Println("MD5:", IWAD.MD5Hash)
		ansi.Println("Version:", yellow(IWAD.Version))

		// Check only once if we need to warn the user, otherwise it'll mess up the final results
		if !bNeedsPatching {
			bNeedsPatching = !IWAD.IsFinal
		}

		if !bUpgradeIWAD && !IWAD.IsFreedoom {
			bUpgradeIWAD = !IWAD.IsFinal
		} else if !bUpgradeFreeDoom && IWAD.IsFreedoom {
			bUpgradeFreeDoom = !IWAD.IsFinal
		}

		// Add an error count if it's not the final version of a wad.
		if !IWAD.IsFinal {
			iErrors = iErrors + 1
		}

		ansi.Println("Needs patching ?", YesorNo(!IWAD.IsFinal))

		// If the IWAD has an additionnal message, please write it so.
		if IWAD.Additionnal != "" {
			color.Cyan(IWAD.Additionnal)
		}

		fmt.Println("")
	} else {

		// At this point, we should dissect the first bytes of the WAD to make sure it's a PWAD.
		// Then, check against the known Addons/Extensions (Hexen:DotDC / NervE)
		// Yet, we only assume this WAD is a PWAD or invalid.
		iErrors = iErrors + 1
		color.Cyan("%s seems unknown. Make sure it's not a modified file or not a PWAD file. (hash:%s)", filename, hash)
		fmt.Println("")
	}
}
