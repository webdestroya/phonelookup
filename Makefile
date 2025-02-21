
refresh-error-codes:
	curl https://www.twilio.com/docs/api/errors/twilio-error-codes.json > internal/generators/errcodegen/twilio_errors.json

.PHONY: test-release
test-release:
	@go tool goreleaser release --skip publish,sign,docker --clean --snapshot

.PHONY: test
test:
	@go test -v -timeout 30s -tags testmode  ./...


.PHONY: outdated
outdated:
	@go list -u -m -f '{{if not .Indirect}}{{if .Update}}{{.}}{{end}}{{end}}' all
