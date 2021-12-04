package model

type User struct {
	*BaseModel
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Status   uint8  `json:"status"`
}

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Role struct {
	*BaseModel
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      uint8  `json:"status"`
	Level       uint8  `json:"level"`
	ParentID    uint32 `json:"parent_id"`
}

type UserRole struct {
	*BaseModel
	UserID uint32 `json:"user_id"`
	RoleID uint32 `json:"role_id"`
	Status uint8  `json:"status"`
}

type Permit struct {
	*BaseModel
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      uint8  `json:"status"`
	Type        uint8  `json:"type"`
	ParentID    uint32 `json:"parent_id"`
	Sort        uint32 `json:"sort"`
}

type RolePermit struct {
	*BaseModel
	RoleID   uint32 `json:"role_id"`
	PermitID uint32 `json:"permit_id"`
	Status   uint8  `json:"status"`
}
