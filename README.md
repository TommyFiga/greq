# greq
> A minimal cURL-inspired HTTP client written in Go.

## Features
- Send HTTP requests with custom methods
- Add custom headers
- Send request bodies
- Optionally include response headers
- Pretty-print JSON
- Write responses directly to a file

## Installation
Make sure that you have Go installed (version 1.20+ recommended). 

Clone the repo and build the project
``` bash
git clone https://github.com/TommyFiga/greq.git
cd greq
go build -o greq ./main.go
```

Now you can run greq locally `./greq`, or you can move it into your `$PATH`
``` bash
mv greq /usr/local/bin/
```

## Usage
Basic request:
``` bash
greq https://httpbin.org/get
```

Specific HTTP method request:
``` bash
greq -X POST https://httpbin.org/post -d '{"name":"greq"}'
```

Custom headers:
``` bash
greq -H "Content-Type: application/json" -H "Authorization: Bearer TOKEN" https://httpbin.org/headers
```

Pretty-print JSON:
``` bash
greq --json https://httpbin.org/json
```

## Options
| Flag        | Description                               | Default |
| ----------- | ----------------------------------------- | ------- |
| `-X`        | HTTP method (GET, POST, PUT, DELETE, …)   | `GET`   |
| `-H`        | Custom header (`Key: Value`) (repeatable) | none    |
| `-d`        | Request body (string)                     | none    |
| `-i`        | Include response headers in output        | false   |
| `-json`     | Pretty-print JSON responses               | false   |
| `-o`        | Write output to file instead of stdout    | none    |

## Project Structure
```
.
├── cmd/
|   └── greq/         # Entry point
└── internal/
    ├── httpclient/   # HTTP client logic
    ├── output/       # File/stdout output handling
    ├── parser/       # Command-line flag parsing
    ├── printer/      # Formatting and pretty-printing
    └── types/        # Shared data structures
```