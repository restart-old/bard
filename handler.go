package bard

import (
	"github.com/df-HCF/class"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/player"
)

type Handler interface {
	HandleBardEffectUse(ctx *event.Context, it class.EnergyEffectItem, affected *[]*player.Player)
}

type NopHandler struct{}

func (NopHandler) HandleBardEffectUse(*event.Context, class.EnergyEffectItem, *[]*player.Player) {}
