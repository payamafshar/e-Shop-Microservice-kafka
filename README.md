Partitions: Well, this one can be a little tricky. A partition is the smallest method of concurrency. Within a consumer group, no two consumers can read off a partition.

Replication factor: To ensure resiliency, Kafka defines a node as the topic leader and copies its data among other Kafka nodes. With this option, we can define how many copies of this topic we need.

ZooKeeper: an application that communicates directly with Kafka nodes, assigns Node leaders in a cluster, and ensures they are alive. In general, it manages Kafka nodes in a distributed environment.

---

1-For writing message with kafka first we neeed create topic with partition with kafka shel or if use kafka-go it has option to create automatic topic without topic creation our program will be panic:
kafka-topics.sh --zookeeper zookeeper:2181 --create --topic twitter.newTweets --replication-factor 1 --partitions 10
error<->

---
