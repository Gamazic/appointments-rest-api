# appointments-rest-api

Simple rest-api project.

Consists of:
* Business logic: service for managing doctor appointments in the clinic.
* DB: simple in-memory key-value database. Written on Golang.
* web server: `gin` has chosen as web framework.

____

Possible requests:
* **`GET`** `/appointments/` - get all appointment
* **`GET`** `/appointments/<id>` - get <id> appointment
* **`DELETE`** `/appointments/<id>` - delete <id> appointment
* **`POST`** `/appointments/` - creates a record and return id of record