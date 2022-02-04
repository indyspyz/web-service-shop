# web-service-shop

***** Command *****


# DOCKER
<!-- ** Get postgres docker image ** -->
docker pull postgres:14

<!-- ** Start docker-compose ** -->
docker-compose up -d db

<!-- ** Stop docker-compose ** -->
docker-compose down

<!-- ** Remove Container docker ** -->
docker rm -f $(docker ps -a -q)

# SERVICE
<!-- ** Start service ** -->
go run main.go
