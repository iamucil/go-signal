package signal

import (
	"github.com/RadicalApp/libsignal-protocol-go/serialize"
)

func newSerializer() *serialize.Serializer {
	serializer := serialize.NewJSONSerializer()

	return serializer
}
