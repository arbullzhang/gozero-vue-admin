package response

import (
	"gozero-vue-admin/service/user/model"
)

type PolicyPathResponse struct {
	Paths []model.CasbinInfo `json:"paths"`
}
