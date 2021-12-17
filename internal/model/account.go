package model

type User struct {
	BaseModel
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
	Nickname string `json:"nickname"` // 昵称
	Phone    string `json:"phone"`    // 电话
	Email    string `json:"email"`    // 邮箱
	Status   uint8  `json:"status"`   // 状态
}

type Role struct {
	BaseModel
	Name        string `json:"name"`        // 名称
	Description string `json:"description"` // 描述
	Status      uint8  `json:"status"`      // 状态
	Level       uint8  `json:"level"`       // 级别
	ParentID    uint32 `json:"parent_id"`   // 父级ID
}

type UserRole struct {
	BaseModel
	UserID uint32 `json:"user_id"` // 用户ID
	RoleID uint32 `json:"role_id"` // 角色ID
	Status uint8  `json:"status"`  // 状态
}

type Permit struct {
	BaseModel
	Name        string `json:"name"`        // 名称
	Description string `json:"description"` // 描述
	Status      uint8  `json:"status"`      // 状态
	Type        uint8  `json:"type"`        // 类型
	Router      string `json:"router"`      // 路由
	Icon        string `json:"icon"`        // 前端页面显示的icon
	ParentID    uint32 `json:"parent_id"`   // 父级ID
	Sort        uint32 `json:"sort"`        // 排序
}

type RolePermit struct {
	BaseModel
	RoleID   uint32 `json:"role_id"`   // 角色ID
	PermitID uint32 `json:"permit_id"` // 权限ID
	Status   uint8  `json:"status"`    // 状态
}

type JwtToken struct {
	ID    uint32 `json:"id"`    // 用户ID
	Token string `json:"token"` // 鉴权令牌
}

type UserToken struct {
	User
	Token string `json:"token"`
}

type RegisterReq struct {
	Username string `form:"username" binding:"required"` // 用户名
	Password string `form:"password" binding:"required"` // 密码
	Phone    string `form:"phone"`                       // 电话
}

type LoginReq struct {
	Username string `form:"username" binding:"required"` // 用户名
	Password string `form:"password" binding:"required"` // 密码
}

type UserReq struct {
}

type RoleReq struct {
}

type PermitReq struct {
}
