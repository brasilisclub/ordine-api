# /bin/bash

docker build -t $USER/ordine-api-app:v1 ./go
docker build -t $USER/ordine-api-database:v1 ./go