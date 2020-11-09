module github.com/iot-for-tillgenglighet/api-problemreport

go 1.14

require (
	github.com/99designs/gqlgen v0.11.3
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/iot-for-tillgenglighet/messaging-golang v0.0.0-20201009211140-579335ad3c49
	github.com/iot-for-tillgenglighet/ngsi-ld-golang v0.0.0-20201029160316-c6fe6d155326
	github.com/rs/cors v1.7.0
	github.com/sirupsen/logrus v1.7.0
	github.com/vektah/gqlparser v1.3.1
	golang.org/x/sys v0.0.0-20200323222414-85ca7c5b95cd // indirect
	gopkg.in/yaml.v2 v2.2.5 // indirect
	gorm.io/driver/postgres v1.0.5
	gorm.io/gorm v1.20.5
)

replace github.com/99designs/gqlgen => github.com/marwan-at-work/gqlgen v0.0.0-20200107060600-48dc29c19314
