package component

import "github.com/dwethmar/vork/entity"

type ComponentType string

// Component is an interface that all components must implement.
type Component interface {
	ID() uint
	SetID(uint)
	Type() ComponentType
	Entity() entity.Entity
}
