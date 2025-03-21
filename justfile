#!/usr/bin/env -S just --justfile

alias i := install
alias r := run

default:
	@just --list

run *argv: install
	uni {{argv}}

install prefix='~/.local/bin':
	go build -o {{prefix}}/uni main.go
