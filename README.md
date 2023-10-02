# Docker
```bash
docker run -d -p 8080:8080 shamimice03/movie-api
```

# golang: movies-api

This code is a simple implementation of a RESTful API for managing movie data. It provides endpoints for retrieving, creating, updating, and deleting movie records. Let's go through the code step by step:

1. The code starts with importing the necessary packages:
   - `encoding/json`: Used for encoding and decoding JSON data.
   - `fmt`: Used for formatted I/O.
   - `log`: Used for logging.
   - `math/rand`: Used for generating random numbers.
   - `net/http`: Used for building HTTP servers and clients.
   - `strconv`: Used for converting string to integer.
   - `github.com/gorilla/mux`: A popular third-party package for building HTTP routers.

2. Two struct types are defined: `Movie` and `Director`. The `Movie` struct represents a movie record with fields like `ID`, `Imdb`, `Title`, and `Director`. The `Director` struct represents the director of a movie with `Firstname` and `Lastname` fields.

3. A slice of `Movie` structs named `movies` is declared to store the movie data.

4. The code defines several functions that will be used as handlers for different HTTP endpoints:
   - `getMovies`: Retrieves all movies from the `movies` slice and sends them as a JSON response.
   - `deleteMovie`: Deletes a movie from the `movies` slice based on the provided movie ID in the URL path.
   - `getMovie`: Retrieves a single movie from the `movies` slice based on the provided movie ID in the URL path.
   - `createMovie`: Creates a new movie by decoding the JSON payload from the request body, assigning it a random ID, and appending it to the `movies` slice.
   - `updateMovie`: Updates an existing movie by replacing it with a new movie object based on the provided movie ID in the URL path.

5. In the `main` function, a new router is created using `mux.NewRouter()`.

6. Two movie records are added to the `movies` slice for demonstration purposes.

7. Different endpoints are registered with their corresponding handler functions using `r.HandleFunc()`. The registered endpoints are:
   - `GET /`: Retrieves all movies.
   - `GET /movies/{id}`: Retrieves a single movie by ID.
   - `POST /movies`: Creates a new movie.
   - `PUT /movies/{id}`: Updates an existing movie by ID.
   - `DELETE /movies/{id}`: Deletes a movie by ID.

8. The server is started by calling `http.ListenAndServe()` with the router and the desired port number (8080). If the server fails to start, it will log the error message.

9. When the server starts successfully, it prints a message indicating the port number it is listening on.

This code provides a basic API for managing movie data. By making HTTP requests to the defined endpoints, you can interact with the movie records stored in memory.
