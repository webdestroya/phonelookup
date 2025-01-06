package config

const (
	CfgName = `.phonelookup`
	CfgType = `yml`
)

const (
	LookupLineTypeIntelligence = `lookup.line_type_intelligence`
	LookupCallerName           = `lookup.caller_name`

	LookupCountryCode = `lookup.country_code`
	LookupExtraFields = `lookup.extra_fields`

	LookupOutputJSON = `lookup.output_json`
	LookupHideErrors = `lookup.hide_errors` // hide data packages that have an error
)

const (
	TwilioUsername = `twilio.username`
	TwilioPassword = `twilio.password`
	TwilioTimeout  = `twilio.timeout`
	// TwilioEdge       = `twilio.edge` // use env var
	// TwilioRegion     = `twilio.region` // use env var
)
