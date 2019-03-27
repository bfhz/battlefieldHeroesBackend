package main

import (
	"context"
	"flag"

	"github.com/sirupsen/logrus"
	"github.com/subosito/gotenv"

	"github.com/Synaxis/battlefieldHeroesBackend/magma/config"
	"github.com/Synaxis/battlefieldHeroesBackend/magma/server"
)

var (
	configFile string
)

func main() {
	initConfig()
	initLogger()

	ctx := context.Background()

	srv, err := server.New()
	if err != nil {
		logrus.Fatal(err)
	}

	srv.ListenAndServe(
		config.Config.HTTPBind,
		config.Config.HTTPSBind,
		config.Config.CertificatePath,
		config.Config.PrivateKeyPath,
	)

	// TODO: Add "github.com/google/gops/agent" to analyze resources
	// if err := agent.Listen(&agent.Options{}); err != nil {
	// 	logrus.Fatal(err)
	// }

	logrus.Info("Listening API requests")
	<-ctx.Done()
	logrus.Info("Exiting...")
}

func initConfig() {
	// Custom path to configuration file
	flag.StringVar(&configFile, "config", ".env", "Path to configuration file")
	flag.Parse()

	// Override env variables
	gotenv.Load(configFile)

	// Initialize config.* public variables
	config.LoadToMemory()
}

func initLogger() {
	logrus.SetLevel(config.LogLevel())
	logrus.SetFormatter(&logrus.JSONFormatter{
		DisableTimestamp: true,
		PrettyPrint: true,
	})
}
