package services

import (
	"fmt"
	"sync"

	"github.com/MovieStoreGuy/lucid/types"
)

var (
	factor *Factory
	once   sync.Once
)

// Factory ...
type Factory struct {
	lock       sync.Mutex
	registered map[string](func() (types.Service, error))
}

func GetFactoryInstance() *Factory {
	once.Do(func() {
		factor = &Factory{
			registered: map[string](func() (types.Service, error)){},
		}
	})
	return factor
}

func (f *Factory) register(name string, create func() (types.Service, error)) {
	f.lock.Lock()
	defer f.lock.Unlock()
	f.registered[name] = create
}

func (f *Factory) String() string {
	return fmt.Sprintf("Factory: [len:%v]", len(f.registered))
}
