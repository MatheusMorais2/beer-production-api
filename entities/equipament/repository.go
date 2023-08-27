package equipament

type EquipamentRepository interface {
	Insert(equipament *Equipament) (*Equipament, error)
}