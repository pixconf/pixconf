package modules

import (
	"encoding/json"
	"sync"
)

type Module interface {
	ListFunctions() []string
	HasFunction(name string) bool
	CallFunction(name string, data ModuleData) (json.RawMessage, error)
}

type ModuleData struct {
	Args   []string
	KWArgs map[string]string
}

var (
	modules map[string]Module
	lock    sync.RWMutex
)

func init() {
	modules = make(map[string]Module)
}

func RegisterModule(name string, module Module) {
	lock.Lock()
	defer lock.Unlock()

	modules[name] = module
}

func HasModule(name string) bool {
	lock.RLock()
	defer lock.RUnlock()

	if _, exists := modules[name]; exists {
		return true
	}

	return false
}

func GetModule(name string) Module {
	lock.RLock()
	defer lock.RUnlock()

	return modules[name]
}
