# Movies Server

This Go application serves as a basic RESTful API for managing movie data. It allows you to perform CRUD (Create, Read, Update, Delete) operations on movie records using HTTP requests.

## Features

- List all movies
- Get a specific movie by ID
- Create a new movie record
- Update an existing movie record
- Delete a movie record

## Prerequisites

To run this server, you need to have:

- [Go](https://golang.org/doc/install) installed on your system.
- Docker installed, if you plan to use the `docker-compose.yml` file to run MariaDB.

## Installation

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/Micah-Shallom/Go-Movie-APP.git
   ```
2. Navigate to the project directory:

   ```bash
   cd Go-Movie-APP
   ```
3. Copy the .env.example to .env and modify the settings:

   ```
   cp .env.example .env
   ```
4. Start MariaDB using Docker Compose(or have a mysql database running):

   ```
   docker-compose up -d
   ```
4. Run the server:

   ```bash
   go run cmd/main/main.go
   ```

   The server will start on port 8000 by default.
5. Access the server on Browser:

```http
http://localhost:8000
```

## Usage

- Use an API client (e.g., [Postman](https://www.postman.com/)) or make HTTP requests to interact with the server's endpoints.
- You can perform various operations such as listing all movies, retrieving a specific movie by ID, creating a new movie record, updating an existing movie record, and deleting a movie record.

## Configuration

The application uses a `.env` file for configuration. The `.env.example` file serves as a template. The following variables are configurable:

- `dbusername`: Database username
- `dbuserpassword`: Database user password
- `dbname`: Database name
- `dbrootpassword`: Database root password
- `DB_HOST`: Database host (default is `localhost`)
- `DB_PORT`: Database port (default is 9000)
- `TIME_ZONE`: Timezone for the database
- `DB_CONFIG_DIRECTORY`: Path to database config directory

Make sure to update these values in the `.env` file before running the application.

## Endpoints

### List all movies

- **GET** `/movies`
- **Mock Up FrontEnd**
- ![get](https://github.com/Micah-Shallom/Go-Movie-APP/assets/64049432/33c6143a-cfdb-4583-883a-3b4f5b4051d9)
### Get a specific movie by ID

- **GET** `/movies/{id}`

### Create a new movie record

- **POST** `/movies`
- **Mock Up FrontEnd**
- ![post](https://github.com/Micah-Shallom/Go-Movie-APP/assets/64049432/05b60567-2588-4fc5-9370-13591e177bb3)


### Update an existing movie record

- **PUT** `/movies/{id}`

### Delete a movie record

- **DELETE** `/movies/{id}`

## JSON Data Format

The server accepts and returns JSON data in the following format:

```json
{
  "id": "1",
  "isbn": "343434",
  "title": "Movie One",
  "director": {
    "firstname": "John",
    "lastname": "Doewy"
  }
}
```

## Contributing

Feel free to contribute to this project by creating issues or submitting pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
