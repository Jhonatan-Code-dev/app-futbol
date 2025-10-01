package schemas

type Rol struct {
	IdRol uint   `gorm:"primaryKey;autoIncrement" json:"id_rol"`
	Rol   string `gorm:"type:varchar(30);unique;not null" json:"rol"`
}
