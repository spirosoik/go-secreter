package encrypt

import (
	"context"
	"encoding/base64"
	"fmt"

	"gocloud.dev/secrets"
	_ "gocloud.dev/secrets/localsecrets" // import local secret
)

type localEncrSvc struct {
	ctx    context.Context
	secret string
}

// NewLocalService factory method
func newLocalService(ctx context.Context, secret string) (Secreter, error) {
	return &localEncrSvc{
		ctx:    ctx,
		secret: secret,
	}, nil
}

func (s *localEncrSvc) Encrypt(opts Options) (string, error) {
	keeper, err := secrets.OpenKeeper(s.ctx, fmt.Sprintf("base64key://%s", s.secret))
	if err != nil {
		return "", err
	}
	defer keeper.Close()

	secret, err := keeper.Encrypt(s.ctx, []byte(opts.Plaintext))
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(secret), nil
}

func (s *localEncrSvc) Decrypt(opts Options) (string, error) {
	keeper, err := secrets.OpenKeeper(s.ctx, fmt.Sprintf("base64key://%s", s.secret))
	if err != nil {
		return "", err
	}
	defer keeper.Close()

	decoded, err := base64.URLEncoding.DecodeString(opts.CipherText)
	if err != nil {
		return "", err
	}
	decrypted, err := keeper.Decrypt(s.ctx, []byte(decoded))
	if err != nil {
		return "", err
	}
	return string(decrypted), nil
}

func (s *localEncrSvc) Type() SecretType {
	return LocalSecret
}
