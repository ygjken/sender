# CLI to Post Program Output to Slack
This is a program that forwards the results returned by command line application to Slack.

## Usage
* First, add Slack WebHook url to `./config.toml`. 
* Use with `nohup`, you can sent output to slack.
```
nohup ./sender echo msg &
```

## TODO
- [x] Add code to read the urls of webhook from config file.
- [x] Receive command and make it easier to execute the `nohop` command.
