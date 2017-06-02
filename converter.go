package iso8583

import (
	"encoding/hex"
	"fmt"
	"strings"
)

// BitmapToBinary converts bitmap string to equivalent binary
func BitmapToBinary(bitmap string) (string, error) {
	h := strings.Replace(bitmap, " ", "", -1)
	return HexToBinary(h)
}

// HexToBinary converts a hexadecimal number string to it's equivalent binary string representation
func HexToBinary(s string) (string, error) {
	b, err := hex.DecodeString(s)
	if err != nil {
		return "", fmt.Errorf("invalid hex: %v", err)
	}
	out := fmt.Sprintf("%.8b", b)
	out = strings.Replace(out, " ", "", -1)
	return strings.Trim(out, "[]"), nil
}
