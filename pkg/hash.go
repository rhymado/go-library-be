package pkg

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type HashConfig struct {
	Time    uint32
	Memory  uint32
	Threads uint8
	KeyLen  uint32
	SaltLen uint32
}

func InitHashConfig() *HashConfig {
	return &HashConfig{}
}

func (h *HashConfig) UseDefaultConfig() *HashConfig {
	return &HashConfig{
		Time:    3,
		Memory:  64 * 1024,
		Threads: 2,
		KeyLen:  32,
		SaltLen: 16,
	}
	// h.Time = 3
	// h.Memory = 64 * 1024
	// h.Threads = 2
	// h.KeyLen = 32
	// h.SaltLen = 16
}

func (h *HashConfig) UseConfig(threads uint8, time, memory, keyLen, saltLen uint32) *HashConfig {
	return &HashConfig{
		Time:    time,
		Memory:  memory,
		Threads: threads,
		KeyLen:  keyLen,
		SaltLen: saltLen,
	}
	// h.Time = time
	// h.Memory = memory
	// h.Threads = threads
	// h.KeyLen = keyLen
	// h.SaltLen = saltLen
}

func (h *HashConfig) genSalt() ([]byte, error) {
	// salt = byte acak sepanjang saltLen
	b := make([]byte, h.SaltLen)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}
	return b, nil
}

func (h *HashConfig) GenHashedPassword(password string) (string, error) {
	salt, err := h.genSalt()
	if err != nil {
		return "", err
	}
	hash := argon2.IDKey([]byte(password), salt, h.Time, h.Memory, h.Threads, h.KeyLen)

	version := argon2.Version
	base64Salt := base64.RawStdEncoding.EncodeToString(salt)
	base64Hash := base64.RawStdEncoding.EncodeToString(hash)

	hashedPassword := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", version, h.Memory, h.Time, h.Threads, base64Salt, base64Hash)
	return hashedPassword, nil
}

func (h *HashConfig) ComparePasswordAndHash(password string, hashedPassword string) (bool, error) {
	salt, hash, err := h.decodeHash(hashedPassword)
	if err != nil {
		return false, err
	}

	// gen hash dari password yang datang
	userHash := argon2.IDKey([]byte(password), salt, h.Time, h.Memory, h.Threads, h.KeyLen)
	if subtle.ConstantTimeCompare(userHash, hash) == 0 {
		return false, nil
	}
	return true, nil
}

func (h *HashConfig) decodeHash(hashedPassword string) (salt []byte, hash []byte, err error) {
	// cek format hash
	values := strings.Split(hashedPassword, "$")
	if len(values) != 6 {
		return nil, nil, errors.New("invalid hashed password format")
	}

	// cek versi argon2id
	var version int
	if _, err := fmt.Sscanf(values[2], "v=%d", &version); err != nil {
		fmt.Println("version error")
		return nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, errors.New("invalid argon2id version")
	}

	// penuhi hashConfig
	if _, err := fmt.Sscanf(values[3], "m=%d,t=%d,p=%d", &h.Memory, &h.Time, &h.Threads); err != nil {
		fmt.Println("config error")
		return nil, nil, err
	}

	salt, err = base64.RawStdEncoding.DecodeString(values[4])
	if err != nil {
		return nil, nil, err
	}
	h.SaltLen = uint32(len(salt))

	hash, err = base64.RawStdEncoding.DecodeString(values[5])
	if err != nil {
		return nil, nil, err
	}
	h.KeyLen = uint32(len(hash))

	// return salt dan hash
	return salt, hash, nil
}
