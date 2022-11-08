.PHONY: dev build ent

dev:
	wails dev

build:
	wails build

ent:
	go generate ./ent
