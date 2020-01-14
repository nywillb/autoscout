# FTC Autoscout

> Automatic scouting for First Tech Challenge Teams

- [⬇️ **Download**](https://github.com/willbarkoff/autoscout/releases) 
- [🐛 **Report a bug**](https://github.com/willbarkoff/issues/new)

---

Autoscout provides data about FTC teams at a given competition.

It currently supports 2 score reporting platforms:
- [Pennsylvania FIRST](http://www.ftcpenn.org/)
- [The Orange Alliance](http://theorangealliance.org/)

While I'd love to support [FTC Scores](https://ftcscores.com/) in the future, currently, it isn't supported.

## Use with [Pennsylvania FIRST](http://www.ftcpenn.org/)

In the directory that you are running the program from, create a file called `config.toml`. This file contains the configuration infromation for the server. A sample config file is available in the [`config.sample.pennfirst.toml`](config.sample.pennfirst.toml) file.

Populate the file as demonstrated below:

```toml
[Stats]
Type = "Penn FIRST" # Must be "Penn FIRST"
URL = "http://detroit.worlds.pennfirst.org/" # The URL used for score reporting
division = "Edison" # The division to scout for.
```

## Use with [The Orange Alliance](http://theorangealliance.org/)

In the directory that you are running the program from, create a file called `config.toml`. This file contains the configuration infromation for the server. A sample config file is available in the [`config.sample.toa.toml`](config.sample.toa.toml) file.

Populate the file as demonstrated below:

```toml
[Stats]
Type = "TOA" # Must be "TOA"
TOAKey = "secret TOA key" # Your TOA Key
TOAOrigin = "Autoscout" # Leave this line as "Autoscout"
TOAEventKey = "TOA event key" # The event key for TOA
```

You can find your TOA key in the your account page. First, register or sign in, then click "Generate API Key," and copy and paste the generated API key.

The TOA Event Key can be found in the URL of the event, for example, the URL of the results page for the [2020 New York City Championship](https://theorangealliance.org/events/1920-NY-NFTCC/rankings) is https://theorangealliance.org/events/1920-NY-NFTCC/rankings, and the event key is `1920-NY-NFTCC`.

## Building from Source
To build from source, you must have [The Go Programming Language](https://golang.org) installed. You can install Go at their website: [golang.org](https://golang.org). 

To download the source, you can type

```shell
$ go get github.com/willbarkoff/autoscout # get the source
$ cd src/github.com/willbarkoff/autoscout # cd into the source
```

Then, install dependencies:

```shell
$ go get
```

Finally, compile and run
```shell
$ go install
$ ../../../../bin/autoscout
```

Pull requests are welcome!