package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/segmentio/ksuid"
)

func GenerateID() string {
	//"github.com/segmentio/ksuid"
	//return ksuid.New().String()
	u, _ := uuid.DefaultGenerator.NewV4()
	return u.String()
}

func GenerateLogID() string {
	//"github.com/segmentio/ksuid"
	return ksuid.New().String()

}

func Base64Decode(str string) []byte {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return data
	}
	return data
}

func Base64Encode(b []byte) string {
	data := base64.StdEncoding.EncodeToString(b)

	return data
}

func BytesToPublicKey(pub []byte) *rsa.PublicKey {
	block, _ := pem.Decode(pub)
	fmt.Println("Decode   ", block)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	fmt.Println(b)
	var err error
	if enc {
		Log.Info("is encrypted pem block")
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			Log.Error(err)
		}
	}
	ifc, err := x509.ParsePKIXPublicKey(b)
	if err != nil {
		Log.Error(err)
	}

	fmt.Println("ifc   ", ifc)
	key, ok := ifc.(*rsa.PublicKey)
	if !ok {
		Log.Error("not ok")
	}

	fmt.Println("key  ", key)
	return key
}

// Decode public key struct from PEM string
func ExportPEMStrToPubKey(pub []byte) *rsa.PublicKey {
	block, _ := pem.Decode(pub)
	key, _ := x509.ParsePKCS1PublicKey(block.Bytes)
	return key
}
