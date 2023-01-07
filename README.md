## Stream SQLite to Minio
This repo contains all the artifacts needed to create a small pipeline containing:
* A dockerized Go application (order-app) that generates random car order records and inserts them to an SQLite database.
* A dockerized [Litestream server](https://litestream.io) that streams all SQLite changes to an S3 compatible storage.
* A dockerized [Minio](https://min.io/) server that acts as the object storage hosting the SQLite replicas/backups sent by Litestream.
## How to use:
```
docker-compose up --build
```
### Note: Configure Litsteam
When running this stack, make sure you provide the IP address of the minio server to Litestream. 
To do this, note the IP address assigned to the minio docker container and provide it to Litstream
by updating `litestream-conf/litestream.yml`

## References:
* https://litestream.io
* https://github.com/benbjohnson/litestream
* https://www.sqlite.org/datatype3.html
* https://7thzero.com/blog/golang-w-sqlite3-docker-scratch-image
