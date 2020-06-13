# dockerps-web
Run docker ps on browser

## How to launch

1. `docker build -t dockerps-web`

2. `docker run -v /var/run/docker.sock:/var/run/docker.sock -p 8080:8080 dockerps-web`
