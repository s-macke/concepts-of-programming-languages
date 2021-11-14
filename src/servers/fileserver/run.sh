docker build -t fs .
echo "Size of Docker image"
docker image inspect fs --format='{{.Size}}'
docker run -p8080:8080 fs