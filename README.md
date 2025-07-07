# Go-Cache-Server

#### For client side please visit  [repo](https://github.com/praveenmahasena/go_cache_client).
#### This project is still on development

Go-Cache-Server acts as server side of this "Redis like database" in memory database system.</br>

## Features
 - In memory database
 - Concurrent
 - No race conditions
 - Has string, set, hash, list data structures to store data

## Tech Stack
 - [Golang](https://go.dev/dl)
 - [Git](https://git-scm.com/downloads)optional

## Server side Setup

 1. ### Clone repo

    ```bash
    git clone https://github.com/praveenmahasena/go_cache_server.git
    ```

 2. ### Change working dir

    ```bash
    cd /go_cache_server
    ```

 3. ### Build binary

    ```bash
    make
    ```

 4. ### Run
    ```bash
    ./bin/go_cache_server -p "<portID_to_server_should_run_on>"
    ```
## TODO
 Impliment graceful shutdown
