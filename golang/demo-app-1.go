package main

import (
	"github.com/gorilla/mux"
	"github.com/JustSteveKing/go-api/pkg/application/registry/ping"
	"github.com/JustSteveKing/go-api/pkg/infrastructure/config"
	"github.com/JustSteveKing/go-api/pkg/infrastructure/kernel"
	application "github.com/JustSteveKing/go-api/pkg/infrastructure/app"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
}

func main() {
	app := kernel.Boot() // Create our application

	ping.BuildPingService(app) // Build our services

	// Run our Application in a coroutine
	go func() { app.Run() }()

	// Wait for termination signals and shut down gracefully
	application.WaitForShutdown(app)
}

// Application is our general purpose Application struct
type Application struct {
	Server *http.Server
	Router *mux.Router
	Logger *zap.Logger
	Config *config.Config
}

// Run will run the Application server
func (app *Application) Run() {
	err := app.Server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}

// WaitForShutdown is a graceful way to handle server shutdown events
func WaitForShutdown(application *Application) {
	// Create a channel to listen for OS signals
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive a signal to our channel
	<-interruptChan

	application.Logger.Info("Received shutdown signal, gracefully terminating")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	application.Server.Shutdown(ctx)
	os.Exit(0)
}

// env is a simple helper function to read an environment variable or return a default value
func env(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
