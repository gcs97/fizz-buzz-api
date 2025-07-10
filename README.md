# fizz-buzz-api

A simple REST API that generates a custom Fizz Buzz sequence. 

It also includes a statistics endpoint to track the most frequent request.

## Usage

Copy the `env.example` into `.env`, updating the values if needed.

```shell
make run            # Run API directly  
make docker-build   # Build Docker image    
make docker-run     # Run Docker image
make test           # Run unit tests
```

## Endpoints

- `GET /api/v1/fizz-buzz`, requiring the query parameters `int1`, `int2`, `limit`, `str1` and `str2`
- `GET /api/v1/stats`