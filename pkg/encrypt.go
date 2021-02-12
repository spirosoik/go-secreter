package encrypt

// Options for encrypt/decrypt
type Options struct {
	// Used only for vault
	Key string

	// Plaintext used for encrypt
	Plaintext string

	// CipherText used for decrypt
	CipherText string
}

// Secreter contract for secret providers
type Secreter interface {
	Encrypt(opts Options) (string, error)
	Decrypt(opts Options) (string, error)
	Type() SecretType
}

// Decrypter allows decryption of secrets
type Decrypter interface {
	Decrypt(opts Options) (string, error)
}

// SecretType the type of the encryption
// secrets will be stored
type SecretType int

const (
	// LocalSecret encryption with base64
	LocalSecret SecretType = iota
	// VaultSecret secrets with hashicorp Vault
	VaultSecret
	// DebugSecret secrets just plain
	DebugSecret
)

func (s SecretType) String() string {
	return [...]string{"local", "vault", "debug"}[s]
}
