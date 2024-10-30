default: test install

test:
  go test -v ./...

install:
  go install ./cmd/gf/
  gf completion zsh > /opt/homebrew/share/zsh-completions/_gf

chlog: changelog
changelog:
  git-chglog -o CHANGELOG.md

wasm:
  GOOS=js GOARCH=wasm go build -o ./cmd/wasm/main.wasm ./cmd/wasm/

serve: wasm
  python3 -m http.server -d ./cmd/wasm/ 2040
