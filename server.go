package main

import (
	"context"
	"net/http"

	"github.com/riadafridishibly/vault-app/internal/imghandler"
	wailsruntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

// AssetServerHack spwans a new server, this one exists
// due the limitations of AssetsHandler in wails 2.1.0
//
//	TODO: Experiment with wails 2.2.0 assets server
type AssetServerHack struct {
	ctx    context.Context
	server *http.Server
	addr   string
}

func NewAssetServer() *AssetServerHack {
	return &AssetServerHack{
		ctx: context.TODO(),
	}
}

func (a *AssetServerHack) updateContext(ctx context.Context) {
	a.ctx = ctx
}

func (a *AssetServerHack) shutdown() {
	a.server.Shutdown(a.ctx)
}

func (a *AssetServerHack) GetAddress() string {
	return a.addr
}

func (a *AssetServerHack) initializeServer() {
	srv, l, err := imghandler.NewServer()
	if err != nil {
		panic(err)
	}

	a.server = srv
	a.addr = "http://" + l.Addr().String()

	go func() {
		err := a.server.Serve(l)
		if err != nil {
			println(err)
		}
	}()
	wailsruntime.LogDebugf(a.ctx, "Internal Image Hanlder Server Running: %s", a.addr)
}
