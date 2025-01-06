package cmd

import (
	"regexp"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/webdestroya/phonelookup/internal/config"
)

var phoneRegexp = regexp.MustCompile(`^\+?[-0-9]+$`)

func isPhoneNumber(value string) bool {
	if value == "" {
		return false
	}

	return phoneRegexp.MatchString(value)
}

func preRunCheckConfig(cmd *cobra.Command, args []string) error {
	return config.Valid(viper.GetViper())
}
