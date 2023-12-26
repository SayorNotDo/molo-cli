package main

import (
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx        context.Context
	killSignal chan struct{}
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.killSignal = make(chan struct{})
}

func (a *App) shutdown(ctx context.Context) {
	/* 中断信号关闭服务器 */
	close(a.killSignal)
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
