docker build -t srv .
docker run --cpus=3 -p8080:8080 srv
