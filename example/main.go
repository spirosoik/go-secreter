package main

import (
	"context"
	"log"

	encrypt "github.com/spirosoik/go-secreter/pkg"
)

func main() {
	secreter, err := encrypt.New(context.Background(), encrypt.Config{
		SecretType: encrypt.LocalSecret,
		SecretKey:  "fyktabT5I8fFK-mkSbbxIsfsbcnP-4QFa5awWmyuGqs=",
	})
	if err != nil {
		log.Fatal(err)
	}

	secret, err := secreter.Encrypt(encrypt.Options{
		Plaintext: "sensitive",
	})

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Cipher Text: %s", secret)

	plaintext, err := secreter.Decrypt(encrypt.Options{
		CipherText: secret,
	})

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Plaintext: %s", plaintext)
}
