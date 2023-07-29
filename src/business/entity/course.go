package entity

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Name    string   `json:"name" gorm:"type:varchar(50)"`
	Credit  int      `json:"credit" gorm:"type:int"`
	Classes []*Class `json:"classes"`
}
