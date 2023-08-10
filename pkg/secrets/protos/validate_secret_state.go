package protos

import (
	"golang.org/x/exp/slices"
)

type SecretState string

func (s SecretState) String() string {
	return string(s)
}

const (
	SecretStateNormal  SecretState = "normal"
	SecretStateHidden  SecretState = "hidden"
	SecretStateDeleted SecretState = "deleted"
)

func IsValidSecretState(name string) bool {
	return slices.Contains([]string{
		string(SecretStateNormal),
		string(SecretStateHidden),
		string(SecretStateDeleted),
	}, name)
}
