package models

import (
	"fmt"

	"github.com/alexhowarth/go-tilt"
	"github.com/davecgh/go-spew/spew"
)

// Active Tilts
type TiltConf struct {
	GormModel
	Colour    tilt.Colour `gorm:"unique_index;auto_increment:false;not null" json:"colour"` // not sure about auto_increment here
	Enabled   bool        `json:"enabled"`
	Name      string      `json:"name"`
	OG        float64     `json:"og"`
	Consumers Consumers   `gorm:"foreignkey:id" json:"consumers"` //`gorm:"no null"`
}

type TiltModel struct{}

func (t *TiltModel) GetTilts() (tilts []TiltConf, err error) {
	err = db.Find(&tilts).Error
	if err != nil {
		fmt.Printf("Error is: %+v\n", err)
	}
	return
}

func (t *TiltModel) GetTilt(id string) (tilt TiltConf, err error) {
	//TODO: I need to do this? .Related(&card)
	err = db.First(&tilt, id).Error
	if err != nil {
		fmt.Printf("Error is: %+v\n", err)
	}
	return
}

func (t *TiltModel) CreateTilt(tilt TiltConf) (TiltConf, error) {
	// TODO first or create and update last seen
	// purge tilts that have not been seen for a while
	err := db.Create(&tilt).Error
	if err != nil {
		fmt.Printf("Error is: %+v\n", err)
		return tilt, err
	}
	return tilt, nil
}

// func (t *TiltModel) UpdateTilt(id string) (tilt TiltConf, err error) {
// 	err = db.First(&tilt, id).Error
// 	if err != nil {
// 		fmt.Printf("Error is: %+v\n", err)
// 	}
// 	return
// }

// TODO is this []TiltConf ?
func (t *TiltModel) UpdateTilt(conf *TiltConf) (*TiltConf, error) {
	// db.First(conf) <-- then save
	err := db.Save(conf).Error
	spew.Dump(conf)
	//err := db.UpdateColumns(conf).Error
	if err != nil {
		fmt.Printf("Error is: %+v\n", err)
	}

	return conf, err
}
