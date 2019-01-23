# IWADverifier

This tool was created, as the original project by Russell has disappeared from the surface of the Internet. 

`IWADverifier` is a small Command Line Interface tool written in Golang. It is used to quickly check if a DOOM-Engine based IWAD is valid or not. 

# Usage
### Windows
`iwadverifier <wad.wad[ wad2.wad ...]>`
### Linux/Mac
`./IWADverifier <wad.wad[ wad2.wad ...]>`

# Features
`IWADverifier` currently supports these features :

### IWAD Identification
IWADverifier can identify these IWADs:
* DOOM v1.0 up to 1.9
* DOOM II v1.6 up to 1.9
* Final DOOM (Plutonia & TNT)
* Heretic v1.0 up to 1.3
* Hexen v1.0 & v1.1 (minus Dark Citadel's Addon)
* Strife v1.0 up to 1.31
* DOOM 3 XBOX IWADs
* DOOM/DOOM2 XBLA IWADs
* DOOM 3 BFG Edition IWADs (minus NERVE.WAD)
* FreeDOOM v0.8 up to 0.11.3

### Latest version detection
IWADverifier looks up if your IWAD is the latest version or not. If it's not, a message tells you what to do to get the latest version of your file !

### Drag & Drop (Windows only ?)
To make it easier for everyone, you can directly drag your IWAD files to IWADverifier !

### Color support
Because having a white-only text in a commandline application is boring, IWADVerifier uses ANSI to color messages. It works perfectly on Windows, Linux and technically on every system !

# Pre-Requisites for compilation
- Golang 1.10 or newer (previous versions weren't tested)
- Package `color` from user Fatih (`go get github.com/fatih/color`)
- Package `go-ansi` from user k0kubun (`go get github.com/k0kubun/go-ansi`)

Then, inside the project folder, write `go build`, and that should be it.

# Screenshot
![ProgramSS](https://i.imgur.com/tviS1Gr.png)

# ToDo List
- [x] Open the WAD to check if it's a IWAD or PWAD before doing the checks
- [ ] Add support for NERVE.WAD, DOOM2 Master levels, and the Hexen Addon.
- [ ] Probably a few fixes here and there.
- [x] Color support for better readability.

# Licence
This program is licenced under GPLv3.