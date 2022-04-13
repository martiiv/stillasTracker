package apiTools

import (
	"encoding/hex"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/ingics/ingics-parser-go/ibs"
	"github.com/ingics/ingics-parser-go/igs"
	"log"
	"strings"
)

type AdvPacket struct {
	msg    *igs.Message
	packet *ibs.Payload
}

var AdvPacketChannel = make(chan AdvPacket, 1000)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, mqttMsg mqtt.Message) {
	// log.Printf("Received MQTT message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	s := string(mqttMsg.Payload())
	if igsMsg := igs.Parse(s); igsMsg != nil {

		if bytes, err := hex.DecodeString(igsMsg.Payload()); err == nil {
			pkt := ibs.Parse(bytes)
			fmt.Println(pkt)

			if model, ok := pkt.ProductModel(); ok && strings.HasPrefix(model, "iBS") {
				fmt.Println(pkt)
				AdvPacketChannel <- AdvPacket{igsMsg, pkt}
			}
		} else {
			log.Printf("invalid payload: %s: %s", err, igsMsg.Payload())
		}
	} else {
		log.Printf("invalid message: %s\n", s)
	}
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Println("MQTT connected")

	go func() {
		topic := "pub"
		token := client.Subscribe(topic, 1, messagePubHandler)
		token.Wait()
		log.Printf("Subbed to topic: " + topic + "\n")
	}()
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	log.Printf("MQTT Connect lost: %v", err)
}

func InitializeMQTTClient() {
	var broker = "broker"
	var port = 1883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("BLE-Parser")
	opts.SetUsername("admin")
	opts.SetPassword("admin")
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

}
