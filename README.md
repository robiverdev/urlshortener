# URL Shortener

A simple URL shortener built with Go. It takes long URLs and generates random 6 character short codes.

## Project structure

urlshortener/
│
├── main.go # Application entry point, server setup, and route registration
├── handlers.go # HTTP handlers for homepage, shortening, and redirecting
├── shortener.go # Logic for generating random short codes
└── store.go # In memory map for storing URL mappings

## How it works

1. POST a long URL to `/shorten`
2. Get back a short code
3. Visit `/{shortCode}` to redirect to the original URL

## Running locally

```bash
git clone https://github.com/robiverdev/urlshortener
cd urlshortener
go run .
```

Server starts on `http://localhost:4000`

## Usage

**Shorten a URL:**

```bash
curl -X POST -d '{"url":"https://google.com"}' http://localhost:4000/shorten
```

Response example: `{"short_code":"x7Km9p"}`

**Visit the short URL:**

```
http://localhost:4000/x7Km9p
```

Redirects to the original URL.

## What I learned

- Building HTTP servers with Go's stdlib
- JSON encoding/decoding
- Maps for in memory storage
- HTTP redirects and status codes
