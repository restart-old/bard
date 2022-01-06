package bard

import (
	"sync"
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
	radius    float64
	*player.Player
	h                Handler
	hMu              sync.RWMutex
	effectCooldown   time.Time
	coolDownDuration time.Duration

	tickers []*tickerFunc.Ticker
}

func New(radius float64, maxEnergy int, coolDownDuration time.Duration) *Bard {
	return &Bard{
		maxEnergy:        maxEnergy,
		radius:           radius,
		coolDownDuration: coolDownDuration,
		h:                NopHandler{},
	}
}

func (b *Bard) Handle(h Handler) {
	if h != nil {
		b.hMu.Lock()
		defer b.hMu.Unlock()
		b.h = h
	}
}

func (b *Bard) handler() Handler {
	b.hMu.RLock()
	defer b.hMu.RUnlock()
	return b.h
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

func (b *Bard) New(p *player.Player) class.Class {
	bard := &Bard{
		coolDownDuration: b.coolDownDuration,
		radius:           b.radius,
		maxEnergy:        b.maxEnergy,
		Player:           p,
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
	return &playerHandler{bard: b}
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
