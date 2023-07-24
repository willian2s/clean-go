# Clean Go Microservice

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

This is a microservice project built using clean architecture principles.

## Table of Contents

- [Clean Go Microservice](#clean-go-microservice)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
  - [Usage](#usage)
  - [Development](#development)
  - [Testing](#testing)
  - [Contributing](#contributing)
  - [License](#license)

## Installation

1. Clone the repository:

```bash
  git clone https://github.com/willian2s/clean-go.git
```

2. Build the microservice:

```bash
  cd clean-go
  go build
```

## Usage

To start the microservice, run the following command:

```bash
  go run main.go
```

or

```bash
  ./main
```

By default, the microservice runs on port 3000. You can access it at `http://localhost:3000`.

## Development

To contribute to the development of this microservice, follow these steps:

1. Install the required development dependencies:

- [Install migrate CLI](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate), and

```bash
  go install github.com/cosmtrek/air@latest
```

2. Start the hot-reloading development server using Air:

```bash
  air
```

This will start the development server and automatically reload the microservice whenever changes are made to the source files.

3. Make your changes and write tests for them.

4. Apply database migrations using golang-migrate/migrate:

```bash
  migrate -source file://path/to/migrations -database "YOURDBCONNECTION_STRING" up
```

Replace `path/to/migrations` with the actual path to your migrations folder and `YOUR_DB_CONNECTION_STRING` with the actual connection string for your database. Adjust the command if you're using a different database engine or migration source.

## Testing

To run the tests, use the following command:

```bash
  go test ./...
```

To generate the coverage report, use the following command:

```bash
  chmod +x ./coverage.sh
  ./coverage.sh
```

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvement, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
