
# Events Service

Convert web hits to pubsub events!


## Build

This is a go app inside of a docker container.
To build

    make build
    make tag
    make push


## Deploy

This simple docker image

    markmims/events

is intended to be deployed on k8s... widely behind a load balancer in order to
help with load testing.


## Usage

The events service is available on tcp/80 from container startup.

This service will accept any GET request path, block while sending to a pubsub
topic (running benchmarks), and then respond with a simple message when done.

The service only generates events when hit with web requests.

The service does not queue multiple requests... it blocks by design to try to
rely on surrounding infrastructure tooling to even out the flow.
