# go-server
Code I tend to rewrite for go servers

## Introduction
Whenever I start working on a new Golang microservice, I find myself writing code that is similar to code I wrote when writing previous services.  The purpose of this repo is to capture how I write a basic REST microservice in go.  This repo is not intended to be imported as a dependency, but rather serve as a template or batch of code that can be copied and then customized.

## Using this code
This repo is published under an [MIT License](LICENSE) which tends to be flexible for the user while still protecting the creator.  Please feel free to use this code and also to submit Pull Requests with any fixes/improvements/suggestions/etc.  Also, if you use code from the repo I request (but do not require) that you star the repo in GitHub.

## Running the code

### Setup development environment
The server uses properties stored in a `.env` file.  To generate said file, from the project root run `scripts/create_env.sh`.

### Running the Server
From the project root run `go run .`



### Prerequisites
- Go (see the [go.mod](https://github.com/NathanBak/go-server/blob/755e067fd4b192641c8478422a49549e316e137c/go.mod#L3) file for the correct version) 
- Git