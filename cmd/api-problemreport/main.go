package main

import (
	log "github.com/sirupsen/logrus"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/iot-for-tillgenglighet/api-problemreport/pkg/database"
	"github.com/iot-for-tillgenglighet/api-problemreport/pkg/handler"
	"github.com/iot-for-tillgenglighet/messaging-golang/pkg/messaging"
	"github.com/iot-for-tillgenglighet/messaging-golang/pkg/messaging/telemetry"
)

func main() {

	serviceName := "api-problemreport"

	log.Infof("Starting up %s ...", serviceName)

	config := messaging.LoadConfiguration(serviceName)
	messenger, _ := messaging.Initialize(config)

	defer messenger.Close()

	database.ConnectToDB()

	messenger.RegisterTopicMessageHandler((&telemetry.Problemreport{}).TopicName(), receiveProblemreport)

	handler.Router()
}
