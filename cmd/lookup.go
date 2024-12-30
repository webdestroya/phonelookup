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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/webdestroya/phonelookup/internal/config"
	"github.com/webdestroya/phonelookup/internal/runners/lookuprunner"
)

// lookupCmd represents the lookup command
var lookupCmd = &cobra.Command{
	Use:   "lookup PHONENUMBER",
	Short: "Lookup CallerID information for one or more phone numbers",
	Long: `Queries Twilio LookupV2 API endpoint to gather information about a list of phone numbers

By default, only the caller_name and line_type_intelligence are queried. You can add more fields if you want.
`,
	Args:    cobra.MinimumNArgs(1),
	Example: `phone`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return (lookuprunner.Runner{}).Run(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(lookupCmd)

	lookupCmd.Flags().Bool("json", false, "Output results in JSON format")
	lookupCmd.Flags().String("country", "", "ISO-3166 country code. Used if the phone number provided is in national format")
	lookupCmd.Flags().StringSlice("fields", []string{}, "Extra fields to include (line_status, sms_pumping_risk, etc)")

	viper.BindPFlag(config.LookupOutputJSON, lookupCmd.Flags().Lookup("json"))
	viper.BindPFlag(config.LookupCountryCode, lookupCmd.Flags().Lookup("country"))
	viper.BindPFlag(config.LookupExtraFields, lookupCmd.Flags().Lookup("fields"))

}
