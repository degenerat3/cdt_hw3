# Crowd_Control
This project is a C2 that utilizes web requests to deliver commands.  This repo contains the server/client/controller.  The server is a docker container which runs a flask app, the controller is a python script which allows users to push commands to the server, and the client(s) are a series of golang programs and bash scripts that will invoke the proper web request to pull commands, then execute them.  

## Server
The server is a docker container that's running alpine linux with flask.  All the server functions/endpoints can be found in the `app.py` file.  

#### Endpoints
The server has the following endpoints:
 - `'/\<ip>/\<typ>' or '/api/callback/\<ip>/\<typ>'` - these are the endpoints the clients will hit in order to receive their commands.  The former is preferred, as the latter will soon be phased out.  The "\<ip>" is the ip of the client that's calling back, and the "\<type>" denotes where the callback is coming from (bash script, golang binary, vimrc) which is used for logging purposes.
  - `'/api/commander/push'` - this is the endpoint that the commander script/CLI will send commands to.  It accepts a JSON POST that contains the target hosts and commands to be executed.
  - `'/api/commander/calls'` - this is an endpoint that serves a log.  It's returns the log that tracks all client callbacks.  Any time a client hits one of the callback endpoints, an entry is made containing "Time | IP | Type", this endpoint returns all entires
  - `'/api/commander/tasks'` - this is the other endpoint that serves a log.  Any time a command is pushed to the server, an entry is added to the task log containing "Time | Targets | Tasks", this endpoint returns all entires.  

## Controller
Command tasks are sent to the server using the `commander.py` file, which is a command line interface.  The CLI allows users to show tracked hosts, then create a new task and set the commands and targets. Example:
```
commander@box$ python3 commander.py
Welcome to Crowd Control Commander
Current server: http://127.0.0.1:80

Type 'help' for CLI options...
Commander> help
Current server: http://127.0.0.1:80
To create new task:  `new task`
New task shortcut:   `t: host[ hosts...]: commands`
Set server IP:       `set server http://0.0.0.0:5000
To view hosts:       `show hosts`
To quit:             `exit`
Unknown command: help
Commander> new task
Commander(TASK)> help
To view hosts:       `show hosts`
To view task info:   `show task`
To add hosts:        `set host = [8.8.8.8]`
To add commands:     `set command = "whoami"`
To launch the task:  `launch`
To quit:             `exit`
Commander(TASK)> set host = [10.1.1.2]
Commander(TASK)> set command = "systemctl stop apache2"
Commander(TASK)> launch
```


## Clients
There is currently only support for Linux clients, although Windows endpoints have been created, so it's a WIP.  The client script can be anything from a golang binary, a python or bash script, etc, as long as it can invoke a web request.  The format is very simple, simply send a GET request to `serverIP/\<ip>\<type>` and the return will be the commands that need to be executed.  

## Future Work:
 - Windows endpoints have been created, we need to write some clients that will use cmd and/or PowerShell (preferably PowerShell).
 - Add a function to `mace.py` that can read hosts from a file, so the arguments would be:  
 `th: targets.txt: echo hello`, where 'targets.txt' would contain a newline separated list of IP addresses.
 - Write a stager script that will pull a client script/binary/whatever, set it up as a scheduled task/service/whatever.  Once this stager script is hosted on the server (must make a new stager endpoint), hosts can be added to the control infrastucture by simply invoking a web request and piping the script into bash/PowerShell.