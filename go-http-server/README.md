# Sample HTTP Server in GoLang

The following go-http-server is developed with the help of this [blog](https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go)

with this project, able to create two servers and listen to them

Start the server in one terminal with

```bash
go run main.go
```

open another terminal and run the following curl commands

```bash
curl http://localhost:3333/
curl http://localhost:4444/
curl http://localhost:3333/hello

curl 'http://localhost:3333?first=1&second='
curl -X POST -d 'This is the body' 'http://localhost:3333?first=1&second='
curl -X POST -F 'myName=Sammy' 'http://localhost:3333/hello'

```

and notice the output and also the server response.

You can stop the server with `CONTROL+C`
