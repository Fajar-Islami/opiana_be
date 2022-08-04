entermysql:
	docker exec -it mysql_opiana mysql -uroot -pSECRET_ROOT

composerestart:
	docker-compose down -v
	docker-compose up -d

compile:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/app
