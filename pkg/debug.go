package encrypt

import "context"

type debugEncrSvc struct {
	ctx context.Context
}

// NewLocalService factory method
func newDebugService(ctx context.Context) (Secreter, error) {
	return &debugEncrSvc{
		ctx: ctx,
	}, nil
}

func (s *debugEncrSvc) Encrypt(opts Options) (string, error) {
	return opts.Plaintext, nil
}

func (s *debugEncrSvc) Decrypt(opts Options) (string, error) {
	return opts.CipherText, nil
}

func (s *debugEncrSvc) Type() SecretType {
	return DebugSecret
}
