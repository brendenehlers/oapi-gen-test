# Open API Code Generation Test

Built this as a quick test for using [OAPI Codegen](https://github.com/deepmap/oapi-codegen) tool and building out a scalable API from the generated code. I used the [Chi](https://github.com/go-chi/chi) library to handle routing. Everything is tied together in the main function, and the server is hosted on port `8080`.

The generated code lives in the `generated` package.

The code for handling requests lives in the `handler` package.

In the `handler` package, the `handler.go` file contains the implementation of the `generated.ServerInterface`, while the `data.go` file contains the logic to interact with the data source (in this example, it's a simple slice on the impl struct)

Let me know if you have any questions!