package main

import (
	"flag"
	"fmt"
	"github.com/vzina/gin-skeleton/middleware"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/vzina/gin-skeleton/config"
	"github.com/vzina/gin-skeleton/router"
)

func main() {

	addr := flag.String("addr", config.Server.Addr, "Address to listen and serve")
	flag.Parse()

	if config.Server.Mode == gin.ReleaseMode {
		gin.DisableConsoleColor()
	}

	// init logger
	if err := middleware.InitLogger(config.LoggerConfig); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}

	app := gin.Default()

	app.Static("/images", filepath.Join(config.Server.StaticDir, "img"))
	app.StaticFile("/favicon.ico", filepath.Join(config.Server.StaticDir, "img/favicon.ico"))
	app.LoadHTMLGlob(config.Server.ViewDir + "/*")
	app.MaxMultipartMemory = config.Server.MaxMultipartMemory << 20

	router.Route(app)

	// Listen and Serve
	app.Run(*addr)
}
