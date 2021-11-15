package models

type God struct {
	Id   int    `json:"id" gorm:"index"`
	Name string `json:"name"`
}

func GetGodsList() ([]*God, error) {
	records := []*God{}

	err := db.Find(&records).Error
	if err != nil {
		return records, err
	}

	db.Find(&records)

	return records, err
}
