package sqldb

type AreaCode struct {
	ID   uint   `gorm:"primary_key"`
	Code string `gorm:"column:code"`
	City Cities
}
