#!/bin/bash

curl localhost:8080/pets -X POST -d '{"name": "Fido"}'

# curl localhost:8080/pets
# curl localhost:8080/pets/123 -X DELETE
# curl localhost:8080/pets/1