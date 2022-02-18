package main

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func extractData(extension pkix.Extension) (data []byte, err error) {
	raw := extension.Value
	_, err = asn1.Unmarshal(raw, &data)
	return data, err
}

func main() {
	var bytes []byte
	if _, err := os.Stat("cert.pem"); err == nil {
		bytes, err = ioutil.ReadFile("cert.pem")
		if err != nil {
			log.Fatalf("%v", err)
		}
	} else {
		fmt.Println("Please download https://raw.githubusercontent.com/spbkaizo/nevergonna/main/cert.pem to current directory")
		os.Exit(1)
	}

	block, _ := pem.Decode(bytes)
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		log.Fatalf("%v", err)
	}
	for _, v := range cert.Extensions {

		switch v.Id.String() {
		case "1.3.6.1.5.5.7.13.37":
			log.Printf("Rick's never gonna give you up...")
			rick, err := extractData(v)
			if err != nil {
				log.Fatalf("FATAL: %v", err)
			}
			file, err := ioutil.TempFile("./", "rick*.mp3")
			if err != nil {
				log.Fatal(err)
			}
			if _, err := file.Write(rick); err != nil {
				file.Close()
				log.Fatal(err)
			}
			if err := file.Close(); err != nil {
				log.Fatal(err)
			}

		}
	}
}
