package encrypt

import (
	"reflect"
	"testing"
)

func TestGetValueKey(t *testing.T) {
	validEpochKey := []byte{
		0xc9, 0xa9, 0xaa, 0x74, 0x5f, 0xe0, 0x9a, 0x29, 0x8f, 0x8, 0x8,
		0xa1, 0xc2, 0xbe, 0xba, 0x5a, 0x16, 0x6b, 0x87, 0x54, 0xb1, 0x55,
		0x7f, 0x15, 0xb4, 0x65, 0xe0, 0xc9, 0xad, 0x5d, 0xfc, 0x31,
	}

	testCases := []struct {
		name        string
		epochKey    []byte
		secretID    string
		version     int64
		expectedSum []byte
		expectedErr error
	}{
		{
			name:     "ValidCase",
			epochKey: validEpochKey,
			secretID: "sec-62dyen4w95kv9jhdnp6bbkb3",
			version:  1111,
			expectedSum: []byte{
				0xb2, 0x85, 0xcf, 0xa8, 0xde, 0x72, 0x5d, 0x23, 0x64, 0x7c, 0xb0,
				0x64, 0xf, 0x7c, 0x73, 0xf0, 0x49, 0x79, 0x8c, 0x13, 0x37, 0x2b,
				0x85, 0x5f, 0x3b, 0x1a, 0xe2, 0x9a, 0x1e, 0x40, 0x93, 0x8,
			},
		},
		{
			name:        "InvalidEpochKey",
			epochKey:    nil,
			secretID:    "sec-62dyen4w95kv9jhdnp6bbkb3",
			version:     12345,
			expectedErr: ErrKeySize,
		},
		{
			name:        "InvalidVersion",
			epochKey:    validEpochKey,
			secretID:    "sec-62dyen4w95kv9jhdnp6bbkb3",
			version:     0,
			expectedErr: ErrWrongVersion,
		},
		{
			name:        "InvalidSecretID",
			epochKey:    validEpochKey,
			secretID:    "62dyen4w95kv9jhdnp6bbkb3",
			version:     1,
			expectedErr: ErrWrongSecretID,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := GetValueKey(tc.epochKey, tc.secretID, tc.version)

			if err != tc.expectedErr {
				t.Errorf("Expected error: %v, but got: %v", tc.expectedErr, err)
			}

			if !reflect.DeepEqual(result, tc.expectedSum) {
				t.Errorf("Expected hash sum: %#v, but got: %#v", tc.expectedSum, result)
			}
		})
	}
}
