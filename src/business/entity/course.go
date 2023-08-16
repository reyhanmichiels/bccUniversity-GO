package entity

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Name    string   `json:"name" gorm:"type:varchar(50);unique;notnull"`
	Credit  int      `json:"credit" gorm:"type:int;notnull"`
	Classes []*Class `json:"classes"`
}
