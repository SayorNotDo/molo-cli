package main

import (
	"context"
	"embed"
	"errors"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"go.uber.org/zap"
	"log"
	Mlog "molo-cli/backend/pkg/log"
	"molo-cli/backend/pkg/proof"
	"molo-cli/backend/router"
	"net/http"
	"os/signal"
	"syscall"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	/* Proof init */
	proof.MustInitProof()

	/* Logger init */
	zap.S().Debug("Initializing Logger")

	Mlog.InitLogger()

	/* 创建应用实例 */
	app := NewApp()

	//AppMenu := menu.NewMenu()

	/* 后端使用Gin来提供服务 */
	go func() {
		ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer stop()

		r := router.Init()

		srv := &http.Server{
			Addr:    ":10280",
			Handler: r,
		}
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal("Server forced to shutdown: ", err)
		}
	}()

	/* 主应用程序由对wails.Run()的调用组成，接受描述应用程序窗口大小、窗口标题、用使用的资源等应用程序配置 */
	err := wails.Run(&options.App{
		Title: "Molo",
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: true,
				UseToolbar:                 true,
				HideTitle:                  true,
			},
		},
		DisableResize: true,
		//Frameless:     true,
		Width:  1280,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets, // 应用程序的前端资产
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,  // 创建窗口并即将开始加载前端资源时的回调
		OnShutdown:       app.shutdown, // 应用程序即将退出时的回调
		//Menu:             AppMenu,
		// 向前端暴露的结构体实例
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
