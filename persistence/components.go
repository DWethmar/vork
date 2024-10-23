package persistence

import (
	"github.com/dwethmar/vork/component"
	"github.com/dwethmar/vork/component/controllable"
	"github.com/dwethmar/vork/component/position"
	"github.com/dwethmar/vork/component/skeleton"
	"github.com/dwethmar/vork/component/velocity"
)

// PersistentComponentTypes returns a list of component types that should be saved.
func PersistentComponentTypes() []component.Type {
	return []component.Type{
		controllable.Type,
		position.Type,
		velocity.Type,
		skeleton.Type,
	}
}
