package config

const (
	LookupLineTypeIntelligence = `lookup.line_type_intelligence`
	LookupCallerName           = `lookup.caller_name`

	LookupCountryCode = `lookup.country_code`
	LookupExtraFields = `lookup.extra_fields`

	LookupOutputJSON = `lookup.output_json`
	LookupHideErrors = `lookup.hide_errors` // hide data packages that have an error
)

const (
	TwilioAccountSid = `twilio.account_sid`
	TwilioAuthToken  = `twilio.auth_token`
	TwilioTimeout    = `twilio.timeout`
)
