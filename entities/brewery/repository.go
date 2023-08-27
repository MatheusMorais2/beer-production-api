package brewery

type BreweryRepository interface {
	Insert(brewery *Brewery) (*Brewery, error)
}