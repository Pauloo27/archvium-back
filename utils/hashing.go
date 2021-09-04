package utils

import (
	"crypto/sha256"
	"encoding/base64"
)

var salt string

func HashPassword(password string) string {
	if salt == "" {
		salt = EnvString("AUTH_SALT")
	}
	hasher := sha256.New()
	hasher.Write([]byte(Fmt(salt, password)))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
