package main

import (
	"flag"

	"github.com/Lofty-Brambles/dis-economy-bot/bot"
)

var (
	pathToConfig *string
	startInDev   *bool
	syncCommands *bool
	syncDatabase *bool
)

func init() {
	pathToConfig = flag.String("config-path", "./config.json", "Use this to set the path to the config file")
	startInDev = flag.Bool("dev", false, "Use this to start the bot in development mode")
	syncCommands = flag.Bool("sync-commands", false, "Use this to sync commands on start")
	syncDatabase = flag.Bool("sync-db", false, "Use this to sync the database on start")
	flag.Parse()
}

func main() {
	
}
