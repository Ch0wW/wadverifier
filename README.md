# IWADverifier

This tool was created, as the original project by Russell has disappeared from the surface of the Internet. 

`IWADverifier` is a small Command Line Interface tool written in Golang. It is used to quickly check if a DOOM-Engine based IWAD is valid or not, and up to date. 

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
* DOOM II v1.666 up to 1.9
* DOOM & DOOM II - Unity versions
* Final DOOM (Plutonia & TNT)
* Master levels of DooM II
* DOOM 3 XBOX
* DOOM/DOOM2 XBLA
* DOOM 3 BFG (+ NERVE.WAD)
* FreeDOOM v0.8 up to 0.12.1
* Heretic v1.0 up to 1.3
* Hexen v1.0 & v1.1 (+ Dark Citadel's Addon)
* Strife v1.0 up to 1.31 (+ Veteran Edition)
* HacX v1.0 up to 1.2
* Chex Quest 1
* Chex Quest 3 v1.0 & 1.4

* And many many more !

### Latest version detection
IWADverifier looks up if your IWAD is the latest version or not. If it's not, a message tells you what to do to get the latest version of your file !

### Drag & Drop (Windows only)
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
- [x] Open the WAD to check if it's a IWAD or PWAD before doing the checks.
- [x] Add support for NERVE.WAD, and the Hexen Addon.
- [x] Add support for HacX and Chex Quest.
- [x] Add support for Master levels for Doom II.
- [ ] Add support for GZDoom .pk3 IWADs (Action DooM / AD2:UB / Adventures of Square / W:BoA / Harmony / Blasphemer / ...)
- [ ] Probably a few code optimizations here and there.
- [x] Color support for better readability.

# Huge thanks to 
* Mike Swanson (Chungy) for adding several IWADs to the list !

# Licence
This program is licenced under GPLv3.
