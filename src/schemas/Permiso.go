package schemas

type Permiso struct {
	IdPermiso   uint   `gorm:"primaryKey;autoIncrement"`
	Permiso     string `gorm:"type:varchar(30);unique;not null"`
	Descripcion string `gorm:"type:varchar(100);not null"`
}
