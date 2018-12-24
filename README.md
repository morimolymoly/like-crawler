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
lc download-latest-pictures [screenName] # download latest liked pictures(since you have downloaded with download-pictures-all subcommand)
```

## Config File
config file exists as `.lc.yaml` in your homedirectory.

```sh
cat ~/.lc.yaml
```
