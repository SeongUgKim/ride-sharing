package util

const (
	XL    = "XL"
	X     = "X"
	BLACK = "BLACK"
)

// IsSupportedCabType returns true if the cab type is supported
func IsSupportedCabType(cabType string) bool {
	switch cabType {
	case XL, X, BLACK:
		return true
	}

	return false
}
