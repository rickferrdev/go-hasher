package hasher

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const (
	MinCost     = bcrypt.MinCost
	MaxCost     = bcrypt.MaxCost
	DefaultCost = bcrypt.DefaultCost
)

type Hasher interface {
	Generate(password []byte) ([]byte, error)
	Compare(hash, password []byte) error
}

type bcryptHasher struct {
	cost int
}

func New(cost int) Hasher {
	if cost < MinCost || cost > MaxCost {
		cost = DefaultCost
	}

	return &bcryptHasher{cost: cost}
}

func (b *bcryptHasher) Generate(password []byte) ([]byte, error) {
	if len(password) > 72 {
		return nil, bcrypt.ErrPasswordTooLong
	}

	hash, err := bcrypt.GenerateFromPassword(password, b.cost)
	if err != nil {
		return nil, fmt.Errorf("hasher: failed to generate hash: %w", err)
	}

	return hash, nil
}

func (b *bcryptHasher) Compare(hash, password []byte) error {
	err := bcrypt.CompareHashAndPassword(hash, password)
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return errors.New("hasher: invalid password")
		}
		return fmt.Errorf("hasher: comparison failed: %w", err)
	}
	return nil
}
