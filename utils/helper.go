package utils

import (
	"log"
	"unsafe"
)

// CheckErrFatal checar o erro
func CheckErrFatal(err error, msg string) {
	if err != nil {
		log.Printf("CheckErr(): %q\n", err)
		log.Fatalf("%s: %s", msg, err)
	}
}

// CheckErr checar o erro
func CheckErr(err error, msg string) string {
	mensagem := ""

	if err != nil {
		log.Printf("CheckErr(): %s: %s", msg, err)
	}

	return mensagem
}

//BytesToString converter bytes para string
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
