# Go, Chi, and PostgreSQL Quests API (refactor from GORM)
<p align="center">
<img src="https://raw.githubusercontent.com/egonelbre/gophers/63b1f5a9f334f9e23735c6e09ac003479ffe5df5/vector/fairy-tale/knight.svg" alt="Knight Gopher" width="300">
</p>


A quest-tracking API built with Go, Chi, and PostgreSQL. This project focuses on clean architecture by decoupling components, using dependency injection, and applying the repository pattern to enhance maintainability and scalability. It supports basic CRUD operations for managing quest data via an HTTP server and uses Docker Compose for service orchestration.

Commit history includes multiple changes to the code. 


## Features
TBD

## Getting Started

### Prerequisites
- Docker
- Docker Compose
- Go (1.22+ recommended)

## Installation

1. Clone this repository:
   ```sh
   git clone https://github.com/travboz/gorm-to-postgres-refactor.git
   cd gorm-to-postgres-refactor
   ```
2. Run docker container:
    ```sh
    make compose-up
   ```
3. Run server:
    ```sh
    make run
    ```
4. Navigate to `http://localhost:8000` and call an endpoint

### `.env` file
This server uses a `.env` file for basic configuration.
See `.env.example` for an example `.env` file.

## API endpoints

| Method   | Endpoint          | Description                         |
|----------|-------------------|-------------------------------------|
| `GET`    | `/quests`         | Retrieve all quests in the game     |
| `GET`    | `/quests/:id`     | Fetch a quest by its ID             |
| `POST`   | `/quests`         | Create a new quest in the game      |
| `PUT`    | `/quests/:id`     | Update a quest with the specified ID|
| `DELETE` | `/quests/:id`     | Delete a quest with the specified ID|



## Example usage

TBD

## Contributing
Feel free to fork and submit PRs!

## License:
`MIT`


This should work for GitHub! Let me know if I can make any tweaks. 


## Image
Image by [Egon Elbre](https://github.com/egonelbre), used under CC0-1.0 license.