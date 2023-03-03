cmd
---

This package contains the top-level main function that runs the service. It follows the same structure suggested by the [go standard project layout](https://github.com/golang-standards/project-layout/tree/master/cmd). The service provides a single executable.

In addition to the `main` function, there is also a `serve` function, which runs the gRPC server and takes care of gracefully shutting it down. Since inserting the ports in the database happens in a streaming fashion, it's likely that the operations take a while to finish executing, and shutting them down in the middle of the execution might have undesirable effects on the database entries. At the first cancellation signal, the gRPC server will no longer respond to any RPC method calls, and wait for a certain amount of time for the ongoing operations to complete. When either the operations complete, the timeout finishes, or another cancellation signal is received, the gRPC server will stop (forcefully if needed).