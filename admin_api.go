package intercom

import (
	"encoding/json"
	"fmt"

	"gopkg.in/intercom/intercom-go.v2/interfaces"
)

// AdminRepository defines the interface for working with Admins through the API.
type AdminRepository interface {
	get(string) (Admin, error)
	list() (AdminList, error)
}

// AdminAPI implements AdminRepository
type AdminAPI struct {
	httpClient interfaces.HTTPClient
}

func (api AdminAPI) list() (AdminList, error) {
	adminList := AdminList{}
	data, err := api.httpClient.Get("/admins", nil)
	if err != nil {
		return adminList, err
	}
	err = json.Unmarshal(data, &adminList)
	return adminList, err
}

func (api AdminAPI) get(id string) (Admin, error) {
	admin := Admin{}
	u := fmt.Sprintf("/admins/%s", id)
	data, err := api.httpClient.Get(u, nil)
	if err != nil {
		return admin, err
	}
	err = json.Unmarshal(data, &admin)
	return admin, err
}
