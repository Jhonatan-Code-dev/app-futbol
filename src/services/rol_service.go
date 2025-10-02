package services

import (
	"app-futbol/src/schemas"

	"gorm.io/gorm"
)

type RolService struct {
	DB *gorm.DB
}

func NewRolService(db *gorm.DB) *RolService {
	return &RolService{DB: db}
}

// Crear un rol
func (s *RolService) Create(rol *schemas.Rol) error {
	return s.DB.Create(rol).Error
}

// Listar todos los roles
func (s *RolService) GetAll() ([]schemas.Rol, error) {
	var roles []schemas.Rol
	err := s.DB.Find(&roles).Error
	return roles, err
}

// Obtener rol por ID
func (s *RolService) GetByID(id uint) (*schemas.Rol, error) {
	var rol schemas.Rol
	err := s.DB.First(&rol, id).Error
	if err != nil {
		return nil, err
	}
	return &rol, nil
}

// Actualizar rol
func (s *RolService) Update(rol *schemas.Rol) error {
	return s.DB.Save(rol).Error
}

// Eliminar rol
func (s *RolService) Delete(id uint) error {
	return s.DB.Delete(&schemas.Rol{}, id).Error
}
