package dIter

type NameConfig struct {
	InterConfig

	name string
}

func NewNameConfig(name string, inter InterConfig) *NameConfig {
	return &NameConfig{
		InterConfig: inter,
		name:        name,
	}
}

func (it *NameConfig) Name() string {
	return it.name
}
