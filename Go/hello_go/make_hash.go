package main

import (
	"crypto/sha256"
	"encoding/hex"
)

// HashString は与えられた文字列をSHA-256でハッシュ化し、16進数文字列として返します
func HashString(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
