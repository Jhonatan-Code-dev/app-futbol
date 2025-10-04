package schemas

type RolPermiso struct {
	IdRolPermiso uint `gorm:"primaryKey;autoIncrement"`

	IDRol     uint `gorm:"not null"`
	IDPermiso uint `gorm:"not null"`

	// Relaciones opcionales para hacer Preload
	Rol     *Rol     `gorm:"foreignKey:IDRol;references:IdRol"`
	Permiso *Permiso `gorm:"foreignKey:IDPermiso;references:IdPermiso"`
}
