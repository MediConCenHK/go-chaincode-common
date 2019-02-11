package go_chaincode_common

import . "github.com/davidkhala/goutils"

func EnsureTransientMap(transient map[string][]byte, property string) []byte {
	if transient[property] == nil {
		PanicString("[" + property + "] not found in transientMap")
	}
	return transient[property]
}
