package main

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
