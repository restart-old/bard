package bard

import (
	"time"

	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/world"
)

type UseFeather struct{}

func (UseFeather) Energy() int           { return 35 }
func (UseFeather) Effect() effect.Effect { return effect.New(effect.JumpBoost{}, 8, 7*time.Second) }
func (UseFeather) Item() world.Item      { return item.Feather{} }

type HeldFeather struct{}

func (HeldFeather) Effect() effect.Effect { return effect.New(effect.JumpBoost{}, 2, 7*time.Second) }
func (HeldFeather) Item() world.Item      { return item.Feather{} }
