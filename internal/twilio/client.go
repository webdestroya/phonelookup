package twilio

import (
	"github.com/spf13/viper"
	tclient "github.com/twilio/twilio-go/client"
	"github.com/webdestroya/phonelookup/internal/config"
)

func NewClient() (tclient.BaseClient, error) {
	c := &tclient.Client{}
	c.Credentials = tclient.NewCredentials(viper.GetString(config.TwilioUsername), viper.GetString(config.TwilioPassword))
	c.SetTimeout(viper.GetDuration(config.TwilioTimeout))

	return c, nil
}
