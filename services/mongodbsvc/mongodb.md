# Infrastructure ([Fork](https://github.com/cemye/hobbyfarm))

The operator, database and the service are deployed under the hobbyfarm namespace. The files are located
in `charts/hobbyfarm/templates/mongodb`.

## mongo-operator.yaml

The MongoDB ReplicaSet is deployed and managed by the MongoDB Kubernetes Commnuity Operator. It creates a stateful set,
a persistent volume claime, a service and a deployment

## mongo-secrets.yaml

Contains the username and password for the database

## mongo-rest.yaml

Contains the deployment for the RESTful service which interfaces with the MongoDB. Will be deprecated when moving to a
microservice infrastructure

## mongo-service.yaml

Exposes the RESTful service via port 30969.

# RESTful Service ([Branch](https://github.com/svalabs/gargantua/tree/mongodb))

The RESTful Service provides an api `api/items` with GET, POST, PUT, DELETE and PATCH.

PATCH request to `65.21.252.65:30969/api/items/someID` will update an existing entry with the ID 'someID' by merging the
data from the request's body into the existing data. Both the old and new data can have varying levels of nesting and
different data types. If there are matching keys, the new data will be appended into the lower levels of the hierarchy,
element by element.

I began converting the structure of the files in `services/mongodbsvc` to the new microservice architecture. This is *
*NOT** finished yet and it should not be functional yet.
The old files are renamed `old_main.go` and `old_Dockerfile`, but remain mostly unchanged. These should be funtional,
but will be phased out.
Like the other services the I already defined the API for Protobuf and ran the script to build it in `protos/mongodb`.

# Installation and usage

## Step 1

Connect to Server via SSH for example:

`ssh root@ssh.i3.learn.sva.dev`

## Step 2

Make sure that Helm is installed and install
the [MongoDB Kubernetes Community Operator](https://github.com/mongodb/mongodb-kubernetes-operator/blob/master/README.md)
with Helm:

`helm repo add mongodb https://mongodb.github.io/helm-charts`

`helm install community-operator mongodb/community-operator`

## Step 3

Deploy MongoDB and Service with Helm

`helm upgrade --install hf hobbyfarm --repo https://cemye.github.io/hobbyfarm --namespace hobbyfarm -f values.yml`

## Step 4

Test the service by sending a http-get-request to `65.21.252.65:30969/api/items`
Profit.
Will be deprecated once jrpc is used instead of REST.

# Last Remarks

If there are any questions left, feel free to contact me.


