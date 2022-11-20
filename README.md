# Discord <-> IRC Relayer

This is a discord bot & webhook combo that relays an IRC Channel into discord, and vice versa.

To set this up, you need a server with a text channel. Add a webhook to that channel. Copy the URL for later.

You also need a bot account. Head to [the discord applications page](https://discord.com/developers/applications), and create a new bot account. Invite the bot, copy the token & enable the message content privilege.

Then get connection information for the IRC (hist, port, password, username, channel name)

Once you have all the information, create a `.env` file with the following content:

```sh
DISCORD_TOKEN="Your Discord token"
GUILD_ID="Your Discord Server (Guild) ID"
RELAY_CHAN="Webhook ID:Webhook Token:Discord Channel ID:IRC Channel name without #"
IRC_HOST="IRC Host:IRC Host token"
IRC_PASSWORD="IRC Password"
IRC_USER="IRC Username"
```

Once done, either use `go install github.com/ShadiestGoat/DiscordIRCRelay@latest`, or compile this repo yourself with `go build`, then run `DiscordIRCRelay` in a directory with the `.env`
