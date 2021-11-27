package models

import "gorm.io/gorm/clause"

type Ability struct {
	Id          int    `json:"id" gorm:"index"`
	Summary     string `json:"summary"`
	URL         string `json:"url"`
	Description string `json:"description"`
	GodID       uint   `json:"god_id"`
}

func (a *Ability) Save() error {
	return db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"summary", "url", "description"}),
	}).Create(a).Error
}

func (a *Ability) InsertMany(records *[]Ability) error {
	return db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"summary", "url", "description"}),
	}).Create(&records).Error
}
