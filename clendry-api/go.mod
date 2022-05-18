module github.com/artomsopun/clendry/clendry-api

go 1.16

require github.com/golang-jwt/jwt v3.2.2+incompatible

require (
	github.com/google/uuid v1.3.0
	github.com/labstack/echo/v4 v4.7.2
	github.com/minio/minio-go/v7 v7.0.23
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/viper v1.10.1
	golang.org/x/time v0.0.0-20201208040808-7e3f01d25324
	gorm.io/driver/mysql v1.3.2
	gorm.io/gorm v1.23.1
)
