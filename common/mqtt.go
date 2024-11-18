package common

import (
	"fmt"
	"net"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func getLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}

	return "", fmt.Errorf("no IP address found")
}

func broadcastServerIP() {

	localIP, err := getLocalIP()
	if err != nil {
		fmt.Println(err)
		return
	}

	broker := GetConfigString("mqtt.broker")
	port := GetConfigInt("mqtt.port")
	topic := GetConfigString("mqtt.topic")

	opts := mqtt.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return
	}

	token := client.Publish(topic, 0, false, localIP)
	token.Wait()

	client.Disconnect(250)
}

func BroadcastServerIP() {
	for {
		broadcastServerIP()
		time.Sleep(2 * time.Second)
	}
}
