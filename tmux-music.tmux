#!/usr/bin/env bash

CURRENT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PATH="/usr/local/bin:$PATH:/usr/sbin:$PATH:$HOME/bin"

main() {
  $(tmux bind-key -T prefix s run -b "source $CURRENT_DIR/scripts/tmux-music")
}

main
