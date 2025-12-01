package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func main() {
	i := 1

	for {
		hash := md5.New()
		hash.Write([]byte(fmt.Sprintf("iwrupvqb%d", i)))

		result := hex.EncodeToString(hash.Sum(nil))
		if strings.HasPrefix(result, "000000") {
			fmt.Println(i)
			break
		}

		i++
	}
}
