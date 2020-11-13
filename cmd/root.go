/*
Copyright © 2020 Hugues Obolonsky <hugh@atosc.org>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

//Package cmd for gotsx root command
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	Log "github.com/apatters/go-conlog"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"

	"gotsx/tsxcommand"
)

//Tsx server config
var Tsx tsxcommand.TsxServer

var (
	//Ret used for function return
	Ret             string
	cfgFile         string
	tsxAddress      string
	tsxPort         int
	cfgFileNotFound = false
	verbosity       = "warn"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gotsx",
	Short: "A simple go program to get status from TheSkyX",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

		Tsx = tsxcommand.TsxServer{
			Addr: tsxAddress,
			Port: tsxPort,
		}
		Log.Debug("Server: %s", Tsx.Addr)
		Log.Debug("Port: %s", Tsx.Port)

	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println(Ret)
	},
	Run: func(cmd *cobra.Command, args []string) {

		Ret = tsxcommand.Send(Tsx, tsxcommand.IsMountAtPark, "Is mount at Park1 position ?")

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is conf/gotsx.yaml)")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	setUpLogs()

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Switch to default program path
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatal(err)
		}
		confdir := fmt.Sprintf("%s/conf", dir)
		// we came from bin directory
		confdir1 := fmt.Sprintf("%s/../conf", dir)
		// Search yaml config file in program path with name "gotsx.yaml".
		viper.AddConfigPath(confdir)
		viper.AddConfigPath(confdir1)
		viper.AddConfigPath(dir)
		viper.SetConfigType("yaml")
		viper.SetConfigName("gotsx")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			cfgFileNotFound = true
			fmt.Println("Config file not found")
		} else {
			Log.Debug("Something look strange")
			Log.Debugf("error: %v\n", err)
		}
	} else {
		Log.Debugf("Using config file: %s\n", viper.ConfigFileUsed())
	}

	manageDefault()
}

func setUpLogs() {

	formatter := Log.NewStdFormatter()
	formatter.Options.LogLevelFmt = Log.LogLevelFormatLongTitle
	Log.SetFormatter(formatter)
	switch verbosity {
	case "debug":
		Log.SetLevel(Log.DebugLevel)
	case "info":
		Log.SetLevel(Log.InfoLevel)
	case "warn":
		Log.SetLevel(Log.WarnLevel)
	case "error":
		Log.SetLevel(Log.ErrorLevel)
	default:
		Log.SetLevel(Log.WarnLevel)
	}

}

func manageDefault() {
	tsxAddress = viper.GetString("tsxAddress")
	tsxPort = viper.GetInt("tsxPort")
}
