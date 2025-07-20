# Domain Security Checker

This is a simple web application built with Go and HTML/JavaScript to check the email security records of any domain.

## Features

- Checks for the presence of **MX**, **SPF**, and **DMARC** DNS records for a given domain.
- Displays the actual SPF and DMARC record values if present.
- Simple web interface for user input and result display.

## How It Works

1. The Go server exposes two endpoints:
   - `/check?domain=example.com`: Returns a JSON response with MX, SPF, and DMARC info for the domain.
   - `/`: Serves the static HTML page (`index.html`).

2. The HTML page lets users enter a domain and view the results.

## Usage

1. **Run the Go server:**
   ```
   go run main.go
   ```

2. **Open the app:**
   - Visit `http://localhost:8080` in your browser.

3. **Check a domain:**
   - Enter a domain (e.g., `example.com`) and click "Check".
   - The result box will show the DNS record status.

## Files

- `main.go` — Go server code for DNS checks and API.
- `index.html` — Web interface for user input and displaying results.

## Requirements

- Go 1.16+ installed
- Internet access for DNS lookups

## License
