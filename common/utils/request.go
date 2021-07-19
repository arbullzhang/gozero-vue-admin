package utils

// Paging common input parameter structure
type PageInfo struct {
	Page     int `json:"page"`     // 页码
	PageSize int `json:"pageSize"` // 每页大小
}

// Find by id structure
type GetById struct {
	ID float64 `json:"id"` // 主键ID
}

type IdsReq struct {
	Ids []int `json:"ids"`
}

// Get role by id structure
type GetAuthorityId struct {
	AuthorityId string // 角色ID
}

type Empty struct{}
