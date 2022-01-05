package bard

import (
	"time"

	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/world"
)

type UseBlazePowder struct{}

func (UseBlazePowder) Energy() int           { return 45 }
func (UseBlazePowder) Effect() effect.Effect { return effect.New(effect.Strength{}, 2, 6*time.Second) }
func (UseBlazePowder) Item() world.Item      { return item.BlazePowder{} }
