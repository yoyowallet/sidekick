# sidekick

You will need to install [GoLang](https://go.dev/doc/install)

## Local Installation

### Installation

```shell
mkdir $GOPATH/src/github.com/yoyowallet

cd $GOPATH/src/github.com/yoyowallet

git clone https://github.com/yoyowallet/sidekick.git

cd sidekick

go get ./...

```

### Local Build (Executable)

#### Build

```shell

go build ./...

```

#### Execute

```shell

./sidekick help

```

```shell

NAME:
   sidekick - A new cli application

USAGE:
   sidekick [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
     env      
     run      
     start    
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --config-source value  (default: "dynamodb") [$SIDEKICK_CONFIG_SOURCE]
   --config-key value     (default: "common") [$SIDEKICK_CONFIG_KEY]
   --config-table value    [$SIDEKICK_CONFIG_TABLE]
   --env value, -e value  set environment variables
   --help, -h             show help
   --version, -v          print the version


```

## Build/Release

### create a release by using goreleaser
```shell

goreleaser release

// OR

goreleaser --rm-dist

```

## Install Build / Release

Supported platforms:

- Darwin_i386
- Darwin_x86_64
- Linux_i386
- Linux_x86_64

```shell


export SIDEKICK_RELEASE_VERSION=v0.0.2
export SIDEKICK_RELEASE_PLATFORM=Darwin_x86_64

curl -L "https://github.com/yoyowallet/sidekick/releases/download/$SIDEKICK_RELEASE_VERSION/sidekick_${SIDEKICK_RELEASE_VERSION}_${SIDEKICK_RELEASE_PLATFORM}.tar.gz" | tar -xz

mv sidekick /usr/local/bin/

```

## Usage Example
```

sidekick --config-table DynamoDB-TableName run -- yoyo shell

// Specific key Value

sidekick --config-table DynamoDB-TableName config-key my-key-value run -- yoyo shell

```