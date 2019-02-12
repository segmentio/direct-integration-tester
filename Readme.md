
# The segment endpoint tester

This endpoint tester submits realistic looking data to your Segment endpoint. 

In particular, the endpoint tester uses a scrubbed version of our production
traffic to generate realistic looking events for you to use. It contains 250
randomized payloads, designed to give you a sense of the variance we see
across customers. 

## Quickstart

First, download the binary for your particular operating system from our
[github releases page][1].

Once you have this file, it's simply a matter of running it from the 
command line:

```
$ direct-endpoint-tester API_KEY=<your-api-key> ENDPOINT=<your-direct-endpoint>
```

This will send roughly 250 different requests to your endpoint. 

## Developing locally

First, install the required dependencies

```
$ goto direct-endpoint-tester
$ go get -u ./...
```

Then run the following

```
$ API_KEY=<your-api-key> ENDPOINT=<your-direct-endpoint> make run
```

## Releasing

If you haven't already, install packr on your system:

```
$ go get -u github.com/gobuffalo/packr/packr
```

To release, then simply run the following make command:

```
$ make release
```

This will generate the necessary binaries to upload to github. 

## License

MIT


[1]: https://github.com/segmentio/direct-integration-tester/releases