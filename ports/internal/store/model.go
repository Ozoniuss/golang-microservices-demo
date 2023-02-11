package store

/*
 Gorm will assume default postgres naming scheme and match the table and column
 names by default: https://gorm.io/docs/models.html#Conventions
*/

// Port models the ports as per their database entity. Gorm struct tags have
// only been used when necessary.
type Port struct {
	Id          string
	Name        string
	City        string
	Country     string
	Alias       []string `gorm:"type:text[]"`
	Regions     []string `gorm:"type:text[]"`
	Coordinates [2]float32
	Province    string
	Timezone    string
	Unlocs      []string
	Code        string
}
