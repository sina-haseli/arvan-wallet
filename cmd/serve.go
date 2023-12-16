package cmd

import (
	"arvanwallet/config"
	"arvanwallet/handler"
	"arvanwallet/repositories"
	"arvanwallet/routes"
	"arvanwallet/services"
	"context"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve arvan wallet application",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serve() {
	ca := config.InitializeConfig()
	rep := repositories.NewRepository(ca.DB, ca.RDB)
	ser := services.NewServices(rep)
	hndl := handler.NewBaseHandler(ser)
	go initializeStreamServer(ser, ca, hndl)
	initializeHttpServer(hndl, ca.PORT)
}

func initializeHttpServer(handler *handler.BaseHandler, port string) {
	e := echo.New()
	e.HideBanner = true
	p := prometheus.NewPrometheus("arvanWallet", nil)
	p.Use(e)
	routes.RegisterRoutes(e, handler)
	e.Logger.Fatal(e.Start(":" + port))
}

func initializeStreamServer(service *services.Services, config *config.ConfiguredApp, handler *handler.BaseHandler) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	eventCh := service.Consumer.Consume(ctx, config.Config.App.ComQueueName, 1000)

	handler.Credit.HandleIncreaseRequestFromChannel(eventCh)
}
