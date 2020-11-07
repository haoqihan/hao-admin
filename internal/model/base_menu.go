package model

type BaseMenu struct {
	*Model
	MenuLevel  uint   `json:"-"`
	ParentId   string `json:"prent_id" gorm:"comment:父菜单ID"`
	Path       string `json:"path" gorm:"comment:路由path"`
	Name       string `json:"name" gorm:"comment:路由name"`
	Hidden     bool   `json:"hidden" gorm:"comment:是否在列表隐藏"`
	Component  string `json:"component" gorm:"comment:对应前端文件路径"`
	Sort       int    `json:"sort" gorm:"comment:前端排序标记"`
	Meta       `json:"meta" gorm:"comment:附加属性"`
	Authoritys []Authority         `json:"authoritys" gorm:"many2many:authority_menus"`
	Children   []BaseMenu          `json:"children" gorm:"-"`
	Parameters []BaseMenuParameter `json:"parameters"`
}

type Meta struct {
	KeepAlive   bool   `json:"keep_alive" gorm:"comment:是否缓存"`
	DefaultMenu bool   `json:"default_menu" gorm:"comment:是否是基础路由"`
	Title       string `json:"title" gorm:"comment:菜单名"`
	Icon        string `json:"icon" gorm:"comment:菜单图标"`
}

type BaseMenuParameter struct {
	BaseMenuId uint
	Type       string `json:"type" gorm:"comment:'地址栏携带参数为params还是query'"`
	Key        string `json:"key" gorm:"comment:地址栏携带的参数的key"`
	Value      string `json:"value" gorm:"comment:地址栏携带参数的值"`
}
