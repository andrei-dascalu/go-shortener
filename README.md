# Hexagonal Architecture Go Workshop #

## Description ##

Based on [Building Hexagonal Microservices with Go](https://www.youtube.com/watch?v=rQnTtQZGpg8) by [Tensor Programming](https://www.youtube.com/channel/UCYqCZOwHbnPwyjawKfE21wg)


### Differences ###

* Uses Fiber v2 framework
* Dockerised with Composer support for backends (redis/mongo)
* Prepared adapter for MySQL (via Gorm)
* Container with development tools (Air for auto-rebuild, Delve for debugging)
* Uses EasyJSON for JSON serialisation


## Resources ##

* https://dev.to/andreidascalu/setup-go-with-vscode-in-docker-for-debugging-24ch
* P1: https://www.youtube.com/watch?v=rQnTtQZGpg8
* P2: https://www.youtube.com/watch?v=xUYDkiPdfWs
* P3: https://www.youtube.com/watch?v=QyBXz9SpPqE


## Setup: MacOS ##

* Brew: `https://brew.sh`
* `brew install go`
* `brew install --cask visual-studio-code`
* `brew install docker`
* `brew install docker-compose`
* VSCode Extensions: LiveShare Pack (w/ Audio), REST Client

## Setup: Windows ##

* Choclatey: `https://chocolatey.org/install`
* `choco install golang`
* `choco install docker-desktop`
* `choco install docker-compose`
* `choco install vscode`
* `choco install vscode-go`
* `choco install vscode-vsliveshare`
* `choco install vscode-live-share-audio`
* VSCode REST client installed via VSCode


## Additional tools ##

These should be installed on-demand, as the workshop progresses

* EasyJSON: `go get -u github.com/mailru/easyjson/...` (`easyjson` should become available in path)
