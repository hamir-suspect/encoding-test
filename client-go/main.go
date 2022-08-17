package main

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/proto"

	"github.com/hamir-suspect/grpc-client-go/api/github.com/hamir-suspect/grpc-client-go/pb"

	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("endpoint", "localhost:50051", "The server address in the format of host:port")
	requestID  = flag.String("request", "", "The request token")
)

func main() {
	fmt.Println("hello world")

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewServerClient(conn)

	msg := &pb.SomeMsg{
		Name:   "something",
		Code:   pb.SomeMsg_OK,
		Number: 23,
	}

	fmt.Println(msg)

	encoded_msg, err := proto.Marshal(msg)
	if err != nil {
		log.Fatalf("could not encode msg", err.Error())
	}
	fmt.Println("encoded msg: ", encoded_msg)

	//	load key from file
	key, err := pubKey()
	if err != nil {
		log.Fatalf("fail to load key: ", err.Error())
	}

	// encrypt data and send

	encrypted := EncryptWithPublicKey(encoded_msg, key)

	base64Encrypted := base64.URLEncoding.EncodeToString(encrypted)
	//base64Encrypted := string(encrypted)

	request := &pb.CreateRequest{EncryptedValue: base64Encrypted}

	log.Println("encrypted: ", encrypted)

	response, err := client.Create(context.Background(), request)
	if err != nil {
		log.Fatalf("%v", err)
	}

	// decrypt data
	log.Printf("Response: \n%+v", response)

	priv, err := privKey()
	if err != nil {
		log.Fatalf("fail to load priv key: ", err.Error())
	}

	base64decoded, err := base64.URLEncoding.DecodeString(response.GetEncryptedValue())
	if err != nil {
		log.Fatalf("failed to decode", err.Error())
	}
	//base64decoded := []byte(response.GetEncryptedValue())

	decrypted := DecryptWithPrivateKey(base64decoded, priv)
	unEncodedResponse := &pb.SomeMsg{}
	proto.Unmarshal(decrypted, unEncodedResponse)

	log.Println("response: ", unEncodedResponse)
}

func privKey() (*rsa.PrivateKey, error) {
	raw, err := ioutil.ReadFile("priv.pem")
	if err != nil {
		return nil, errors.New("key file not found")
	}

	block, rest := pem.Decode(raw)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		log.Fatal("failed to decode PEM block containing public key")
	}

	k, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Got a %T, with remaining data: %q \n", k, rest)

	return k, nil
}

func pubKey() (*rsa.PublicKey, error) {
	raw, err := ioutil.ReadFile("pub.pem")
	if err != nil {
		return nil, errors.New("key file not found")
	}

	block, rest := pem.Decode(raw)
	if block == nil || block.Type != "RSA PUBLIC KEY" {
		log.Fatal("failed to decode PEM block containing public key")
	}

	k, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Got a %T, with remaining data: %q \n", k, rest)

	return k, nil
}

func EncryptWithPublicKey(msg []byte, pub *rsa.PublicKey) []byte {
	en, err := rsa.EncryptPKCS1v15(rand.Reader, pub, msg)
	if err != nil {
		log.Fatal("not encrypted: ", err.Error())
	}
	return en
}

func DecryptWithPrivateKey(msg []byte, priv *rsa.PrivateKey) []byte {
	en, err := rsa.DecryptPKCS1v15(rand.Reader, priv, msg)
	if err != nil {
		log.Fatal("not encrypted: ", err.Error())
	}
	return en
}

//	hash := sha512.New()
//	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, pub, msg, nil)
//	if err != nil {
//		log.Println(err)
//	}
//	return ciphertext
//}
