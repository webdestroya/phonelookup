/*
Copyright © 2024 Mitch Dempsey <webdestroya@users.noreply.github.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/webdestroya/phonelookup/internal/config"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "phonelookup",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	setDefaultCommandIfNonePresent()
	err := rootCmd.Execute()
	if err != nil {

		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// viper.SetDefault("twilio.timeout", 5*time.Second)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.phonelookup.yml)")

	rootCmd.PersistentFlags().String("twilio-sid", "", "Twilio Account SID")
	rootCmd.PersistentFlags().String("twilio-token", "", "Twilio Auth Token")

	viper.BindPFlag(config.TwilioAccountSid, rootCmd.PersistentFlags().Lookup("twilio-sid"))
	viper.BindPFlag(config.TwilioAuthToken, rootCmd.PersistentFlags().Lookup("twilio-token"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".phonelookup" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".phonelookup")
	}

	viper.AutomaticEnv() // read in environment variables that match

	config.SetDefaults(viper.GetViper())

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {

		// fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func setDefaultCommandIfNonePresent() {
	if len(os.Args) > 1 {
		potentialCommand := os.Args[1]
		for _, command := range subCommands() {
			if command == potentialCommand {
				return
			}
		}
		os.Args = append([]string{os.Args[0], lookupCmd.Name()}, os.Args[1:]...)
	}

}

func subCommands() (commandNames []string) {
	for _, command := range rootCmd.Commands() {
		commandNames = append(commandNames, append(command.Aliases, command.Name())...)
	}
	return
}
