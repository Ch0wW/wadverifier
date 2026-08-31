package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"wadverifier/gameinfo"
	"wadverifier/wad"
	"wadverifier/wad/flags"
	"wadverifier/wad/games"
	"wadverifier/wad/status"
	"wadverifier/wadapi"

	emoji "github.com/enescakir/emoji"
	"github.com/fatih/color"
	ansi "github.com/k0kubun/go-ansi"
)

const (
	mRelease      = 0
	mPointRelease = 8
	mMinorRelease = 0
	IWADbytes     = 1145132873
	PWADbytes     = 1145132880
)

var (
	patchflag        games.PatchType
	noenter          bool
	PWADDefinitions  []wad.Entry
	bFoundUnknownWAD bool
)

// Just a quick function to require the user to press ENTER.
// Now, it only happens on Windows. (for the drag & drop feature)
func PressEnter() {

	if runtime.GOOS != "windows" || noenter {
		return
	}

	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func OutputVersion(b status.Status, f flags.Flags) string {
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	magenta := color.New(color.FgMagenta).SprintFunc()

	if b.IsFinal() {
		if f.IsRerelease() {
			return fmt.Sprintf("%v %s", emoji.CheckMark, green("Latest version"))
		} else {
			return fmt.Sprintf("%v%v %s", emoji.CheckMark, emoji.CheckMark, green("Latest Original release"))
		}

	}
	if b.IsNotFinal() {
		return fmt.Sprintf("%v %s", emoji.CrossMark, red("Outdated release"))
	}

	return magenta("❔ Unknown release")
}

func CheckIWAD(filename string, hash wad.HashInfo, wadlist []wad.Entry) bool {

	yellow := color.New(color.FgYellow).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	fmt.Println("Checking file :", filename)

	var WAD wad.Entry
	var err error

	for _, wadlist := range wadlist {
		WAD, err = CompIWADData(wadlist, hash)
		if err == nil {
			break // Success, exit loop
		}
	}

	// Check the other mods
	/*if err != nil {
		if len(PWADInfo_Custom) > 0 {
			IWAD, err = CompIWADData(PWADInfo_Custom, hash)
		}
	}*/

	// STILL NOTHING??? UGH. OK. ERROR TIME.
	if err != nil {
		// At this point, we should dissect the first bytes of the WAD to make sure it's a PWAD.
		// Then, check against the known Addons/Extensions (HEXDD / Nerve)
		// If still nothing, we assume this WAD is unknown.
		iErrors = iErrors + 1
		color.Red("Wad is currently unknown to the database!")
		color.Red("MD5 Hash : %s / SHA-1 Hash: %s", hash.MD5, hash.SHA1)

		fmt.Println("")
		return true
	}

	// However we are lucky
	fmt.Printf("MD5: %s / SHA-1: %s\n", hash.MD5, hash.SHA1)

	if WAD.Version != "" {
		ansi.Println("WAD :", green(WAD.Name), cyan(fmt.Sprintf("(%s)", WAD.Version)))
	} else {
		ansi.Println("WAD :", green(WAD.Name))
	}

	if WAD.PWADRequires != "" {
		ansi.Println("WAD Requires:", yellow(WAD.PWADRequires))
	}

	// Add an error count if it's not the final version of a wad.
	if WAD.Patchinfo != games.NONE && WAD.Status.IsNotFinal() {
		iErrors += 1
		patchflag |= WAD.Patchinfo // Now, flag our messages if our IWAD is older
	}

	// Hide information if unnecessary to the end-user.
	if WAD.Flags&flags.RERELEASE == 1 {
		color.Yellow("  - This WAD comes from a Re-release and might not be sourceport-compatible or retro-compatible with the original file!")
	}
	if WAD.Flags&flags.PRERELEASE == 1 {
		color.Yellow("  - This WAD comes from a pre-release.")
	}

	// Don't output the versioning status if the WAD is not recognized...
	if WAD.Patchinfo != games.NONE || WAD.Flags&flags.HIDDEN == 0 {
		ansi.Println("Status: ", OutputVersion(WAD.Status, WAD.Flags))
	}

	// If the WAD has an additionnal message, please write it so.
	if WAD.Additional != "" {
		ansi.Println("Additional info:", cyan(WAD.Additional))
	}

	fmt.Println("")
	return false
}

func main() {

	color.Cyan("WAD Verifier %d.%d.%d", mRelease, mPointRelease, mMinorRelease)
	color.Cyan("https://github.com/ch0ww/wadverifier")
	color.Cyan("---------------------------------------")
	fmt.Println("")

	var verbose bool
	var customjson string

	flag.BoolVar(&verbose, "v", false, "Add verbose messages")
	flag.BoolVar(&noenter, "no-enter", false, "Remove the need to press ENTER at the end of the program.")
	flag.StringVar(&customjson, "resfile", "pwaddata.json", "specify a custom json file containing custom PWAD definitions.")
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

	// Initialize IWAD/Addon lists
	iwadlist := gameinfo.PopulateWadInfo()
	if customjson != "" {
		err, PWADList := wadapi.LoadCustomPWADFile(customjson)
		if err != nil {
			color.Yellow("Unable to open or read %s (%s)", customjson, err)
		} else {
			iwadlist = append(iwadlist, PWADList.WadEntries...)
		}
	}

	// Put the colors
	color.Output = ansi.NewAnsiStdout()

	for _, fFile := range args {

		// Check if the user omitted the extension.
		// If so, assume the file is a .wad
		// ToDo: Check later for .pk3 files !
		if filepath.Ext(strings.ToLower(fFile)) == "" {
			fFile = fmt.Sprintf("%s.wad", fFile)
			fmt.Println(fFile)
		}

		// Try to check if file is a .wad
		if filepath.Ext(strings.ToLower(fFile)) != ".wad" {

			if verbose {
				color.Yellow("%s is not a .wad file! Skipping...", fFile)
				fmt.Println("")
			}
			continue
		}

		valid := OpenCheckWADValid(fFile)
		if !valid {
			color.Yellow("%s is not a valid WAD file! Skipping...", fFile) // Need to call SPA 1-800-388-PIR8 ?! Memories...
			iErrors = iErrors + 1
			continue
		}

		// Now, try to get the MD5 hash from the file
		hash_md5, err := wad.GetMD5Hash(fFile)
		if err != nil {
			color.Yellow("Error getting the MD5 hash (Reason: %s). Skipping... ", err)
			iErrors = iErrors + 1
			fmt.Println("")
			continue
		}

		hash_sha1, err := wad.GetSHA1Hash(fFile)
		if err != nil {
			color.Yellow("Error getting the MD5 hash (Reason: %s). Skipping... ", err)
			iErrors = iErrors + 1
			fmt.Println("")
			continue
		}

		hashes := wad.HashInfo{
			MD5:  hash_md5,
			SHA1: hash_sha1,
		}

		bValue := CheckIWAD(fFile, hashes, iwadlist)

		if !bFoundUnknownWAD && bValue {
			bFoundUnknownWAD = true
		}
	}

	// This is ugly, but I have to find a way to make
	// If there's some patching needed, warn the user how to do it.
	if patchflag != 0 {

		type flagMessage struct {
			flag games.PatchType
			msg  string
		}

		// Define all flag-message pairs
		messages := []flagMessage{
			{games.IWAD, "To patch your IWAD to the latest version, please use IWADPatcher 1.2 by Peter Vaskovics:\n• Windows binaries: http://downloads.zdaemon.org/iwadpatcher-1.2-bin.zip\n• Source code: https://github.com/petervas/iwadpatcher"},
			{games.DOOM_SHAREWARE, "Your Shareware version of Doom is outdated. Please get the latest version below :\n|-> https://www.doomworld.com/idgames/idstuff/doom/doom19s"},
			{games.FREEDOOM, "Your version of FreeDOOM/FreeDM is outdated. Please get the latest one below :\n|-> https://github.com/freedoom/freedoom/releases"},
			{games.HACX, "Your version of HacX is outdated. Please get the latest one below :\n|-> http://www.drnostromo.com/hacx/page.php?content=download"},
			{games.CHEX_QUEST_3, "Your version of Chex Quest 3 is outdated. Please get the latest one below :\n|-> http://www.chucktropolis.com/gamers.htm"},
			{games.STRIFE_SHAREWARE, "Download the latest shareware of Strife at https://www.doomworld.com/idgames/roguestuff/strife11"},
			{games.STRIFE_VETERAN_EDITION, "Your version of Strife: Veteran Edition is outdated.\n• If you bought it on Steam, S:VE should be updated automatically.\n• If you bought it on GOG, you will need to redownload it (Latest version is 1.2) or to use GOG Galaxy"},
			{games.SIGIL, "Your version of SIGIL is outdated. Please get the latest one below :\n|-> https://romero.com/sigil"},
			{games.SIGIL_2, "Your version of SIGIL II is outdated. Please get the latest one below :\n|-> https://romero.com/sigil"},
			{games.REKKR, "Your version of REKKR is outdated. Please get the latest one below :\n|-> http://manbitesshark.com/"},
			{games.KEX_DOOM2024, "The wad used in Doom + Doom II looks outdated. Please update your binaries to the latest version on STEAM or GOG."},
			{games.KEX_HERETIC_HEXEN2025, "The wad used in Heretic + Hexen looks outdated. Please update your binaries to the latest version on STEAM or GOG."},
			{games.DOOM_UNITY, "DOOM Unity is outdated. It is greatly recommended to install the Nightdive Studio's updated port instead (free upgrade)"},
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
		color.Red("1 outdated or unknown WAD has been found.")
	} else if iErrors > 1 {
		color.Red("%d outdated or unknown WADs have been found.", iErrors)
	} else {
		color.Green("No problem detected. Have fun!")
	}

	PressEnter()
}
