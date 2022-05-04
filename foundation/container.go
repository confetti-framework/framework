package foundation

import (
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/framework/contract/inter"
	"github.com/confetti-framework/framework/support"
	"reflect"
)

type Container struct {
	// The container created at boot time
	bootContainer inter.Container

	// A key value store. Callbacks are not automatically executed and stored.
	bindings inter.Bindings

	// A key value store. If the value is a callback, it will be executed
	// once per request and the result will be saved.
	singletons inter.Bindings
}

func NewContainer() *Container {
	containerStruct := Container{}
	containerStruct.bindings = make(inter.Bindings)
	containerStruct.singletons = make(inter.Bindings)

	return &containerStruct
}

func NewContainerByBoot(bootContainer inter.Container) inter.Container {
	container := NewContainer()
	container.bootContainer = bootContainer

	return container
}

func NewTestApp(decorator func(container inter.Container) inter.Container) inter.App {
	app := NewApp()
	container := inter.Container(NewContainer())
	container = decorator(container)
	app.SetContainer(NewContainerByBoot(container))
	return app
}

// Determine if the given abstract type has been bound.
func (c *Container) Bound(abstract string) bool {
	_, bound := c.bindings[abstract]
	_, boundSingleton := c.singletons[abstract]
	return bound || boundSingleton
}

// Register a binding with the container.
func (c *Container) Bind(abstract interface{}, concrete interface{}) {
	abstractString := support.Name(abstract)
	c.bindings[abstractString] = concrete
}

// Register a shared binding in the container.
func (c *Container) Singleton(abstract interface{}, concrete interface{}) {
	abstractString := support.Name(abstract)
	c.singletons[abstractString] = concrete
}

// Register an existing instance as shared in the container without an abstract
func (c *Container) Instance(concrete interface{}) interface{} {
	c.Bind(concrete, concrete)

	return concrete
}

// GetE the container's bindings.
func (c *Container) Bindings() inter.Bindings {
	result := inter.Bindings{}

	if c.bootContainer != nil {
		result = c.bootContainer.Bindings()
	}

	for abstract, concrete := range c.singletons {
		result[abstract] = concrete
	}

	for abstract, concrete := range c.bindings {
		result[abstract] = concrete
	}

	return result
}

// MakeE the given type from the container.
func (c *Container) Make(abstract interface{}) interface{} {
	concrete, err := c.MakeE(abstract)
	if nil != err {
		panic(err)
	}
	return concrete
}

// MakeE the given type from the container.
func (c *Container) MakeE(abstract interface{}) (interface{}, error) {
	var concrete interface{}
	var err error = nil
	var abstractName = support.Name(abstract)

	kind := support.Kind(abstract)
	if support.Kind(abstract) == reflect.Ptr && abstract == nil {
		return nil, errors.New("can't resolve interface. To resolve an interface, " +
			"use the following syntax: (*interface)(nil), use a string or use the struct itself")
	}

	if object, present := c.bindings[abstractName]; present {
		concrete = object

	} else if object, present := c.singletons[abstractName]; present {
		concrete, err = c.getConcreteBinding(concrete, object, abstractName)

	} else if c.bootContainer != nil && c.bootContainer.Bound(abstractName) {
		// Check the container that was created at boot time
		concrete, err = c.bootContainer.MakeE(abstract)
		c.bindings[abstractName] = concrete

	} else if kind == reflect.Struct {
		// If struct cannot be found, we simply have to use the struct itself.
		concrete = abstract
	} else if kind == reflect.String {
		var instances support.Map
		instances, err = support.NewMapE(c.bindings)
		if err == nil {
			var value support.Value
			if c.bootContainer != nil {
				bootBindings := c.bootContainer.Bindings()
				bootInstances, err := support.NewMapE(bootBindings)
				if err != nil {
					return nil, err
				}
				instances.Merge(bootInstances)
			}
			value, err = instances.GetE(abstract.(string))
			//goland:noinspection GoNilness
			concrete = value.Raw()
		}
	}

	if err != nil {
		err = errors.Wrap(err, "get instance '%s' from container", abstractName)
	}

	resolvePointerValue(abstract, concrete)

	return concrete, err
}

func (c *Container) getConcreteBinding(
	concrete interface{},
	object interface{},
	abstractName string,
) (interface{}, error) {
	// If abstract is bound, use that object.
	concrete = object
	value := reflect.ValueOf(concrete)

	// If concrete is a callback, run it and save the result.
	if value.Kind() == reflect.Func {
		if value.Type().NumIn() != 0 {
			return nil, errors.WithStack(CanNotInstantiateCallbackWithParameters)
		}
		concrete = value.Call([]reflect.Value{})[0].Interface()
	}

	// Don't save result in bootContainer. We don't want to share the result across multiple requests
	if c.bootContainer != nil {
		c.bindings[abstractName] = concrete
	}

	return concrete, nil
}

// Extend an abstract type in the container.
func (c *Container) Extend(abstract interface{}, function func(service interface{}) interface{}) {
	concrete := c.Make(abstract)

	newConcrete := function(concrete)

	c.Bind(abstract, newConcrete)
}

func resolvePointerValue(abstract interface{}, concrete interface{}) {
	if support.Kind(abstract) == reflect.Ptr {
		of := reflect.ValueOf(abstract)
		if !of.IsNil() {
			of.Elem().Set(reflect.ValueOf(concrete))
		}
	}
}
