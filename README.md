# IWADverifier

`IWADverifier` is a small Command Line Interface tool written in Golang. It is used to quickly check if a DOOM-Engine based IWAD is valid or not, and up to date. 

# Usage
```iwadverifier [-v] [-no-enter] <wad.wad[ wad2.wad ...]>``` 

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
* Chex Quest 1, 2, and 3 (v1.0 & 1.4)
* SIGIL v1.0 up to 1.21 
* DOOM 64 (NightDive Studios)
* And many many more !

### Latest version detection
IWADverifier looks up if your IWAD is the latest version or not. If it's not, a message tells you what to do to get the latest version of your file !

### Drag & Drop (Windows only)
Windows users can directly drag their IWAD files to the IWADverifier executable to quickly verify its validity!

### Color support
Because having a white-only text in a commandline application is boring, IWADVerifier uses ANSI to color messages. All systems should support it without any issue.

# Pre-Requisites for compilation
- Golang 1.12 or newer (previous versions weren't tested)
- Package `color` from user Fatih (`go get github.com/fatih/color`)
- Package `go-ansi` from user k0kubun (`go get github.com/k0kubun/go-ansi`)

Then, inside the project folder, write `go build`, and that should be it.

# Screenshot
![ProgramSS](https://i.imgur.com/tviS1Gr.png)

# ToDo List
- [ ] Add support for GZDoom .pk3 IWADs (Action DooM / AD2:UB / Adventures of Square / W:BoA / Harmony / Blasphemer / ...)
- [ ] Get the earlier "Unity Doom" wad versions from the September 2020 update.
- [ ] Probably a few code optimizations here and there.

# Huge thanks to 
* Mike Swanson (Chungy) for adding several IWADs to the list !

# Licence
This program is licenced under GPLv3.
