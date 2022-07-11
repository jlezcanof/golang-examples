package main

type Encrypter interface {
	Encrypt([]byte) string
}

type Decrypter interface {
	Decrypt(string, interface{}) error
}

type EncrypDecrypter interface {
	Encrypter
	Decrypter
}
