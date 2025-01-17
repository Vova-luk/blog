## Project Overview

Backend application designed for blog management. The project implements a clean architecture, utilizing repositories, services, and handlers, with PostgreSQL and Redis as its core databases.

## Features

+ User authentication, registration and authorization.
+ The project implements the CREATE, READ, UPDATE and DELETE (CRUD) of records.
+ Email verification with code
+ Password hashing

## Stack
<ins>Programming language</ins>: Golang

<ins>Database</ins>: PostgreSQL (storage users, posts and comments)\
          Redis (storage verification codes and session IDs)
          
<ins>ORM</ins>: gorm

<ins>Libraries</ins>: go-chi/chi (Request routing)
           bcrypt (password hashing)
           
Tools :

## License

This project is licensed under the MIT License. See the LICENSE file for more information.
