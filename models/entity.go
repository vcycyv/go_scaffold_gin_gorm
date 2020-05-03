package models

type {entity_capital} struct {
	Base

	Name string
}

func Get{entity_capital_plural}(pageNum int, pageSize int, maps interface{}) ([]{entity_capital}, error) {
	var {entity_plural} []{entity_capital}
	err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&{entity_plural}).Error
	if err != nil {
		return []{entity_capital}{}, err
	}
	return {entity_plural}, nil
}

func Get{entity_capital}(id string) (*{entity_capital}, error) {
	var {entity} {entity_capital}
	err := db.Where("id = ?", id).First(&{entity}).Error
	if err != nil {
		return &{entity_capital}{}, err
	}

	return &{entity}, err
}

func Add{entity_capital}({entity} *{entity_capital}) (*{entity_capital}, error) {
	err := db.Create({entity}).Error

	return {entity}, err
}

func Edit{entity_capital}({entity} *{entity_capital}) (*{entity_capital}, error) {
	err := db.Save({entity}).Error

	return {entity}, err
}

func Delete{entity_capital}(id string) error {
	return db.Where("id = ?", id).Delete(&{entity_capital}{}).Error
}
