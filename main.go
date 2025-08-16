package main

import (
	"bufio"
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	emoji "github.com/enescakir/emoji"
	"github.com/fatih/color"
	ansi "github.com/k0kubun/go-ansi"
)

const (
	mRelease      = 0
	mPointRelease = 7
	IWADbytes     = 1145132873
	PWADbytes     = 1145132880
)

var (
	patchflag        GPatch
	noenter          bool
	customdata       CustomData
	bFoundUnknownWAD bool
)

func CustomData_Init(filename string) (error, bool) {

	cfg, err := os.Open(filename)
	if err != nil {
		return err, false
	}

	err = json.NewDecoder(cfg).Decode(&customdata)
	if err != nil {
		return err, true
	}

	cfg.Close()

	return nil, false
}

// Just a quick function to require the user to press ENTER.
// Now, it only happens on Windows. (for the drag & drop feature)
func PressEnter() {

	if runtime.GOOS != "windows" || noenter {
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

func CompIWADData(data []WadInfo, hash string) (WadInfo, error) {
	for i := range data {
		if hash == data[i].MD5Hash {
			return data[i], nil
		}
	}
	return WadInfo{}, errors.New("unknown WAD")
}

func OutputVersion(b GStatus, f GFlags) string {
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	magenta := color.New(color.FgMagenta).SprintFunc()

	if b == IS_FINAL {
		if f&FL_RERELEASE == 1 {
			return fmt.Sprintf("%v %s", emoji.CheckMark, green("Latest version"))
		} else {
			return fmt.Sprintf("%v%v %s", emoji.CheckMark, emoji.CheckMark, green("Latest Original release"))
		}

	}
	if b == IS_NOTFINAL {
		return fmt.Sprintf("%v %s", emoji.CrossMark, red("Outdated release"))
	}

	return magenta("❔ Unknown release")
}

func CheckIWAD(filename string, hash string) bool {

	iwadOrder := [][]WadInfo{
		IWADInfo_Doom,
		IWADInfo_Doom2,
		IWADInfo_FinalDoom,
		IWADInfo_MasterLevels,
		IWADInfo_Heretic,
		IWADInfo_Hexen,
		IWADInfo_Strife,
		IWADInfo_SVE,
		IWADInfo_FreeDoom,
		IWADInfo_Misc,
	}

	yellow := color.New(color.FgYellow).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	fmt.Println("Checking file :", filename)

	var IWAD WadInfo
	var err error

	for _, wadlist := range iwadOrder {
		IWAD, err = CompIWADData(wadlist, hash)
		if err == nil {
			break // Success, exit loop
		}
	}

	// Check the other mods
	if err != nil {
		if len(PWADInfo_Custom) > 0 {
			IWAD, err = CompIWADData(PWADInfo_Custom, hash)
		}
	}

	// STILL NOTHING??? UGH. OK. ERROR TIME.
	if err != nil {
		// At this point, we should dissect the first bytes of the WAD to make sure it's a PWAD.
		// Then, check against the known Addons/Extensions (Hexen:DotDC / NervE)
		// If still nothing, we assume this WAD is unknown.
		iErrors = iErrors + 1
		color.Red("Wad is currently unknown to the database!")
		color.Red("MD5 Hash of file: %s", hash)

		fmt.Println("")
		return true
	}

	// However we are lucky
	fmt.Println("MD5:", IWAD.MD5Hash)
	if IWAD.Version != "" {
		ansi.Println("WAD :", green(IWAD.Name), cyan(fmt.Sprintf("(%s)", IWAD.Version)))
	} else {
		ansi.Println("WAD :", green(IWAD.Name))
	}

	if IWAD.PWADRequires != "" {
		ansi.Println("WAD Requires:", yellow(IWAD.PWADRequires))
	}

	// Add an error count if it's not the final version of a wad.
	if IWAD.Status == IS_NOTFINAL {
		iErrors += 1
		patchflag |= IWAD.Game // Now, flag our messages if our IWAD is older
	}

	// Hide information if unnecessary to the end-user.
	if IWAD.Flags&FL_RERELEASE == 1 {
		color.Yellow("  - This WAD is from a Commercial Re-release/Remakster and should not be used in sourceports !")
	}
	if IWAD.Flags&FL_PRERELEASE == 1 {
		color.Yellow("  - This WAD is from a beta release.")
	}

	// Don't output the versioning status if it's a Pre-Release...
	if IWAD.Flags&FL_PRERELEASE == 0 && IWAD.Flags&FL_HIDDEN == 0 {
		ansi.Println("Status: ", OutputVersion(IWAD.Status, IWAD.Flags))
	}

	// If the IWAD has an additionnal message, please write it so.
	if IWAD.Additional != "" {
		ansi.Println("Additional info :", cyan(IWAD.Additional))
	}

	fmt.Println("")
	return false
}

func main() {

	color.Cyan("WAD Verifier %d.%d", mRelease, mPointRelease)
	color.Cyan("https://github.com/ch0ww/wadverifier")
	color.Cyan("---------------------------------------")
	fmt.Println("")

	var nocheck, verbose bool
	var filename string

	flag.BoolVar(&verbose, "v", false, "Be more verbose")
	flag.BoolVar(&noenter, "no-enter", false, "Remove the need to press ENTER at the end of the program.")
	flag.BoolVar(&nocheck, "offline", false, "Check the Github project for the latest release")
	flag.StringVar(&filename, "resfile", "pwaddata.json", "opens a custom WAD resources file.")
	flag.Parse()

	// Get the arguments
	args := flag.Args()

	if len(args) == 0 {
		color.Yellow("No wad file specified.")
		color.Yellow("Usage: iwadverifier [-v] [-no-enter] [-resfile file.json] <wad.wad[ wad2.wad ...]>")

		fmt.Println("== Flags ==")
		fmt.Println("-v : Be more verbose in case of warning messages")
		fmt.Println("-no-enter : Removes the check to press ENTER at the end of the program")
		fmt.Printf("\n\n== Arguments ==\n")
		fmt.Println("-resfile <filename>: opens a custom WAD resources file.")
		return
	}

	err, errtype := CustomData_Init(filename)
	if err != nil {
		if errtype {
			color.Yellow("Unable to open or read %s. Make sure it's a properly formatted JSON file.", filename)
		}
	}

	// Initialize IWAD/Addon lists
	PopulateIWADInfos()

	// Put the colors
	color.Output = ansi.NewAnsiStdout()

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

			if verbose {
				color.Yellow("%s is not a .wad file ! Ignoring...", args[i])
				fmt.Println("")
			}
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

		bValue := CheckIWAD(args[i], hash)

		if !bFoundUnknownWAD && bValue {
			bFoundUnknownWAD = true
		}
	}

	// This is ugly, but I have to find a way to make
	// If there's some patching needed, warn the user how to do it.
	if patchflag != 0 {

		type flagMessage struct {
			flag GPatch
			msg  string
		}

		// Define all flag-message pairs
		messages := []flagMessage{
			{GAME_IWAD, "To patch your IWAD to the latest version, please use IWADPatcher 1.2 by Peter Vaskovics:\n• Windows binaries: http://downloads.zdaemon.org/iwadpatcher-1.2-bin.zip\n• Source code: http://downloads.zdaemon.org/iwadpatcher-1.2.zip"},
			{GAME_SHAREWARE, "Your Shareware version of Doom is outdated. Please get the latest version below :\n|-> https://www.doomworld.com/idgames/idstuff/doom/doom19s"},
			{GAME_FREEDOOM, "Your version of FreeDOOM/FreeDM is outdated. Please get the latest one below :\n|-> https://github.com/freedoom/freedoom/releases"},
			{GAME_HACX, "Your version of HacX is outdated. Please get the latest one below :\n|-> http://www.drnostromo.com/hacx/page.php?content=download"},
			{GAME_CHEX_QUEST_3, "Your version of Chex Quest 3 is outdated. Please get the latest one below :\n|-> http://www.chucktropolis.com/gamers.htm"},
			{GAME_STRIFE_VE, "Your version of Strife: Veteran Edition is outdated.\n• If you bought it on Steam, S:VE should be updated automatically.\n• If you bought it on GOG, you will need to redownload it (Latest version is 1.2) or to use GOG Galaxy"},
			{GAME_SIGIL, "Your version of SIGIL is outdated. Please get the latest one below :\n|-> https://romero.com/sigil"},
			{GAME_SIGIL_2, "Your version of SIGIL II is outdated. Please get the latest one below :\n|-> https://romero.com/sigil"},
			{GAME_REKKR, "Your version of REKKR is outdated. Please get the latest one below :\n|-> http://manbitesshark.com/"},
			{GAME_KEXDOOM2024, "The wad used in Doom + Doom II looks outdated. Please update your binaries to the latest version on STEAM or GOG."},
			{GAME_KEXHEREXEN2025, "The wad used in Heretic + Hexen looks outdated. Please update your binaries to the latest version on STEAM or GOG."},
		}
		color.Cyan("==================================================================================")
		color.Cyan("")

		// Loop through all messages
		// Loop through all messages
		for _, msg := range messages {
			if patchflag&msg.flag != 0 {
				color.Cyan(msg.msg)
			}
		}
		color.Cyan("==================================================================================")
	}

	if bFoundUnknownWAD {
		color.Red("If you noticed some WADs entries were missing, please create an issue : https://github.com/Ch0wW/wadverifier/issues/")
	}

	if iErrors == 1 {
		color.Red("1 error has been found.")
	} else if iErrors > 1 {
		color.Red("%d errors have been found.", iErrors)
	} else {
		color.Green("Everything looks fine. Have fun!")
	}

	PressEnter()
}
