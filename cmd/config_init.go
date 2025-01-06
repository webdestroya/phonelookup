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
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/webdestroya/phonelookup/internal/config"
)

// cfgInitCmd represents the init command
var cfgInitCmd = &cobra.Command{
	Use:          "init",
	Short:        "Initialize a default configuration file",
	SilenceUsage: true,
	Args:         cobra.NoArgs,
	RunE:         runCfgInitCmd,
}

func init() {
	configCmd.AddCommand(cfgInitCmd)

	cfgInitCmd.Flags().String("username", "", "Twilio Username")
	cfgInitCmd.Flags().String("password", "", "Twilio Password")
	cfgInitCmd.Flags().String("country", "", "ISO-3166 country code. Used if the phone number provided is in national format")

	cfgInitCmd.Flags().Bool("force", false, "Overwrite existing config if one exists")

	cfgInitCmd.SetHelpFunc(func(c *cobra.Command, s []string) {
		c.Flags().MarkHidden("config")
		c.Parent().HelpFunc()(c, s)
	})
}

func runCfgInitCmd(cmd *cobra.Command, args []string) error {

	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	tmp := viper.New()
	tmp.AddConfigPath(home)
	tmp.SetConfigType(config.CfgType)
	tmp.SetConfigName(config.CfgName)
	tmp.AutomaticEnv()

	cfgPath := path.Join(home, config.CfgName+"."+config.CfgType)

	tmp.BindPFlag(config.TwilioUsername, cmd.Flags().Lookup("username"))
	tmp.BindPFlag(config.TwilioPassword, cmd.Flags().Lookup("password"))

	// attempt to read it, but ignore any errors
	_ = tmp.ReadInConfig()

	if len(tmp.ConfigFileUsed()) > 0 {
		cfgPath = tmp.ConfigFileUsed()
	}

	if !tmp.IsSet(config.TwilioUsername) || !tmp.IsSet(config.TwilioPassword) {
		return errors.New("You must provide both --username and --password for a valid config")
	}

	if country, ce := cmd.Flags().GetString("country"); ce == nil && country != "" {
		if len(country) != 2 {
			return errors.New("Country code should be exactly two characters")
		}
		tmp.Set(config.LookupCountryCode, strings.ToUpper(country))
	}

	if force, fe := cmd.Flags().GetBool("force"); fe == nil && force {
		err = tmp.WriteConfigAs(cfgPath)
	} else {

		err = tmp.SafeWriteConfigAs(cfgPath)
	}

	cobra.CheckErr(err)

	fmt.Fprintf(cmd.OutOrStdout(), "Wrote configuration file to: %s\n", cfgPath)

	return nil
}
