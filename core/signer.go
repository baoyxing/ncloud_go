package core

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"strconv"
	"time"
)

type Signer struct {
	Credentials Credential
}

func NewSigner(credsProvider Credential) *Signer {
	return &Signer{
		Credentials: credsProvider,
	}
}

func (s *Signer) Sign(path string) (signer, timestamp string) {
	t := time.Now().Unix()
	tStr := strconv.FormatInt(t, 10)
	msg := fmt.Sprintf("%s%s/%s", s.Credentials.AccessKey, path, tStr)
	h := hmac.New(sha1.New, []byte(s.Credentials.SecretKey))
	h.Write([]byte(msg))
	sign := h.Sum(nil)
	return fmt.Sprintf("%x", sign), tStr
}
