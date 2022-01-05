package bard

import (
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/player"
)

type handler struct {
	player.NopHandler
	bard *Bard
}

func (h *handler) HandleItemUse(ctx *event.Context) {
	h.bard.p.Message(h.bard.energy)
}
