package invasion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	_cityBar = &City{
		Name:  "Bar",
		South: "Foo",
	}
	_cityBaz = &City{
		Name: "Baz",
		East: "Foo",
	}
	_cityFoo = &City{
		Name:  "Foo",
		North: "Bar",
		South: "Qu-ux",
		West:  "Baz",
	}
	_cityLee = &City{
		Name:  "Lee",
		North: "Yee",
		West:  "Loo",
	}
	_cityLoo = &City{
		Name: "Loo",
		East: "Lee",
	}
	_cityQuux = &City{
		Name:  "Qu-ux",
		North: "Foo",
	}
	_cityYee = &City{
		Name:  "Yee",
		South: "Lee",
	}
)

func TestCity_hasMovement(t *testing.T) {
	tests := []struct {
		name string
		city *City
		want bool
	}{
		{"test Bar", _cityBar, true},
		{"test Baz", _cityBaz, true},
		{"test Foo", _cityFoo, true},
		{"test Lee", _cityLee, true},
		{"test Loo", _cityLoo, true},
		{"test Qu-ux", _cityQuux, true},
		{"test Yee", _cityYee, true},
		{"test Yee", &City{Name: "Test"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.city.hasMovement()
			assert.EqualValues(t, tt.want, got)
		})
	}
}

func TestPath_destroy(t *testing.T) {
	tests := []struct {
		name string
		path Path
	}{
		{"test destroy 1", Path("test1")},
		{"test destroy 2", Path("test2")},
		{"test destroy 3", Path("")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.path.destroy()
			assert.Equal(t, Path(""), tt.path)
		})
	}
}

func TestPath_hasPath(t *testing.T) {
	tests := []struct {
		name string
		path Path
		want bool
	}{
		{"test path 1", Path("test1"), true},
		{"test path 2", Path("test2"), true},
		{"test no path ", Path(""), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.path.hasPath()
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_getRandomRouteAndCity(t *testing.T) {
	tests := []struct {
		name    string
		city    City
		wantErr bool
	}{
		{"test random route and city Yee", *_cityYee, false},
		{"test random route and city for Qu-ux", *_cityQuux, false},
		{"test random route and city for Loo", *_cityLoo, false},
		{"test random route and city for Lee", *_cityLee, false},
		{"test random route and city for Baz", *_cityBaz, false},
		{"test random route and city for Bar", *_cityBar, false},
		{"test random route and city for empty", City{Name: "test"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := getRandomRouteAndCity(tt.city)
			if tt.wantErr {
				assert.Equal(t, got1, RouteInvalid)
				assert.Equal(t, len(got2), 0)
				return
			}
			assert.Greater(t, len(got1), 0)
			assert.Greater(t, len(got2), 0)
		})
	}
}

func Test_newCity(t *testing.T) {
	tests := []struct {
		name    string
		row     string
		want    *City
		wantErr bool
	}{
		{"row 1", "Foo north=Bar west=Baz south=Qu-ux", &City{Name: "Foo", North: "Bar", West: "Baz", South: "Qu-ux"}, false},
		{"row 2", "Bar south=Foo west=Bee north=Lee", &City{Name: "Bar", North: "Lee", West: "Bee", South: "Foo"}, false},
		{"row 3", "Lee north=Yee west=Loo", &City{Name: "Lee", North: "Yee", West: "Loo"}, false},
		{"row 4", "Mee north=Yoo west=Poo", &City{Name: "Mee", North: "Yoo", West: "Poo"}, false},
		{"row 5", "Poo south=Ooo", &City{Name: "Poo", South: "Ooo"}, false},
		{"row 6", "Bury", &City{Name: "Bury"}, false},
		{"row error 1", "Poo south:Ooo", &City{}, true},
		{"row error 2", "Lee north=Yee west:Loo", &City{}, true},
		{"row error 3", "Lee north=Yee west", &City{}, true},
		{"row error 4", "Poo south Ooo", &City{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newCity(tt.row)
			if tt.wantErr && err != nil {
				return
			}
			assert.Nil(t, err)
			assert.EqualValues(t, tt.want, got)
		})
	}
}
