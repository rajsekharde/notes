#!/usr/bin/bash

set -euo pipefail

echo ""

echo "Listing directory contents:"
ls -l
echo ""

echo "System disk usage:"
df -h
echo ""

echo "System ram usage:"
free -h

echo ""
