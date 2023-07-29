package entity

type UserClass struct {
	UserID  uint `gorm:"primaryKey"`
	ClassID uint `gorm:"primaryKey"`
}
