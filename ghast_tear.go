package bard

import (
	"time"

	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/world"
)

type UseGhastTear struct{}

func (UseGhastTear) Energy() int { return 45 }
func (UseGhastTear) Effect() effect.Effect {
	return effect.New(effect.Regeneration{}, 3, 6*time.Second)
}
func (UseGhastTear) Item() world.Item { return item.GhastTear{} }

type HeldGhastTear struct{}

func (HeldGhastTear) Effect() effect.Effect {
	return effect.New(effect.Regeneration{}, 1, 6*time.Second)
}
func (HeldGhastTear) Item() world.Item { return item.GhastTear{} }
