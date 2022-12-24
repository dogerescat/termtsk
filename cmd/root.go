package cmd

import (
	"fmt"
	"os"
	"os/user"
	"termtsk/database/firebase"
	"termtsk/ui/form"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	Firebase firebase.Config
}

type DB interface {
	Create(task *form.Task) error
	GetAll() []*form.Task
	Update(task *form.Task) error
}

var cfgFile string
var config Config
var database DB
var rootCmd = &cobra.Command{
	Use:   "termtsk",
	Short: "root",
	Long:  "root",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/termtsk/config.toml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		cur, err := user.Current()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		viper.AddConfigPath(cur.HomeDir + "/termtsk")
		viper.SetConfigName("config")
		viper.SetConfigType("toml")
	}
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	database = firebase.NewFirestore(config.Firebase)
}
