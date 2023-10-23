# aliases
alias cls='clear'

# Install golangci-lint
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.54.2

# Install taskfile
go install github.com/go-task/task/v3/cmd/task@latest

# Install delve
go install github.com/go-delve/delve/cmd/dlv@master