package signal

import (
	"fmt"

	"github.com/RadicalApp/libsignal-protocol-go/fingerprint"
	"github.com/RadicalApp/libsignal-protocol-go/keys/prekey"
	"github.com/RadicalApp/libsignal-protocol-go/logger"
	"github.com/RadicalApp/libsignal-protocol-go/session"
	"github.com/RadicalApp/libsignal-protocol-go/state/record"
	"github.com/kr/pretty"
)

func Serializing() {
	serializer := newSerializer()

	alice := newUser("Alice", 1, serializer)
	bob := newUser("Bob", 2, serializer)

	// create our users who will talk to each other
	alice.buildSession(bob.address, serializer)
	bob.buildSession(alice.address, serializer)
	fmt.Printf("bob's registration id: %# v\n", bob.registrationID)
	fmt.Printf("bob's device id: %# v\n", bob.deviceID)
	fmt.Printf("bob's prekey: %# v\n", bob.preKeys[0].ID())
	// fmt.Printf("%# v\n", bob)
	retrivedPreKey := prekey.NewBundle(
		bob.registrationID,
		bob.deviceID,
		bob.preKeys[0].ID(),
		bob.signedPreKey.ID(),
		bob.preKeys[0].KeyPair().PublicKey(),
		bob.signedPreKey.KeyPair().PublicKey(),
		bob.signedPreKey.Signature(),
		bob.identityKeyPair.PublicKey(),
	)
	// fmt.Printf("Retrieved key: %# v\n", retrivedPreKey)
	// process Bob's retrieved prekey to establish a session
	alice.sessionBuilder.ProcessBundle(retrivedPreKey)

	// create a session ciper to encrypt messages to bob
	plainTextMessage := []byte("Hello")
	sessionCipher := session.NewCipher(alice.sessionBuilder, bob.address)
	fmt.Printf("Session Cipher: % #v\n", sessionCipher)
	sessionCipher.Encrypt(plainTextMessage)
	fmt.Printf("%# v\n", plainTextMessage)

	// Serialize our session so it can be stored
	loadedSession := alice.sessionStore.LoadSession(bob.address)
	serializedSession := loadedSession.Serialize()
	logger.Debug(string(serializedSession))

	// // try deserializing our session back into an object
	deserializedSession, err := record.NewSessionFromBytes(serializedSession, serializer.Session, serializer.State)
	if err != nil {
		logger.Error("Failed to deserilize session")
	}

	fmt.Printf("Original Session Record: %# v\n", pretty.Formatter(loadedSession))
	fmt.Printf("Deserialized Session Record: %# v\n", pretty.Formatter(deserializedSession))

}

func Fingerprint() {
	serializer := newSerializer()

	alice := newUser("Alice", 1, serializer)
	bob := newUser("Bob", 1, serializer)

	fp := fingerprint.NewDisplay(
		alice.identityKeyPair.PublicKey().Serialize(),
		bob.identityKeyPair.PublicKey().Serialize(),
	)

	fmt.Println(fp.DisplayText())
}

//func init() {
// Logger.setup(Debug)
//logger.Configure("")
//logger.Setup({Debug})
//}
