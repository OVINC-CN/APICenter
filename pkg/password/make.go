package password

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"hash"
	"math/rand"

	"golang.org/x/crypto/pbkdf2"
)

const (
	iter         = 100000
	saltSize     = 22
	hashAlg      = "pbkdf2_sha256"
	allowedChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	unusableHash = "!"
)

func Iter() int {
	return iter
}

func SaltSize() int {
	return saltSize
}

func HashAlg() string {
	return hashAlg
}

func AllowedChars() string {
	return allowedChars
}

func UnusableHash() string {
	return unusableHash
}

func MakePassword(password string) (string, error) {
	if password == "" {
		return unusableHash, nil
	}
	return makePbkdf2(password, sha256.Size, sha256.New)
}

func makePbkdf2(password string, keyLen int, h func() hash.Hash) (string, error) {
	salt := getRandomSalt(saltSize)
	dk := pbkdf2.Key([]byte(password), salt, iter, keyLen, h)
	b64Hash := base64.StdEncoding.EncodeToString(dk)
	return fmt.Sprintf("%s$%d$%s$%s", hashAlg, iter, salt, b64Hash), nil
}

func getRandomSalt(size int) []byte {
	salt := make([]byte, size)
	l := len(allowedChars)
	for i := range salt {
		salt[i] = allowedChars[rand.Intn(l)]
	}
	return salt
}
