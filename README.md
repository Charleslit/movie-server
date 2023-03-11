# movie-server
This API allows you to manage movies through CRUD (Create, Read, Update, Delete) operations. The API endpoints are as follows:
Get All Movies

GET /movies

Retrieves all movies from the server.

Response

    200 OK on success

json

[
    {
        "id": "1",
        "isbn": "456756",
        "title": "movie one",
        "Director": {
            "firstname": "john",
            "lastname": "ogwoka "
        }
    },
    {
        "id": "2",
        "isbn": "4567898",
        "title": "movietwoo",
        "Director": {
            "firstname": "charles",
            "lastname": "duoe"
        }
    }
]

Get a Movie by ID

GET /movies/{id}

Retrieves a specific movie from the server based on its ID.

Parameters

    id - The ID of the movie to retrieve.

Response

    200 OK on success
    404 Not Found if the movie with the specified ID does not exist

json

{
    "id": "1",
    "isbn": "456756",
    "title": "movie one",
    "Director": {
        "firstname": "john",
        "lastname": "ogwoka "
    }
}

Create a Movie

POST /movies

Adds a new movie to the server.

Request Body

json

{
    "isbn": "978-0-596-52068-7",
    "title": "Building Web Applications with Flask",
    "Director": {
        "firstname": "Ore",
        "lastname": "Ogundipe"
    }
}

Response

    201 Created on success

json

{
    "id": "123456",
    "isbn": "978-0-596-52068-7",
    "title": "Building Web Applications with Flask",
    "Director": {
        "firstname": "Ore",
        "lastname": "Ogundipe"
    }
}

Update a Movie

PUT /movies/{id}

Updates an existing movie in the server based on its ID.

Parameters

    id - The ID of the movie to update.

Request Body

json

{
    "isbn": "978-1-449-39091-9",
    "title": "Building RESTful Web Services with Go",
    "Director": {
        "firstname": "Nic",
        "lastname": "Jackson"
    }
}

Response

    200 OK on success
    404 Not Found if the movie with the specified ID does not exist

json

{
    "id": "123456",
    "isbn": "978-1-449-39091-9",
    "title": "Building RESTful Web Services with Go",
    "Director": {
        "firstname": "Nic",
        "lastname": "Jackson"
    }
}

Delete a Movie

DELETE /movies/{id}

Removes an existing movie from the server based on its ID.

Parameters

    id - The ID of the movie to remove.

Response

    200 OK on success
    404 Not Found if the movie with the specified ID does not exist

Implementation Details

The API is implemented in Go language using the Gorilla mux library for HTTP routing and handling. The API supports JSON format for both request and response bodies.

The movie struct defines the properties of a movie,
