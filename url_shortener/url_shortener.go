package url_shortener

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/big"
)

func GenerateLink(userUrl string) string {
	hashBytes := sha256url(userUrl)
	generatedNumber := new(big.Int).SetBytes(hashBytes).Uint64()
	return encodeBase64([]byte(fmt.Sprintf("%d", generatedNumber)))
}

func sha256url(userUrl string) []byte {
	hash := sha256.New()
	hash.Write([]byte(userUrl))
	return hash.Sum(nil) //get hash for data by Sum function, to append nothing append nil
}

func encodeBase64(bytes []byte) string {
	sEnc := base64.StdEncoding.EncodeToString(bytes)
	return sEnc
}
