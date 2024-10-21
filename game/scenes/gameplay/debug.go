package gameplay

import (
	"fmt"

	"github.com/dwethmar/vork/ecsys"
	"github.com/dwethmar/vork/entity"
)

// debugHierarchy prints the hierarchy of entities in the ECS.
func debugHierarchy(h *ecsys.ECS) {
	// Start rendering from the root entity
	renderHierarchy(h, h.Root(), "", true)
}

// renderHierarchy recursively prints the hierarchy of entities in the ECS.
func renderHierarchy(h *ecsys.ECS, e entity.Entity, prefix string, isLast bool) {
	// Choose the appropriate branch character
	if isLast {
		fmt.Printf("%s└── Entity %v\n", prefix, e)
		// Adjust prefix for the next level of children
		prefix += "    "
	} else {
		fmt.Printf("%s├── Entity %v\n", prefix, e)
		// Adjust prefix for the next level of children
		prefix += "│   "
	}

	// Get the children of the current entity
	children := h.Children(e)

	// Recursively print each child with the new prefix
	for i, child := range children {
		// Determine if this child is the last one in the list
		isLastChild := i == len(children)-1
		renderHierarchy(h, child, prefix, isLastChild)
	}
}
