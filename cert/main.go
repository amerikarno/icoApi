package main

import (
	"crypto/tls"
	"log"
)

func main() {
	cert, err := tls.LoadX509KeyPair("myCA.pem", "myCA.crt")
	if err != nil {
		log.Fatal(err)
	}
	cfg := &tls.Config{Certificates: []tls.Certificate{cert}}
	listener, err := tls.Listen("tcp", ":2000", cfg)
	if err != nil {
		log.Fatal(err)
	}
	_ = listener
}