.PHONY: generate
## Install tools
all:
#region CODE_REGION(project)
	go install github.com/onsi/ginkgo/v2/ginkgo@v2.23.3
	go install github.com/matryer/moq@v0.5.3
	go install github.com/git-chglog/git-chglog/cmd/git-chglog@v0.15.4
	go install golang.org/x/vuln/cmd/govulncheck@v1.1.4
	go install github.com/abice/go-enum@v0.6.1
#endregion
#region CODE_REGION(service)
#endregion

# those are here because we generate a service in the tests.
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.35.2
	go install code.cestus.io/blaze/cmd/protoc-gen-blaze@v0.7.2
	go install github.com/google/wire/cmd/wire@v0.6.0
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.26.0
	