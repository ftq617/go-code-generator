package cmd

import (
	"code/gen/pkg/cmd/code"
	"code/gen/pkg/cmd/table"
	"code/gen/util/conf"
	"code/gen/util/logger"
	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// initConfig reads in config file and ENV variables if set.
func initConfig(cmd *cobra.Command, _ []string) error {
	cfgFile, err := cmd.Flags().GetString("config")
	if err != nil {
		return err
	}
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".sitekeeper_client" (without extension).
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
	}

	viper.SetEnvPrefix("x_edge") // set environment variables prefix to avoid conflict
	viper.AutomaticEnv()         // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		logger.Log.WithFields(logrus.Fields{"err": err}).Error("Use default config;using config file fail:", viper.ConfigFileUsed())
		viper.SetDefault("database.url", "localhost")
		viper.SetDefault("database.ip", "localhost")
		viper.SetDefault("database.port", "3306")
		viper.SetDefault("database.username", "root")
		viper.SetDefault("database.password", "root")
		viper.SetDefault("database.database", "")
	}
	return conf.ResetData()
}

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "autocode",
		Short:         "autocode CLI",
		Long:          "A command line tool for autocode.",
		SilenceErrors: true,
		SilenceUsage:  true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := initConfig(cmd, args); err != nil {
				return err
			}
			return nil
		},
	}


	persistentFlags := cmd.PersistentFlags()
	persistentFlags.StringP("config", "c", "", "config file (default is $HOME/config.yml)")


	// Register child command
	cmd.AddCommand(table.NewCmdTable())
	cmd.AddCommand(code.NewCmdCode())
	return cmd
}
