package mock

import "github.com/stretchr/testify/assert"

type ErrorReader struct{}

func (er ErrorReader) Read(_ []byte) (n int, err error) {
	return 0, assert.AnError
}
