docker build -t srv_reactive .
docker run --cpus=3 -p8080:8080 srv_reactive