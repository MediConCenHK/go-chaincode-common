package go_chaincode_common

import (
	"testing"
	"github.com/testify/assert"
)

const signKey = `-----BEGIN PRIVATE KEY-----
MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgvGTjKl43i7pvc4pV
n1IqHRQQ4aQSa6ZsvX8Gg59IZlKhRANCAAT5xEqfR/tm6xAb9mWvyWHVN/xuVRnw
mug9Nc9QIMvCiP4uzObrIkHYK6CTTwmzuddFID1lic+KWArkiyRrfWEg
-----END PRIVATE KEY-----`

const signCert  = `-----BEGIN CERTIFICATE-----
MIICFDCCAbugAwIBAgIRAJB97xiyiE+/XxA7QvfJUdMwCgYIKoZIzj0EAwIwbDEL
MAkGA1UEBhMCQ04xETAPBgNVBAgTCFNoYW5naGFpMREwDwYDVQQHEwhTaGFuZ2hh
aTEZMBcGA1UEChMQb3JnMS5leGFtcGxlLmNvbTEcMBoGA1UEAxMTY2Eub3JnMS5l
eGFtcGxlLmNvbTAeFw0xOTAxMjEwODU1MjhaFw0yOTAxMTgwODU1MjhaMG0xCzAJ
BgNVBAYTAkNOMREwDwYDVQQIEwhTaGFuZ2hhaTERMA8GA1UEBxMIU2hhbmdoYWkx
DzANBgNVBAsTBmNsaWVudDEnMCUGA1UEAwweVXNlclNpZ25Ub2tlbkBvcmcxLmV4
YW1wbGUuY29tMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE+cRKn0f7ZusQG/Zl
r8lh1Tf8blUZ8JroPTXPUCDLwoj+Lszm6yJB2Cugk08Js7nXRSA9ZYnPilgK5Isk
a31hIKM9MDswDAYDVR0TAQH/BAIwADArBgNVHSMEJDAigCAIu4B9SOwh/WYVBaX7
FxZafDJVwv0gnua2r/21iMDVuzAKBggqhkjOPQQDAgNHADBEAiAnus+Ojn+ODxCh
dXPuPZKRbauzSwgp38xvfNujJTpoEgIgRoQSZAiDYUKupO+pLioOmB+S+B+pXiZK
msq/aA3ro+0=
-----END CERTIFICATE-----`

const otherCert = `-----BEGIN CERTIFICATE-----
MIICADCCAaegAwIBAgIRAKCvy/ZyABcmrZu1ibL6R+QwCgYIKoZIzj0EAwIwZDEL
MAkGA1UEBhMCQ04xETAPBgNVBAgTCFNoYW5naGFpMREwDwYDVQQHEwhTaGFuZ2hh
aTEVMBMGA1UEChMMb3JnMS5hYWEuY29tMRgwFgYDVQQDEw9jYS5vcmcxLmFhYS5j
b20wHhcNMTgxMTEzMDgxNTQ5WhcNMjgxMTEwMDgxNTQ5WjBhMQswCQYDVQQGEwJD
TjERMA8GA1UECBMIU2hhbmdoYWkxETAPBgNVBAcTCFNoYW5naGFpMQ8wDQYDVQQL
EwZjbGllbnQxGzAZBgNVBAMMElVzZXIxQG9yZzEuYWFhLmNvbTBZMBMGByqGSM49
AgEGCCqGSM49AwEHA0IABOAKIQdTuUAt3xMojGGfVuk3nB3uwslvyCRl32CpHWMB
nC7Y8AwaVlWYUJmAOie/FEmW6oyeoxJsE7M4GGcgXyGjPTA7MAwGA1UdEwEB/wQC
MAAwKwYDVR0jBCQwIoAgugDycinD9SEBVvx6Xi6LtlEvQEetzlkIuq9v33qsp3Aw
CgYIKoZIzj0EAwIDRwAwRAIgIqOLjefqPzLtuzA96neBmMeWqn2kFuhFJu9QjaaB
VeYCICLzw4FCcrr1BWooEOOiJTNdzove8UgfIEpg5PloRQ+T
-----END CERTIFICATE-----`

func TestSignAndVerify(t *testing.T) {

	// sign
	signedMsg, err := Sign("ID", []byte("MSG1"), []byte(signKey))
	assert.NoError(t, err)
	assert.NotNil(t, signedMsg.Sig)

	// verify
	valid, err := Verify(signedMsg, []byte(signCert))
	assert.NoError(t, err)
	assert.True(t, valid)
}

func TestVerifyWithModifiedMessage(t *testing.T) {

	// sign
	signedMsg, err := Sign("ID", []byte("MSG1"), []byte(signKey))
	assert.NoError(t, err)
	assert.NotNil(t, signedMsg.Sig)


	// verify
	signedMsg.Payload = []byte("MSG2") // msg is modified
	valid, err := Verify(signedMsg, []byte(signCert))
	assert.NoError(t, err)
	assert.False(t, valid)
}

func TestVerifyWithOtherCert(t *testing.T) {

	// sign
	signedMsg, err := Sign("ID", []byte("MSG1"), []byte(signKey))
	assert.NoError(t, err)
	assert.NotNil(t, signedMsg.Sig)

	// verify
	valid, err := Verify(signedMsg, []byte(otherCert))
	assert.NoError(t, err)
	assert.False(t, valid)
}