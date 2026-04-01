# curl

The curl "Client URL" command is a powerful command-line tool in Linux used for transferring data to or from a server using a wide range of network protocols. It is widely used for tasks such as downloading files, testing APIs, and automating web requests in scripts.

## Basic Usage

1. curl http://localhost:8080
- http:// : Protocol
- localhost : IP Address - here, host machine (127.0.0.1)
- 8080 : Port

Will return whatever server responds with

2. curl -I http://localhost:8080
- Show headers only

3. curl -v http://localhost:8080
- Shows connection steps, IP, port, request, and response
- Great for debugging network/server issues

4. curl -X GET http://localhost:8080/
- -X flag used to set HTTP Method