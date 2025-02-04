package sqldb

type Cities struct {
	ID    uint   `gorm:"primary_key"`
	Name  string `gorm:"column name"`
	State uint   `gorm:"column state"`
}
