# RKVS - A Remote Key Value Store
Barebone, single BadgerDB instance exposed over gRPC.

Installation:
install --ip <your_ip> --port <grpc_port> --http_port <http_port>

Exposes its own proto file over web using the http_port. Useful for easy client connection in distributed systems.
