1. Overview
Apache Kafka is a renowned distributed streaming platform for streaming records in real-time. It permits developers to publish and subscribe to streams of records using Kafka topics. We can use docker-compose to manage Kafka streams using multi-container Docker applications.

In this tutorial, we’ll learn to create a Kafka topic using Docker Compose. Also, we’ll publish and consume messages from that topic using a single Kafka broker.

2. Setup Kafka using Docker Compose
Docker Compose tool is handy for managing multiple service setups. The Kafka cluster requires Zookeeper and Kafka brokers, which is very useful in such a case. To set up a Kafka cluster, we need to run two services: Zookeeper and Kafka Brokers.

Let’s look at the docker-compose.yml to set up and run the Kafka server:

version: '3'
services:
  zookeeper:
    image: wurstmeister/zookeeper
    container_name: zookeeper
    ports:
      - "2181:2181"
    networks:
      - kafka-net
  kafka:
    image: wurstmeister/kafka
    container_name: kafka
    ports:
      - "9092:9092"
    environment:
      KAFKA_ADVERTISED_LISTENERS: INSIDE://kafka:9092,OUTSIDE://localhost:9093
      KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: INSIDE:PLAINTEXT,OUTSIDE:PLAINTEXT
      KAFKA_LISTENERS: INSIDE://0.0.0.0:9092,OUTSIDE://0.0.0.0:9093
      KAFKA_INTER_BROKER_LISTENER_NAME: INSIDE
      KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181
      KAFKA_CREATE_TOPICS: "baeldung:1:1"
    networks:
      - kafka-net
networks:
  kafka-net:
    driver: bridge
Copy

freestar

In the above docker-comose.yml, we run two different services. The zookeeper service uses “wurstmeister/zookeeper” as base image. Furthermore, we exposed the 2181 port of the container to the host machine, which allows external access to the Zookeeper. Additionally, we run the Kafka service with “wurstmeister/Kafka” as the base image. To allow external access to the Kafka server, we exposed the 9092 port.

We provided essential ENV variables to the Kafka service. These ENV variables configure listener settings inside and outside the Docker network. It uses the PLAINTEXT protocol, defines bind addresses for both listeners (0.0.0.0), specifies “INSIDE” as the inter-broker listener name, provides the connection string to Zookeeper (using container name “zookeeper” and port 2181), and sets up a topic named “baeldung” with replication factor 1 and partition count 1.

3. Start the Kafka Cluster
Till now, we looked at creating services to run Zookeeper and Kafka using docker-compose.yml. To demonstrate, let’s look at the command to start the Kafka cluster:

$  docker-compose up -d   
[+] Running 2/2
 ✔ Container zookeeper  Running  0.0s 
 ✔ Container kafka      Started  
Copy
The above command will simply start the Zookeeper and Kafka containers in detached mode. This means we can interact with Kafka and Zookeeper while the terminal remains free from other tasks. The -d mode ensures that the container runs in the background, which helps us manage the Kafka environment properly.

3.1. Verifying the Kafka Cluster
In order to verify that the Kafka cluster is running successfully, let’s run the following command to see the running containers:

$ docker-compose ps
NAME                IMAGE                    COMMAND                  SERVICE        CREATED          STATUS          PORTS
kafka               wurstmeister/kafka       "start-kafka.sh"         kafka          27 seconds ago   Up 27 seconds   0.0.0.0:9092->9092/tcp
zookeeper           wurstmeister/zookeeper   "/bin/sh -c '/usr/sb…"   zookeeper      3 minutes ago    Up 3 minutes    22/tcp, 2888/tcp, 3888/tcp, 0.0.0.0:2181->2181/tcp
Copy
In the above output, we can see that the Zookeeper and Kafka containers are up and running. This verification step confirms that the Kafka cluster is ready to handle topics and messages. Therefore data from various sources can be ingested.

3.2. Creating a Kafka Topic
Till now, the Kafka cluster is up and running. Now, let’s create a Kafka topic:

$ docker-compose exec kafka kafka-topics.sh --create --topic baeldung_linux
  --partitions 1 --replication-factor 1 --bootstrap-server kafka:9092
Copy
In the above command, a new topic named baeldung_linux is created with 1 partition and 1 replica set using a Kafka broker on port 9092. Now using this topic, we can streamline Kafka data. This will allow us to exchange messages and events for a variety of applications.

3.3. Publishing and Consuming Messages
With the Kafka topic in place, let’s publish and consume some messages. First, start a consumer by running the following command:

$ docker-compose exec kafka kafka-console-consumer.sh --topic baeldung_linux
  --from-beginning --bootstrap-server kafka:9092
Copy
Using the above command, we’ll be able to consume all the messages sent over to this topic. Additionally, we used –from-beginning to consume all messages sent over the topic from the beginning. Let’s also look at publishing the data to this Kafka topic:

$ docker-compose exec kafka kafka-console-producer.sh --topic baeldung_linux
  --broker-list kafka:9092
Copy
By using the above command, we can generate and send messages to the baeldung_linux topic. Sending messages to Kafka topics using Docker is simple and efficient.

4. Conclusion
In this article, we explored how to create a Kafka topic using Docker Compose. First, we set up a full-fledged Kafka cluster using two different services. We run Kafka Zookeeper and Kafka Broker using Docker Compose. After that, we created a Kafka topic to publish and subscribe to messages using the Kafka cluster.

In short, we created a simple Kafka cluster using docker-compose services. This whole setup helps us distribute and communicate data in real-time within our Kafka-based applications via docker-compose services. Additionally, it helps us save on resources since we don’t have to run the Kafka cluster all the time.

 Subscribe 
guest


{}[+]
0 COMMENTS
