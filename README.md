# Postgres_CRUD

This is a CRUD API in go with jwt authentication, it uses Postgres and GORM. 

# Setting up locally

To run this locally git clone the repository, add a .env file

Here is an example .env file:
```
# Database credentials
DB_HOST="<<DB_HOST>>"
DB_USER="<<DB_USER>>"
DB_PASSWORD="<<DB_PASSWORD>>"
DB_NAME="diary_app"
DB_PORT="<<DB_PORT>>"

# Authentication credentials
TOKEN_TTL="2000"
JWT_PRIVATE_KEY="SECRET_HERE"
```

Generally its a good idea to not push any .env or .env.* files(these should be added to the gitignore), but for the sake of 
the demo I have added a .env.example file for reference

Then do:
```
go run main.go
```

This will run our server on localhost port 3000

Note: change the .env file name in the loadEnv() function in main.go as needed.

# The routes

There are 3 auth routes in total:
- /auth/register
- /auth/login
- /auth/logout

All of these use the POST method

There are 4 protected routes for CRUD operations:
- /api/post (POST method)
- /api/post (GET method)
- /api/post (PUT method)
- /api/post (DELETE method)

The database file has a function that we use to connect to our database

The middleware file contains a handler to Validate JWT tokens.

The models file contains all the models we use throughout the code, it has 3 models:
- AuthenticationInput model
- User model
- Post model

AuthenticationInput model is used for auth

There are also a few helper functions that help us in protecting routes in the helper file

The controllers file contains most of the handler functions we use to handle various routes

Note: Currently the backend sets a jwt as a cookie
I think while connecting with frontend it should be alright since 
The frontend should be handling sending the jwt automatically within the header
The cookie as of now is set and logout just invalidates the cookie.

This was the simplest way to implement logout, something I had done at the very end when all
the other routes were handled. JWT tokens are stateless and setting it as cookie makes sense 
to be able to handle logout. Time crunches didn't allow for extensive testing with postman, most
of the testing was done using curl from the terminal

# Version details
- OS: 6.2.13-arch1-1 (arco linux)
- Go: go1.20.3 linux/amd64
- PostgreSQL: 15.2

