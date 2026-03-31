package verify

import "errors"

type JSONStorage map[string]string

func NewJSONStorage() JSONStorage {
	return make(JSONStorage)
}

func (j JSONStorage) Save(email string, hash string) error {
	j[hash] = email
	return nil
}

func (j JSONStorage) Find(hash string) (string, error) {
	if email, exists := j[hash]; exists {
		return email, nil
	}
	return "", errors.New("email not found for hash: " + hash)
}

func (j JSONStorage) Delete(hash string) error {
	delete(j, hash)
	return nil
}
