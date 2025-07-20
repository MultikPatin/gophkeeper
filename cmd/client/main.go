package main

import (
	"main/internal/client/cli"
	//l "main/internal/logger"
)

func main() {
	//fmt.Printf("Build version: %s\n", buildVersion)
	//fmt.Printf("Build date: %s\n", buildDate)
	//fmt.Printf("Build commit: %s\n", buildCommit)
	//
	//logger := l.GetLogger()
	//defer l.SyncLogger()
	//
	//logger.Info("Initializing client...")

	cli.Execute()
}
