package enum

type RegionalEnum int

const (
	// North -
	North = iota
	// South -
	South
	// Central -
	Central
	// East -
	East
	// Toshima -
	Toshima
	// Any -
	Any
)

func (action RegionalEnum) String() string {
	return [...]string{
		"北部",
		"南部",
		"中部",
		"東部",
		"外島",
		"不分",
	}[action]
}
