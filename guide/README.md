## Basic Setup

- Fork the repository [The-Origin-Labs/petrovedge](https://github.com/The-Origin-Labs/petrovedge)

- ```shell
    git clone https://github.com/<username>/petrovedge
    cd petrovedge
    ```

## Setting up Python-dev environment

- Install Python3 and above [python.org](https://www.python.org/)

- ```javascript
    cd services
    mkdir <service-name>
    python -m venv .venv
    source .venv/Scripts/activate
    ```

## Setting up Go/Golang-dev environment

- Install Go/Golang 1.17 and above [go.dev](https://go.dev/)

- ```javascript
    cd api
    go get .
    go run cmd/main.go
    ```
## Setting up Nodejs-dev environment


- Install [NodeJs](https://nodejs.org/en/)

- ```javascript
    cd services
    mkdir <service-name>; cd <service-name>
    npm init -y
    ```

