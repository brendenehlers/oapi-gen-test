#!/bin/bash

echo "POST"
curl localhost:8080/pets -X POST -d '{"name": "Fido"}'

echo "GET ALL"
curl localhost:8080/pets

echo "GET 1"
curl localhost:8080/pets/1

echo "DELETE"
curl localhost:8080/pets/1 -X DELETE