dev:
	echo "development" > .version
	go run .


build:
	go generate
	GOOS=linux GOARCH=amd64 go build -o socrates-discord-bot .

release:
	go generate
	docker build -t "registry.digitalocean.com/lucianonooijen/socrates:latest" .
	docker push "registry.digitalocean.com/lucianonooijen/socrates:latest"
