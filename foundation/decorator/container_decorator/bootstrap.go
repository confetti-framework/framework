package container_decorator

import (
	"github.com/confetti-framework/framework/contract/inter"
)

type Handler struct {
	Bootstraps []inter.Bootstrap
}

func (d Handler) BootstrapWith(container inter.Container) inter.Container {
	for _, bootstrapper := range d.Bootstraps {
		container = bootstrapper.Bootstrap(container)
	}

	return container
}
