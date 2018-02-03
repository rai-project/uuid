package uuid

import (
	"crypto/sha1"
	"encoding/hex"

	"github.com/davecgh/go-spew/spew"
)

func NewV4() string {
	u4 := Must(global.NewV4())
	return u4.String()
}

func NewV5(u, name string) string {
	u5 := global.NewV5(FromStringOrNil(u), name)
	return u5.String()
}

func New(obj interface{}) string {
	u4 := Must(global.NewV4())
	h := sha1.New()
	spew.Fdump(h, obj)
	u5 := global.NewV5(u4, hex.EncodeToString(h.Sum(nil)))
	return u5.String()
}
