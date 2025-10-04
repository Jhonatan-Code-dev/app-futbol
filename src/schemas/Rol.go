package schemas

type Rol struct {
	IdRol uint   `gorm:"primaryKey;autoIncrement"`
	Rol   string `gorm:"size:30;unique;not null"`
}
