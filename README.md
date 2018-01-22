
# RedFull API

Open-source api created with Go based on [Widgets Single Page App][]

[Widgets Single Page App]: https://github.com/RedVentures/widgets-spa

## How to run

### With Docker

#### Prerequisites

* [Docker]
* [Docker Compose 1.8]

Remember to check if both are updated to the last stable version.

[docker]: https://www.docker.com
[docker compose 1.8]: https://docs.docker.com/compose/install/

To run the application via Docker, follow the commands below:

```shell
go get github.com/velrino/RedFull
cd $GOPATH/src/github.com/velrino/RedFull
docker-compose up
```

### Compiling

To compile the application and run in your machine, you must first create a `.env` file containing the secret used by the application to generate authentication tokens. You can use the `.env.example` file as an example:

```shell
go get github.com/velrino/RedFull
cd $GOPATH/src/github.com/velrino/RedFull
go install
go run main.go
```

In both cases, the application will run in the `7878` port, so you can access it in `http://localhost:7878`.

## Documentation

The documentation was generated with [ApiDocs][] and is accessible in `http://localhost:7878/docs`

[ApiDocs]: http://apidocjs.com/
