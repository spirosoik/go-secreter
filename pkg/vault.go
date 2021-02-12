package encrypt

import (
	"context"

	vaultapi "github.com/hashicorp/vault/api"
	"github.com/pkg/errors"
	"gocloud.dev/secrets/hashivault"
	_ "gocloud.dev/secrets/hashivault" // import local secret
)

type vaultEncrSvc struct {
	client *vaultapi.Client
	ctx    context.Context
	cfg    Config
}

// NewVaultService factory method
func newVaultService(ctx context.Context, cfg Config) (Secreter, error) {
	client, err := hashivault.Dial(ctx, &hashivault.Config{
		Token: cfg.VaultConfig.Token,
		APIConfig: vaultapi.Config{
			Address: cfg.VaultConfig.Token,
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "hashivault.Dial")
	}
	return &vaultEncrSvc{
		client: client,
		ctx:    ctx,
		cfg:    cfg,
	}, nil
}

func (s *vaultEncrSvc) Encrypt(opts Options) (string, error) {
	keeper := hashivault.OpenKeeper(s.client, s.cfg.SecretKey, nil)
	defer keeper.Close()

	secret, err := keeper.Encrypt(s.ctx, []byte(opts.Plaintext))
	if err != nil {
		return "", errors.Wrap(err, "keeper.Encrypt")
	}
	return string(secret), nil
}

func (s *vaultEncrSvc) Decrypt(opts Options) (string, error) {
	keeper := hashivault.OpenKeeper(s.client, s.cfg.SecretKey, nil)
	defer keeper.Close()

	decrypted, err := keeper.Decrypt(s.ctx, []byte(opts.CipherText))
	if err != nil {
		return "", errors.Wrap(err, "keeper.Decrypt")
	}
	return string(decrypted), nil
}

func (s *vaultEncrSvc) Type() SecretType {
	return VaultSecret
}
