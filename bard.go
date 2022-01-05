package bard

import (
	"time"

	"github.com/RestartFU/tickerFunc"
	"github.com/df-HCF/class"
	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/item/armour"
	"github.com/df-mc/dragonfly/server/player"
)

// Bard ...
type Bard struct {
	energy    int
	maxEnergy int
	*player.Player

	effectCooldown time.Time

	tickers []*tickerFunc.Ticker
}

func (b *Bard) EffectCoolDown() bool              { return b.effectCooldown.After(time.Now()) }
func (b *Bard) SetEffectCoolDown(d time.Duration) { b.effectCooldown = time.Now().Add(d) }

func (b *Bard) energyTicker() *tickerFunc.Ticker {
	return tickerFunc.NewTicker(1*time.Second, func() {
		if b.maxEnergy > b.energy {
			b.energy++
		}
	})
}

func (*Bard) New(p *player.Player) class.Class {
	bard := &Bard{
		energy:    10000,
		maxEnergy: 120,
		Player:    p,
	}
	bard.tickers = append(bard.tickers, bard.energyTicker())
	return bard
}

func (*Bard) Armour() class.Armour {
	return class.Armour{
		Helmet:     armour.TierGold,
		Chestplate: armour.TierGold,
		Leggings:   armour.TierGold,
		Boots:      armour.TierGold,
	}
}

func (b *Bard) Handler(p *player.Player) player.Handler {
	return &handler{bard: b}
}

func (b *Bard) Tickers(p *player.Player) []*tickerFunc.Ticker {
	return b.tickers
}
func (b *Bard) Effects() []effect.Effect {
	return []effect.Effect{
		effect.New(effect.Speed{}, 2, 730*time.Hour),
		effect.New(effect.Regeneration{}, 1, 730*time.Hour),
		effect.New(effect.Resistance{}, 2, 730*time.Hour),
	}
}
