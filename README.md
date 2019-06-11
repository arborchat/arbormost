# ArborMost

A one-way bridge that relays arbor messages from an arbor server onto a mattermost server.
There is no way to reply to the arbor messages except by joining the arbor server on which
they were sent.

You can get information about the Arbor project [here](https://man.sr.ht/~whereswaldon/arborchat/).

Install:
```
# the mattermost codebase is huge, and this uses part of it. Download may take some time
go get -v -u github.com/arborchat/arbormost
```

Example usage:

```
arbormost -username <mattermost username> -password <mattermost password> -team <mattermost team> -channel <mattermost channel> -url <mattermost server address> -arbor-address <arbor server IP:PORT>
```

Note that `team` and `channel` are specified in their URL-safe format. Look at the URL of a
channel in Mattermost to find its team and channel. They'll look something like:
`https://example.com/teamName/channels/channelName`
