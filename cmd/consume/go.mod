module mainKafka

go 1.21.6

replace kafka/consume => ../../pkg/kafka

require kafka/consume v0.0.0-00010101000000-000000000000

require github.com/confluentinc/confluent-kafka-go v1.9.2 // indirect
