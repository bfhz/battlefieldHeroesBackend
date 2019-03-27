package rank

import (
	"github.com/gocraft/dbr"
	"github.com/sirupsen/logrus"

	"github.com/Synaxis/battlefieldHeroesBackend/backend/model"
	"github.com/Synaxis/battlefieldHeroesBackend/backend/network"
	"github.com/Synaxis/battlefieldHeroesBackend/backend/network/codec"
	"github.com/Synaxis/battlefieldHeroesBackend/backend/ranking"
	"github.com/Synaxis/battlefieldHeroesBackend/backend/storage/database"
)

const (
	rankGetStats          = "GetStats"
	rankUpdateStats       = "UpdateStats"
	rankGetStatsForOwners = "GetStatsForOwners"
)

// Ranking probably stands for Ranking
type Ranking struct {
	DB database.Adapter
}

func (r *Ranking) answer(client *network.Client, pnum uint32, payload interface{}) {
	client.WriteEncode(&codec.Answer{
		Type:         codec.FeslRanking,
		PacketNumber: pnum,
		Payload:      payload,
	})
}

func (r *Ranking) fetchStats(heroStats *model.HeroStats, keys []string) ([]statsPair, error) {
	stats := make([]statsPair, len(keys))

	statsValues, err := ranking.GetStats(heroStats, keys...)
	if err != nil {
		return nil, err
	}

	for i, k := range keys {
		s := statsPair{Key: k}
		switch k {
		case "c_apr", "c_emo", "c_eqp", "c_items":
			s.Text = statsValues[k]
		default:
			s.Value = statsValues[k]
		}
		stats[i] = s
	}

	return stats, nil
}

func (r *Ranking) changeStats(p *model.HeroStats, key, value, updateType, pointType string) error {
	logrus.Debugf("Updating stats %s with (%v, %v, %v)", key, value, updateType, pointType)
	err := ranking.UpdateStatValue(p, key, value, updateType, pointType)
	if err != nil {
		logrus.Errorf("Cannot update %s with (%v, %v, %v)", key, value, updateType, pointType)
	}

	return err
}

func (r *Ranking) commitStats(sess *dbr.Session, p *model.HeroStats, heroID int) error {
	tx, err := r.DB.Begin(sess)
	if err != nil {
		return err
	}

	err = r.DB.UpdateStats(tx, heroID, p)
	if err != nil {
		return err
	}

	err = r.DB.Commit(tx)
	return err
}
