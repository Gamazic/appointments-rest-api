# appointments-rest-api

Simple rest-api project.

Consists of:
* Business logic: service for managing doctor appointments in the clinic.
* DB: simple in-memory key-value database. Written on Golang.
* web server: `gin` has been chosen as web framework.

____

Possible requests:
* **`GET`** `/appointments/` - get all appointments
* **`GET`** `/appointments/<id>` - get an appointment with the `<id>` key
* **`DELETE`** `/appointments/<id>` - delete an appointment with the `<id>` key
* **`POST`** `/appointments/` - creates a record and returns an id of the record
    BODY:
    ```json
    {
      "name": String,
      "disease": String
    }
    ```