package main

//--------------------------
type GPatch int64

const (
	GAME_NONE           GPatch = iota
	GAME_IWAD                  = 1 << 0 // The official, classic IWADs for Doom, DooM 2, Heretic, Hexen, HexenDD, Plutonia, TNT, Strife.
	GAME_SHAREWARE             = 1 << 1 // All the Doom 1 sharewares.
	GAME_FREEDOOM              = 1 << 2 // i.e. for Freedoom Phase 1/2 and FreeDM
	GAME_HACX                  = 1 << 3 // i.e. for HacX 1.0 - 1.2 (no support for 2.0 yet as it's still not released)
	GAME_CHEX_QUEST_3          = 1 << 4
	GAME_STRIFE_VE             = 1 << 5  // i.e. for Strife: Veteran Edition (heard that 1.0/1.1 are still around out there...)
	GAME_SIGIL                 = 1 << 6  // SIGIL by John Romero
	GAME_SIGIL_2               = 1 << 7  // SIGIL2 by John Romero
	GAME_REKKR                 = 1 << 8  // REKKR by Revae
	GAME_KEXDOOM2024           = 1 << 9  // KEXDoom/Doom + Doom II/OdaKEX (2024 re-release)
	GAME_KEXHEREXEN2025        = 1 << 10 // KEXen/Heretic + Hexen/OdaKEX (2025 re-release)
	GAME_DOOMUNITY             = 1 << 11 // DOOM (UNITY) - In case of, because it's still part of a re-release.
)

type GStatus int

const (
	IS_FINAL    GStatus = iota
	IS_NOTFINAL         = 1 << 0 // IS NOT the final release
	IS_UNKNOWN          = 1 << 1 // Cannot be fully disclosed of being the latest version or not.
)

type GFlags int

const (
	FL_NONE       GFlags = iota
	FL_RERELEASE         = 1 << 0 // Is part of a re-releases but is not used in communities.
	FL_HIDDEN            = 1 << 1 // Hidden - No need to display it
	FL_PRERELEASE        = 1 << 1 // Is part of a betas
)

//--------------------------

type CustomData struct {
	Name string    `json:"name"`
	Data []WadInfo `json:"wadinfo"`
}

// WadInfo : all WAD data returned from the program
type WadInfo struct {
	MD5Hash      string  `json:"md5"`
	Name         string  `json:"name"`
	Version      string  `json:"version"`
	Status       GStatus `json:"status"` // Is it the final release of this?
	Flags        GFlags  `json:"flags"`  // Special Flags
	Game         GPatch  `json:"game"`
	PWADRequires string  `json:"requires"`   // If the official PWAD requires an IWAD to run
	Additional   string  `json:"additional"` // If I need to display an additionnal message for this IWAD.
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

	PWADInfo_Custom []WadInfo

	// Patching messages
	iErrors = 0
)
