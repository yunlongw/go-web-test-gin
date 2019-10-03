package tag_service

import "go-web-test/models"

type Tag struct {
	ID int
	Name string
	CreatedBy string
	ModifiedBy string
	State int

	PageNum int
	PageSize int
}

func (t *Tag) ExistByName() (bool, error)  {
	return models.ExistTagByName(t.Name)
}

func (t *Tag) ExistByID() (bool, error)  {
	return models.ExistTagByID(t.ID)
}

func (t *Tag) Add() error  {
	return models.AddTag(t.Name, t.State, t.CreatedBy)
}

func (t *Tag) Delete() interface{} {
	return models.DeleteTag(t.ID)
}

func (t *Tag) Edit() error {
	data := make(map[string]interface{})
	data["modified_by"] = t.ModifiedBy
	data["name"] = t.Name
	if t.State >= 0 {
		data["state"] = t.State
	}

	return models.EditTag(t.ID, data)
}