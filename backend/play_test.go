package internal_test

// import (
// 	"bytes"
// 	"net"
// 	"testing"

// 	"github.com/Synaxis/battlefieldHeroesBackend/backend/theater"

// 	"github.com/Synaxis/battlefieldHeroesBackend/backend/fesl/pnow"
// 	"github.com/Synaxis/battlefieldHeroesBackend/backend/fesl/rank"
// "github.com/Synaxis/battlefieldHeroesBackend/backend/network"
// "github.com/Synaxis/battlefieldHeroesBackend/backend/network/codec"
// )

// func TestPlayNow(t *testing.T) {
// 	ipAddr, _ := net.ResolveIPAddr("", "26.49.251.103")
// 	client := &net.Client{
// 		IpAddr:   ipAddr,
// 		IsActive: true,
// 		Conn:     &fakeConn{buf: new(bytes.Buffer)},
// 	}

// 	pnowCtrl := &pnow.Pnow{}
// 	rankCtrl := &rank.Rank{}
// 	thtrCtrl := &theater.Theater{}

// 	// 21:47:23 : [09:47:23 PM] ->N req pnow 0xc000000d {TXN=Start, partition.partition=/eagames/bfwest-dedicated, debugLevel=off, version=1, players.[]=1, players.0.ownerId=3, players.0.ownerType=1, players.0.props.{sessionType}=listServers, players.0.props.{name}=s0u1m4st3r[09:47:23 PM] , players.0.props.{firewallType}=unknown, players.0.props.{poolMaxPlayers}=1, players.0.props.{poolTimeout}=30, players.0.props.{poolTargetPlayers}=0:1, players.0.props.{availableServerCount}=1, players.0.props.{maxListServersResult}=20[09:47:23 PM] , players.0.props.{filter-version}="1.42.217478.0 ", players.0.props.{filterToGame-version}=version, players.0.props.{filter-avail_slots_royal}=yes, players.0.props.{filterToGame-avail_slots_royal}=U-avail_slots_royal[09:47:23 PM] , players.0.props.{filter-data_center}=iad, players.0.props.{filterToGame-data_center}=U-data_center, players.0.props.{filter-map}=both, players.0.props.{filterToGame-map}=U-map, players.0.props.{filter-ranked}=yes[09:47:23 PM] , players.0.props.{filterToGame-ranked}=U-ranked, players.0.props.{filter-server_state}=has_players, players.0.props.{filterToGame-server_state}=U-server_state, players.0.props.{filter-servertype}=public[09:47:23 PM] , players.0.props.{filterToGame-servertype}=U-servertype, players.0.props.{pref-army_balance}=Allies, players.0.props.{prefVotingMethod-army_balance}=lottery, players.0.props.{fitValues-army_balance}="MaxAxis,Axis,Balanced,Allies,MaxAllies"[09:47:23 PM] , players.0.props.{fitTable-army_balance}=0;0;0;0;0|-1;0.1;0.5;0.9;1|0;0;0;0;0|1;0.9;0.5;0.1;-1|0;0;0;0;0, players.0.props.{fitWeight-army_balance}=200, players.0.props.{fitThresholds-army_balance}=0:0[09:47:23 PM] , players.0.props.{prefToGame-army_balance}=U-army_balance, players.0.props.{pref-lvl_avg}=30, players.0.props.{aggrPref-lvl_avg}=0, players.0.props.{fitScale-lvl_avg}=15, players.0.props.{fitWeight-lvl_avg}=200[09:47:23 PM] , players.0.props.{fitThresholds-lvl_avg}=0:100, players.0.props.{prefToGame-lvl_avg}=U-lvl_avg, players.0.props.{pref-lvl_sdv}=0, players.0.props.{aggrPref-lvl_sdv}=0, players.0.props.{fitScale-lvl_sdv}=2, players.0.props.{fitWeight-lvl_sdv}=120[09:47:23 PM] , players.0.props.{fitThresholds-lvl_sdv}=0:0, players.0.props.{prefToGame-lvl_sdv}=U-lvl_sdv, players.0.props.{pref-percent_full}=80, players.0.props.{aggrPref-percent_full}=0, players.0.props.{fitScale-percent_full}=30[09:47:23 PM] , players.0.props.{fitWeight-percent_full}=200, players.0.props.{fitThresholds-percent_full}=0:0, players.0.props.{prefToGame-percent_full}=U-percent_full, players.0.props.{}=44}
// 	startPacket := &codec.Cmd{
// 		Msg: codec.Fields{
// 			"TXN": `Start`,
// 			"partition.partition":                              `/eagames/bfwest-dedicated`,
// 			"debugLevel":                                       `off`,
// 			"version":                                          `1`,
// 			"players.[]":                                       `1`,
// 			"players.0.ownerId":                                `3`,
// 			"players.0.ownerType":                              `1`,
// 			"players.0.props.{sessionType}":                    `listServers`,
// 			"players.0.props.{name}":                           `s0u1m4st3r[09:47:23 PM]`,
// 			"players.0.props.{firewallType}":                   `unknown`,
// 			"players.0.props.{poolMaxPlayers}":                 `1`,
// 			"players.0.props.{poolTimeout}":                    `30`,
// 			"players.0.props.{poolTargetPlayers}":              `0:1`,
// 			"players.0.props.{availableServerCount}":           `1`,
// 			"players.0.props.{maxListServersResult}":           `20`,
// 			"players.0.props.{filter-version}":                 `"1.42.217478.0 "`,
// 			"players.0.props.{filterToGame-version}":           `version`,
// 			"players.0.props.{filter-avail_slots_royal}":       `yes`,
// 			"players.0.props.{filterToGame-avail_slots_royal}": `U-avail_slots_royal`,
// 			"players.0.props.{filter-data_center}":             `iad`,
// 			"players.0.props.{filterToGame-data_center}":       `U-data_center`,
// 			"players.0.props.{filter-map}":                     `both`,
// 			"players.0.props.{filterToGame-map}":               `U-map`,
// 			"players.0.props.{filter-ranked}":                  `yes`,
// 			"players.0.props.{filterToGame-ranked}":            `U-ranked`,
// 			"players.0.props.{filter-server_state}":            `has_players`,
// 			"players.0.props.{filterToGame-server_state}":      `U-server_state`,
// 			"players.0.props.{filter-servertype}":              `public`,
// 			"players.0.props.{filterToGame-servertype}":        `U-servertype`,
// 			"players.0.props.{pref-army_balance}":              `Allies`,
// 			"players.0.props.{prefVotingMethod-army_balance}":  `lottery`,
// 			"players.0.props.{fitValues-army_balance}":         `"MaxAxis,Axis,Balanced,Allies,MaxAllies"`,
// 			"players.0.props.{fitTable-army_balance}":          `0;0;0;0;0|-1;0.1;0.5;0.9;1|0;0;0;0;0|1;0.9;0.5;0.1;-1|0;0;0;0;0`,
// 			"players.0.props.{fitWeight-army_balance}":         `200`,
// 			"players.0.props.{fitThresholds-army_balance}":     `0:0`,
// 			"players.0.props.{prefToGame-army_balance}":        `U-army_balance`,
// 			"players.0.props.{pref-lvl_avg}":                   `30`,
// 			"players.0.props.{aggrPref-lvl_avg}":               `0`,
// 			"players.0.props.{fitScale-lvl_avg}":               `15`,
// 			"players.0.props.{fitWeight-lvl_avg}":              `200`,
// 			"players.0.props.{fitThresholds-lvl_avg}":          `0:100`,
// 			"players.0.props.{prefToGame-lvl_avg}":             `U-lvl_avg`,
// 			"players.0.props.{pref-lvl_sdv}":                   `0`,
// 			"players.0.props.{aggrPref-lvl_sdv}":               `0`,
// 			"players.0.props.{fitScale-lvl_sdv}":               `2`,
// 			"players.0.props.{fitWeight-lvl_sdv}":              `120`,
// 			"players.0.props.{fitThresholds-lvl_sdv}":          `0:0`,
// 			"players.0.props.{prefToGame-lvl_sdv}":             `U-lvl_sdv`,
// 			"players.0.props.{pref-percent_full}":              `80`,
// 			"players.0.props.{aggrPref-percent_full}":          `0`,
// 			"players.0.props.{fitScale-percent_full}":          `30`,
// 			"players.0.props.{fitWeight-percent_full}":         `200`,
// 			"players.0.props.{fitThresholds-percent_full}":     `0:0`,
// 			"players.0.props.{prefToGame-percent_full}":        `U-percent_full`,
// 			"players.0.props.{}":                               `44`,
// 		},
// 	}

