# acectl - ACE Project command line control

## Install

### Build It Yourself

You have to install golang v1.7 or above

```shell
$ go install github.com/kkpoon/acectl
$ acectl help
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
$ acectl timesheet input -t 12345 --hours 8,8,8,8,8,0,0
```

### More usage

```shell
$ acectl help
```