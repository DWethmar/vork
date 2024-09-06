package component

import "github.com/dwethmar/vork/entity"

type ComponentType string

// Component is an interface that all components must implement.
type Component interface {
	ID() uint32
	SetID(uint32)
	Type() ComponentType
	Entity() entity.Entity
}

// BaseComponent is a base struct that all components should embed.
type BaseComponent struct {
	I uint32
	E entity.Entity
	T ComponentType
}

func (c *BaseComponent) ID() uint32            { return c.I }
func (c *BaseComponent) SetID(i uint32)        { c.I = i }
func (c *BaseComponent) Type() ComponentType   { return c.T }
func (c *BaseComponent) Entity() entity.Entity { return c.E }
