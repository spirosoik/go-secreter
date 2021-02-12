package encrypt

import (
	"context"
)

// VaultConfig it is used in order to be able to interact
// with transit engine of vault. Token is the one which is needed
// in order to interact with Vault API
type VaultConfig struct {
	Enabled bool
	Token   string
	Host    string
}

// Config config for secreter
type Config struct {
	SecretType  SecretType
	SecretKey   string
	VaultConfig VaultConfig
}

// validate checks if the required configuration
// is there in order to be used.
func (c Config) validate() error {
	if c.SecretKey == "" {
		return ErrSecretKeyMissing
	}

	if !c.VaultConfig.Enabled {
		return nil
	}
	if c.VaultConfig.Host == "" {
		return ErrVaultHostMissing
	}

	if c.VaultConfig.Token == "" {
		return ErrVaultTokenMissing
	}
	return nil
}

// New factory method to get the proper secreter
func New(ctx context.Context, cfg Config) (Secreter, error) {
	switch cfg.SecretType {
	case VaultSecret:
		return newVaultService(ctx, cfg)
	case LocalSecret:
		return newLocalService(ctx, cfg.SecretKey)
	case DebugSecret:
		return newDebugService(ctx)
	}
	return nil, ErrNotSupportedSecret
}
