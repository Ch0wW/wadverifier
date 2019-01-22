# IWADverifier

This tool was created, as the original project by Russell has disappeared from the surface of the Internet.

IWADverifier is a small Command Line Interface tool written in Golang. It is used to quickly check if a DOOM IWAD is valid or not. <br />
It supports Drag & Drop support for Windows (and probably other systems, if it allows it)<br />
It can quickly also check if your IWADs are outdated or not. 
If so, it'll redirect you to another open-source tool called IWADPatcher 1.2, written by Phenex2, or, for the case of FreeDOOM, to the official Github repository.

# Usage
### Windows
`iwadverifier <wad.wad[ wad2.wad ...]>`
### Linux/Mac
`./IWADverifier <wad.wad[ wad2.wad ...]>`

# IWAD Support
IWADverifier currently supports these IWADs:
* DOOM v1.0 up to 1.9
* DOOM II v1.6 up to 1.9
* Final DOOM (Plutonia & TNT)
* Heretic v1.0 up to 1.3
* Hexen v1.0 & v1.1 (minus Dark Citadel's Addon)
* DOOM 3 XBOX IWADs
* DOOM/DOOM2 XBLA IWADs
* DOOM 3 BFG Edition IWADs (minus NERVE.WAD)
* FreeDOOM v0.8 up to 0.11.3

# Pre-Requisites for compilation
- Golang 1.10 or newer (previous versions weren't tested)
- Package "Color" from user Fatih (`go get github.com/fatih/color`)
- Package "go-ansi" from user k0kubun (`go get github.com/k0kubun/go-ansi`)

Then, inside the project folder, write `go build`, and that should be it.

# ToDo:
- [ ] Open the WAD to check if it's a IWAD or PWAD before doing the checks
- [ ] Add support for NERVE.WAD, DOOM2 Master levels, and the Hexen Addon.
- [ ] Probably a few fixes here and there.
- [x] Color support for better readability.
