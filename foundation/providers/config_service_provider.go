package providers

import (
	"github.com/confetti-framework/framework/inter"
)

type ConfigServiceProvider struct {
	Index map[string]interface{}
}

func (c ConfigServiceProvider) Register(container inter.Container) inter.Container {
	container.Bind("config", c.Index)

	return container
}
