package core

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"testing"
)

func TestSigner_Sign(t *testing.T) {
	msg := "abc123/refresh/1568953410"
	h := hmac.New(sha1.New, []byte("sk"))
	h.Write([]byte(msg))
	sign := h.Sum(nil)
	fmt.Println("sign:", fmt.Sprintf("%x", sign))

}
