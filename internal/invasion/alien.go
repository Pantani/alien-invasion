package invasion

import (
	"github.com/Pantani/logger"
	"github.com/goombaio/namegenerator"
)

type (
	// AlienStatus represents status of alien.
	AlienStatus string
	// Alien represents the Alien object.
	Alien struct {
		Number uint
		Name   string
		City   string
		Status AlienStatus
	}
	// Aliens represents a list of Alien objects.
	Aliens []*Alien
)

const (
	// Alien status.
	AlienStatusLive AlienStatus = "live"
	AlienStatusDead AlienStatus = "dead"
)

// newAlien creates a new Alien object
// It returns an Alien object.
func newAlien(number uint, city string) *Alien {
	// Generate a random name for Alien.
	nameGenerator := namegenerator.NewNameGenerator(int64(number))
	return &Alien{
		Number: number,
		Name:   nameGenerator.Generate(),
		City:   city,
		Status: AlienStatusLive,
	}
}

// HasAlive verify have alive aliens inside the alien list.
func (a *Aliens) HasAlive() bool {
	for _, alien := range *a {
		if alien.Status == AlienStatusLive {
			return true
		}
	}
	return false
}

// Names return a list of aliens name from a list of aliens object.
func (a *Aliens) Names() []string {
	names := make([]string, len(*a))
	for i, alien := range *a {
		names[i] = alien.Name
	}
	return names
}

// kill kill all aliens from a list.
func (a *Aliens) kill() {
	for _, alien := range *a {
		alien.Status = AlienStatusDead
	}
}

// MoveAliens move aliens to the next city randomly.
func (a *Aliens) MoveAliens(world World) {
	for _, alien := range *a {
		// verify if alien is alive.
		if alien.Status == AlienStatusDead {
			logger.Info("Alien is dead", logger.Params{"number": alien.Number, "name": alien.Name})
			continue
		}
		// get the city
		city := world[alien.City]
		// verify if the city have routes.
		if !city.hasMovement() {
			logger.Info("Alien is locked, no more routes",
				logger.Params{"city": alien.City, "alien": alien.Name})
			continue
		}
		// get a new city to travel.
		route, newCityName := getRandomRouteAndCity(*city)
		logger.Info("moving alien",
			logger.Params{"old city": alien.City, "route": route, "new city": newCityName})
		// set a new city
		alien.City = newCityName
	}
}

// StartWars verify have aliens in the same city, if yes, destroy the city and the aliens.
func (a *Aliens) StartWars(world World) {
	// group aliens by city.
	cities := make(map[string]Aliens)
	for _, alien := range *a {
		cities[alien.City] = append(cities[alien.City], alien)
	}
	// verify have more them one alien in the same city.
	for city, aliens := range cities {
		if len(aliens) <= 1 || !aliens.HasAlive() {
			continue
		}
		logger.Info("war started", logger.Params{"aliens": aliens.Names(), "city": city})
		world.destroyCity(city)
		aliens.kill()
	}
}
