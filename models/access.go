package models

type Access struct {
	Id int
	ModuleName string
	ActionName string
	Type int
	Url string
	ModuleId int
	Sort int
	Description string
	Status int
	AddTime int
	AccessItem []Access `gorm:"foreignKey:ModuleId";references:Id`
	Check bool `gorm:"-"`
}

func (Access) TableName() string {
	return "access"
}