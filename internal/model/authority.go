package model

type Authority struct {
	CreatedBy       string      `json:"created_by"`
	ModifiedBy      string      `json:"modified_by"`
	CreatedOn       uint32      `json:"created_on"`
	ModifiedOn      uint32      `json:"modified_on"`
	DeletedOn       uint32      `json:"deleted_on"`
	IsDel           uint8       `json:"is_del"`
	AuthorityId     string      `json:"authorityId" gorm:"not null;unique;primary_key;comment:'角色ID';size:90"`
	AuthorityName   string      `json:"authority_name" gorm:"comment:'角色名称'"`
	ParentId        string      `json:"parent_id" gorm:"comment:'父角色ID'"`
	DataAuthorityId []Authority `json:"data_authority_id" gorm:"many2many:data_authority_id;association_jointable_foreignkey:data_authority_id"`
	Children        []Authority `json:"children" gorm:"-"`
	BaseMenu        []BaseMenu  `json:"base_menu" gorm:"many2many:authority_menus"`
}
