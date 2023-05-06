package model

type SysMenu struct {
	BasicModel
	PID       int    `gorm:"type:int" json:"pid"`
	Name      string `json:"name"`
	Path      string `json:"path"`
	Component string `json:"component"`
	Title     string `json:"title"`
	IsHidden  bool   `json:"isHidden"`
	Sequence  uint8  `gorm:"type:tinyint" json:"sequence"`
	Option    string `json:"option"`
	Icon      string `json:"icon"`
	Redirect  string `json:"redirect"`
	IsUnfold  bool   `json:"isUnfold"`
}

type SysRole struct {
	BasicModel
	Name string `json:"name"`
	Rs   string `json:"rs"`
	Memo string `json:"memo"`
}

type SysUser struct {
	BasicModel
	UserName string `json:"username"`
	RoleId   string `json:"roleId"`
	Password string `json:"-"`
}

type MenuTree struct {
	SysMenu
	Children []MenuTree
}
