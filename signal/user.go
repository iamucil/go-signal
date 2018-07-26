package signal

import (
	"github.com/RadicalApp/libsignal-protocol-go/groups"
	"github.com/RadicalApp/libsignal-protocol-go/keys/identity"
	"github.com/RadicalApp/libsignal-protocol-go/protocol"
	"github.com/RadicalApp/libsignal-protocol-go/serialize"
	"github.com/RadicalApp/libsignal-protocol-go/session"
	"github.com/RadicalApp/libsignal-protocol-go/state/record"
	"github.com/RadicalApp/libsignal-protocol-go/util/keyhelper"
)

type user struct {
	name     string
	deviceID uint32
	address  *protocol.SignalAddress

	identityKeyPair *identity.KeyPair
	registrationID  uint32

	preKeys      []*record.PreKey
	signedPreKey *record.SignedPreKey

	sessionStore      *InMemorySession
	preKeyStore       *InMemoryPreKey
	signedPreKeyStore *InMemorySignedPreKey
	identityStore     *InMemoryIdentityKey
	senderKeyStore    *InMemorySenderKey

	sessionBuilder *session.Builder
	groupBuilder   *groups.SessionBuilder
}

func (u *user) buildSession(address *protocol.SignalAddress, serializer *serialize.Serializer) {
	u.sessionBuilder = session.NewBuilder(
		u.sessionStore,
		u.preKeyStore,
		u.signedPreKeyStore,
		u.identityStore,
		address,
		serializer,
	)
}

func (u *user) buildGroupSession(serializer *serialize.Serializer) {
	u.groupBuilder = groups.NewGroupSessionBuilder(u.senderKeyStore, serializer)
}

func newUser(name string, deviceID uint32, serializer *serialize.Serializer) *user {
	signalUser := &user{}

	// generate identity keyPair
	signalUser.identityKeyPair, _ = keyhelper.GenerateIdentityKeyPair()

	// generate a registration id
	signalUser.registrationID = keyhelper.GenerateRegistrationID()

	//generate preKeys
	signalUser.preKeys, _ = keyhelper.GeneratePreKeys(0, 100, serializer.PreKeyRecord)

	// generate Signed PreKey
	signalUser.signedPreKey, _ = keyhelper.GenerateSignedPreKey(signalUser.identityKeyPair, 0, serializer.SignedPreKeyRecord)

	signalUser.sessionStore = NewInMemorySession(serializer)
	signalUser.preKeyStore = NewInMemoryPreKey()
	signalUser.signedPreKeyStore = NewInMemorySignedPreKey()
	signalUser.identityStore = NewInMemoryIdentityKey(signalUser.identityKeyPair, signalUser.registrationID)
	signalUser.senderKeyStore = NewInMemorySenderKey()

	signalUser.signedPreKeyStore.StoreSignedPreKey(
		signalUser.signedPreKey.ID(),
		record.NewSignedPreKey(
			signalUser.signedPreKey.ID(),
			signalUser.signedPreKey.Timestamp(),
			signalUser.signedPreKey.KeyPair(),
			signalUser.signedPreKey.Signature(),
			serializer.SignedPreKeyRecord,
		),
	)

	// create a remote address that we'll be building our session with.
	signalUser.name = name
	signalUser.deviceID = deviceID
	signalUser.address = protocol.NewSignalAddress(name, deviceID)

	// create a group session builder
	signalUser.buildGroupSession(serializer)

	return signalUser
}
