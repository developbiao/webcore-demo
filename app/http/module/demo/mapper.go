package demo

import (
	demoService "github.com/developbiao/webcore-demo/app/provider/demo"
)

func UserModelsToUserDTOs(modes []UserModel) []UserDTO {
	ret := []UserDTO{}
	for _, model := range modes {
		t := UserDTO{
			ID:   model.UserId,
			Name: model.Name,
		}
		ret = append(ret, t)
	}
	return ret
}

func StudentsToUserDTOs(students []demoService.Student) []UserDTO {
	ret := []UserDTO{}
	for _, student := range students {
		t := UserDTO{
			ID:   student.ID,
			Name: student.Name,
		}
		ret = append(ret, t)
	}
	return ret
}
