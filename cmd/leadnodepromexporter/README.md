To build `registry.gitlab.com/omt-gateway/ticker/leadnodepromexp:latest`, issue the following command sequence in this directory:

1. `docker login registry.gitlab.com` - uses your Gitlab creds, same as for git operations
2. `docker build -t registry.gitlab.com/omt-gateway/ticker/leadnodepromexp:latest .`
3. `docker push registry.gitlab.com/omt-gateway/ticker/leadnodepromexp:latest`

To run `registry.gitlab.com/omt-gateway/ticker/leadnodepromexp:latest`, issue the following sequence of commands

1. `docker pull registry.gitlab.com/omt-gateway/ticker/leadnodepromexp:latest`
2. `docker run -p 2112:2112 leadnodepromexporter:latest` (please note that port 2112 is exposed in the container and this command binds to host port 2112)

To test, issue the following on the same machine on which the container is running:

1. curl http://localhost:2112/metrics

You should get output such as is shown [here](https://prometheus.io/docs/guides/go-application)