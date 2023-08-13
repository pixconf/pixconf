package test

import (
	"encoding/json"

	"github.com/pixconf/pixconf/cmd/agent/agent/modules"
)

func init() {
	modules.RegisterModule("test", &Test{})
}

type Test struct{}

func (t *Test) ListFunctions() []string {
	return []string{"ping"}
}

func (t *Test) HasFunction(name string) bool {
	switch name {
	case "ping":
		return true
	default:
		return false
	}
}

func (t *Test) CallFunction(name string, data modules.ModuleData) (json.RawMessage, error) {
	switch name {
	case "ping":
		return t.ping(data)
	default:
		return nil, modules.ErrUnknownFunction
	}
}

func (t *Test) ping(_ modules.ModuleData) (json.RawMessage, error) {
	return []byte("pong"), nil
}
