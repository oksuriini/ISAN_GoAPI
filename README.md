# ISAN API Golang

ISAN API written with golang.

## Database structure

Below simple structure for SQL database.

| Field             | Type         | Null | Key | Default | Extra |
| ----------- | -----------|-----------|-----------|-----------|-----------|
| record_type       | text         | YES  |     | NULL    |       |
| title             | text         | YES  |     | NULL    |       |
| work_type         | text         | YES  |     | NULL    |       |
| release_date      | date         | YES  |     | NULL    |       |
| director          | text         | YES  |     | NULL    |       |
| duration_min      | int          | YES  |     | NULL    |       |
| original_language | text         | YES  |     | NULL    |       |
| isan              | varchar(255) | NO   | PRI | NULL    |       |


## Handlers

TODO

## Golang structs and other stuff

TODO

### ISAN

### Queries