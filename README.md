# URL Shortener (Go)

Same project, different languages. Learning by rebuilding the same thing multiple times in different programming languages.

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

Shorten a URL:
```bash
curl -X POST -d '{"url":"https://google.com"}' http://localhost:4000/shorten
```

Response example: `{"short_code":"x7Km9p"}`

Visit the short URL:
```
http://localhost:4000/x7Km9p
```

Redirects to the original URL.

## What I learned

* Building HTTP servers with Go's standard library
* JSON encoding and decoding with the `encoding/json` package
* Using maps for in-memory storage
* HTTP redirects and status codes
* Go's explicit error handling approach
* Structuring Go projects with proper package organization
* Working with Go's type system and structs

## Other implementations

- [JavaScript/Express](https://github.com/robiverdev/urlshortenerjs)
- [Rust](link) *(coming soon)*
- [C](link) *(coming soon)*
