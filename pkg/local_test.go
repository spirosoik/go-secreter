package encrypt

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type localServiceSuite struct {
	suite.Suite

	encService Secreter
}

func (s *localServiceSuite) SetupSuite() {
	svc, _ := newLocalService(context.Background(), "smGbjm71Nxd1Ig5FS0wj9SlbzAIrnolCz9bQQ6uAhl4=")
	s.encService = svc
}

func TestLocalInit(t *testing.T) {
	suite.Run(t, new(localServiceSuite))
}

func (s *localServiceSuite) TestEncrypt() {
	opts := Options{
		Plaintext: "test",
	}
	_, err := s.encService.Encrypt(opts)

	assert.Equal(s.T(), nil, err)
}

func (s *localServiceSuite) TestDecrypt() {
	opts := Options{
		Plaintext: "test",
	}
	secret, _ := s.encService.Encrypt(opts)
	plaintext, err := s.encService.Decrypt(Options{
		CipherText: secret,
	})

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), opts.Plaintext, plaintext)
}
