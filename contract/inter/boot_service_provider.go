package inter

type BootServiceProvider interface {
	Boot(container Container) Container
}
