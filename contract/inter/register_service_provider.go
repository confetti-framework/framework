package inter

type RegisterServiceProvider interface {
	Register(container Container) Container
}
