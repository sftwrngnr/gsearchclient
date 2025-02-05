package sqldb

type Areacodes struct {
	ID    uint `gorm:"primary_key"`
	Code  uint `gorm:"column:code"`
	State uint `gorm:"column:state"`
}
