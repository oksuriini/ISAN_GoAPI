# ISAN API Golang

ISAN API written with golang.
The purpose is to create simple api, which will fetch ISAN (International Standard Audiovisual Number) information of an audiovisual asset.

The api uses simple api path:
hostname:8080/api/v1/<path>/<to>/<query>

Currently it has two paths implimented:

`/test` Is used to test the connectivity of the api itself.

`/api/v1/all` Is used to fetch all records from mysql database.

## Future features (priority in order)

- Impliment rest of the api queries and their paths
- Simple frontend which can be used to make queries
- Allow inserting into database
- Improve frontend's appearance

## Possible features

- API-key option and registration system
- Whatever???

## Database structure

Below simple structure for SQL database.

| Field             | Type         | Null | Key | Default | Extra |
| ---               | ---          | ---  | --- | ---     | ---   |
| record_type       | text         | NO   |     | NULL    |       |
| title             | text         | NO   |     | NULL    |       |
| work_type         | text         | NO   |     | NULL    |       |
| release_date      | date         | NO   |     | NULL    |       |
| director          | text         | NO   |     | NULL    |       |
| duration_min      | int          | NO   |     | NULL    |       |
| original_language | text         | YES  |     | NULL    |       |
| isan              | varchar(255) | NO   | PRI | NULL    |       |


## Handlers

TODO

## Golang structs and other stuff

TODO

### ISAN

TODO

### Queries

TODO

### Docker support

Currently the project includes two dockerfiles. Notice that these dockerfiles are by no means well optimized.

'Dockerfile.db', which can be used to build an image for mysql database.
'Dockerfile.gofile' which can be used to build an image for the api application.

Both of these files can also be found from my Dockerhub:

Mysql image:
hub.docker.com/repository/docker/oksuriini/goapi-mysqldb/

GolangApp:
hub.docker.com/repository/docker/oksuriini/goapi-app/

The project also includes docker compose file, which can be used to make the whole package work out of the box.
Notice that the compose file is also not well optimized and might be security risk currently as it uses ```network_mode: "host"``` option.