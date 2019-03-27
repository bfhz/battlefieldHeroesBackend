package main

import (
	"encoding/json"
	"flag"

	"github.com/gocraft/dbr"
	"github.com/sirupsen/logrus"
	"github.com/subosito/gotenv"

	"github.com/Synaxis/battlefieldHeroesBackend/backend/config"
	"github.com/Synaxis/battlefieldHeroesBackend/backend/model"
	"github.com/Synaxis/battlefieldHeroesBackend/backend/storage/database"
)

func initConfig() {
	var (
		configFile string
	)
	flag.StringVar(&configFile, "config", ".env", "Path to configuration file")
	flag.Parse()

	gotenv.Load(configFile)
	config.Initialize()
}

func initLogger() {
	logrus.SetLevel(config.LogLevel())
}

func main() {
	initConfig()
	initLogger()

	db, err := database.New()
	if err != nil {
		logrus.Fatal(err)
	}

	// err = CreateSchema(db.NewSession())
	// if err != nil {
	// 	logrus.Fatal(err)
	// }

	tx, err := db.NewSession().Begin()
	if err != nil {
		logrus.Fatal(err)
	}

	err = CreateServer(tx)
	if err != nil {
		logrus.Fatal(err)
	}

	p := Player{
		Username:  "SomeUser",
		Password:  "admin1",
		sessionID: "topsecret",
		HeroNames: []string{"FirstHero", "SecondHero"},
	}
	err = CreatePlayer(&p, tx)
	if err != nil {
		logrus.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		logrus.Fatal(err)
	}
}

func CreateSchema(sess *dbr.Session) (err error) {
	_, err = sess.Exec(`CREATE TABLE heroes (
		heroID INT(11) NOT NULL AUTO_INCREMENT,
		hero_name VARCHAR(50) NOT NULL,
		user_id INT(11) NOT NULL,
		hero_stats TEXT NOT NULL,
		PRIMARY KEY (heroID)
		)
		COLLATE='latin1_swedish_ci'
		;
	`)
	if err != nil {
		return err
	}

	_, err = sess.Exec(`CREATE TABLE players (
			user_id INT(11) NOT NULL AUTO_INCREMENT,
			username VARCHAR(50) NULL DEFAULT NULL,
			password VARCHAR(50) NULL DEFAULT NULL,
			sessionID VARCHAR(50) NULL DEFAULT NULL,
			PRIMARY KEY (user_id)
		)
		COLLATE='latin1_swedish_ci'
		;
	`)
	if err != nil {
		return err
	}

	_, err = sess.Exec(`CREATE TABLE servers (
			server_id INT(11) NOT NULL AUTO_INCREMENT,
			soldier_name VARCHAR(50) NULL DEFAULT NULL,
			account_username VARCHAR(50) NULL DEFAULT NULL,
			account_password VARCHAR(50) NULL DEFAULT NULL,
			api_key VARCHAR(50) NULL DEFAULT NULL,
			PRIMARY KEY (server_id)
		)
		COLLATE='latin1_swedish_ci'
		;
	`)
	if err != nil {
		return err
	}

	return nil
}

func CreateServer(tx *dbr.Tx) error {
	server := model.Server{
		SoldierName:     "Test-Server",
		AccountUsername: "Test-Server",
		AccountPassword: "Test-Server",
		APIKey:          "SERVER-APIKEY",
	}

	err := model.InsertServer(tx, &server)
	if err != nil {
		logrus.Error(tx.Rollback())
		return err
	}
	return nil
}

type Player struct {
	ID int

	Username  string
	Password  string
	sessionID string
	HeroNames []string
}

func CreatePlayer(p *Player, tx *dbr.Tx) error {
	player := model.Player{
		Username:  p.Username,
		Password:  p.Password,
		sessionID: p.sessionID,
	}

	err := model.InsertPlayer(tx, &player)
	if err != nil {
		logrus.Error(tx.Rollback())
		return err
	}

	for _, heroName := range p.HeroNames {
		err = CreateHero(tx, player.ID, heroName)
		if err != nil {
			logrus.Error(tx.Rollback())
			return err
		}
	}

	p.ID = player.ID
	return nil
}

func CreateHero(tx *dbr.Tx, playerID int, heroName string) error {
	by, err := json.Marshal(model.NewHeroStats())
	if err != nil {
		return err
	}

	hero := model.Hero{
		heroID:    playerID,
		HeroName:  heroName,
		HeroStats: string(by),
	}

	err = model.InsertHero(tx, &hero)
	if err != nil {
		return err
	}

	return nil
}
