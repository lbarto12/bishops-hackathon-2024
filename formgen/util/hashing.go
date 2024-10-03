package util

import (
	"fmt"
	"hash/fnv"
)

func HashFunc(s string) string {
	h := fnv.New64a()
	_, _ = h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum64())
}
