# go-movies-backend
Backend for full-stack Golang React movies project

# How to run this project
Run go run ./cmd/api from within the root directory of the folder

# How the api works
This api runs on PostgreSQL for the database and is currently configured to the localhost setting on postgres. There are currently a few routes available for use.
The routes that are available are:  
1. localhost:4000/v1/movies.  
This route is for getting all of the movies that are in the Postgres database.
3. localhost:4000/v1/movie/:id
This route is for getting one movie inside of the Postgres database based on the id that is inputted when calling the api.  
For example, localhost:4000/v1/movie/1 will return the information for the movie with id 1, the same with 2, 3, etc.

# Packages used in this project
This project uses the httprouter from julien schmidt's httprouter. The way that it works in this project is by creating an application variable that gets passed as a pointer receiver to a receiver function in routes. This acts a bit like OOP but the difference is there is no inheritance and the upper layer doesn't have any information of the lower implementation. (Link: https://github.com/julienschmidt/httprouter)
