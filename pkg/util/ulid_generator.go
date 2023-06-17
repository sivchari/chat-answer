//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
//go:generate goimports -w --local "github.com/sivchari/chat-answer" mock_$GOPACKAGE/mock_$GOFILE
package util

import (
	"crypto/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

type ULIDGenerator interface {
	Generate() (string, error)
}

type ulidGenerator struct{}

func NewUILDGenerator() ULIDGenerator {
	return &ulidGenerator{}
}

func (*ulidGenerator) Generate() (string, error) {
	id, err := ulid.New(ulid.Timestamp(time.Unix(1000000, 0)), rand.Reader)
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
