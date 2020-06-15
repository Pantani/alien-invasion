package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/Pantani/alien-invasion/internal/invasion"

	"github.com/Pantani/logger"
)

var (
	_world      string
	_aliens     uint
	_iterations uint
)

func init() {
	flag.StringVar(&_world, "w", "test/world_1.txt", "world file name")
	flag.UintVar(&_aliens, "a", 10, "number of aliens")
	flag.UintVar(&_iterations, "i", 10000, "number of iterations")
	flag.Parse()
}

func main() {
	// load the world.
	world, err := invasion.LoadWorld(_world)
	if err != nil {
		logger.Error(err)
		return
	}
	// generate new aliens.
	aliens := world.GenerateAliens(_aliens)
	// verify have more them one alien in the same city.
	aliens.StartWars(world)
	for i := uint(0); i < _iterations; i++ {
		// verify the world have possible movements.
		if !world.HasMovement() {
			logger.Info("world doesn't have more movements")
			break
		}
		// verify if have aliens alive.
		if !aliens.HasAlive() {
			logger.Info("world doesn't have invasion alive")
			break
		}
		logger.Info("***************************************************")
		logger.Info("-----------------")
		logger.Info(fmt.Sprintf("ITERATION %d", i))
		logger.Info("-----------------")
		// move aliens to the next city.
		aliens.MoveAliens(world)
		// verify have more them one alien in the same city.
		aliens.StartWars(world)
		time.Sleep(time.Millisecond * 1)
	}
	logger.Info("simulation is finished")
}
