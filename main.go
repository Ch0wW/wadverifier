package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
	"io"
	"crypto/md5"
	"encoding/hex"
	"path/filepath"
)

// WadInfo : all WAD data returned from the program
type WadInfo struct {
	MD5Hash string
	Version string
	IsFinal bool
	PWADRequires string			// If the official PWAD requires an IWAD to run
}

var (
	IWADInfo_Doom []WadInfo	
	IWADInfo_Doom2 []WadInfo
	IWADInfo_FinalDoom []WadInfo
	IWADInfo_Heretic []WadInfo
	IWADInfo_Hexen []WadInfo
	IWADInfo_Misc []WadInfo

	bNeedsPatching = false
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
	
	//==============================
	// HEADER OF THE PROGRAM
	//==============================
	fmt.Println("IWAD Verifier", m_version," - By Ch0wW")
	fmt.Println("https://github.com/ch0ww/iwadverifier")
	fmt.Println("---------------------------------------")
	fmt.Println("")

	// Get the arguments
	args := os.Args[1:]

	if (len(args) == 0) {
		fmt.Println("No argument specified.");
		fmt.Println("Usage: iwadverifier <wad.wad[ wad2.wad ...]>")
		return
	}

	for i := range args {

		// Check if the user omitted the extension.
		// If so, assume the file is a .wad
		if (filepath.Ext(strings.ToLower(args[i])) == "") {
			args[i] = fmt.Sprintf("%s.wad", args[i])
			fmt.Println(args[i])
		}

		// Try to check if file is a .wad
		if ( filepath.Ext(strings.ToLower(args[i])) != ".wad"){
			fmt.Println(args[i], "is not a .wad file ! Ignoring...")
			fmt.Println("")
			continue
		}

		// Now, try to get the MD5 hash from the file
		hash, err := hash_file_md5(args[i])
		if (err != nil) {
			fmt.Println("Error getting the MD5 hash: Skipping... (",err,")")
			fmt.Println("")
			continue
		}

		CheckIWAD(args[i], hash)
	}

	// If there's some patching needed, warn the user how to do it.
	if (bNeedsPatching) {
		fmt.Println("==================================================================================")
		fmt.Println("")
		fmt.Println("To patch your IWAD to the latest version, please use IWADPatcher 1.2 by Phenex :")
		fmt.Println("Windows binaries: http://downloads.zdaemon.org/iwadpatcher-1.2-bin.zip")
		fmt.Println("Source code: http://downloads.zdaemon.org/iwadpatcher-1.2.zip")
		fmt.Println("")
		fmt.Println("==================================================================================")
	}

	PressEnter()
}