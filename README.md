# Socrates Discord bot

## Dependencies

- Golang

## How to change the status of the Discord bot

```py
import discord
from discord.ext import commands

client = commands.Bot(command_prefix = '.')

@client.event
async def on_ready():
    await client.change_presence(status=discord.Status.online,
                                 activity=discord.Game('debating plebs at the agora'))
    print('done')

client.run('token here')
```
