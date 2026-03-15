# Feature Ideas for Snip

This list is a menu of possible features you can add as you grow the project. Each item is written to be concrete and actionable.

 **To update status**  
 `[ ]` = Not started  
 `[x]` = Done  
 `[~]` = In progress

## Product and UX

`[ ]` Custom aliases (let users choose their own short code)

`[ ]` Branded domains (support multiple domains like `go.yoursite.com`)

`[ ]` Bulk shorten endpoint (upload a list of URLs and get results back)

`[ ]` QR code generation for each short link

`[ ]` Edit or deactivate a short link from the UI

`[ ]` Click-to-copy buttons in the UI

## Analytics

`[ ]` Per-link analytics dashboard

`[ ]` Referrer tracking (where clicks come from)

`[ ]` Device and browser breakdown

`[ ]` Time-series charts for clicks

`[ ]` Export analytics as CSV

## Reliability and Correctness

`[ ]` Collision handling with automatic retries

`[ ]` Idempotent create endpoint (same URL returns same code if desired)

`[ ]` Background cleanup for expired links

`[ ]` Graceful shutdown and request timeouts

## Security and Abuse Prevention

`[ ]` Rate limiting for create and redirect endpoints

`[ ]` Blocklist and allowlist for domains

`[ ]` Basic URL safety checks (phishing and malware flags)

`[ ]` CAPTCHA or bot protection on the create form

`[ ]` CSRF protection if you add sessions

## API and Platform

`[ ]` Public JSON API with API keys

`[ ]` Webhooks for click events

`[ ]` OpenAPI documentation and a client SDK

`[ ]` Feature flags for experimental changes

## Operations

`[ ]` Structured logging and request IDs

`[ ]` Metrics and tracing (Prometheus and OpenTelemetry)

`[ ]` Health and readiness endpoints

`[ ]` Dockerfile and docker-compose for local dev

`[ ]` CI pipeline for linting, tests, and migration checks
