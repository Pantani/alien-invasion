package invasion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	_alien1 = &Alien{
		Name:   "divine-cloud",
		Number: 0,
		City:   "Lee",
		Status: AlienStatusLive,
	}
	_alien2 = &Alien{
		Name:   "bold-glitter",
		Number: 1,
		City:   "Loo",
		Status: AlienStatusLive,
	}
	_alien3 = &Alien{
		Name:   "little-sky",
		Number: 2,
		City:   "Qu-ux",
		Status: AlienStatusLive,
	}
	_alien4 = &Alien{
		Name:   "lingering-fog",
		Number: 3,
		City:   "Foo",
		Status: AlienStatusLive,
	}
	_alien5 = &Alien{
		Name:   "green-sun",
		Number: 4,
		City:   "Baz",
		Status: AlienStatusLive,
	}
	_alien6 = &Alien{
		Name:   "sparkling-surf",
		Number: 5,
		City:   "Bar",
		Status: AlienStatusLive,
	}
	_alien7 = &Alien{
		Name:   "frosty-meadow",
		Number: 6,
		City:   "",
		Status: AlienStatusDead,
	}
	_alien8 = &Alien{
		Name:   "long-violet",
		Number: 7,
		City:   "",
		Status: AlienStatusDead,
	}
	_alien9 = &Alien{
		Name:   "lingering-paper",
		Number: 8,
		City:   "",
		Status: AlienStatusDead,
	}
	_aliens = Aliens{_alien1, _alien2, _alien3, _alien4, _alien5, _alien6, _alien7, _alien8, _alien9}
)

func TestAliens_HasAlive(t *testing.T) {
	tests := []struct {
		name   string
		aliens Aliens
		want   bool
	}{
		{"test alien alive 1", _aliens, true},
		{"test alien alive 2", Aliens{_alien1, _alien2}, true},
		{"test alien alive 3", Aliens{_alien1, _alien9}, true},
		{"test no aliens alive 1", Aliens{_alien9}, false},
		{"test no aliens alive 2", Aliens{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.aliens.HasAlive()
			assert.EqualValues(t, tt.want, got)
		})
	}
}

func TestAliens_Names(t *testing.T) {
	tests := []struct {
		name   string
		aliens Aliens
		want   []string
	}{
		{"test names 1", _aliens, []string{"divine-cloud", "bold-glitter", "little-sky", "lingering-fog", "green-sun", "sparkling-surf", "frosty-meadow", "long-violet", "lingering-paper"}},
		{"test names 2", Aliens{_alien1, _alien2, _alien3, _alien4}, []string{"divine-cloud", "bold-glitter", "little-sky", "lingering-fog"}},
		{"test names 3", Aliens{_alien1, _alien2, _alien3}, []string{"divine-cloud", "bold-glitter", "little-sky"}},
		{"test names 4", Aliens{_alien1, _alien2}, []string{"divine-cloud", "bold-glitter"}},
		{"test names 5", Aliens{_alien1}, []string{"divine-cloud"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.aliens.Names()
			assert.EqualValues(t, tt.want, got)
		})
	}
}

func TestAliens_kill(t *testing.T) {
	tests := []struct {
		name   string
		aliens Aliens
	}{
		{"test kill 1", Aliens{_alien1, _alien2, _alien3}},
		{"test kill 2", Aliens{_alien1, _alien2}},
		{"test kill 3", Aliens{_alien1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.aliens.kill()
			for _, a := range tt.aliens {
				assert.Equal(t, AlienStatusDead, a.Status)
			}
		})
	}
}

func Test_newAlien(t *testing.T) {
	type args struct {
		number uint
		city   string
	}
	tests := []struct {
		name string
		args args
		want *Alien
	}{
		{"test create 1", args{0, "Bar"}, &Alien{Number: 0, City: "Bar", Status: AlienStatusLive}},
		{"test create 2", args{10, "Foo"}, &Alien{Number: 10, City: "Foo", Status: AlienStatusLive}},
		{"test create 2", args{100, "Lee"}, &Alien{Number: 100, City: "Lee", Status: AlienStatusLive}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newAlien(tt.args.number, tt.args.city)
			assert.EqualValues(t, tt.want.Status, got.Status)
			assert.EqualValues(t, tt.want.City, got.City)
			assert.EqualValues(t, tt.want.Number, got.Number)
		})
	}
}

func TestAliens_MoveAliens(t *testing.T) {
	tests := []struct {
		name   string
		aliens Aliens
		world  World
	}{
		{"move aliens 1", Aliens{_alien1}, _defaultWorld},
		{"move aliens 2", Aliens{_alien1, _alien2}, _defaultWorld},
		{"move aliens 3", Aliens{_alien1, _alien2, _alien3}, _defaultWorld},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var aliens Aliens
			copy(aliens, tt.aliens)
			tt.aliens.MoveAliens(tt.world)
			assert.NotEqualValues(t, aliens, tt.aliens)
		})
	}
}

func TestAliens_StartWars(t *testing.T) {
	tests := []struct {
		name   string
		aliens Aliens
		world  World
		finish bool
	}{
		{"start wars 1", _aliens, _defaultWorld, false},
		{
			"start wars 2",
			Aliens{
				&Alien{
					Number: 0,
					Name:   "Test1",
					City:   "Foo",
					Status: AlienStatusLive,
				},
				&Alien{
					Number: 1,
					Name:   "Test2",
					City:   "Foo",
					Status: AlienStatusLive,
				},
			},
			World{"Foo": &City{Name: "Foo"}},
			true,
		}, {
			"start wars 3",
			Aliens{
				&Alien{
					Number: 0,
					Name:   "Test1",
					City:   "Foo",
					Status: AlienStatusLive,
				},
				&Alien{
					Number: 1,
					Name:   "Test2",
					City:   "Foo",
					Status: AlienStatusLive,
				},
				&Alien{
					Number: 2,
					Name:   "Test3",
					City:   "Bar",
					Status: AlienStatusDead,
				},
			},
			World{"Foo": &City{Name: "Foo"}, "Bar": &City{Name: "Bar"}},
			true,
		},
		{
			"start wars 4",
			Aliens{
				&Alien{
					Number: 0,
					Name:   "Test1",
					City:   "Foo",
					Status: AlienStatusLive,
				},
				&Alien{
					Number: 1,
					Name:   "Test2",
					City:   "Foo",
					Status: AlienStatusLive,
				},
				&Alien{
					Number: 2,
					Name:   "Test3",
					City:   "Bar",
					Status: AlienStatusLive,
				},
			},
			World{"Foo": &City{Name: "Foo"}, "Bar": &City{Name: "Bar"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.aliens.StartWars(tt.world)
			assert.Equal(t, tt.finish, !tt.aliens.HasAlive() && !tt.world.HasMovement())
		})
	}
}
