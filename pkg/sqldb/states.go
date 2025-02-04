package sqldb

type States struct {
	ID      uint   `gorm:"primary_key"`
	Abbrev  string `gorm:"column:abbrev"`
	Name    string `gorm:"column:name"`
	Capital string `gorm:"column:capital"`
	Region  string `gorm:"column:region"`
}
