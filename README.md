# uni-cli
My in-house university workflow tool.

See `uni --help` for all CLI options.

## Uni directory
The uni directory is the main directory where this tool operates.
Per default it is expected to be at `~/uni/`.

You can set another uni directory with the `--directory`/`-d` flag.

## Configuration
The mandatory configuration file is `UNI-DIRECTORY/uni-cli.toml`.
It is generally managed by the CLI tool itself but can be edited manually.

## `unicd`
`unicd` is a shell function that uses `uni path` to provide a smarter `cd` for stuff in the uni directory.
It a *very* simple script but needed because `cd` functionality is a shell-builtin feature.

There are shell function scripts provided for POSIX-compliant shells (like `sh`, `bash`, `zsh`, ...) and `fish`.
