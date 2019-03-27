package ranking

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/Synaxis/battlefieldHeroesBackend/backend/model"
	"strconv"
)

const (
	PointTypeFloat = iota
)

type Stats map[string]string

type Adder func(p *model.HeroStats, value string, pointType int) error

type Getter func(p *model.HeroStats) (string, error)

type Setter func(p *model.HeroStats, value string, pointType int) error

func GetStats(p *model.HeroStats, keys ...string) (Stats, error) {
	s := Stats{}

	for _, key := range keys {
		val, err := getStatsValue(p, key)
		if err != nil {
			return nil, err
		}
		s[key] = val
	}

	return s, nil
}

func getStatsValue(p *model.HeroStats, key string) (string, error) {
	if getter, ok := getters[key]; ok {
		val, err := getter(p)
		if err != nil {
			return "", fmt.Errorf("ranking: cant fetch value for key='%s', %v", key, err)
		}

		return val, nil
	}

	return "", fmt.Errorf("ranking: cant get value for key='%s'", key)
}

const (
	UpdateTypeReplace = iota //0
	UpdateTypeAdd1           //iota increased 1
	UpdateTypeAdd2           // 2
	UpdateTypeAdd3           //3
)

func UpdateStatValue(p *model.HeroStats, key, value string, updateType, pointType string) error {
	ut, err := strconv.Atoi(updateType)
	if err != nil {
		return fmt.Errorf(
			"ranking: unknown updateType for key='%s' (%s)",
			key,
			updateType,
		)
	}

	pt, err := strconv.Atoi(pointType)
	if err != nil {
		return err
	}

	switch ut {
	case UpdateTypeReplace:
		return setStatsValue(p, key, value, pt)
	case UpdateTypeAdd1:
		return addStatsValue(p, key, value, pt)
	case UpdateTypeAdd2:
		return addStatsValue(p, key, value, pt)
	case UpdateTypeAdd3:
		return addStatsValue(p, key, value, pt)
	default:
		return fmt.Errorf(
			"ranking: cannot update stats value key='%s' (method=%s, value=%s)",
			key,
			updateType,
			value,
		)
	}
}

func addStatsValue(p *model.HeroStats, key, val string, pt int) error {
	logrus.Println("added stats value")
	if adder, ok := adders[key]; ok {
		if err := adder(p, val, pt); err != nil {
			return fmt.Errorf("ranking: cant add value for key='%s'", key)
		}
	}

	return nil
}

func setStatsValue(p *model.HeroStats, key, val string, pt int) error {
	logrus.Println("replaced stats value")
	if setter, ok := setters[key]; ok {
		if err := setter(p, val, pt); err != nil {
			return fmt.Errorf("ranking: cant set value for key='%s'", key)
		}
	}

	return nil
}
