package main

import (
	"bufio"
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/fatih/color"
	ansi "github.com/k0kubun/go-ansi"
)

const (
	m_version = "0.4"
	IWADbytes = 1145132873
	PWADbytes = 1145132880
)

// Just a quick function to require the user to press ENTER.
// Now, it only happens on Windows. (for the drag & drop feature)
func PressEnter() {

	if runtime.GOOS != "windows" {
		return
	}

	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func hash_file_md5(filePath string) (string, error) {
	//Initialize variable returnMD5String now in case an error has to be returned
	var returnMD5String string

	//Open the passed argument and check for any error
	file, err := os.Open(filePath)
	if err != nil {
		return returnMD5String, err
	}

	//Tell the program to call the following function when the current function returns
	defer file.Close()

	//Open a new hash interface to write to
	hash := md5.New()

	//Copy the file in the hash interface and check for any error
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}

	//Get the 16 bytes hash
	hashInBytes := hash.Sum(nil)[:16]

	//Convert the bytes to a string
	returnMD5String = hex.EncodeToString(hashInBytes)

	return returnMD5String, nil

}

// Adapted from https://github.com/XerTheSquirrel/go2it/blob/master/wad.go
// We only need the first Long.
func isWADvalid(filepath string) bool {

	// Opening it AGAIN
	file, err := os.Open(filepath)
	if err != nil {
		return false
	}

	// And AGAIN, don't forget to close it !
	defer file.Close()

	data := make([]byte, 4)
	_, err = io.ReadFull(file, data)
	if err != nil {
		return false
	}

	// Close the file as it is not needed anymore
	file.Close()

	// Need the magic number
	magic := binary.LittleEndian.Uint32(data[0:4])
	if magic != IWADbytes && magic != PWADbytes {
		return false
	}

	return true
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

	if b == false {
		return red("No")
	}

	return green("Yes")
}