// 	// 18:54:04 : ->N req rank 0xc000000e {TXN=UpdateStats, u.0.o=5, u.0.ot=1, u.0.s.[]=5, u.0.s.0.ut=0, u.0.s.0.k=c_apr, u.0.s.0.v=0.0000, u.0.s.0.t=10;979;981, u.0.s.0.pt=0, u.0.s.1.ut=0, u.0.s.1.k=c_emo, u.0.s.1.v=0.0000, u.0.s.1.t=5000;5007;5016;0;0;0;0;0;0, u.0.s.1.pt=0, u.0.s.2.ut=0, u.0.s.2.k=c_eqp, u.0.s.2.v=0.0000, u.0.s.2.t=3002;3014;2141;3155;2005;0;0;0;0;0, u.0.s.2.pt=0, u.0.s.3.ut=0, u.0.s.3.k=c_ltp, u.0.s.3.v=9294.0000, u.0.s.3.t=, u.0.s.3.pt=0, u.0.s.4.ut=0, u.0.s.4.k=c_wmid0, u.0.s.4.v=6000.0000, u.0.s.4.t=, u.0.s.4.pt=0, u.[]=1}
// 	rankUpdateStats := &codec.Cmd{
// 		Msg: codec.Fields{
// 			"TXN":        `UpdateStats`,
// 			"u.0.o":      `5`,
// 			"u.0.ot":     `1`,
// 			"u.0.s.[]":   `5`,
// 			"u.0.s.0.ut": `0`,
// 			"u.0.s.0.k":  `c_apr`,
// 			"u.0.s.0.v":  `0.0000`,
// 			"u.0.s.0.t":  `10;979;981`,
// 			"u.0.s.0.pt": `0`,
// 			"u.0.s.1.ut": `0`,
// 			"u.0.s.1.k":  `c_emo`,
// 			"u.0.s.1.v":  `0.0000`,
// 			"u.0.s.1.t":  `5000;5007;5016;0;0;0;0;0;0`,
// 			"u.0.s.1.pt": `0`,
// 			"u.0.s.2.ut": `0`,
// 			"u.0.s.2.k":  `c_eqp`,
// 			"u.0.s.2.v":  `0.0000`,
// 			"u.0.s.2.t":  `3002;3014;2141;3155;2005;0;0;0;0;0`,
// 			"u.0.s.2.pt": `0`,
// 			"u.0.s.3.ut": `0`,
// 			"u.0.s.3.k":  `c_ltp`,
// 			"u.0.s.3.v":  `9294.0000`,
// 			"u.0.s.3.t":  ``,
// 			"u.0.s.3.pt": `0`,
// 			"u.0.s.4.ut": `0`,
// 			"u.0.s.4.k":  `c_wmid0`,
// 			"u.0.s.4.v":  `6000.0000`,
// 			"u.0.s.4.t":  ``,
// 			"u.0.s.4.pt": `0`,
// 			"u.[]":       `1`,
// 		},
// 	}

