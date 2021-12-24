// Copyright © 2021 NAME HERE <EMAIL ADDRESS>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/WalterWj/tidb-tools-ops/common/version"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	cfgFile      string
	printVersion bool // not using cobra.Command.Version to make it possible to show component versions
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tidb-tools-ops",
	Short: "tidb tools",
	Long:  `tidb tools`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if printVersion && len(args) == 0 {
			fmt.Printf("version: %s\n", version.NewToolsVersion()["version"])
			fmt.Println(version.NewToolsVersion()["goVersion"])
			fmt.Printf("GitHash: %s\n", version.NewToolsVersion()["build"])
			fmt.Printf("Build Time: %s\n", version.NewToolsVersion()["buildTime"])
		} else if len(args) == 0 {
			fmt.Println(`Use "tidb-tools-ops --help" or "tidb-tools-ops -h" for more information about a command.`)
		}
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

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tidb-tools-ops.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolVarP(&printVersion, "version", "v", false, "Print the version of tidb-tools-ops")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".tidb-tools-ops" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".tidb-tools-ops")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

// other var
var (
	// statusdump & analyze
	dbhost, dbname, dbusername, dbpassword, dbtable   string
	dbport, dbStatusPort, mode, thread, stats_healthy int
	// export
	host, username, password string
	port                     int
)
