package rank

import (
	"fmt"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/Synaxis/battlefieldHeroesBackend/backend/network"
)

type ansUpdateStats struct {
	Txn   string      `fesl:"TXN"`
	Users []userStats `fesl:"u"`
}

type userStats struct {
	OwnerID   int          `fesl:"o"`  // 3
	OwnerType int          `fesl:"ot"` // 1
	Stats     []updateStat `fesl:"s"`
}

type updateStat struct {
	Key        string `fesl:"k"`  // c_ltp
	PointType  int    `fesl:"pt"` // 0
	Text       string `fesl:"t"`  // "" 
	UpdateType int    `fesl:"ut"` // 0
	Value      string `fesl:"v"`  // 9025.0000
}

type stat struct {
	text  string
	value float64
}

// UpdateStats - updates stats about a soldier
func (r *Ranking) UpdateStats(event network.EventClientCommand) {
	switch event.Client.GetClientType() {
	case "server":
		r.serverUpdateStats(&event)
	default:
		r.clientUpdateStats(&event)
	}
}

func (r *Ranking) clientUpdateStats(event *network.EventClientCommand) {
	r.updateStats(event)
}

func (r *Ranking) serverUpdateStats(event *network.EventClientCommand) {
	r.updateStats(event)
}

func (r *Ranking) updateStats(event *network.EventClientCommand) {
	users, _ := strconv.Atoi(event.Command.Message["u.[]"])
	sess := r.DB.NewSession()

	for i := 0; i < users; i++ {
		heroID, _ := event.Command.Message.IntVal(fmt.Sprintf("u.%d.o", i))
		p, err := r.DB.GetHeroStats(sess, heroID)
		if err != nil {
			logrus.
				WithError(err).
				WithField("heroID", event.Command.Message[fmt.Sprintf("u.%d.o", i)]).
				Warn("Cannot resolve hero stats when updating stats")
			return
		}

		numKeys, _ := event.Command.Message.IntVal(fmt.Sprintf("u.%d.s.[]", i))
		for j := 0; j < numKeys; j++ {
			prefix := fmt.Sprintf("u.%d.s.%d.", i, j)

			key := event.Command.Message.Get(prefix + "k")
			ut := event.Command.Message.Get(prefix + "ut")
			pt := event.Command.Message.Get(prefix + "pt")
			val := event.Command.Message.Get(prefix + "v")
			text := event.Command.Message.Get(prefix + "t")

						//ChangeStats in both cases
						if text != "" {
							// c_items, c_eqp..
							val = text
							logrus.Println("--UpdateStat replace ut 0--"+key, val, ut)
							r.changeStats(&p, key, val, ut, pt)
							//GOTO LN102
						} else {
							logrus.Println("--UpdateStat sum ut 3"+key, val, ut)
							r.changeStats(&p, key, val, ut, pt)
						}
			
					}
			
					if err = r.commitStats(sess, &p, heroID); err != nil {
						logrus.Error(err)
					}
				}
			
				r.answer(event.Client, event.Command.PayloadID, ansUpdateStats{
					Txn: rankUpdateStats,
				})
			}