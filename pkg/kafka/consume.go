package consumer

import (
	"database/sql"
	"encoding/json"
	// "encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	// "path/filepath"
)

var db *sql.DB

type Margin struct {
	// id            int
	amount        int
	currency_id   int
	user_id       int
	commission_id int
	// commission_id *int

	// account_id    int
	// executed_id   int
	// created_id    int
}

func Consume() {
	fmt.Println("start Consumer...")
	// brokers := [3]string{"b1.kafka.ramzinex.net:9092", "b2.kafka.ramzinex.net:9092", "b3.kafka.ramzinex.net:9092"}
	// certPath := filepath.Join("..", "..", "config", "kafka", "cert.pem") // Adjust path based on your project structure
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		// "bootstrap.servers":        "b3.kafka.ramzinex.net:9092",
		// "group.id":                 "aliId",
		// "auto.offset.reset":        "earliest",
		// "security.protocol":        "SSL",
		// "ssl.ca.location":          "../../config/kyc-kafka/club.pem",
		// "ssl.certificate.location": "../../config/kyc-kafka/club.pem",
		// "ssl.key.location":         "../../config/kyc-kafka/club.pem",
		"bootstrap.servers":        "b3.kafka.ramzinex.net:9092",
		"group.id":                 "secondaliId",
		"auto.offset.reset":        "earliest",
		"security.protocol":        "SSL",
		"ssl.certificate.location": "../../config/kafka/cert.pem",
		"ssl.ca.location":          "../../config/kafka/cacert.pem",
		"ssl.key.location":         "../../config/kafka/key.pem",
	})
	if err != nil {
		panic(err)
	}
	// c.SubscribeTopics([]string{"kyc_userinfo"}, nil)
	c.SubscribeTopics([]string{"margin"}, nil)
	for {
		fmt.Println("start for...")
		msg, err := c.ReadMessage(-1)
		fmt.Println("here1")
		if err == nil {
			fmt.Println("No error in kafkaConsumer!")
			// fmt.Printf(string(msg.Value))

			var myMsg map[any]int
			unerr := json.Unmarshal(msg.Value, &myMsg)
			if unerr != nil {
				fmt.Errorf("unmarshal error: %s", unerr)
			}
			fmt.Println("myMsg is: %s", myMsg)

			// fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))

			var marginMsg Margin
			marginMsg.amount = myMsg["id"]
			marginMsg.currency_id = myMsg["currency_id"]
			marginMsg.user_id = myMsg["user_id"]
			marginMsg.commission_id = myMsg["commission_id"]

			messageProccess(marginMsg)
		} else {
			fmt.Println("here2")
			fmt.Printf("error is: %s", err)
		}
		fmt.Println(("here 3"))
	}
	c.Close()
}

func messageProccess(msg Margin) {

	// if msg.commission_id != nil {
	result, err := db.Exec("INSERT INTO margin (amount, currency_id, user_id) VALUES (?, ?, ?)", msg.amount, msg.currency_id, msg.user_id)
	if err != nil {
		fmt.Errorf("addMargin failed: %v", err)
	} else {
		fmt.Println("result is: %s", result)
	}
	// }

}
