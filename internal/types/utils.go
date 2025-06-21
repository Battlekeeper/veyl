package types

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"log"
)

func DecodeRelayAuth(body []byte) (*RelayAuth, error) {
	var auth RelayAuth
	err := json.Unmarshal(body, &auth)
	if err != nil {
		return nil, err
	}
	return &auth, nil
}

func Base64ToHex(base64Key string) string {
	decodedKey, err := base64.StdEncoding.DecodeString(base64Key)
	if err != nil {
		log.Panic("Failed to decode base64 key:", err)
	}
	hexKey := hex.EncodeToString(decodedKey)
	return hexKey
}

func HexToBase64(hexKey string) (string, error) {
	decodedKey, err := hex.DecodeString(hexKey)
	if err != nil {
		return "", err
	}
	base64Key := base64.StdEncoding.EncodeToString(decodedKey)
	return base64Key, nil
}

func GenerateAuthenticationToken(length int) string {
	token := make([]byte, length)
	_, err := rand.Read(token)
	if err != nil {
		log.Panic("Failed to generate authentication token:", err)
	}
	return base64.StdEncoding.EncodeToString(token)
}
