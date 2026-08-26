package main

import "wadverifier/wad"

//--------------------------

type CustomData struct {
	Name string      `json:"name"`
	Data []wad.Entry `json:"wadinfo"`
}

var (
	IWADInfo_Doom         []wad.Entry
	IWADInfo_Doom2        []wad.Entry
	IWADInfo_FinalDoom    []wad.Entry
	IWADInfo_Heretic      []wad.Entry
	IWADInfo_Hexen        []wad.Entry
	IWADInfo_MasterLevels []wad.Entry
	IWADInfo_Strife       []wad.Entry
	IWADInfo_SVE          []wad.Entry
	IWADInfo_FreeDoom     []wad.Entry
	IWADInfo_Misc         []wad.Entry // PWAD and addons

	PWADInfo_Custom []wad.Entry

	// Patching messages
	iErrors = 0
)
