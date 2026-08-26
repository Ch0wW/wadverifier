package status

type Status int

const (
	FINAL    Status = iota
	NOTFINAL        = 1 << 0 // IS NOT the final release
	UNKNOWN         = 1 << 1 // Cannot be fully disclosed of being the latest version or not.
)

func (s Status) EncodeToString() string {
	switch s {
	case FINAL:
		return "Final"
	case NOTFINAL:
		return "Not Final"
	}

	return "Unknown version"
}

func (s Status) IsFinal() bool {
	return s == FINAL
}

func (s Status) IsNotFinal() bool {
	return s == NOTFINAL
}
