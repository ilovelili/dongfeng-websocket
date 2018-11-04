package app

import (
	"sync"
)

var (
	mu = sync.Mutex{}
	wg = sync.WaitGroup{}
)

// App app object wraps the neccessary domains
type App struct {
	// TBD
	Config *natsws.Config
	Server *natsws.NatsWebSocket
}

// Bootstarp Bootstarp the service
func (app *App) Bootstarp() error {
	myapp, err := app.init()
	if err != nil {
		return err
	}

	return myapp.Server.Start()
}

func (app *App) init() (*App, error) {
	myapp, err := app.initializeConfig()
	if err != nil {
		return myapp, err
	}

	myapp, err = app.initializeServer()
	return myapp, err
}

// initializeConfig init config
func (app *App) initializeConfig() (*App, error) {
	config := util.GetConfig()

	app.Config = &natsws.Config{
		ListenInterface: config.WebSocket.Host,
		JWKS:            config.Auth.JWKS,
		URLPattern:      "/",
		NatsAddress:     config.Nats.Host,
		NatsPoolSize:    config.Nats.GetMaxConnectionCount(),
		NatsTopics:      util.GetNatsTopics(config.Nats.Topics),
	}

	return app, nil
}

// initializeServer init nats websocket server
func (app *App) initializeServer() (*App, error) {
	natsWebsocket := natsws.New(app.Config)
	app.Server = natsWebsocket
	return app, nil
}
