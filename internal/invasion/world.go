package invasion

import (
	"bufio"
	"math/rand"
	"os"
	"reflect"
	"strings"

	"github.com/Pantani/logger"
)

type (
	// World represents the map of world name and city object.
	World map[string]*City
)

// addCity add a city and your routes to the world.
func (w *World) addCity(city *City) {
	// get city.
	(*w)[city.Name] = city
	// verify north has path.
	if city.North.hasPath() {
		key := string(city.North)
		if _, ok := (*w)[key]; !ok {
			// add city from path if not exist.
			(*w)[key] = &City{South: Path(city.Name), Name: key}
		} else {
			// add city to path.
			(*w)[key].South = Path(city.Name)
		}
	}
	// verify south has path.
	if city.South.hasPath() {
		key := string(city.South)
		if _, ok := (*w)[key]; !ok {
			// add city from path if not exist.
			(*w)[key] = &City{North: Path(city.Name), Name: key}
		} else {
			// add city to path.
			(*w)[key].North = Path(city.Name)
		}
	}
	// verify east has path.
	if city.East.hasPath() {
		key := string(city.East)
		if _, ok := (*w)[key]; !ok {
			// add city from path if not exist.
			(*w)[key] = &City{West: Path(city.Name), Name: key}
		} else {
			// add city to path.
			(*w)[key].West = Path(city.Name)
		}
	}
	// verify west has path.
	if city.West.hasPath() {
		key := string(city.West)
		if _, ok := (*w)[key]; !ok {
			// add city from path if not exist.
			(*w)[key] = &City{East: Path(city.Name), Name: key}
		} else {
			// add city to path.
			(*w)[key].East = Path(city.Name)
		}
	}
}

// destroyCity destroy a city from world passing the city name.
func (w *World) destroyCity(name string) {
	// get city to destroy.
	city, ok := (*w)[name]
	// verify city is already destroyed.
	if !ok {
		return
	}
	// destroy north path.
	if city.North.hasPath() {
		(*w)[string(city.North)].South.destroy()
	}
	// destroy south path.
	if city.South.hasPath() {
		(*w)[string(city.South)].North.destroy()
	}
	// destroy east path.
	if city.East.hasPath() {
		(*w)[string(city.East)].West.destroy()
	}
	// destroy west path.
	if city.West.hasPath() {
		(*w)[string(city.West)].East.destroy()
	}
	// delete city.
	city = nil
	delete(*w, name)
	logger.Info("City destroyed", logger.Params{"city": name})
}

// HasMovement verify the world have possible movements.
func (w *World) HasMovement() bool {
	for _, city := range *w {
		if city.hasMovement() {
			return true
		}
	}
	return false
}

// getRandomCity get a random city from world.
// It returns a city pointer.
func (w *World) getRandomCity() *City {
	keys := reflect.ValueOf(*w).MapKeys()
	key := keys[rand.Intn(len(keys))].String()
	return (*w)[key]
}

// GenerateAliens generate aliens based in the number.
// It returns a list of aliens.
func (w *World) GenerateAliens(number uint) Aliens {
	aliens := make(Aliens, 0)
	for i := uint(0); i < number; i++ {
		c := w.getRandomCity()
		alien := newAlien(i, c.Name)
		aliens = append(aliens, alien)
		logger.Info("New alien landed", logger.Params{"number": alien.Number, "name": alien.Name, "city": c.Name})
	}
	return aliens
}

// LoadWorld load world and cities from file name located inside the test folder.
// It returns a map of world and an error if occurs.
func LoadWorld(filename string) (World, error) {
	filename = strings.Trim(filename, "\n")
	logger.Info("Loading world....", logger.Params{"filename": filename})

	//open file.
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// parse cities.
	world := make(World)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if !scanner.Scan() {
			break
		}
		c, err := newCity(text)
		if err != nil {
			return nil, err
		}
		world.addCity(c)
	}
	logger.Info("New world loaded", logger.Params{"cities": len(world)})
	return world, nil
}
