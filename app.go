package main

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/logger"
)

var wailslog = logger.NewDefaultLogger()

// App struct
type App struct {
	ctx              context.Context
	dialogService    *DialogService
	passwordService  *PasswordService
	clipboardService *ClipboardService
	databaseService  *DatabaseService
	assetServer      *AssetServerHack
	userService      *UserService
}

// NewApp creates a new App application struct
func NewApp() (*App, error) {
	db, err := NewDatabaseService()
	if err != nil {
		return nil, err
	}

	clip, err := NewClipboardService()
	if err != nil {
		return nil, err
	}

	return &App{
		dialogService:    NewDialog(context.Background()),
		passwordService:  NewPasswordService(),
		assetServer:      NewAssetServer(),
		databaseService:  db,
		clipboardService: clip,
		userService:      &UserService{},
	}, nil
}

func (a *App) shutdown(ctx context.Context) {
	a.databaseService.close()
	a.assetServer.shutdown()
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.dialogService.updateContext(ctx)
	a.databaseService.updateContext(ctx)
	a.assetServer.updateContext(ctx)
	a.assetServer.initializeServer()
	a.userService.setDB(a.databaseService)
}
