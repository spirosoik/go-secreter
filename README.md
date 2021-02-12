[![PkgGoDev](https://pkg.go.dev/badge/github.com/spirosoik/go-secreter)](https://pkg.go.dev/github.com/spirosoik/go-secreter)
[![report card](https://img.shields.io/badge/report%20card-a%2B-ff3333.svg?style=flat-square)](http://goreportcard.com/report/spirosoik/go-secreter)


# Secreter
A secreter interface to provide encryption in transit.

## Supports

- Local (simple encryption)
- Vault transit engine

In Progress:
- AWS KMS
- Azure Key Vault

`note: for Local must be a base64-encoded key, of length 32 bytes when decoded`

## Example

```bash
make run-example
```

Encrypt and decrypt sensitive info example:

```golang

// Secreter initialise
secreter, err := encrypt.New(context.Background(), encrypt.Config{
  SecretType: encrypt.LocalSecret,
  SecretKey:  "fyktabT5I8fFK-mkSbbxIsfsbcnP-4QFa5awWmyuGqs=",
})

// Encrypt
secret, err := secreter.Encrypt(encrypt.Options{
  Plaintext: "sensitive",
})

// Decrypt
plaintext, err := secreter.Decrypt(encrypt.Options{
  CipherText: secret,
})
```

Check the full example [here](example/main.go) 

If you want to generate keys for encryption you can use the following [binaries here](key-generator/)