// 	rankCtrl.UpdateStats(network.EClientCmd{client, rankUpdateStats})

// 	// 21:47:23 : [09:47:23 PM] <-N req pnow 0xc000000d {id.partition=/eagames/bfwest-dedicated, TXN=Start, id.id=1}
// 	pnowCtrl.Start(network.EClientCmd{client, startPacket})
// 	// 21:47:23 : [09:47:23 PM] <-A res pnow 0x80000000 {id.partition=/eagames/bfwest-dedicated, props.{}.[]=2, props.{games}.0.lid=1, props.{games}.0.fit=1001, TXN=Status, id.id=1, sessionState=COMPLETE, props.{resultType}=JOIN, props.{games}.0.gid=7, props.{games}.[]=1}
// 	pnowCtrl.Status(network.EClientCmd{client, startPacket})

// 	// 21:47:23 : [09:47:23 PM] ->N GDAT 0x40000000 {LID=1, GID=7, TID=3}
// 	gdatPacket := &codec.Cmd{
// 		Msg: codec.Fields{
// 			"LID": `1`,
// 			"GID": `7`,
// 			"TID": `3`,
// 		},
// 	}

// 	// 21:47:23 : [09:47:23 PM] <-N GDAT {INT-IP=26.49.251.103, IP=26.49.251.103, B-U-map_name=village, B-U-easyzone=no, INT-PORT=18567, B-U-elo_rank=1000, B-U-avail_slots_national=yes, NAME=[iad]A}
// 	thtrCtrl.GameData(network.EClientCmd{client, gdatPacket})

