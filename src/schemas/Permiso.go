package schemas

type Permiso struct {
	IdPermiso   uint   `gorm:"primaryKey;autoIncrement" json:"id_permiso"`
	Permiso     string `gorm:"type:varchar(30);unique;not null" json:"permiso"`
	Descripcion string `gorm:"type:varchar(100);not null" json:"descripcion"`
}
