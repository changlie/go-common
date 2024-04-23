package codes

import (
	"encoding/base64"
	"encoding/hex"
)

func Hex(raw []byte) string {
	return hex.EncodeToString(raw)
}

func HexEncode(raw string) string {
	return hex.EncodeToString([]byte(raw))
}

func HexRaw(encode string) []byte {
	bytes, err := hex.DecodeString(encode)
	if err != nil {
		return nil
	}
	return bytes
}

func Base64(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}
func Base64Encode(src string) string {
	return base64.StdEncoding.EncodeToString([]byte(src))
}
func Base64Raw(src string) []byte {
	bytes, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return nil
	}
	return bytes
}
