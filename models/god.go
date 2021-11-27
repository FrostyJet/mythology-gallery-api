package models

type God struct {
	Id          int       `json:"id" gorm:"index"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Title       string    `json:"title"`
	Free        string    `json:"free"`
	New         string    `json:"new"`
	Pantheon    string    `json:"pantheon"`
	Pros        string    `json:"props"`
	Type        string    `json:"type"`
	Role        string    `json:"role"`
	Card        string    `json:"card"`
	PantheonEN  string    `json:"pantheon_en"`
	GodNameEN   string    `json:"god_name_en"`
	RoleEN      string    `json:"role_en"`
	HeaderImage string    `json:"header_image"`
	Lore        string    `json:"lore"`
	Skins       []Skin    `json:"skins"`
	Abilities   []Ability `json:"abilities"`
}

func GetGodsList() ([]*God, error) {
	records := []*God{}

	err := db.Preload("Abilities").Preload("Skins").Find(&records).Error
	if err != nil {
		return records, err
	}

	return records, err
}

func InsertGodsList(records *[]God) error {
	err := db.Create(&records).Error
	return err
}

func SelectGodsFromWithLimit(id int, limit int) ([]*God, error) {
	records := []*God{}

	err := db.Order("id asc").Limit(limit).Where("id >= ?", id).Find(&records).Error
	if err != nil {
		return records, err
	}

	return records, err
}

func (g *God) Update(record *God) error {
	return db.Model(g).Updates(&record).Error
}

func (g *God) GetBySlug(slug string) (God, error) {
	var record God
	err := db.Model(&g).Preload("Abilities").Preload("Skins").Where("slug = ?", slug).First(&record).Error

	return record, err
}
