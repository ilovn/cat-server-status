package util

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"time"
)

func Token(t string) (token string) {
	crutime := time.Now().Unix()
	h := md5.New()
	if len(t) <= 0 {
		io.WriteString(h, strconv.FormatInt(crutime, 10))
	} else {
		io.WriteString(h, t)
	}
	token = fmt.Sprintf("%x", h.Sum(nil))
	return
}
