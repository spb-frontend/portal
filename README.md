# portal

This is a tool for inter-chat syncronisation. It's currenlty work-in-progress.
For now only available Gitter with Slack sync.

## Configuration

Place file ```config.json``` within executable's directory.

Example configuration:
```json
{
  "gitter": {
    "token": "TOKEN",
    "room": "ROOM"
  },
  "slack": {
    "token": "TOKEN",
    "channel": "CHANNEL",
    "channelId": "SLACK_CHANNEL_ID"
  }
}
```

More services and more flexible configuration is yet to come.
If you have a question, please open an [issue](https://github.com/spb-frontend/portal/issues).

PRs welcome!
