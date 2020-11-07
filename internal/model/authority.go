package model

type Authority struct {
	*Model
	AuthorityId     string      `json:"authority_id" gorm:"not null;unique;primary_key;comment:角色ID;size:90"`
	AuthorityName   string      `json:"authority_name" gorm:"comment:角色名称"`
	ParentId        string      `json:"parent_id" gorm:"comment:父角色ID"`
	DataAuthorityId []Authority `json:"data_authority_id" gorm:"many2many:data_authority_id"`
	Children        []Authority `json:"children" gorm:"-"`
	BaseMenu        []BaseMenu  `json:"base_menu" gorm:"many2many:authority_menus"`
}
