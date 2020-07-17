package articletest

import (
	"encoding/base64"
	"fmt"
)

func Encode(name string) string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%v", name)))
}

func Decode(cursor string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}
