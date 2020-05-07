module github.com/iot-for-tillgenglighet/api-problemreport

go 1.13

require (
	github.com/99designs/gqlgen v0.11.3
	github.com/agnivade/levenshtein v1.0.3 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/go-chi/chi v4.1.0+incompatible
	github.com/gorilla/mux v1.7.3 // indirect
	github.com/iot-for-tillgenglighet/api-temperature v0.0.0-20200413202351-bb27821b1e29 // indirect
	github.com/iot-for-tillgenglighet/messaging-golang v0.0.0-20200124165843-f64c6b8239e8
	github.com/iot-for-tillgenglighet/ngsi-ld-golang v0.0.0-20200507095135-ffc0edb4751b
	github.com/jinzhu/gorm v1.9.12
	github.com/rs/cors v1.7.0
	github.com/sirupsen/logrus v1.5.0
	github.com/urfave/cli v1.22.1 // indirect
	github.com/vektah/dataloaden v0.3.0 // indirect
	github.com/vektah/gqlparser v1.2.1
	golang.org/x/sys v0.0.0-20200107162124-548cf772de50 // indirect
	golang.org/x/tools v0.0.0-20191115202509-3a792d9c32b2 // indirect
)

replace github.com/99designs/gqlgen => github.com/marwan-at-work/gqlgen v0.0.0-20200107060600-48dc29c19314
