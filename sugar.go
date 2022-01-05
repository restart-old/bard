package bard

import (
	"time"

	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/world"
)

type UseSugar struct{}

func (UseSugar) Energy() int           { return 25 }
func (UseSugar) Effect() effect.Effect { return effect.New(effect.Speed{}, 3, 8*time.Second) }
func (UseSugar) Item() world.Item      { return item.Sugar{} }
