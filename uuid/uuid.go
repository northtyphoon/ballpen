package uuid

import "fmt"

const (
	// Bits is the number of bits in a UUID
	Bits = 128

	// Size is the number of bytes in a UUID
	Size = Bits / 8

	stringFormat = "%08x-%04x-%04x-%04x-%012x"

	lengthInStringFormat = 36
)

var (
	// ErrInvalidUUID indicates uuid string is not valid
	ErrInvalidUUID = fmt.Errorf("Invalid uuid string")
)

// UUID represents a UUID value
type UUID [Size]byte

// Parse attempts to extract a uuid from the string or returns an error.
func Parse(str string) (UUID, error) {
	var uuid UUID

	if len(str) != lengthInStringFormat {
		return uuid, ErrInvalidUUID
	}

	segments := make([][]byte, 5)

	if _, err := fmt.Sscanf(str, stringFormat, &segments[0], &segments[1], &segments[2], &segments[3], &segments[4]); err != nil {
		return uuid, err
	}

	copy(uuid[0:4], segments[0])
	copy(uuid[4:6], segments[1])
	copy(uuid[6:8], segments[2])
	copy(uuid[8:10], segments[3])
	copy(uuid[10:16], segments[4])

	return uuid, nil
}
