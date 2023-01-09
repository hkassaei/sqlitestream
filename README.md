## Overview
This repository explores the use of sqlite and sqlitestream for resource-constrained deployments, such as edge, where a more lightweight data management solution, e.g. compared to deploying PostgreSQL or MySQL, might be desired.

## docker-compose
docker-compose directory deploys a small local pipeline with 3 containers: 
* car-orders app: a simple Go app that generates light traffic by periodically inserting car order records into the db
* sqlitestream: a lightweight agent that replicates all SQLite changes to an object store, such as minio or aws s3.
* minio: an s3 compatible object storage.

The idea is to be able to very quickly bring up the stack using one command, thanks to docker-compose.

## kubernetes
kubernetes directory contains the artifacts that deploy the car-orders app and litestream in a Kubernete cluster.
Litestream replicates all SQLite changes to an s3 bucket. The combination of lightweight components such as SQLite and litestream
with a lightweight Kubernetes cluster such as k3s is particularly interesting for resource-constrained edge deployments. 