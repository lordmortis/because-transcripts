package httpServer

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/errgo.v2/errors"
	"html/template"
	"net/http"
	"os"

	"BecauseLanguageBot/config"
	"BecauseLanguageBot/httpServer/templateData"
)

type HttpServer struct {
	defaultRouter *gin.Engine
	bindAddress   string
	httpServer    *http.Server
	devMode       bool
}

type MiddlewareFunc func() (gin.HandlerFunc, error)

func Init(config config.HttpConfig) (*HttpServer, error) {

	var server HttpServer

	server.bindAddress = fmt.Sprintf("%s:%d", config.BindAddress, config.Port)
	server.defaultRouter = gin.Default()

	return &server, nil
}

func (server *HttpServer) SetDevelopmentMode(devMode bool) error {
	server.devMode = devMode
	if devMode {
		server.defaultRouter.LoadHTMLGlob("httpServer/templates/*")
	} else {
		t, err := loadBinTemplates()
		if err != nil {
			return errors.Because(err, nil, "could not load binary templates")
		}
		server.defaultRouter.SetHTMLTemplate(t)
	}

	return nil
}

func loadBinTemplates() (*template.Template, error) {
	t := template.New("")

	for _, name := range templateData.AssetNames() {
		file, err := templateData.AssetInfo(name)
		if err != nil {
			os.Stderr.WriteString(fmt.Sprintf("Unable to get info on template file '%s' - %s\n", name, err))
			continue
		}

		if file.IsDir() {
			continue
		}

		data, err := templateData.AssetString(name)
		if err != nil {
			os.Stderr.WriteString(fmt.Sprintf("Unable to read template file '%s' - %s\n", name, err))
			continue
		}
		t, err = t.New(name).Parse(data)
		if err != nil {
			os.Stderr.WriteString(fmt.Sprintf("Unable to parse template file '%s' into template - %s\n", name, err))
			continue
		}
	}

	return t, nil
}

func (server *HttpServer) Start(middlewares ...MiddlewareFunc) error {
	for index, middleware := range middlewares {
		instancedMiddleware, err := middleware()
		if err != nil {
			return errors.Because(err, nil, fmt.Sprintf("unable to instance middleware %d", index))
		}
		server.defaultRouter.Use(instancedMiddleware)
	}

	server.defaultRouter.GET("/", handleIndex)
	server.defaultRouter.GET("/episode/:id", handleEpisode)
	server.defaultRouter.GET("/speaker/:id", handleSpeaker)
	server.defaultRouter.GET("/search", handleSearch)

	server.httpServer = &http.Server{
		Addr:    server.bindAddress,
		Handler: server.defaultRouter,
	}

	go func() {
		if err := server.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			os.Stderr.WriteString(fmt.Sprintf("Http server listen error: %s\n", err))
		}
	}()

	return nil
}

func (server *HttpServer) Stop() error {
	if err := server.httpServer.Close(); err != nil {
		return errors.Because(err, nil, "Could not stop server")
	}

	return nil
}
