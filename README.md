# RaceConditioner
A tool to exploit a race condition exploit in an web application. 
Similar syntax to curl for ease of use.

## NOTICE
This tool should only be used as a CTF-tool, and should not be used against non-consenting targets.

## Download
Linux/Mac:
```
git clone https://github.com/andytandy24/RaceConditioner.git
cd RaceConditioner
./setup.sh
```
You might need to give the setup script permission to execute

## Before use
If you have a local instance of the website you should first try it there with a simulated process like sleep to see if it is exploitable using this tool.

## Usage
```
Usage of RaceConditioner:
  -D    Toggles debug mode, is false by default
  -H string
        Specifies the headers of the request
  -X string
        Specifies the http method that should be used (default "GET")
  -cookie string
        Specifies the cookie of the request
  -data string
        Specifies the data that is meant to be sent
  -n int
        Defines the amount of simultaneous requests sent to a target (default 20)
  -u string
        Specifies the target url (default "http://127.0.0.1")
```

## Examples
```
RaceC -X "GET" -u "http://127.0.0.1?do=buy"
```
```
RaceC -X "POST" -u "http://127.0.0.1" -H "Content-Type=application/json;Accept=jpeg" -data "username=atand&points=999"
```