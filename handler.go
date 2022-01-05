package bard

import (
	"math"
	"time"

	"github.com/RestartFU/dfutil"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/player"
)

type handler struct {
	player.NopHandler
	bard *Bard
}

func (h *handler) HandleItemUse(ctx *event.Context) {
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
		bard.SetEffectCoolDown(10 * time.Second)
		bard.energy -= i.Energy()

		// Give effect message
		eLVL, _ := dfutil.Itor(i.Effect().Level())
		bard.Messagef("§eYou have given §9%s %s§e to §a%v(TODO) §eteammates", dfutil.EffectName(i.Effect()), eLVL)

		// Add the effect
		if e, ok := bard.Effect(i.Effect().Type()); ok {
			dfutil.NewEffectNoLoss(i.Effect(), e).Add(bard.Player)
			return
		}
		bard.AddEffect(i.Effect())
	}
}
