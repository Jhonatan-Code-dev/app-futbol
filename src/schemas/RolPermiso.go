package schemas

type RolPermiso struct {
	IdRolPermiso uint `gorm:"primaryKey;autoIncrement" json:"id_rol_permiso"`

	IDRol     uint `gorm:"not null" json:"id_rol"`
	IDPermiso uint `gorm:"not null" json:"id_permiso"`

	// Relaciones opcionales para hacer Preload
	Rol     *Rol     `gorm:"foreignKey:IDRol;references:IdRol" json:"rol,omitempty"`
	Permiso *Permiso `gorm:"foreignKey:IDPermiso;references:IdPermiso" json:"permiso,omitempty"`
}
