Snip Review Notes (2026-03-15)
==============================

These notes are focused on helping you level up quickly. They are structured for quick scanning and follow-up.

High-Impact Issues (Fix First)

- internal/api/handlers/url_handler.go:42 DB insert errors are ignored, yet the handler still returns 201 and a success row. This can show a successful short URL that was never saved (for example: unique constraint or DB outage).
- internal/api/handlers/url_handler.go:54 and internal/api/handlers/url_handler.go:121 Raw user input is written directly into HTML without escaping. That is an XSS risk because users can submit script or malformed HTML in URLs.
- cmd/snip/main.go:12 DATABASE_URL is printed at startup. That is a credential leak in logs.

Correctness / Behavior Issues

- internal/shortener/shortener.go:27 The base URL is hard-coded to <http://localhost:3000>. In production this will generate broken short links unless you are running on localhost.
- internal/api/handlers/url_handler.go:119 The get all handler returns only code, while the create handler returns a full URL. The table will show inconsistent data.
- internal/validator/validator.go:9 The TLD allowlist is very short. This rejects many valid domains (.dev, .ai, .co, country TLDs, etc.) and makes the validator feel randomly strict.
- internal/api/server.go:53 The server uses http.ListenAndServe with no timeouts. That can allow slow-client resource exhaustion.
- internal/api/handlers/url_handler.go:101 rows.Err() is never checked after scanning, so mid-iteration errors are silently ignored.
- internal/api/handlers/url_handler.go:28 Decoding into &payload is unusual for a pointer-to-struct. It works, but it is not idiomatic and can hide mistakes.
- internal/generator/generator.go:24 The generator prints each code. That is noisy and can leak data if logs are captured.
- web/templates/index.html:33 The hx-get URL is absolute (<http://localhost:3000/shorten-urls>), which breaks when deployed on another host.

What You Nailed

- Clear modularity for core logic: validator, generator, and shortener are separated and easy to reason about.
- Good use of crypto/rand for unpredictable short codes in internal/generator.
- Database connection setup with pgxpool and Ping is solid in internal/database.
- Migrations are clean and include indexes that make lookup fast (migrations/000001_create_urls_table.up.sql).
- HTMX integration is a nice, minimal approach that keeps the UI responsive without complex JS.
- Tests exist for both generator and validator, which is great for a first iteration.

Improvements You Can Make Soon (Practical)

- Escape HTML when writing rows, or render rows with html/template instead of raw fmt.Fprintf. This removes the XSS risk.
- Use a BASE_URL or SERVER_PUBLIC_URL env var and build short URLs from it.
- Add collision handling: if the insert fails on unique code, generate a new code and retry a few times.
- Improve URL validation: consider net/url plus publicsuffix, or a simpler has host plus scheme rule instead of a fixed TLD allowlist.
- Return explicit Content-Type: text/html for HTMX responses, and application/json if you switch to JSON.
- Replace http.StatusMovedPermanently with 302 or 307 to avoid caching issues when you might want to change redirects later.
- Limit request body size for /shorten to avoid abuse (http.MaxBytesReader).
- Add pagination or a limit for /shorten-urls to avoid huge table responses.
- Pre-parse templates or embed them instead of parsing on each request in HandleHomeRender.
- Update .example.env to include DATABASE_MIGRATION_URL so the migrator is documented.

Architecture / Project Structure Notes

- Handlers mix HTTP, DB queries, and HTML rendering in a single place. This makes testing hard and tightens coupling. Recommendation: introduce a UrlRepository interface (DB access) and a ShortenerService (business rules), then keep handlers thin.
- internal/shortener cannot detect collisions because it does not interact with storage. Recommendation: move generate plus check plus store into a service layer that depends on a repository interface.
- Configuration is scattered across cmd/snip/main.go and internal/api/server.go. Recommendation: create a small config struct in cmd/snip and pass dependencies to api.NewServer.
- pkg is empty. Either remove it or move reusable packages into it so it has a purpose.
- web assets are served from disk. For deployment, consider embed.FS so the binary is self-contained.
- migrations include expires_at and last_accessed, but these fields are unused. Recommendation: either implement expiration and last-access tracking or remove those columns until needed.
- Ensure .env and the snip binary are not committed to version control. They are in .gitignore, but verify they are not tracked.

Testing Gaps to Fill

- No tests for handlers or database interactions. Even a few integration tests would catch the insert failed but returned 201 issue.
- No tests for collisions, invalid payloads, or redirect behavior.

Professional-Level Polish Ideas

- Add structured logging (request ID, status, latency).
- Add graceful shutdown with context cancellation and a server timeout.
- Use a consistent error style (lowercase error strings in Go).
