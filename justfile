default: test install

test:
  go test -v ./...

install:
  go install ./cmd/gf/
  gf completion zsh > /opt/homebrew/share/zsh-completions/_gf

chlog: changelog
changelog:
  git-chglog -o CHANGELOG.md
