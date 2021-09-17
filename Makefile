dev:
	echo "development" > .version
	go run .


release:
	go generate
	GOOS=linux GOARCH=amd64 go build -o socrates-discord-bot .
	scp ./socrates-discord-bot lucianonooijen.com:/home/bytecode/socrates-discord-bot/socrates-discord-bot-temp
	ssh lucianonooijen.com 'mv /home/bytecode/socrates-discord-bot/socrates-discord-bot-temp /home/bytecode/socrates-discord-bot/socrates-discord-bot'
	ssh lucianonooijen.com 'sudo systemctl restart socratesbot.service'
