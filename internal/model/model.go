package model

type BaseModel struct {
	ID         uint32 `json:"id"`          // ID
	CreatedAt  uint32 `json:"created_at"`  // 创建时间
	CreatedBy  string `json:"created_by"`  // 创建人
	ModifiedAt uint32 `json:"modified_at"` // 修改时间
	ModifiedBy string `json:"modified_by"` // 修改人
	IsDel      uint8  `json:"is_del"`      // 是否已删除
	DeletedAt  uint32 `json:"deleted_at"`  // 删除时间
}
