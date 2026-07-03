package server

import (
	"context"
	"napoleon-email/src/app/infrastructure"
	config "napoleon-email/src/config/app"
	"napoleon-email/src/config/firebase"
	"napoleon-email/src/pkg/logger"
	"napoleon-email/src/routes"
	"os"
	"os/signal"
	"syscall"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/cobra"
)

var RunServerCmd = &cobra.Command{
	Use: "server",
	Short: "Run the Server",
	Run: func(cmd *cobra.Command, args []string) {
		firestore, err := firebase.ConnectionFirestore(context.Background(), config.GoogleApplicationCredentials())
		if err != nil {
			logger.LogError("error connecting to firestore", err, logger.LogStruct{Action: "error_connecting_firestore", User: 0})
			os.Exit(1)
		}
		infrastructure.InitKernel(firestore)
		port := config.AppPort()
		app := fiber.New(fiber.Config{
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  60 * time.Second,
		})

		app.Use(cors.New(cors.Config{
			AllowOrigins:     "http://localhost:5500, http://localhost:3000, http://127.0.0.1:5500, http://localhost:5173",
			AllowMethods:     "GET,POST,OPTIONS",
			AllowHeaders:     "Content-Type,Authorization",
			AllowCredentials: true,
		}))

		routes.Router(app)
		StartServer(app, ":"+port, firestore)
	},
}

func StartServer(app *fiber.App, addr string, fb *firestore.Client) {
	go waitForShutdown(app, fb)

	logger.LogInfo("server starting", logger.LogStruct{Action: "server_starting"})

	if err := app.Listen(addr); err != nil {
		logger.LogError("server stopped with error", err, logger.LogStruct{Action: "server_stopped"})
		os.Exit(1)
	}
	logger.LogInfo("server stopped gracefully", logger.LogStruct{Action: "server_stopped_gracefully"})
}

func waitForShutdown(app *fiber.App, fb *firestore.Client) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit,
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGINT,
	)

	sig := <-quit
	logger.LogInfo("signal received, shutting down", logger.LogStruct{Action: "signal_received", Data: sig.String()})

	shutdown(app, fb)
}

func shutdown(app *fiber.App, fb *firestore.Client) {
	if err := app.ShutdownWithTimeout(20 * time.Second); err != nil {
		logger.LogError("error shutting down server", err, logger.LogStruct{Action: "error_shutting_down_server"})
	} else {
		logger.LogInfo("server shut down cleanly", logger.LogStruct{Action: "server_shut_down_cleanly"})
	}

	if err := fb.Close(); err != nil {
		logger.LogError("error closing firestore connection", err, logger.LogStruct{Action: "error_closing_firestore_connection"})
	} else {
		logger.LogInfo("firestore connection closed", logger.LogStruct{Action: "firestore_connection_closed"})
	}
}
