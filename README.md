# Bookstore API

## Overview

The Bookstore API is a backend service developed in Go (Golang) for managing a virtual bookstore. It provides RESTful endpoints for performing CRUD (Create, Read, Update, Delete) operations on a collection of books. This project is designed to showcase backend development skills and the use of modern technologies like the Gin framework and GORM ORM.

## Features

- Add new books to the catalog.
- Retrieve all books with options for filtering and sorting.
- Retrieve details of a specific book by its ID.
- Update or delete book entries.
- Pagination support for large datasets.

## Technologies Used

| Technology       | Purpose                                |
|------------------|---------------------------------------|
| **Go (Golang)**  | Programming language                  |
| **Gin**          | Web framework for building the API    |
| **GORM**         | ORM for interacting with the database |
| **SQLite**       | Lightweight database                  |
| **Postman**      | Testing and interacting with the API  |

## Prerequisites

- Go installed on your machine (version 1.18 or later).
- SQLite installed (optional, as the database file will be generated automatically).
- A tool like Postman or curl for testing the API.

## Getting Started

# Endpoints
# Base URL
- http://localhost:8081

# List of Endpoints

| Method       | Endpoint       | Description                     |
|--------------|----------------|---------------------------------|
| **GET**	     |  /books	      | Retrieve all books              |
| **GET**	     |  /books/:id    |	Retrieve a specific book by ID  |
| **POST**	   |  /books        | Add a new book                  |
| **DELETE**	 |  /books/:id    |	Delete a specific book          |


## Clone the Repository

```bash
git clone https://github.com/mliooh/bookstore-api
cd bookstore-api
go mod tidy
go run main.go


