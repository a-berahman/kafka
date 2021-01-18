# Kafka Learning Path
Kafka repository is an educational project for persion developer community.

You can follow Video Series in [Youtube](https://youtube.com/playlist?list=PLKmjtHAd1JMe3F-Q9j4UY9qVWwT9HGM05)
## What is Kafka? ([Kafka official document](https://kafka.apache.org/documentation/))
Kafka is a streaming message platform. Breaking it down a bit further:

- “Streaming”: Lots of messages (think tens or hundreds of thousands) being sent frequently by publishers ("producers"). Message polling occurring frequently by lots of subscribers ("consumers").

- “Message”: From a technical standpoint, a key value pair. From a non-technical standpoint, a relatively small number of bytes (think hundreds to a few thousand bytes).

- If the above isn’t your planned use case, Kafka may not be the solution you are looking for. Contact your favorite Cloudera representative to discuss and find out. It is better to understand what you can and cannot do upfront than to go ahead based on some enthusiastic arbitrary vendor message with a solution that will not meet your expectations in the end.


## Prerequisite 

- Java 1.8
- Kafka binary, You can download the file [here](https://kafka.apache.org/downloads).

## FAQ for Setup Problems([Kafka official FAQ](https://cwiki.apache.org/confluence/display/KAFKA/FAQ))


> Zookeeper - java.net.BindException: Address already in use

Something is already occupying your port 2181. Figure out which application it is and stop it

> Kafka - org.apache.kafka.common.KafkaException: Socket server failed to bind to 0.0.0.0:9092: Address already in use.

Something is already occupying your port 9092. Figure out what it is and stop it.
Otherwise, if you really insist, you can change the Kafka port by adding the following line to server.properties

> I have launched Kafka in a VM or in the Cloud, and I can't produce to Kafka

If you can't produce to Kafka, it's possible you are using a VM and this can break the Kafka behaviour. Please look at the annex lectures for solutions of how to deal with that. I strongly recommend doing this tutorial using the Kafka binaries and localhost

> The topics list is disappearing

Make sure you have changed the Zookeeper dataDir=/path/to/data/zookeeper , and Kafka log.dirs=/path/to/data/kafka

### example for port 9093
listeners=PLAINTEXT://:9093
> My topics are losing their data after a while

This is how Kafka works. Data is only retained for 7 days.



## License
[MIT](https://choosealicense.com/licenses/mit/)
