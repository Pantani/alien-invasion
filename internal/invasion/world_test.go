package invasion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	_defaultWorld = World{
		"Bar":   _cityBar,
		"Baz":   _cityBaz,
		"Foo":   _cityFoo,
		"Lee":   _cityLee,
		"Loo":   _cityLoo,
		"Qu-ux": _cityQuux,
		"Yee":   _cityYee,
	}
)

func TestLoadWorld(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     World
		wantErr  bool
	}{
		{
			"test file 1",
			"world_1.txt",
			World{
				"Bar":   _cityBar,
				"Baz":   _cityBaz,
				"Foo":   _cityFoo,
				"Qu-ux": _cityQuux,
			},
			false,
		}, {
			"test file 2",
			"world_2.txt",
			World{
				"Bar":   _cityBar,
				"Baz":   _cityBaz,
				"Foo":   _cityFoo,
				"Qu-ux": _cityQuux,
			},
			false,
		}, {
			"test file 3",
			"world_3.txt",
			World{
				"Bar":   _cityBar,
				"Baz":   _cityBaz,
				"Foo":   _cityFoo,
				"Qu-ux": _cityQuux,
			},
			false,
		}, {
			"test file 4",
			"world_4.txt",
			World{
				"Bar":   _cityBar,
				"Baz":   _cityBaz,
				"Foo":   _cityFoo,
				"Lee":   _cityLee,
				"Loo":   _cityLoo,
				"Qu-ux": _cityQuux,
				"Yee":   _cityYee,
			},
			false,
		},
		{"test file error", "world_err.txt", nil, true},
		{"test no file", "world_not_exist.txt", nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadWorld("./../../test/" + tt.filename)
			if tt.wantErr && err != nil {
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestWorld_GenerateAliens(t *testing.T) {
	tests := []struct {
		name   string
		world  World
		number uint
	}{
		{"generate 10 aliens", _defaultWorld, 10},
		{"generate 100 aliens", _defaultWorld, 100},
		{"generate 1000 aliens", _defaultWorld, 1000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.world.GenerateAliens(tt.number)
			assert.Equal(t, int(tt.number), len(got))
		})
	}
}

func TestWorld_HasMovement(t *testing.T) {
	tests := []struct {
		name  string
		world World
		want  bool
	}{
		{"test not moves 1", World{}, false},
		{"test not moves 2", World{"Bar": {Name: "Bar"}}, false},
		{"test not moves 3", World{"Bar": {Name: "Bar"}, "Foo": {Name: "Foo"}}, false},
		{"test not moves 4", World{"Bar": {Name: "Bar", North: "Foo"}}, true},
		{"test moves 1", World{"Bar": {Name: "Bar", North: "Foo"}, "Foo": {Name: "Foo", South: "Bar"}}, true},
		{"test moves 2", _defaultWorld, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.world.HasMovement()
			assert.EqualValues(t, tt.want, got)
		})
	}
}

func TestWorld_addCity(t *testing.T) {
	tests := []struct {
		name  string
		world World
		city  *City
		want  World
	}{
		{
			"add one city to empty world",
			World{},
			_cityYee,
			World{"Lee": &City{Name: "Lee", North: "Yee"}, "Yee": _cityYee},
		}, {
			"add one city to world 1",
			World{"Bar": _cityBar},
			_cityYee,
			World{"Lee": &City{Name: "Lee", North: "Yee"}, "Bar": _cityBar, "Yee": _cityYee},
		}, {
			"add one city to world 2",
			World{"Bar": _cityBar, "Yee": _cityYee},
			_cityLoo,
			World{"Lee": &City{Name: "Lee", West: "Loo"}, "Bar": _cityBar, "Yee": _cityYee, "Loo": _cityLoo},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.world.addCity(tt.city)
			for key, value := range tt.world {
				want := tt.want[key]
				assert.EqualValues(t, want, value)
			}
		})
	}
}

func TestWorld_destroyCity(t *testing.T) {
	tests := []struct {
		name  string
		world World
		args  string
		want  World
	}{
		{
			"remove one city from world 1",
			World{"Lee": &City{Name: "Lee", North: "Yee"}, "Yee": _cityYee},
			"Yee",
			World{"Lee": &City{Name: "Lee"}},
		}, {
			"remove one city from world 2",
			World{"Lee": &City{Name: "Lee", North: "Yee"}, "Bar": _cityBar, "Yee": _cityYee},
			"Yee",
			World{"Lee": &City{Name: "Lee"}, "Bar": _cityBar},
		}, {
			"remove one city form world 3",
			World{"Lee": &City{Name: "Lee", West: "Loo"}, "Bar": _cityBar, "Yee": _cityYee, "Loo": _cityLoo},
			"Loo",
			World{"Lee": &City{Name: "Lee"}, "Bar": _cityBar, "Yee": _cityYee},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.world.destroyCity(tt.args)
			for key, value := range tt.world {
				want := tt.want[key]
				assert.EqualValues(t, want, value)
			}
		})
	}
}

func TestWorld_getRandomCity(t *testing.T) {
	tests := []struct {
		name  string
		world World
	}{
		{"test random city 1", _defaultWorld},
		{"test random city 2", _defaultWorld},
		{"test random city 3", _defaultWorld},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.world.getRandomCity()
			assert.NotNil(t, got)
			assert.Greater(t, len(got.Name), 0)
		})
	}
}
