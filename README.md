## Project Overview

A backend application designed for managing blogs. The project implements a clean architecture using repositories, services, and handlers, with PostgreSQL and Redis as the underlying databases. It is possible to create, edit, and delete posts and comments on them. The project implements an authorization and authentication system using sessions and cookies. When a user logs in, a unique session is created and stored in Redis for easy access and management. Each request is checked for session validity using middleware.

## Features

+ User authentication, registration and authorization.
+ The project implements the CREATE, READ, UPDATE and DELETE (CRUD) of records.
+ Email verification with code
+ Password hashing

## Stack
<ins>Programming language</ins>: Golang

<ins>Database</ins>:\
PostgreSQL (storage users, posts and comments)\
Redis (storage verification codes and session IDs)
          
<ins>ORM</ins>: \
gorm

<ins>Libraries</ins>: \
go-chi/chi (Request routing)\
bcrypt (password hashing)\
gomail (sending email)

## Setup instructions

1. Clone the repository: `https://github.com/Vova-luk/blog.git`
2. Set up PostgreSQL and Redis.
3. Navigate to the project folder: `cd blog/cmd`
4. Run the server: `go run main.go`

## License

This project is licensed under the MIT License. See the LICENSE file for more information.
