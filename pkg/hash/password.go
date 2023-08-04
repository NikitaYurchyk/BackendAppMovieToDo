package hash

import "crypto"

type SHA256Hasher struct {
	word string
}

func NewSHA256Hasher(_word string) *SHA256Hasher {
	return &SHA256Hasher{
		word: _word,
	}
}

func (h *SHA256Hasher) Hash(password string) (string, error) {
	hash := crypto.SHA256.New()

	if _, err := hash.Write([]byte(password)); err != nil {
		return "", err
	}

	return string(hash.Sum([]byte(h.word))), nil
}
