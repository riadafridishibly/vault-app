.PHONY: dev build ent

dev:
	wails dev

module:
	wails generate module

build:
	wails build

ent:
	go generate ./ent
