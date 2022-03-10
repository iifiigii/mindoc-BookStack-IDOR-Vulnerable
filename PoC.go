package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"time"
)

type CookieRemember struct {
	MemberId int
	Account  string
	Time     time.Time
}

func main() {
	//mindoc
	//m := CookieRemember{MemberId: 1, Account: "2"}
	//var network bytes.Buffer
	//enc := gob.NewEncoder(&network)
	//enc.Encode(m)
	//encoded := base64.URLEncoding.EncodeToString([]byte(network.String()))
	//fmt.Println("base64编码，得到vs:", encoded)
	//s := "mindoc"
	//h := hmac.New(sha256.New, []byte(s))
	//fmt.Fprintf(h, "%s%s", encoded, "123")
	//a := fmt.Sprintf("%02x", h.Sum(nil))
	//fmt.Println("sha256加密，得到sig:", a)
	//bookstock
	m := CookieRemember{MemberId: 1, Account: "2"}
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	enc.Encode(m)
	encoded := base64.URLEncoding.EncodeToString([]byte(network.String()))
	fmt.Println("base64编码，得到vs:", encoded)
	s := "godoc"
	h := hmac.New(sha1.New, []byte(s))
	fmt.Fprintf(h, "%s%s", encoded, "123")
	a := fmt.Sprintf("%02x", h.Sum(nil))
	fmt.Println("sha1加密，得到sig:", a)
}
