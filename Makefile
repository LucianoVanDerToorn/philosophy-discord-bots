dev:
	echo "development" > .version
	go run .


release:
	go generate
	GOOS=linux GOARCH=amd64 go build -o socrates-discord-bot .
	scp -P 4125 ./socrates-discord-bot bytecode@164.90.202.218:/home/bytecode/socrates-discord-bot/socrates-discord-bot-temp
	ssh bytecode@164.90.202.218 -p 4125 'mv /home/bytecode/socrates-discord-bot/socrates-discord-bot-temp /home/bytecode/socrates-discord-bot/socrates-discord-bot'
	ssh bytecode@164.90.202.218 -p 4125 'sudo systemctl restart socratesbot.service'
