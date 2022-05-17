package database

import "rest-get-api/pkg/model"

func GetUsers(name string) []model.User {
	if name != "" {
		user, ok := model.Users[name]
		if ok {
			return []model.User{user}
		} else {
			return []model.User{}
		}
	}
	output := []model.User{}
	for key := range model.Users {
		output = append(output, model.Users[key])
	}
	return output
}
