default: install

install:
  go install
  gf completion zsh > /opt/homebrew/share/zsh-completions/_gf
