docker ps -l
docker ps -l -a | grep 'name'

docker logs -f id

docker exec -it id sh

docker container rm -f id

net addr show docker0           查看素主机ip

