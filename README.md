# Let's RUN PbType


## What to do
This repository defines the API gateway and internal platform service/message structure using protocol buffer syntax. \
The `gateway` directory contains the public API endpoints and services described for HTTP and GRPC clients. \
The `platform` directory contains the messages and services described for GRPC clients. These interfaces in the `platform`
directory will be shared among the micro services dedicated to its specific ecommerce platform.

## How to use
- Run `./generate` to generate Golang code based on the protocol buffer declaration.