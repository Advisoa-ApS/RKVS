# RKVS - Remote Key Value Store
Just a single BadgerDB instance exposed over gRPC.

Can be installed as a systemd service on Linux:
install --ip <your_ip> --port <grpc_port> --http_port <http_port>

Exposes its own proto file over web using the http_port (useful for easy client connection in distributed systems incompatible with go modules).

Supported methods:
* ExecuteTransaction - a batch of set/delete operations.
* Get.