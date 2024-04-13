package planet_type

type PlanetType int

const (
	Undefined PlanetType = iota
	Terrestrial
	GasGiant
	IceGiant
	DwarfPlanet
)
