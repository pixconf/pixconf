package authkey

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pixconf/pixconf/internal/buildinfo"
)

var (
	Base64Encoding = base64.RawURLEncoding
)

type AuthHeader struct {
	Algorithm string `json:"alg"`
	PublicKey string `json:"pk"`
}

type AuthPayload struct {
	Issuer   string `json:"iss"`
	JwtID    string `json:"jti"`
	IssuedAT int64  `json:"iat"`
	Version  string `json:"ver"`
}

func (a *AuthKey) generateAuthHeader() (string, error) {
	header := AuthHeader{
		Algorithm: "ed25519",
		PublicKey: Base64Encoding.EncodeToString(a.pub),
	}

	headerJSON, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	return Base64Encoding.EncodeToString(headerJSON), nil
}

func (a *AuthKey) generateAuthPayload(agentID string) (string, error) {
	payload := AuthPayload{
		Issuer:   agentID,
		JwtID:    uuid.New().String(),
		IssuedAT: time.Now().Unix(),
		Version:  buildinfo.Version,
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return "", nil
	}

	return Base64Encoding.EncodeToString(payloadJSON), nil
}

func (a *AuthKey) GenerateAuthKey(agentID string) ([]byte, error) {
	encodedHeader, err := a.generateAuthHeader()
	if err != nil {
		return nil, fmt.Errorf("failed to generate auth header: %w", err)
	}

	encodedPayload, err := a.generateAuthPayload(agentID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate auth payload: %w", err)
	}

	dataToSign := fmt.Sprintf("%s.%s", encodedHeader, encodedPayload)

	signature := a.Sign([]byte(dataToSign))
	encodedSignature := Base64Encoding.EncodeToString(signature)

	return []byte(fmt.Sprintf("%s.%s", dataToSign, encodedSignature)), nil
}
