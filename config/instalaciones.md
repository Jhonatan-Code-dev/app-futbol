# Librerias que instale para el proyecto
instalar fiber:

go get github.com/gofiber/fiber/v2
go get github.com/gofiber/swagger

instalar swagger:
go install github.com/swaggo/swag/cmd/swag@latest
go get github.com/swaggo/fiber-swagger
go get github.com/swaggo/files/v2

swag init


INSTALAR GORM PARA CONTROLAR BASE DE DATOS:
go get -u gorm.io/gorm

CONECTOR DE MYSQL:
go get -u gorm.io/driver/mysql


COMANDOS PARA GENERAR DOCUMENTACION:
swag init


para variables de entorno:
go get github.com/joho/godotenv

go get github.com/kelseyhightower/envconfig


Para seguridad:

go get github.com/golang-jwt/jwt/v5


instalar para seguridad de contraseñas:
go get golang.org/x/crypto/bcrypt


UTILIZAMOS AIR INIT PARA COMPILAR CODIGO :

air init