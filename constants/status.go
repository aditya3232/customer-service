package constants

type CustomerStatusString string

const (
	Active   CustomerStatusString = "ACTIVE"
	InActive CustomerStatusString = "INACTIVE"
)

func (s CustomerStatusString) IsValid() bool {
	switch s {
	case Active, InActive:
		return true
	default:
		return false
	}
}
