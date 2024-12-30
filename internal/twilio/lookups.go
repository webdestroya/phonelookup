package twilio

import (
	"github.com/twilio/twilio-go/client"
	twilioLookups "github.com/twilio/twilio-go/rest/lookups/v2"
)

func NewLookupsClient(c client.BaseClient) (*twilioLookups.ApiService, error) {

	if c == nil {
		cl, err := NewClient()
		if err != nil {
			return nil, err
		}
		c = cl
	}

	return twilioLookups.NewApiServiceWithClient(c), nil
}

type FetchPhoneNumberParams = twilioLookups.FetchPhoneNumberParams
type LookupsV2PhoneNumber = twilioLookups.LookupsV2PhoneNumber
