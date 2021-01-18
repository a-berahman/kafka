1-kafka-topics
    kafka-topics --bootstrap-server localhost:9092 --topic <name> --create --partition 3 --replication-factor 1
    kafka-topics --bootstrap-server localhost:9092 --list
    kafka-topics --bootstrap-server localhost:9092 --topic <name> --describe 
    kafka-topics --bootstrap-server localhost:9092 --topic <name> --delete   https://issues.apache.org/jira/browse/KAFKA-1194

2-kafka-console-producer
    kafka-console-producer --bootstrap-server localhost:9092 --topic <name>
    kafka-console-producer --bootstrap-server localhost:9092 --topic <name> --producer-property acks=all
    kafka-console-producer --bootstrap-server localhost:9092 --topic <name> 

3-kafka-console-consumer
    kafka-console-consumer --bootstrap-server localhost:9092 --topic <name>
    kafka-console-consumer --bootstrap-server localhost:9092 --topic <name> --from-beginning
    kafka-console-consumer --bootstrap-server localhost:9092 --topic <name> --group <name>

5-kafka-consumer-groups
    kafka-consumer-groups --bootstrap-server localhost:9092 --list
    kafka-consumer-groups --bootstrap-server localhost:9092 --topic <name> --group <name> --reset-offsets --to-earliest --execute
    kafka-consumer-groups --bootstrap-server localhost:9092 --describe --group <name>
    kafka-consumer-groups --bootstrap-server localhost:9092 --topic <name> --group <name> --reset-offsets --shift-by -2 --execute


