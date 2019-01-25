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

//--------------------------
type GPatch int

const (
	GAME_IWAD      GPatch = iota // i.e. for Doom, DooM 2, Heretic, Hexen, HexenDD, Plutonia, TNT, Strife
	GAME_SHAREWARE               // All the Doom 1 sharewares.
	GAME_FREEDOOM                // i.e. for Freedoom Phase 1/2 and FreeDM
	GAME_HACX                    // i.e. for HacX 1.0 - 1.2 (no support for 2.0 yet as it's still not released)
	GAME_CHEX_QUEST_3
	GAME_STRIFE_VE // i.e. for Strife: Veteran Edition (heard that 1.0/1.1 are still around out there...)
)

//--------------------------

// WadInfo : all WAD data returned from the program
type WadInfo struct {
	MD5Hash      string
	Version      string
	IsFinal      bool
	Game         GPatch
	PWADRequires string // If the official PWAD requires an IWAD to run
	Additional   string // If I need to display an additionnal message for this IWAD.
}

var (
	IWADInfo_Doom         []WadInfo
	IWADInfo_Doom2        []WadInfo
	IWADInfo_FinalDoom    []WadInfo
	IWADInfo_Heretic      []WadInfo
	IWADInfo_Hexen        []WadInfo
	IWADInfo_MasterLevels []WadInfo
	IWADInfo_Strife       []WadInfo
	IWADInfo_SVE          []WadInfo
	IWADInfo_FreeDoom     []WadInfo
	IWADInfo_Misc         []WadInfo // PWAD and addons

	// Patching messages
	bNeedsPatching   = false
	bUpgradeIWAD     = false
	bUpgradeDoomSW   = false
	bUpgradeFreeDoom = false
	bUpgradeHacX     = false
	bUpgradeSVE      = false
	bUpgradeCQ3      = false
	iErrors          = 0
)

const (
	m_version = "0.2"
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
			color.Yellow("%s is not a valid WAD file !", args[i])
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
			color.Cyan("To patch your IWAD to the latest version, please use IWADPatcher 1.2 by Phenex :")
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
			color.Cyan("Your version of FreeDOOM/FreeDM seems outdated. Please get the latest one below :")
			color.Cyan("|-> https://github.com/freedoom/freedoom/releases")
			color.Cyan("")
		}
		if bUpgradeHacX {
			color.Cyan("Your version of HacX seems outdated. Please get the latest one below :")
			color.Cyan("|-> http://www.drnostromo.com/hacx/page.php?content=download")
			color.Cyan("")
		}
		if bUpgradeCQ3 {
			color.Cyan("Your version of Chex Quest 3 seems outdated. Please get the latest one below :")
			color.Cyan("|-> http://www.chucktropolis.com/gamers.htm")
			color.Cyan("")
		}
		if bUpgradeSVE {
			color.Cyan("Your version of Strife: Veteran Edition seems outdated.")
			color.Cyan("• If you bought it on Steam, S:VE should be updated automatically.")
			color.Cyan("• If you bought it on GOG, you will need to redownload it (Latest version is 1.2) or to use GOG Galaxy")
			color.Cyan("")
		}
		color.Cyan("==================================================================================")
	}

	if iErrors == 1 {
		color.Red("1 error has been found. Check it !")
	} else if iErrors > 1 {
		color.Red("%d errors have been found. Check them !", iErrors)
	} else {
		color.Green("Everything looks fine. Happy fragging !")
	}

	PressEnter()
}
