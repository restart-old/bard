package main

import (
	"github.com/df-HCF/bard"
	"github.com/df-HCF/class"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/armour"
)

func main() {
	c := server.DefaultConfig()
	c.Players.SaveData = false
	s := server.New(&c, nil)
	s.Start()
	class.Register(&bard.Bard{})
	for {
		p, err := s.Accept()
		if err != nil {
			return
		}

		p.Armour().Inventory().Handle(class.NewHandler(p))
		p.Inventory().AddItem(item.NewStack(item.Sugar{}, 64))
		p.Inventory().AddItem(item.NewStack(item.SpiderEye{}, 64))

		p.Inventory().AddItem(item.NewStack(item.Helmet{Tier: armour.TierGold}, 1))
		p.Inventory().AddItem(item.NewStack(item.Chestplate{Tier: armour.TierGold}, 1))
		p.Inventory().AddItem(item.NewStack(item.Leggings{Tier: armour.TierGold}, 1))
		p.Inventory().AddItem(item.NewStack(item.Boots{Tier: armour.TierGold}, 1))
	}
}
