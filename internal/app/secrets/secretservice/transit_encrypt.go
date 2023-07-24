package secretservice

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/pixconf/pixconf/internal/encrypt"
	"github.com/pixconf/pixconf/internal/protos"
)

func (s *Service) TransitEncrypt(ctx context.Context, request *protos.SecretTransitEncryptRequest) (*protos.SecretTransitEncryptResponse, error) {
	epoch, err := s.db.GetCurrentEpoch(ctx, true)
	if err != nil {
		return nil, err
	}

	encrypted, err := s.db.Encrypt.Encrypt([]byte(request.GetData()))
	if err != nil {
		return nil, err
	}

	response := &protos.SecretTransitEncryptResponse{
		Data: fmt.Sprintf("secrets:v1:%s:%s", strconv.FormatInt(epoch.ID, 16), encrypt.EncodeToString(encrypted)),
	}

	return response, nil
}

func (s *Service) TransitDecrypt(ctx context.Context, request *protos.SecretTransitDecryptRequest) (*protos.SecretTransitDecryptResponse, error) {
	if !strings.HasPrefix(request.GetData(), "secrets:v1:") {
		return nil, errors.New("wrong encrypted format")
	}

	splited := strings.Split(request.GetData(), ":")
	if len(splited) != 4 {
		return nil, errors.New("wrong encrypted format")
	}

	epoch, err := strconv.ParseInt(splited[2], 16, 64)
	if err != nil {
		return nil, err
	}

	rawData, err := encrypt.DecodeFromString(splited[3])
	if err != nil {
		return nil, err
	}

	epochInfo, err := s.db.GetEpoch(ctx, epoch)
	if err != nil {
		return nil, err
	}

	if epochInfo == nil {
		return nil, errors.New("epoch not found")
	}

	enc, err := encrypt.New(epochInfo.PrivateKey, epochInfo.EncryptionType)
	if err != nil {
		return nil, err
	}

	decrypted, err := enc.Decrypt(rawData)
	if err != nil {
		return nil, err
	}

	return &protos.SecretTransitDecryptResponse{
		Data: string(decrypted),
	}, nil
}
