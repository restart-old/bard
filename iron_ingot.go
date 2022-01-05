package bard

import (
	"time"

	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/world"
)

type UseIronIngot struct{}

func (UseIronIngot) Energy() int           { return 35 }
func (UseIronIngot) Effect() effect.Effect { return effect.New(effect.Resistance{}, 3, 6*time.Second) }
func (UseIronIngot) Item() world.Item      { return item.IronIngot{} }

type HeldIronIngot struct{}

func (HeldIronIngot) Effect() effect.Effect { return effect.New(effect.Resistance{}, 1, 6*time.Second) }
func (HeldIronIngot) Item() world.Item      { return item.IronIngot{} }
