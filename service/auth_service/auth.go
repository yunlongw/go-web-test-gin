package auth_service

import "go-web-test-gin/models"

type Member struct {
	Id        int
	UserName  string
	PassWord  string
	createdOn int
}

func (m Member) AddMember() (err error) {
	maps := make(map[string]interface{})
	maps["username"] = m.UserName
	maps["password"] = m.PassWord
	if err = models.AddMember(maps); err != nil {
		return err
	}
	return nil
}

func (m Member) ExistByUserName() (bool, error) {
	return models.ExistMemberByUserName(m.UserName)
}
