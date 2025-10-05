package password

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"hash"
	"strconv"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

func CheckPassword(password, encoded string) (bool, error) {
	if strings.Index(encoded, unusableHash) == 0 {
		_, _ = MakePassword(password)
		return false, nil
	}
	return checkPbkdf2(password, encoded, sha256.Size, sha256.New)
}

func checkPbkdf2(password, encoded string, keyLen int, h func() hash.Hash) (bool, error) {
	parts := strings.SplitN(encoded, "$", 4)
	if len(parts) != 4 {
		return false, errors.New("hash must consist of 4 segments")
	}
	iter, err := strconv.Atoi(parts[1])
	if err != nil {
		return false, fmt.Errorf("wrong number of iterations: %v", err)
	}
	salt := []byte(parts[2])
	k, err := base64.StdEncoding.DecodeString(parts[3])
	if err != nil {
		return false, fmt.Errorf("wrong hash encoding: %v", err)
	}
	dk := pbkdf2.Key([]byte(password), salt, iter, keyLen, h)
	return bytes.Equal(k, dk), nil
}
