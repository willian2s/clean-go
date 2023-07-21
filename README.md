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

```bash
  go install github.com/cosmtrek/air@latest
```

2. Start the hot-reloading development server using Air:

```bash
  air
```

## Testing

To run the tests, use the following command:

```bash
  go test ./...
```

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvement, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
