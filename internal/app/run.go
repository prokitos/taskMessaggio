package app

import (
	"fmt"
	"module/internal/server"
	"time"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	Server *fiber.App
}

// создание сервера
func (a *App) NewServer(port string) {
	a.Server = server.ServerStart(port)
}

// пытается выключить сервер, а если не получится, то через 30 секунд экстренно сбросит соединение
func (a *App) Stop() {

	fmt.Println("Gracefully shutting down...")
	a.Server.ShutdownWithTimeout(30 * time.Second)

}
