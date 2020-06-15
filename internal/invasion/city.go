package invasion

import (
	"math/rand"
	"strings"

	"github.com/Pantani/errors"
)

type (
	// Path represents the path between two cities.
	Path string
	// City represents the city object.
	City struct {
		Name  string
		North Path
		South Path
		East  Path
		West  Path
	}
	// Route represents the name of the route.
	Route string
)

const (
	// RouteInvalid invalid route.
	RouteInvalid Route = "invalid"
	// RouteNorth north route.
	RouteNorth Route = "north"
	// RouteSouth south route.
	RouteSouth Route = "south"
	// RouteEast east route.
	RouteEast Route = "east"
	// RouteWest west route.
	RouteWest Route = "west"
)

// getRandomRouteAndCity get a random city with route.
func getRandomRouteAndCity(c City) (Route, string) {
	// return invalid if doesn't have movement.
	if !c.hasMovement() {
		return RouteInvalid, ""
	}
	// create a list of route to choose one randomly.
	routes := make([]Route, 0)
	names := make([]Path, 0)
	if c.North.hasPath() {
		routes = append(routes, RouteNorth)
		names = append(names, c.North)
	}
	if c.South.hasPath() {
		routes = append(routes, RouteSouth)
		names = append(names, c.South)
	}
	if c.West.hasPath() {
		routes = append(routes, RouteWest)
		names = append(names, c.West)
	}
	if c.East.hasPath() {
		routes = append(routes, RouteEast)
		names = append(names, c.East)
	}
	randomIndex := rand.Intn(len(routes))
	return routes[randomIndex], string(names[randomIndex])
}

// destroy destroy the path.
func (p *Path) destroy() {
	*p = ""
}

// hasPath verify the path exist.
func (p *Path) hasPath() bool {
	return len(*p) > 0
}

// hasMovement verify the city have possible movements.
func (c *City) hasMovement() bool {
	if c.North.hasPath() {
		return true
	}
	if c.South.hasPath() {
		return true
	}
	if c.West.hasPath() {
		return true
	}
	if c.East.hasPath() {
		return true
	}
	return false
}

// newCity creates a new City object based in the row of the file.
// It returns a City object and an error if occurs.
func newCity(row string) (*City, error) {
	// split the parameters of row.
	s := strings.Split(row, " ")
	if len(s) < 1 {
		return nil, errors.E("invalid parse for city", errors.Params{"row": row})
	}

	// parse row to parameters.
	c := &City{Name: s[0]}
	for i := 1; i < len(s); i++ {
		text := s[i]
		path := strings.Split(text, "=")
		if len(path) <= 1 {
			return nil, errors.E("invalid path for city", errors.Params{"row": row, "path": text})
		}
		switch strings.ToLower(path[0]) {
		case "north":
			c.North = Path(path[1])
		case "south":
			c.South = Path(path[1])
		case "east":
			c.East = Path(path[1])
		case "west":
			c.West = Path(path[1])
		default:
			return nil, errors.E("invalid path for city", errors.Params{"row": row, "path": text})
		}
	}
	return c, nil
}
