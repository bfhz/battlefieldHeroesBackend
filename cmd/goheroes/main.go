package main

// Generate setters and getters for stats
// $ go generate ./cmd/goheroes
//go:generate go run ../stats-codegen/main.go -scan ../../backend/model --getters ../../backend/ranking/getters.go --setters ../../backend/ranking/setters.go --adders ../../backend/ranking/adders.go

import (
	"context"
	"flag"

	"github.com/google/gops/agent"
	"github.com/sirupsen/logrus"
	"github.com/subosito/gotenv"

	"github.com/Synaxis/battlefieldHeroesBackend/backend/config"
	"github.com/Synaxis/battlefieldHeroesBackend/backend/network"
	"github.com/Synaxis/battlefieldHeroesBackend/backend/fesl"
	"github.com/Synaxis/battlefieldHeroesBackend/backend/matchmaking"
	"github.com/Synaxis/battlefieldHeroesBackend/backend/storage/database"
	"github.com/Synaxis/battlefieldHeroesBackend/backend/storage/kvstore"
	"github.com/Synaxis/battlefieldHeroesBackend/backend/theater"
)

func main() {
	initConfig()
	initLogger()

	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()
	startServer(ctx)

	// Use "github.com/google/gops/agent"
	if err := agent.Listen(agent.Options{}); err != nil {
		logrus.Fatal(err)
	}

	logrus.Println("Serving...")
	<-ctx.Done()
}

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

	// logrus.SetFormatter(&logrus.JSONFormatter{
	// 	DisableTimestamp: true,
	// })
	// logrus.SetFormatter(new(prefixed.TextFormatter))
	// logrus.SetFormatter(&prefixed.TextFormatter{
	// 	DisableTimestamp: true,
	// 	DisableColors:    true,
	// })
}

func startServer(ctx context.Context) {
	db, err := database.New()
	if err != nil {
		logrus.Fatal(err)
	}

	network.InitClientData()
	kvs := kvstore.NewInMemory()
	mm := matchmaking.NewPool()

	fesl.New(config.FeslClientAddr(), false, db, mm).ListenAndServe(ctx)
	fesl.New(config.FeslServerAddr(), true, db, mm).ListenAndServe(ctx)

	theater.New(config.ThtrClientAddr(), db, kvs, mm).ListenAndServe(ctx)
	theater.New(config.ThtrServerAddr(), db, kvs, mm).ListenAndServe(ctx)
}
