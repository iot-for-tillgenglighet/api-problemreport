package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/iot-for-tillgenglighet/api-problemreport/pkg/database"
	"github.com/iot-for-tillgenglighet/api-problemreport/pkg/handler"
	"github.com/iot-for-tillgenglighet/messaging-golang/pkg/messaging"
)

func main() {

	serviceName := "api-problemreport"

	log.Infof("Starting up %s ...", serviceName)

	config := messaging.LoadConfiguration(serviceName)
	messenger, _ := messaging.Initialize(config)

	defer messenger.Close()

	db, _ := database.ConnectToDB()

	//	messenger.RegisterTopicMessageHandler((&telemetry.Snowdepth{}).TopicName(), receiveProblemreport)

	handler.CreateRouterAndStartServing(db)
}
