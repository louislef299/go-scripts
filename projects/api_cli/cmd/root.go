/*
Copyright © 2022 Louis Lefebvre <lefeb073@umn.com>
*/
package cmd

import (
	"log"
	"os"
	"strings"

	mdtlog "github.com/louislef299/go-scripts/projects/mlctl/internal/log"
	"github.com/louislef299/go-scripts/projects/mlctl/internal/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile, logLvl, logFile string
	Log                      *mdtlog.Logger
)

func NewCmdRoot() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "clctl",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
	examples and usage of using your application. For example:
	
	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		Version: version.String(),
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			lvl := levelPrecedence(logLvl)
			Log = mdtlog.SetLogLevel(lvl, logFile)
			Log.Trace.Println("loglevel", lvl)
		},
	}

	// initialize viper config
	cobra.OnInitialize(initConfig)

	viper.BindEnv("loglvl")
	viper.BindEnv("logfile")
	viper.BindPFlag("loglevel", cmd.PersistentFlags().Lookup("loglevel"))
	viper.BindPFlag("logfile", cmd.PersistentFlags().Lookup("logfile"))
	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.clctl/config.yaml)")
	cmd.PersistentFlags().StringVar(&logLvl, "loglevel", "ERROR", "the log level for the given command [trace, info, warning, error]")
	cmd.PersistentFlags().StringVar(&logFile, "logfile", "", "path to log file, if none, outputs to screen")

	return cmd
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetEnvPrefix("MDT")
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		if _, err = os.Stat(home + "/.clctl"); os.IsNotExist(err) {
			os.Mkdir(home+"/.clctl", 0744)
		}
		if _, err = os.Stat(home + "/.clctl/config"); os.IsNotExist(err) {
			f, err := os.Create(home + "/.clctl/config")
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
		}
		// Search config in home directory with name ".clctl" (without extension).
		viper.AddConfigPath(home + "/.clctl")
		viper.SetConfigType("toml")
		viper.SetConfigName("config")
	}
	viper.AutomaticEnv() // read in environment variables that match
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("could not locate the config file")
	}
}

// Compares the environment variable to the user flag to find which log level to choose.
// The lowest log level takes precedence.
func levelPrecedence(lvl string) string {
	viperLvl := viper.Get("loglvl")
	usrLvl := strings.ToUpper(lvl)
	switch {
	case viperLvl == "TRACE" || usrLvl == "TRACE":
		return "TRACE"
	case viperLvl == "INFO" || usrLvl == "INFO":
		return "INFO"
	case viperLvl == "WARNING" || usrLvl == "WARNING":
		return "WARNING"
	case viperLvl == "ERROR" || usrLvl == "ERROR":
		return "ERROR"
	default:
		return usrLvl
	}
}
