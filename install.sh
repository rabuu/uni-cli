#!/bin/sh

set -xe

cwd=$(basename $PWD)
[ "$cwd" = "uni-cli" ] || exit 1

go build -o bin/uni
cp bin/uni ~/.local/bin/uni
