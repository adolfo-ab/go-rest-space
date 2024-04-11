package planet

type PlanetType int

const (
	Undefined PlanetType = iota
	Terrestrial
	GasGiant
	IceGiant
	DwarfPlanet
)
