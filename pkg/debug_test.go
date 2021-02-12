package encrypt

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type debugServiceSuite struct {
	suite.Suite

	encService Secreter
}

func (s *debugServiceSuite) SetupSuite() {
	svc, _ := newDebugService(context.Background())
	s.encService = svc
}

func TestDebugInit(t *testing.T) {
	suite.Run(t, new(debugServiceSuite))
}

func (s *debugServiceSuite) TestEncrypt() {
	opts := Options{
		Plaintext: "test",
	}
	secret, err := s.encService.Encrypt(opts)

	assert.Equal(s.T(), nil, err)
	assert.Equal(s.T(), opts.Plaintext, secret)
}

func (s *debugServiceSuite) TestDecrypt() {
	opts := Options{
		CipherText: "hash",
	}
	secret, err := s.encService.Decrypt(opts)

	assert.Equal(s.T(), nil, err)
	assert.Equal(s.T(), opts.CipherText, secret)
}
