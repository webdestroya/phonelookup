# Phone Lookup

Just a simple utility to return caller ID info from a phone number.

## Usage

```
$ phonelookup +12345678900

Phone Number: +12345678900
  Country: US

  Caller Name
    Caller Type: UNDETERMINED

  Line Type Intelligence
    Carrier: Bandwidth/13 - Bandwidth.com - SVR
    Type: nonFixedVoip
    Mobile Country Code: 313
    Mobile Network Code: 981
```


## Config
Configuration is done via a config file, `~/.phonelookup.yml`

```yaml

# Twilio auth settings. REQUIRED
twilio:
  # This is your Account SID
  username: AC.....

  # This will likely be the Auth Token
  password: xxxxxxxx

lookup:
  # OPTIONAL
  # Specify the default country code.
  # if not provided, will use Twilio's default (usually US)
  country_code: US

  # Extra fields you wish to include in addition to caller_name
  # and line_type_intelligence
  #
  # Available values are on: https://www.twilio.com/docs/lookup/v2-api#data-packages
  extra_fields:
    - sms_pumping_risk

  # If true, then only the JSON responses will be returned
  # You can optionally set this using the --json command flag
  #
  output_json: false

  # If you provide extra fields that you do not have access to,
  # or you provide a number in a country where a package is not available,
  # Twilio will return an error.
  #
  # Setting this field to true will prevent it from appearing in the results
  # if the response includes an error.
  #
  hide_errors: false

  
  # Whether to perform the 'line_type_intelligence' query
  # You can disable this if you live outside the US
  line_type_intelligence: true

  # Whether to perform the 'caller_name' query
  # You can disable this if you live outside the US
  caller_name: true

```