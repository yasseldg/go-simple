package dIter

import (
	"fmt"
	"sync"

	"github.com/yasseldg/go-simple/logs/sLog"
)

type IterConfig struct {
	Inter

	mu sync.Mutex

	name  string
	count int
	index int

	configs []InterNameConfig
}

func NewIterConfig(name string) *IterConfig {
	return &IterConfig{
		Inter: New(),

		configs: []InterNameConfig{},

		name: name,
	}
}

func (it *IterConfig) String(name string) string {
	configs := ""
	for _, config := range it.configs {
		configs += fmt.Sprintf("\n %s", config.String(config.Name()))
	}

	return fmt.Sprintf("%s %s ( %d / %d ): %s", it.name, name, it.index, it.count, configs)
}

func (it *IterConfig) Log(name string) {
	it.mu.Lock()
	defer it.mu.Unlock()

	sLog.Info("IterConfig: %s ", it.String(name))
}

func (it *IterConfig) Name() string {
	return it.name
}

func (it *IterConfig) Add(configs ...InterNameConfig) {
	it.mu.Lock()
	defer it.mu.Unlock()

	it.configs = append(it.configs, configs...)

	it.setCount()
}

func (it *IterConfig) Count() int {
	return it.count
}

func (it *IterConfig) Next() bool {
	it.mu.Lock()

	i_first := -1

	var i_last int
	for i, config := range it.configs {
		if config.Count() == 0 {
			continue
		}

		if i_first < 0 {
			i_first = i
		}

		if config.Next() {
			config.Log(config.Name())

			if i == i_first {
				it.index++

				it.mu.Unlock()
				return true
			}

			it.configs[i_last].Reset()

			it.mu.Unlock()
			return it.Next()
		}

		if config.Count() > 1 {
			i_last = i
		}
	}

	it.mu.Unlock()
	return false
}

func (it *IterConfig) Reset() {
	it.mu.Lock()
	defer it.mu.Unlock()

	skip := true
	for _, config := range it.configs {

		config.Reset()

		if skip {
			skip = false
			continue
		}

		config.Next()
	}

	it.index = 0
}

// private methods

func (it *IterConfig) setCount() {
	if len(it.configs) == 0 {
		it.count = 0
		return
	}

	it.count = 1
	for _, config := range it.configs {
		count := config.Count()
		if count == 0 {
			continue
		}
		it.count *= config.Count()
	}
}
