package di

import (
	"go.uber.org/dig"
)

type Container struct {
	container *dig.Container
}

func NewContainer() *Container {
	return &Container{
		container: dig.New(),
	}
}

func (c *Container) Register(constructor interface{}) error {
	return c.container.Provide(constructor)
}

func (c *Container) Invoke(function interface{}) error {
	return c.container.Invoke(function)
}
