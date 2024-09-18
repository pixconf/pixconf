package authkey

import (
	"crypto/ed25519"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

func Unmarshal(data string) (*AuthHeader, *AuthPayload, error) {
	payloadSplited := strings.Split(data, ".")
	if len(payloadSplited) != 3 {
		return nil, nil, errors.New("invalid payload")
	}

	header, err := unmarshalHeader(payloadSplited[0])
	if err != nil {
		return nil, nil, err
	}

	payload, err := unmarshalPayload(payloadSplited[1])
	if err != nil {
		return nil, nil, err
	}

	signRaw, err := Base64Encoding.DecodeString(payloadSplited[2])
	if err != nil {
		return nil, nil, err
	}

	signedData := fmt.Sprintf("%s.%s", payloadSplited[0], payloadSplited[1])

	if err := unmarshalCheckSign(header.PublicKey, []byte(signedData), signRaw); err != nil {
		return nil, nil, err
	}

	return header, payload, nil
}

func unmarshalHeader(data string) (*AuthHeader, error) {
	headerJSON, err := Base64Encoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	header := &AuthHeader{}

	if err := json.Unmarshal(headerJSON, &header); err != nil {
		return nil, err
	}

	return header, nil
}

func unmarshalPayload(data string) (*AuthPayload, error) {
	payloadJSON, err := Base64Encoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	payload := &AuthPayload{}

	if err := json.Unmarshal(payloadJSON, &payload); err != nil {
		return nil, err
	}

	return payload, nil
}

func unmarshalCheckSign(pubKey string, signedData, signRaw []byte) error {
	ed25519PublicKeyRaw, err := Base64Encoding.DecodeString(pubKey)
	if err != nil {
		return err
	}

	if len(ed25519PublicKeyRaw) != ed25519.PublicKeySize {
		return errors.New("invalid public key size")
	}

	ed25519PublicKey := ed25519.PublicKey(ed25519PublicKeyRaw)

	if !ed25519.Verify(ed25519PublicKey, signedData, signRaw) {
		return errors.New("invalid signature")
	}

	return nil
}
