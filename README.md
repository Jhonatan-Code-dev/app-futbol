# API de Gestión de Usuarios y Roles

Este proyecto es una **API RESTful** desarrollada en **Golang** utilizando el framework **Fiber**, destinada a gestionar usuarios, roles y autenticación con **JWT**. Además, integra **Swagger** para la documentación y utiliza **GORM** para la conexión a la base de datos.

---

## Características principales

- Gestión de **usuarios**: registro, solicitud de acceso, login.
- Gestión de **roles** con permisos diferenciados.
- **Autenticación JWT** para proteger rutas.
- **Hash de contraseñas** usando `bcrypt`.
- Documentación automática con **Swagger**.
- Migraciones automáticas de la base de datos con GORM.

---

## Tecnologías

- **Go** (Golang)
- **Fiber** (framework web)
- **GORM** (ORM para Go)
- **JWT** (Autenticación)
- **bcrypt** (Hash de contraseñas)
- **Swagger** (Documentación de API)
- **PostgreSQL / MySQL / SQLite** (según configuración)
- **Air** (Recarga automática en desarrollo)

---

## Estructura del proyecto

