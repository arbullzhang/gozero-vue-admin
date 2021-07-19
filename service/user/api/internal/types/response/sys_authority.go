package response

import "gozero-vue-admin/service/user/model"

type SysAuthorityResponse struct {
	Authority model.SysAuthority `json:"authority"`
}

//type SysAuthorityCopyResponse struct {
//	Authority      model.SysAuthority `json:"authority"`
//	OldAuthorityId string             `json:"oldAuthorityId"` // 旧角色ID
//}
