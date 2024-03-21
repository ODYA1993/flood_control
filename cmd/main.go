package main

import (
	"context"
	"flood_control/internal/config"
	"flood_control/internal/floodcontrol"
	"log"
	"time"
)

func main() {
	cfg := config.New()

	var fc FloodControl
	fc = floodcontrol.NewRedisFloodControl(cfg.RedisAddr, cfg.Window, cfg.Limit)
	defer fc.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	id := int64(1)

	for i := 0; i < 15; i++ {
		passed, err := fc.Check(ctx, id)
		if err != nil {
			log.Printf("error checking Flood_Control: %v", err)
		}
		if !passed {
			log.Println("превышен лимит запросов")
		}
	}

}

// FloodControl интерфейс, который нужно реализовать.
// Рекомендуем создать директорию-пакет, в которой будет находиться реализация.
type FloodControl interface {
	// Check возвращает false если достигнут лимит максимально разрешенного
	// кол-ва запросов согласно заданным правилам флуд контроля.
	Check(ctx context.Context, userID int64) (bool, error)
	Close() error
}
