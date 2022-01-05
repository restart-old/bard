package bard

import (
	"sync"

	"github.com/df-HCF/class"
	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/world"
)

type EffectHeldItem interface {
	Effect() effect.Effect
	Item() world.Item
}

var items []class.EnergyEffectItem
var itemsMu sync.RWMutex

func RegisterItem(i class.EnergyEffectItem) {
	itemsMu.Lock()
	defer itemsMu.Unlock()
	items = append(items, i)
}

func EffectItemByItem(i world.Item) (class.EnergyEffectItem, bool) {
	itemsMu.RLock()
	defer itemsMu.RUnlock()
	for _, item := range items {
		if item.Item() == i {
			return item, true
		}
	}
	return nil, false
}
