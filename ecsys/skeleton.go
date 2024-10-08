package ecsys

import (
	"fmt"

	"github.com/dwethmar/vork/component/skeleton"
	"github.com/dwethmar/vork/entity"
)

func (s *ECS) Skeleton(e entity.Entity) (skeleton.Skeleton, error) {
	c, err := s.skeletonStore.FirstByEntity(e)
	if err != nil {
		return skeleton.Skeleton{}, fmt.Errorf("could not get skeleton: %w", err)
	}
	return *c, nil
}

func (s *ECS) Skeletons() []skeleton.Skeleton {
	c := s.skeletonStore.List()
	r := make([]skeleton.Skeleton, len(c))
	for i, v := range c {
		r[i] = *v
	}
	return r
}
