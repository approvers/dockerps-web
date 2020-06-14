# dockerps-web
Run `docker ps -a` on browser.

Tested on Linux and Docker for Windows.

## How to launch

### Use GitHub packages
1. Make sure you logged into GitHub package registry on your Docker client.

2. `docker run -v /var/run/docker.sock:/var/run/docker.sock -p 8080:8080 docker.pkg.github.com/approvers/dockerps-web/image:latest`

### Build on local

#### Use Go directly
1. `go run main.go`

#### Use docker
1. `docker build . -t dockerps-web`

2. `docker run -v /var/run/docker.sock:/var/run/docker.sock -p 8080:8080 dockerps-web`
