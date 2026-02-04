#!/usr/bin/env bash

set -euo pipefail

RED='\e[31m'
GREEN='\e[32m'
YELLOW='\e[33m'
BLUE='\e[34m'
RESET='\e[0m'

for i in {1..10}; do
  case $((i % 4)) in
    0) color="$RED" ;;
    1) color="$GREEN" ;;
    2) color="$YELLOW" ;;
    3) color="$BLUE" ;;
  esac

  printf "\r%bNumber: %2d%b" "$color" "$i" "$RESET"
  sleep 0.5
done

printf "\n"
