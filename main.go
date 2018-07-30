package main

import (
	"fmt"

	"github.com/RadicalApp/libsignal-protocol-go/util/keyhelper"
	"github.com/iamucil/go-signal/signal"
)

func main() {
	regID := keyhelper.GenerateRegistrationID()
	fmt.Println(regID)
	signal.Serializing()
}
