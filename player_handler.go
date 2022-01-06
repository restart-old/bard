package bard

import (
	"math"
	"time"

	"github.com/RestartFU/dfutil"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/player"
)

type playerHandler struct {
	player.NopHandler
	bard *Bard
}

func (h *playerHandler) HandleItemUse(ctx *event.Context) {
	bard := h.bard
	held, _ := bard.HeldItems()

	if i, ok := EffectItemByItem(held.Item()); ok {
		// Making sure the bard has enough energy
		if bard.energy <= i.Energy() {
			bard.Messagef("§cYou do not have enough energy for this! You need %v energy. but you only have %v", i.Energy(), bard.energy)
			return
		}
		// Making sure the bard is not on effect cooldown
		if bard.EffectCoolDown() {
			bard.Messagef("§cYou cannot use this for another %v seconds!", math.Floor(float64(time.Until(bard.effectCooldown).Seconds()*10))/10)
			return
		}

		// Set the effect cooldown and remove the energy used
		bard.SetEffectCoolDown(bard.coolDownDuration)
		bard.energy -= i.Energy()

		ctx := event.C()

		pls := dfutil.PlayersInRadius(bard.Player, bard.radius)
		players := &pls

		bard.handler().HandleBardEffectUse(ctx, i, players)
		ctx.Continue(func() {
			for _, p := range *players {
				// Add the effect
				if e, ok := bard.Effect(i.Effect().Type()); ok {
					dfutil.NewEffectNoLoss(i.Effect(), e).Add(p)
					return
				}
				p.AddEffect(i.Effect())
			}
		})
	}
}
