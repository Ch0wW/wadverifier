package flags

type Flags int

const (
	NONE       Flags = iota
	RERELEASE        = 1 << 0 // Is part of a re-releases but is not used in communities.
	HIDDEN           = 1 << 1 // Hidden - No need to display it
	PRERELEASE       = 1 << 2 // Is part of a betas
	EXTRADATA        = 1 << 3 // Extra Data (New Music, etc etc).
)

func (f Flags) IsRerelease() bool {
	return f == RERELEASE
}

func (f Flags) IsPrerelease() bool {
	return f == PRERELEASE
}
