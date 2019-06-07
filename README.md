# Simple HTTP test server

This is a very simple HTTP server which dumps all details of every incoming request for testing purposes.

You need to specify on which port(s) it should listen, the `-port` parameter can be specified multiple times.

The main use for this is to run this as a docker image, to mock backend HTTP services for the API gateway.