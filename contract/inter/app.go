package inter

const AppProvider = "app_provider"

type App interface {
	AppReader

	// Container returns the service container
	Container() *Container

	// SetContainer set the service container
	SetContainer(container Container)

	// Singleton registered a shared binding in the container.
	Singleton(abstract interface{}, concrete interface{})

	// Bind registered an existing instance as shared in the container.
	Bind(abstract interface{}, concrete interface{})

	// Instance registered an existing instance as shared in the container without an abstract
	Instance(concrete interface{}) interface{}

	Environment() (string, error)

	IsEnvironment(environments ...string) bool

	// The Log method gives you an instance of a logger. You can write your log
	// messages to this instance.
	Log(channels ...string) LoggerFacade

	// Db returns the database facade. If no parameters are provided, it will use
	// the default connection.
	Db(connection ...string) Database
}
