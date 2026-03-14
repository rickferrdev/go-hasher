# Bcrypt Hasher Wrapper (Go-Hasher)

This project is a simple wrapper for the `golang.org/x/crypto/bcrypt` library. It was developed exclusively for personal study purposes, focusing on abstracting hashing logic and implementing interfaces in Go.

## Objective

The package provides a standardized interface for generating and comparing hashes, making it easier to replace implementations or mock behaviors in unit tests.

## Features

* Generate bcrypt hashes from byte slices.
* Validate passwords against existing hashes.
* Custom error handling.
* Automatic validation of the algorithm's cost.

## Interface

The package exposes the following interface:

```go
type Hasher interface {
    Generate(password []byte) ([]byte, error)
    Compare(hash, password []byte) error
}

```

## Usage Example

### Initialization

To create a new instance with the default cost:

```go
h := hasher.New(hasher.DefaultCost)

```

### Generating a Hash

```go
password := []byte("my_secure_password")
hash, err := h.Generate(password)
if err != nil {
    // handle error
}

```

### Comparing a Password

```go
err := h.Compare(hash, []byte("my_secure_password"))
if err != nil {
    // if the error is "hasher: invalid password", the password is incorrect
}

```

## Cost Constants

The wrapper exposes bcrypt cost constants for easier configuration:

* `MinCost`: 4
* `MaxCost`: 31
* `DefaultCost`: 10

## Limitations

As per the Go bcrypt implementation, the maximum password length is 72 bytes. If a larger password is sent to the `Generate` method, the `bcrypt.ErrPasswordTooLong` error will be returned.

## Contributing

Even though this is a personal study project, contributions are highly appreciated! Feel free to submit Pull Requests at any time to improve features, refactor code, or add new capabilities.
