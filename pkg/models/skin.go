package models

import "gorm.io/gorm/clause"

type Skin struct {
	GodID uint   `json:"god_id" gorm:"index:skin_god,unique"`
	Id    int    `json:"id" gorm:"index"`
	Name  string `json:"name" gorm:"index:skin_god,unique"`
	Type  string `json:"type"`
	Image string `json:"image"`
}

func (s *Skin) InsertMany(records *[]Skin) error {
	return db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}, {Name: "god_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "type", "image"}),
	}).Create(&records).Error
}
