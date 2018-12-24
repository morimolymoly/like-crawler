# like-crawler

## Build
```sh
dep ensure && go build
```

## Usage
```sh
lc setcreds # set credentials
lc setsavepath # set savepath
lc download-pictures-all [screenName] # download all liked pictures
lc download-latest-pictures [screenName] # download latest liked pictures(config file know where to start to download)
```

## Config File
config file exists as `.lc.yaml` in your homedirectory.

```sh
cat ~/.lc.yaml
```

### content
* accesssecret
* accesstoken
* consumerkey
* consumersecret
* savepath - path
* sinceid - download pictures which is included in tweet after this tweet id


