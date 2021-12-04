package model

type BaseModel struct {
	ID         uint32 `json:"id"`
	CreatedAt  uint32 `json:"created_at"`
	CreatedBy  string `json:"created_by"`
	ModifiedAt uint32 `json:"modified_at"`
	ModifiedBy string `json:"modified_by"`
	IsDel      uint8  `json:"is_del"`
	DeletedAt  uint32 `json:"deleted_at"`
}
