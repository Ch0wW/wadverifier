package games

// --------------------------
type PatchType int64

const (
	NONE                   PatchType = iota   // No Patch available
	DOOM_SHAREWARE                   = 1 << 0 // All the Doom 1 sharewares.
	IWAD                             = 1 << 1 // The official, classic IWADs for Doom, DooM 2, Heretic, Hexen, HexenDD, Plutonia, TNT, Strife.
	FREEDOOM                         = 1 << 2 // i.e. for Freedoom Phase 1/2 and FreeDM
	HERETIC_SHAREWARE                = 1 << 3 // i.e. for Freedoom Phase 1/2 and FreeDM
	HACX                             = 1 << 4 // i.e. for HacX 1.0 - 1.2 (no support for 2.0 yet as it's still not released)
	CHEX_QUEST_3                     = 1 << 5
	STRIFE_SHAREWARE                 = 1 << 6  //
	STRIFE_VETERAN_EDITION           = 1 << 8  //
	SIGIL                            = 1 << 9  // SIGIL by John Romero
	SIGIL_2                          = 1 << 10 // SIGIL2 by John Romero
	REKKR                            = 1 << 11 // REKKR by Revae
	KEX_DOOM2024                     = 1 << 12 // KEXDoom/Doom + Doom II/OdaKEX (2024 re-release)
	KEX_HERETIC_HEXEN2025            = 1 << 13 // KEXen/Heretic + Hexen (2025 re-release)
	DOOM_UNITY                       = 1 << 14 // DOOM (UNITY) - In case of, because it's still part of a re-release.
)
