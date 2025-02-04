package sqldb

type Cities struct {
	id    uint   `gorm:"primary_key"`
	Name  string `gorm:"column name"`
	State States `gorm:"column state"`
}
