docker exec -it 3197 kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 3 --topic test-tp
docker exec -it 3197 kafka-topics.sh --bootstrap-server localhost:9092 --list
docker exec -it 3197 kafka-console-producer.sh --bootstrap-server localhost:9092 --topic test-tp
docker exec -it 3197 kafka-console-consumer.sh --from-beginning --bootstrap-server localhost:9092 --topic test-tp