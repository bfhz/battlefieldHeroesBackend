package dbtest

import "github.com/Synaxis/battlefieldHeroesBackend/backend/model"

type FakeDB struct {
	QueriesAdapter
	TxAdapter
}

func NewFakeDB() *FakeDB {
	hero := model.Hero{
		ID:       10,
		HeroName: "SecondHero",
		PlayerID: 8,
	}

	otherHero := model.Hero{
		ID:       9,
		HeroName: "FirstHero",
		PlayerID: 8,
	}

	return &FakeDB{
		QueriesAdapter: QueriesAdapter{
			Server: model.Server{
				ID:              123,
				APIKey:          "SERVER-APIKEY",
				SoldierName:     "Test-Server",
				AccountUsername: "Test-Server",
				AccountPassword: "Test-Server",
			},

			Player: model.Player{
				ID:        8,
				Username:  "SomeUser",
				Password:  "admin1",
				GameToken: "topsecret",
			},

			Heroes:    []model.Hero{hero, otherHero},
			Hero:      hero,
			HeroStats: model.NewHeroStats(),
		},
	}
}
