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
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/webdestroya/phonelookup/internal/config"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Shows the current configuration setup that will be used",
	Args:  cobra.NoArgs,
	RunE:  runConfigCmd,
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().Bool("unmask", false, "Do not mask password when displaying configuration")
}

func runConfigCmd(cmd *cobra.Command, args []string) error {

	unmask := false
	if v, err := cmd.Flags().GetBool("unmask"); err == nil {
		unmask = v
	}

	cfgFileName := viper.ConfigFileUsed()
	if len(cfgFileName) == 0 {
		cfgFileName = "NO CONFIG FILE FOUND"
	}

	out := cmd.OutOrStdout()

	fmt.Fprintf(out, "Config File: %s\n\n", cfgFileName)

	pass := viper.GetString(config.TwilioPassword)
	if !unmask && len(pass) > 4 {
		pass = pass[0:4] + strings.Repeat("*", len(pass)-4)
	}

	fmt.Fprintln(out, "Twilio Information:")
	fmt.Fprintf(out, "  Username: %s\n", viper.GetString(config.TwilioUsername))
	fmt.Fprintf(out, "  Password: %s\n", pass)

	return nil
}
