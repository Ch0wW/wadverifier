package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	ansi "github.com/k0kubun/go-ansi"
)

// WadInfo : all WAD data returned from the program
type WadInfo struct {
	MD5Hash      string
	Version      string
	IsFinal      bool
	IsFreedoom   bool   // If FreeDOOM, link for the upgrade to the latest version
	PWADRequires string // If the official PWAD requires an IWAD to run
	Additionnal  string // If I need to display an additionnal message for this IWAD.
}

var (
	IWADInfo_Doom      []WadInfo
	IWADInfo_Doom2     []WadInfo
	IWADInfo_FinalDoom []WadInfo
	IWADInfo_Heretic   []WadInfo
	IWADInfo_Hexen     []WadInfo
	IWADInfo_Strife    []WadInfo
	IWADInfo_FreeDoom  []WadInfo

	bNeedsPatching   = false
	bUpgradeIWAD     = false
	bUpgradeFreeDoom = false
	iErrors          = 0
)

const (
	m_version = "0.1"
)

// Just a quick function to require the user to press ENTER.
func PressEnter() {
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

func main() {

	// Initialize IWAD/Addon lists
	PopulateIWADInfos()

	color.Output = ansi.NewAnsiStdout()

	//==============================
	// HEADER OF THE PROGRAM
	//==============================
	color.Cyan("IWAD Verifier %s - By Ch0wW", m_version)
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

	// If there's some patching needed, warn the user how to do it.
	if bNeedsPatching && (bUpgradeFreeDoom || bUpgradeIWAD) {
		color.Cyan("==================================================================================")
		color.Cyan("")

		if bUpgradeIWAD {
			color.Cyan("To patch your IWAD to the latest version, please use IWADPatcher 1.2 by Phenex :")
			color.Cyan("• Windows binaries: http://downloads.zdaemon.org/iwadpatcher-1.2-bin.zip")
			color.Cyan("• Source code: http://downloads.zdaemon.org/iwadpatcher-1.2.zip")
			color.Cyan("")
		}
		if bUpgradeFreeDoom {
			color.Cyan("Your version of FreeDOOM/FreeDM seems outdated. Please get the latest one below :")
			color.Cyan("|-> https://github.com/freedoom/freedoom/releases")
			color.Cyan("")
		}
		color.Cyan("==================================================================================")
	}

	if iErrors == 1 {
		color.Red("1 error has been found. Check it !")
	} else if iErrors > 1 {
		color.Red("%d errors have been found. Check them !", iErrors)
	} else {
		color.Green("Everything looks fine. Happy gaming !")
	}

	PressEnter()
}
