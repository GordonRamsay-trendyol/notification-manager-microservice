container_id=$(input)

# notification server topics
docker exec -t {container_id} kafka-topics.sh --bootstrap-server :9092 --create --topic sms_notification --partitions 3 --replication-factor 1
docker exec -t {container_id} kafka-topics.sh --bootstrap-server :9092 --create --topic email_notification --partitions 3 --replication-factor 1
docker exec -t {container_id} kafka-topics.sh --bootstrap-server :9092 --create --topic push_notification --partitions 3 --replication-factor 1

# user update topic 
# use this topic whenever user changes their email, firstname, lastname or phone number.
docker exec -t {container_id} kafka-topics.sh --bootstrap-server :9092 --create --topic user_update --partitions 3 --replication-factor 1

# list topics if they are created well or not
docker exec -t {container_id} kafka-topics.sh --bootstrap-server :9092 --list