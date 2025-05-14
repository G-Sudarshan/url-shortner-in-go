# URL Shortener in Go

A simple and efficient URL shortener service built with Go. This project demonstrates how to create a lightweight web service for shortening URLs and redirecting users to the original links.

## Table of Contents
- [Features](#features)
- [Technologies](#technologies)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Contributing](#contributing)
- [License](#license)

---

## Features

- Shorten long URLs into concise, shareable links.
- Redirect users to the original URL using the short link.
- Efficient and fast, built with Go.
- Dockerized for easy deployment and scalability.

---

## Technologies

This project is primarily built using:
- **Go** (94.5%)
- **Docker** (5.5%)

---

## Installation

### Prerequisites
- [Go](https://golang.org/doc/install) 
- [Docker](https://www.docker.com/get-started) (for containerized deployment)

### Setup Instructions

1. Clone the repository:
   ```bash
   git clone https://github.com/G-Sudarshan/url-shortner-in-go.git
   cd url-shortner-in-go
   ```

2. Build the project:
   ```bash
   go build
   ```

3. Run the project:
   ```bash
   go run main.go
   ```

4. Alternatively, use Docker to build and run:
   ```bash
   docker build -t url-shortner .
   docker run -p 8080:8080 url-shortner
   ```

---

## Usage

1. Start the application locally or using Docker.
2. Access the service at `http://localhost:8080`.
3. Use the API endpoints to shorten URLs and redirect users.

---

## API Endpoints

### `POST /shorten`
Shortens a URL.
- **Request Body:** JSON object with `url` field.
  ```json
  {
    "url": "https://example.com"
  }
  ```
- **Response:**
  ```json
  {
    "shortUrl": "http://localhost:8080/abcd"
  }
  ```

### `GET /{shortUrl}`
Redirects to the original URL.
- **Example:** `GET http://localhost:8080/abcd`

