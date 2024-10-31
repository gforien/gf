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

current_version := `gf git getVersion`
next_major := `gf git releaseMajor`
next_minor := `gf git releaseMinor`
next_patch := `gf git releasePatch`

major:
  @just tag {{ next_major }}
minor:
  @just tag {{ next_minor }}
patch:
  @just tag {{ next_patch }}

tag next_version:
  @echo "Current version: {{ current_version}}"
  @echo "Releasing new minor: {{ next_version}}"
  git-chglog -o CHANGELOG.md --next-tag {{ next_version}}
  git add CHANGELOG.md
  git commit -m "chore: release {{ next_version}}"

release:
  gh release create {{ current_version }} --generate-notes
