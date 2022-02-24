test_statik
=======

## Description

test http server with statik asset

## How to run

Run below command and start HTTP server on listening port `8080`

```console
statik -f && go run main.go
```

After that, see [http://localhost:8080/?body=Hello%20World](http://localhost:8080/?body=Hello%20World) on this browser 

or run below command.

```console
curl http://localhost:8080/?body=Hello%20World
```

## Author

[warawara28](https://github.com/warawara28)
