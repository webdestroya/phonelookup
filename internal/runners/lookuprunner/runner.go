package lookuprunner

import (
	"encoding/json"
	"fmt"
	"io"
	"maps"
	"reflect"
	"slices"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/webdestroya/phonelookup/internal/config"
	"github.com/webdestroya/phonelookup/internal/twilio"
)

var (
	headingStyle = lipgloss.NewStyle().Bold(true)
	errorStyle   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#ff0000"))

	resultType = reflect.TypeOf(twilio.LookupsV2PhoneNumber{})
)

type Runner struct{}

func (lr Runner) Run(cmd *cobra.Command, args []string) error {

	client, err := twilio.NewLookupsClient(nil)
	if err != nil {
		return err
	}

	multiple := len(args) > 1

	results := make([]*twilio.LookupsV2PhoneNumber, 0, len(args))

	hideIfError := viper.GetBool(config.LookupHideErrors)

	fields := viper.GetStringSlice(config.LookupExtraFields)
	if viper.GetBool(config.LookupCallerName) {
		fields = append(fields, "caller_name")
	}
	if viper.GetBool(config.LookupLineTypeIntelligence) {
		fields = append(fields, "line_type_intelligence")
	}
	slices.Sort(fields)
	fields = slices.Compact(fields)

	params := &twilio.FetchPhoneNumberParams{}
	params.SetFields(strings.Join(fields, ","))
	params.SetCountryCode(viper.GetString(config.LookupCountryCode))

	for _, phoneNumber := range args {

		resp, err := client.FetchPhoneNumber(phoneNumber, params)
		if err != nil {
			return err
		}

		results = append(results, resp)
	}

	if viper.GetBool(config.LookupOutputJSON) {

		out, err := json.MarshalIndent(results, "", "  ")
		if err != nil {
			return err
		}

		fmt.Fprintln(cmd.OutOrStdout(), string(out))

		return nil
	}

	extraFields := viper.GetStringSlice(config.LookupExtraFields)
	slices.Sort(extraFields)

	for i, resp := range results {

		if multiple && i > 0 {
			fmt.Fprint(cmd.OutOrStdout(), "\n\n")
		}

		fmt.Fprintln(cmd.OutOrStdout(), headingStyle.Render(fmt.Sprintf("Phone Number: %s", *resp.PhoneNumber)))

		if resp.ValidationErrors != nil && len(*resp.ValidationErrors) > 0 {
			fmt.Fprintln(cmd.OutOrStdout(), errorStyle.Render("  Error: "+strings.Join(*resp.ValidationErrors, ", ")))
			continue
		}

		if resp.CountryCode != nil && *resp.CountryCode != "" {
			fmt.Fprintf(cmd.OutOrStdout(), "  Country: %s\n", *resp.CountryCode)
		}

		lr.printSection(cmd, "Caller Name", resp.CallerName, false, printCallerName)
		lr.printSection(cmd, "Line Type Intelligence", resp.LineTypeIntelligence, false, printLineTypeIntel)

		r := reflect.ValueOf(resp).Elem()

		for i := 0; i < r.NumField(); i++ {
			field := r.Field(i)
			tagName, _, _ := strings.Cut(resultType.Field(i).Tag.Get("json"), ",")
			if slices.Contains(extraFields, tagName) {
				lr.printSection(cmd, tagName, field.Interface().(*any), hideIfError, nil)
			}
		}

	}

	return nil
}

func (Runner) printSection(cmd *cobra.Command, label string, data *any, hideIfError bool, displayFunc func(io.Writer, map[string]any)) {
	if data == nil || *data == nil {
		return
	}

	stdout := cmd.OutOrStdout()

	heading := "\n" + headingStyle.Render("  "+label)

	entry, ok := (*data).(map[string]any)
	if !ok {
		fmt.Fprintln(stdout, heading)
		fmt.Fprintf(stdout, "    Error: Invalid data type in response: %T\n", *data)
		return
	}

	if e, ok := entry["error_code"]; ok && e != nil {
		if hideIfError {
			return
		}
		fmt.Fprintln(stdout, heading)
		fmt.Fprintf(stdout, "    Error Code: %s\n", twilio.GetErrorString(e))
		return
	}

	if displayFunc != nil {
		fmt.Fprintln(stdout, heading)
		displayFunc(stdout, entry)
		return
	}

	keys := slices.Collect(maps.Keys(entry))
	slices.Sort(keys)

	fmt.Fprintln(stdout, heading)
	for _, key := range keys {
		switch v := entry[key].(type) {
		case string:
			fmt.Fprintf(stdout, "   %s: %s\n", key, v)
		case nil:
			// nothing
		default:
			fmt.Fprintf(stdout, "   %s: %v\n", key, v)
		}
	}

}

func printCallerName(out io.Writer, entry map[string]any) {
	if v, ok := entry["caller_name"].(string); ok {
		fmt.Fprintf(out, "    Caller Name: %s\n", v)
	}

	if v, ok := entry["caller_type"].(string); ok {
		fmt.Fprintf(out, "    Caller Type: %s\n", v)
	}
}

func printLineTypeIntel(out io.Writer, entry map[string]any) {
	if v, ok := entry["carrier_name"].(string); ok {
		fmt.Fprintf(out, "    Carrier: %s\n", v)
	}

	if v, ok := entry["type"].(string); ok {
		fmt.Fprintf(out, "    Type: %s\n", v)
	}

	if v, ok := entry["mobile_country_code"].(string); ok {
		fmt.Fprintf(out, "    Mobile Country Code: %s\n", v)
	}

	if v, ok := entry["mobile_network_code"].(string); ok {
		fmt.Fprintf(out, "    Mobile Network Code: %s\n", v)
	}
}
