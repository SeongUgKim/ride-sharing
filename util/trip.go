package util

const (
	Creating = "creating"
	Running  = "running"
	Done     = "done"
)

func IsSupportedTripStatus(tripStatus string) bool {
	switch tripStatus {
	case Creating, Running, Done:
		return true
	}

	return false
}

