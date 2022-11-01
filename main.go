package main

import (
	"code/gen/pkg/cmd"
	"code/gen/util/logger"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	rootCmd := cmd.NewCmdRoot()
	if err := rootCmd.Execute(); err != nil {
		logger.Log.WithFields(logrus.Fields{"err": err}).Error("client start fail")
		fmt.Println(err)
		os.Exit(1)
	}
}