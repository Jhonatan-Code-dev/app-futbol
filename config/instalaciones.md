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

para validacion de datos:

go get github.com/go-playground/validator/v10



quiero agregar FX o Wire

utilize WIRE:
para instar 
go install github.com/google/wire/cmd/wire@latest 

para agregar al projecto:
go get github.com/google/wire

comando para generar codigo puro de go con wire:

wire ./di ---> se especifica la carpeta y al final obtienes wire_gen.go


