package config

import (
	"time"

	"github.com/spf13/viper"
)

func SetDefaults(v *viper.Viper) *viper.Viper {
	v.SetDefault(TwilioTimeout, 5*time.Second)

	v.SetDefault(LookupCallerName, true)
	v.SetDefault(LookupLineTypeIntelligence, true)

	v.SetDefault(LookupOutputJSON, false)
	v.SetDefault(LookupHideErrors, false)

	v.SetDefault(LookupCountryCode, "US")

	return v
}
