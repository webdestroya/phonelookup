package config

import (
	"errors"

	"github.com/spf13/viper"
)

func Valid(v *viper.Viper) error {
	if !v.IsSet(TwilioUsername) {
		return errors.New("missing Twilio username")
	}

	if !v.IsSet(TwilioPassword) {
		return errors.New("missing Twilio password")
	}

	return nil
}
