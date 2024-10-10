default: test install

test:
  go test -v ./...

install:
  go install
  gf completion zsh > /opt/homebrew/share/zsh-completions/_gf
