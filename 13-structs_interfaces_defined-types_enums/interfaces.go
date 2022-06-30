package main

type Encrypter interface {
	Encrypt([]byte) string
}

type Decrypter interface {
	Decrypt(string, interface{}) error
}

func send(encrypter Encrypter) {
	// method implementation
}

func receive(decrypter Decrypter) {
	// method implementation
}
