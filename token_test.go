package main

import (
	"testing"
)

func TestTokenHashing(t *testing.T) {

	token := Token("thisisatest")
	hash, err := token.Hash()
	if err != nil {
		t.Error(err)
		return
	}

	if hash.IsValid("foo") != false {
		t.Error("Should return false with invalid 'password'")
	}

	if hash.IsValid("thisisatest") != true {
		t.Error("Should return true with valid 'password'")
	}

}

func TestEmptyHashes(t *testing.T) {

}
