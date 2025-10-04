package schemas

type Permiso struct {
	IdPermiso   uint   `gorm:"primaryKey;autoIncrement"`
	Permiso     string `gorm:"size:30;unique;not null"`
	Descripcion string `gorm:"size:100;not null"`
}
