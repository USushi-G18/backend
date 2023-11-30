package auth

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

var (
	ErrWrongPassword = errors.New("wrong password")
	ErrInvalidHash   = errors.New("invalid hash")
)

type Argon2Params struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

var StdArgon2Params *Argon2Params = &Argon2Params{
	Memory:      64 * 1024,
	Iterations:  3,
	Parallelism: 4,
	SaltLength:  16,
	KeyLength:   32,
}

// u-sushi hash
// $argon2id$v=19$m=65536,t=3,p=4$vb4AULkt8aSGh/Rfq5BwOQ$jzmrQ/jv6I6e2MZhfejoIvUzaarV4874fgZ6cJS6eF4
func HashPassword(password string, p *Argon2Params) (string, error) {
	wrapErr := func(err error) error {
		return fmt.Errorf("hash password: %v", err)
	}
	salt := make([]byte, p.SaltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return "", wrapErr(err)
	}
	hash := argon2.IDKey([]byte(password), salt, p.Iterations, p.Memory, p.Parallelism, p.KeyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, p.Memory, p.Iterations, p.Parallelism, b64Salt, b64Hash)
	return encodedHash, nil
}

func VerifyPassword(password string, encodedHash string) (bool, error) {
	wrapErr := func(err error) error {
		return fmt.Errorf("verify password: %v", err)
	}
	p, salt, hash, err := DecodeHash(encodedHash)
	if err != nil {
		return false, wrapErr(err)
	}

	otherHash := argon2.IDKey([]byte(password), salt, p.Iterations, p.Memory, p.Parallelism, p.KeyLength)

	// subtle.ConstantTimeCompare() function to help prevent timing attacks.
	return subtle.ConstantTimeCompare(hash, otherHash) == 1, nil
}

func DecodeHash(encodedHash string) (p *Argon2Params, salt []byte, hash []byte, err error) {
	wrapErr := func(err error) error {
		return fmt.Errorf("decode hash: %v", err)
	}
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, wrapErr(ErrInvalidHash)
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, wrapErr(ErrInvalidHash)
	}

	p = &Argon2Params{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.Memory, &p.Iterations, &p.Parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	p.SaltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.KeyLength = uint32(len(hash))

	return p, salt, hash, nil
}
