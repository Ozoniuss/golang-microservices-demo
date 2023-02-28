# Protobuf

The `protobuf` folder contains the protobuf representation of the messages that can be used to communicate between the APIs. Protobuf is a flexible and efficient solution to encoding data into binary format. The gRPC protocol is built on top of HTTP and can use the protobuf data format to enable communication via RPC-style calls. The serialization happens very quickly with a small payload, which makes it ideal for low-latency communication between microservices.

The folder contains a typical protobuf structure, coming with a `ports` folder holding all the protobuf messages of the ports service. The go code will be generated inside this folder, and thus a go module has been provided in order to enable access to it inside other services. The folder can easily be extended to contain protobufs for multiple services, by simply adding a new folder for any new service.

Usually, I split the protos for every service into `model` and `api`. The former contains logical domain-specific protobuf messages defined for each entity. These messages are meant to define the protobuf representation of the entity, thus allowing to serialize the entity data and send it over the wire. The latter defines the RPC methods that can be called on the server, often involving the defined models. This may contain additional messages that are not domain specific and thus not part of `model`, but rather operation specific such as list filters.

Note that when microservices interact with each other, they are usually not aware of their internal resource representation. Instead, they use the resource's protobuf model to exchange data via the defined RPC procedures. The protobuf compiler generates code that can be imported in each service, where a conversion between the protobuf representation and the internal representation usually takes place.

google/rpc
----------

This is a special package with code imported form [googleapis](https://github.com/googleapis/googleapis/blob/master/google/rpc/status.proto) which in this case is used to communicate errors in rpc streams. Error handling is different with streams, because unlike unary calls where an error is returned as a response, an error that occurs in a stream terminates the stream connection. For this reason, specific errors must be sent as messages through the stream itself. The technique used in this case is to wrap a stream response in a `oneof` containing the usual message(s) and a `google.grpc.Status` message, representing the error.

Unfortunately, the `status.proto` cannot be imported directly, at least not in a way that I know of, so the `status.proto` file has to be included in the project, in the directory `./google/rpc` starting from the root to match the package name. The content of the file can be found at the provided link.
