# sidekick

## Install
```
curl -L https://github.com/yoyowallet/sidekick/releases/download/v0.0.1/sidekick_0.0.1_Darwin_x86_64.tar.gz | tar -xz
mv sidekick /usr/local/bin/
```

## Usage Example
```
# connect to staging VPN
cd path/to/platform
workon platform
sidekick --config-table Staging1-Config run -- yoyo shell
```
