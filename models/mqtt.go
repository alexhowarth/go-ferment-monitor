package models

import (
	"fmt"
)

type MQTTConf struct {
	GormModel
	Enabled  bool   `json:"enabled"`
	Url      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type MQTTModel struct{}

func (m *MQTTModel) GetMQTTConfig() (conf MQTTConf, err error) {
	println("hello")
	err = db.First(&conf).Error
	if err != nil {
		fmt.Printf("Error is: %+v\n", err)
	}
	return
}

func (t *MQTTModel) CreateMQTTConfig(conf *MQTTConf) (*MQTTConf, error) {
	err := db.FirstOrCreate(conf).Error
	if err != nil {
		fmt.Printf("Error is: %+v\n", err)
	}
	return conf, err
}

func (t *MQTTModel) UpdateMQTTConfig(conf *MQTTConf) (*MQTTConf, error) {
	err := db.Save(conf).Error
	if err != nil {
		fmt.Printf("Error is: %+v\n", err)
	}
	return conf, err
}
