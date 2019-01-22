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

	if bFound {
		fmt.Println("MD5:", IWAD.MD5Hash)
		ansi.Println("Version:", yellow(IWAD.Version))

		// Check only once if we need to warn the user, otherwise it'll mess up the final results
		if !bNeedsPatching {
			bNeedsPatching = !IWAD.IsFinal
		}
		ansi.Println("Needs patching ?", YesorNo(!IWAD.IsFinal))
		fmt.Println("")
	} else {

		// At this point, we should dissect the first bytes of the WAD to make sure it's a PWAD.
		// Then, check against the known Addons/Extensions (Hexen:DotDC / NervE)
		// Yet, we only assume this WAD is a PWAD or invalid.
		color.Cyan("%s seems unknown. Make sure it's not a modified file or not a PWAD file. (hash:%s)", filename, hash)
		fmt.Println("")
	}
}
