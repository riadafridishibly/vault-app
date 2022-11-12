package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/iwpnd/pw"
	"github.com/riadafridishibly/vault-app/pkg/breachchecker"
	"github.com/riadafridishibly/vault-app/pkg/dao"
	"github.com/riadafridishibly/vault-app/pkg/generator"
	"github.com/riadafridishibly/vault-app/pkg/generator/common"
	"github.com/riadafridishibly/vault-app/pkg/imghandler"
	"github.com/wailsapp/wails/v2/pkg/logger"
	wailsruntime "github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.design/x/clipboard"
)

var lg = logger.NewDefaultLogger()

// App struct
type App struct {
	ctx    context.Context
	server *http.Server
	addr   string
	dao    *dao.Dao
}

// NewApp creates a new App application struct
func NewApp() (*App, error) {
	dao, err := dao.NewDao()

	if err != nil {
		return nil, err
	}
	return &App{dao: dao}, nil
}

func (a *App) initializeServer() {
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

func (a *App) shutdown(ctx context.Context) {
	a.dao.Close()
	a.server.Shutdown(ctx)
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	err := clipboard.Init()
	if err != nil {
		lg.Error("Couldn't initialize clipboard")
	}
	a.initializeServer()
}

func (a *App) GetAddress() string {
	return a.addr
}

func (a *App) CopyToClipboard(value string) {
	clipboard.Write(clipboard.FmtText, []byte(value))
}

func (a *App) PasteFromClipboard() string {
	return string(clipboard.Read(clipboard.FmtText))
}

func (a *App) CheckPasswordBreach(password string) (int, error) {
	return breachchecker.GetBreach(password)
}

func (a *App) GenerateRandomPassword(options []string) string {
	var opts []pw.Option
	for _, opt := range options {
		switch opt {
		case "uppercase":
			opts = append(opts, pw.WithUpperCase())
		case "lowercase":
			opts = append(opts, pw.WithLowerCase())
		case "numbers":
			opts = append(opts, pw.WithNumbers())
		case "special":
			opts = append(opts, pw.WithSpecial())
		}
	}

	return pw.NewPassword(20, opts...)
}

func (a *App) OpenDialog() ([]string, error) {
	return wailsruntime.OpenMultipleFilesDialog(a.ctx, wailsruntime.OpenDialogOptions{
		DefaultDirectory:           os.Getenv("HOME"),
		DefaultFilename:            "",
		Title:                      "Select Files",
		ShowHiddenFiles:            false,
		CanCreateDirectories:       false,
		ResolvesAliases:            false,
		TreatPackagesAsDirectories: false,
	})
}

func (a *App) GenerateNewPassword(seed, password string) (string, error) {
	if seed == "" && password == "" {
		return "", nil
	}

	g := generator.Generator{
		Options: common.Options{
			AllowRepeat:    true,
			IncludeSymbols: true,
			IncludeSpace:   true,
			Length:         23,
		},
	}
	pw, err := g.Generate([]byte(seed), []byte(password))
	if err != nil {
		return "", err
	}

	return pw, nil
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
