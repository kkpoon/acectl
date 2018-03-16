# acectl - ACE Project command line control

[![Github All Releases](https://img.shields.io/github/downloads/kkpoon/acectl/total.svg)](https://github.com/kkpoon/acectl/releases)
[![GitHub (pre-)release](https://img.shields.io/github/release/kkpoon/acectl/all.svg)](https://github.com/kkpoon/acectl/releases)

## Install

### Download Binary

You could download binary from github [release page](https://github.com/clustertech-dev/acectl/releases)

### Build It Yourself

1. You have to install [golang](https://golang.org)
2. Install [dep](https://github.com/golang/dep)

```shell
$ dep ensure
$ go build
$ ./acectl help
```

## usage

### Login

The login command helps you to get the authentication token from ACEProject
and store it in your home directory, default: `$HOME/.acectl.json`

```shell
$ acectl login
```

### Timesheet

```shell
# list the task belongs to you, find the task id of your task
$ acectl task list
# Log time to timesheet on task id=12345, mon=8, tue=8..., sat=0, sun=0
$ acectl timesheet input --hours 8,8,8,8,8,0,0 -t 12345
saved at Timesheet Line ID: 7689
# oh no, the time is not correct, lets update it
$ acectl timesheet input --hours 0,7,0,0,0,0,0 -t 12345 -u 7689
```

### More usage

```shell
$ acectl help
```