func SetFlag(typegame *bool, IWAD WadInfo) {
	if !*typegame {
		*typegame = !IWAD.IsFinal
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

	// Check Master Levels
	if bFound == false {
		IWAD, bFound = CompIWADData(IWADInfo_MasterLevels, hash)
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

	// Check SVE
	if bFound == false {
		IWAD, bFound = CompIWADData(IWADInfo_SVE, hash)
	}

	// Check FreeDoom/FreeDM
	if bFound == false {
		IWAD, bFound = CompIWADData(IWADInfo_FreeDoom, hash)
	}

	// Check the other mods
	if bFound == false {
		IWAD, bFound = CompIWADData(IWADInfo_Misc, hash)
	}

	if bFound {
		fmt.Println("MD5:", IWAD.MD5Hash)
		ansi.Println("WAD File:", yellow(IWAD.Version))

		// Check only once if we need to warn the user, otherwise it'll mess up the final results
		if !bNeedsPatching {
			bNeedsPatching = !IWAD.IsFinal
		}

		// Now, flag our messages if our IWAD is older
		switch IWAD.Game {
		case GAME_IWAD:
			SetFlag(&bUpgradeIWAD, IWAD)
			break
		case GAME_FREEDOOM:
			SetFlag(&bUpgradeFreeDoom, IWAD)
			break
		case GAME_HACX:
			SetFlag(&bUpgradeHacX, IWAD)
			break
		case GAME_CHEX_QUEST_3:
			SetFlag(&bUpgradeCQ3, IWAD)
			break
		case GAME_STRIFE_VE:
			SetFlag(&bUpgradeSVE, IWAD)
			break
		case GAME_SIGIL:
			SetFlag(&bUpgradeSigil, IWAD)
			break
		}

		// Add an error count if it's not the final version of a wad.
		if !IWAD.IsFinal {
			iErrors = iErrors + 1
		}

		ansi.Println("Latest version?", YesorNo(IWAD.IsFinal))

		// If the IWAD has an additionnal message, please write it so.
		if IWAD.Additional != "" {
			color.Cyan(IWAD.Additional)
		}

		fmt.Println("")
	} else {

		// At this point, we should dissect the first bytes of the WAD to make sure it's a PWAD.
		// Then, check against the known Addons/Extensions (Hexen:DotDC / NervE)
		// Yet, we only assume this WAD is a PWAD or invalid.
		iErrors = iErrors + 1
		color.Cyan("%s seems unknown. Make sure it's not a modified file or not a PWAD file. (hash:%s )", filename, hash)
		color.Cyan("If this IWAD wasn't found, please write an issue on https://github.com/Ch0wW/iwadverifier/issues/")

		fmt.Println("")
	}
}

func main() {

	// Initialize IWAD/Addon lists
	PopulateIWADInfos()

	// Put the colors
	color.Output = ansi.NewAnsiStdout()

	color.Cyan("IWAD Verifier %s", m_version)
	color.Cyan("https://github.com/ch0ww/iwadverifier")
	color.Cyan("---------------------------------------")
	fmt.Println("")

	// Get the arguments
	args := os.Args[1:]

	if len(args) == 0 {
		color.Yellow("No argument specified.")
		color.Yellow("Usage: iwadverifier <wad.wad[ wad2.wad ...]>")
		return
	}

	for i := range args {

		// Check if the user omitted the extension.
		// If so, assume the file is a .wad
		// ToDo: Check later for .pk3 files !
		if filepath.Ext(strings.ToLower(args[i])) == "" {
			args[i] = fmt.Sprintf("%s.wad", args[i])
			fmt.Println(args[i])
		}

		// Try to check if file is a .wad
		if filepath.Ext(strings.ToLower(args[i])) != ".wad" {
			color.Yellow("%s is not a .wad file ! Ignoring...", args[i])
			iErrors = iErrors + 1
			fmt.Println("")
			continue
		}

		valid := isWADvalid(args[i])
		if !valid {
			color.Yellow("%s is not a valid WAD file !", args[i]) // Need to call SPA 1-800-388-PIR8 ?! Memories...
			iErrors = iErrors + 1
			continue
		}

		// Now, try to get the MD5 hash from the file
		hash, err := hash_file_md5(args[i])
		if err != nil {
			color.Yellow("Error getting the MD5 hash: Skipping... (%s)", err)
			iErrors = iErrors + 1
			fmt.Println("")
			continue
		}

		CheckIWAD(args[i], hash)
	}

	// This is ugly, but I have to find a way to make
	// If there's some patching needed, warn the user how to do it.
	if bNeedsPatching {
		color.Cyan("==================================================================================")
		color.Cyan("")

		if bUpgradeIWAD {
			color.Cyan("To patch your IWAD to the latest version, please use IWADPatcher 1.2 by Peter Vaskovics:")
			color.Cyan("• Windows binaries: http://downloads.zdaemon.org/iwadpatcher-1.2-bin.zip")
			color.Cyan("• Source code: http://downloads.zdaemon.org/iwadpatcher-1.2.zip")
			color.Cyan("")
		}
		if bUpgradeDoomSW {
			color.Cyan("Your Shareware version of Doom is outdated. Please get the latest version below :")
			color.Cyan("|-> https://www.doomworld.com/idgames/idstuff/doom/doom19s")
			color.Cyan("")
		}
		if bUpgradeFreeDoom {
			color.Cyan("Your version of FreeDOOM/FreeDM is outdated. Please get the latest one below :")
			color.Cyan("|-> https://github.com/freedoom/freedoom/releases")
			color.Cyan("")
		}
		if bUpgradeHacX {
			color.Cyan("Your version of HacX is outdated. Please get the latest one below :")
			color.Cyan("|-> http://www.drnostromo.com/hacx/page.php?content=download")
			color.Cyan("")
		}
		if bUpgradeCQ3 {
			color.Cyan("Your version of Chex Quest 3 is outdated. Please get the latest one below :")
			color.Cyan("|-> http://www.chucktropolis.com/gamers.htm")
			color.Cyan("")
		}
		if bUpgradeSVE {
			color.Cyan("Your version of Strife: Veteran Edition is outdated.")
			color.Cyan("• If you bought it on Steam, S:VE should be updated automatically.")
			color.Cyan("• If you bought it on GOG, you will need to redownload it (Latest version is 1.2) or to use GOG Galaxy")
			color.Cyan("")
		}
		if bUpgradeSigil {
			color.Cyan("Your version of SIGIL is outdated. Please get the latest one below :")
			color.Cyan("|-> https://romero.com/")
			color.Cyan("")
		}
		color.Cyan("==================================================================================")
	}

	if iErrors == 1 {
		color.Red("1 error has been found. Check it!")
	} else if iErrors > 1 {
		color.Red("%d errors have been found. Check them!", iErrors)
	} else {
		color.Green("Everything looks fine. Have fun!")
	}

	PressEnter()
}