// 	// 21:47:23 : [09:47:23 PM] ->N EGAM 0x40000000 {R-U-accid=3, R-U-category=3, R-U-dataCenter=iad, R-U-elo=979, R-U-externalIp=26.49.251.103, R-U-kit=2, R-U-lvl=30, R-U-team=2, PORT=61774, R-INT-PORT=61774, R-INT-IP=192.168.1.104, PTYPE=P, LID=1, GID=7, TID=4}
// 	egamPacket := &codec.Cmd{
// 		Msg: codec.Fields{
// 			"R-U-accid":      `3`,
// 			"R-U-category":   `3`,
// 			"R-U-dataCenter": `iad`,
// 			"R-U-elo":        `979`,
// 			"R-U-externalIp": `26.49.251.103`,
// 			"R-U-kit":        `2`,
// 			"R-U-lvl":        `30`,
// 			"R-U-team":       `2`,
// 			"PORT":           `61774`,
// 			"R-INT-PORT":     `61774`,
// 			"R-INT-IP":       `192.168.1.104`,
// 			"PTYPE":          `P`,
// 			"LID":            `1`,
// 			"GID":            `7`,
// 			"TID":            `4`,
// 		},
// 	}

// 	// 21:47:23 : [09:47:23 PM] <-N EGAM {LID=1, GID=7, TID=4}
// 	thtrCtrl.EnterGame(network.EClientCmd{client, egamPacket})

// 	// 21:47:23 : [09:47:23 PM] <-U EGEG {LID=1, PID=3, UGID=Server, P=18567, EKEY=O65zZ2D2A58mNrZw1hmuJw%3d%3d, INT-PORT=18567, TID=4, PL=pc, INT-IP=26.49.251.103, SECRET=2587913, HUID=1, GID=7, TICKET=2018751182, I=26.49.251.103}
// 	egegPacket := &codec.Cmd{
// 		Msg: codec.Fields{
// 			"LID":      `1`,
// 			"PID":      `3`,
// 			"UGID":     `Server`,
// 			"P":        `18567`,
// 			"EKEY":     `O65zZ2D2A58mNrZw1hmuJw%3d%3d`,
// 			"INT-PORT": `18567`,
// 			"TID":      `4`,
// 			"PL":       `pc`,
// 			"INT-IP":   `26.49.251.103`,
// 			"SECRET":   `2587913`,
// 			"HUID":     `1`,
// 			"GID":      `7`,
// 			"TICKET":   `2018751182`,
// 			"I":        `26.49.251.103`,
// 		},
// 	}

// 	// 21:47:23 : [09:47:23 PM] GM: Making host connection.
// 	// 21:47:23 : [09:47:23 PM]
// 	// MakeConnId created connident: 0x173
// 	// 21:47:23 : [09:47:23 PM] GM: Sent hello to host.
// 	// 21:47:24 : [09:47:24 PM] GM: Received host hello.
// 	// 21:47:24 : [09:47:24 PM] GM: Received roster element.
// 	// 21:47:24 : [09:47:24 PM] GM: Processed roster notice for player id 3 (1 of 1).
// 	// 21:47:24 : [09:47:24 PM] GM: Sent roster ack to host.
// 	// 21:47:24 : [09:47:24 PM] GM: Received join for player 3.
// 	// 21:47:24 : [09:47:24 PM] GM: Received join complete for player 3.
// 	// /// ...
// 	// 21:48:51 : [09:48:51 PM] <-A req fsys 0xc0000000 {TXN=MemCheck, memcheck.[]=0, salt=5}
// 	// 21:48:51 : [09:48:51 PM] ->R res fsys 0x80000000 {TXN=MemCheck, result=}
// 	// 21:48:56 : [09:48:56 PM] ->O ECNL 0x40000000 {LID=1, GID=7, TID=5}
// 	// 21:48:56 : [09:48:56 PM] ->N req rank 0xc000000e {TXN=GetStats, owner=3, ownerType=1, periodId=0, periodPast=0, keys.0=c_items, keys.1=c_wallet_hero, keys.2=c_wallet_valor, keys.[]=3}
// 	// 21:49:00 : [09:49:00 PM] <-N req rank 0xc000000e {stats.[]=3, ownerId=3, stats.0.key=c_items, stats.0.text=2026;2027;2028;2031;2032;2033;2046;2047;2048;2055;2056;2057;2091, stats.1.value=0.0000, stats.1.text=0.0000, stats.2.key=c_wallet_valor, stats.2.text=59, TXN=GetStats[09:49:00 PM] , ownerType=1, stats.0.value=2026;2027;2028;2031;2032;2033;2046;2047;2048;2055;2056;2057;2091, stats.1.key=c_wallet_hero, stats.2.value=59}
// 	// 21:49:00 : [09:49:00 PM] <-O ECNL {TID=5, GID=7, LID=1}

// 	// var (
// 	// 	gameID  int    = 7
// 	// 	lobbyID string = "1"
// 	// 	ticket  string = "2018751182"
// 	// )
// 	// thtrCtrl.EGEG(&network.EClientCmd{client, egegPacket}, gameID, lobbyID, ticket)
// }
