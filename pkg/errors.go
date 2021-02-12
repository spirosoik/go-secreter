package encrypt

import "errors"

// ErrNotSupportedSecret an error which is occurred if the type
// of secreter is missing
var ErrNotSupportedSecret = errors.New("secret type is required")

// ErrSecretKeyMissing an error which is occurred if the type
// of secret key for encrytpion is missing
var ErrSecretKeyMissing = errors.New("secret key is required")

// ErrVaultHostMissing an error which is occurred if the vault is enabled
// and the Vault host is missing
var ErrVaultHostMissing = errors.New("vault host is required")

// ErrVaultTokenMissing an error which is occurred if the vault is enabled
// and the Vault token is missing
var ErrVaultTokenMissing = errors.New("vault token is required")
