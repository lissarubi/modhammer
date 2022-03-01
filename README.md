# Mod Hammer

Mod hammer do a same action in a list of Twitch channels. This program is very useful to ban users from a list of channels.

# Installation

## With Golang

If you alredy has [Golang](https://golang.org/) installed in your machine, run:

```
go install github.com/edersonferreira/modhammer
```

## Without Golang

If you don't have [Golang](https://golang.org/) installed in your machine, do these commands:

```
git clone https://github.com/edersonferreira/modhammer
cd modhammer
sudo cp bin/modhammer /usr/bin/modhammer
```

# Configuration

Run `modhammer --setup` to set the username, Twitch IRC Token, and the channels what do you want to connect.

```
Put your Twitch Username:
> myusername
Put your Twitch Token:
> oauth:token
Put the channels on which do you want to send the messages (split with comma)
> channel01,channel02
```

# Usage

Run `modhammer [message]` to send this message in all the channels. Like:

```
modhammer /ban certainuser
```

or

```
modhammer /clear
```
