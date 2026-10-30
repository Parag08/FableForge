# FableForge

FableForge is intended to illustrate novels. The planned pipeline:

1. **Chapter in** — the author provides a chapter of text.
2. **Scenes + character bible** — the chapter is broken into illustratable
   scenes, and recurring characters (appearance, clothing, distinguishing
   features) and settings are captured once in a shared character bible.
3. **Consistent illustrations** — each scene is illustrated using the
   character bible so characters look the same from image to image.
4. **EPUB out** — the text and illustrations are assembled into an
   illustrated EPUB.

## Status

**Scaffold only.** None of the pipeline above is implemented yet. What exists today:

- `backend/` — Go HTTP server with `GET /healthz` and a demo `POST /login`
  that issues a JWT cookie for a single env-configured account. Config is
  env-only and validated at startup. Tests cover login and health.
- `frontend/` — React app with a Google sign-in landing page and a
  placeholder home page describing the workflow. It does not call the
  backend yet.

## Repository layout

```
backend/
  main.go                 entry point: load config, start server
  server.go               routes, CORS, handlers
  server_test.go          httptest tests for /healthz and /login
  internal/config/        env-based config, fails fast on missing keys
  swagger.yml             OpenAPI description of the real endpoints
  .env.example
frontend/
  src/App.jsx             routes: / (landing), /home (placeholder)
  src/pages/              Landing.jsx, Home.jsx
  .env.example
```

## Running locally

Prerequisites: Go 1.22+, Node.js 18+.

### Backend

```sh
cd backend
cp .env.example .env              # fill in JWT_SECRET (>= 32 chars) and DEMO_PASSWORD
set -a; . ./.env; set +a
go run .                          # listens on $PORT
go test ./...
```

### Frontend

```sh
cd frontend
cp .env.example .env              # set REACT_APP_GOOGLE_CLIENT_ID
npm install
npm start                         # http://localhost:3000
```

## Environment variables

Backend (all required, no defaults — the server refuses to start if any is missing):

| Key | Meaning |
| --- | --- |
| `PORT` | Listen port |
| `JWT_SECRET` | HMAC secret for JWTs, at least 32 characters |
| `TOKEN_TTL` | Token lifetime, Go duration (e.g. `15m`) |
| `ALLOWED_ORIGIN` | Frontend origin allowed by CORS |
| `DEMO_USERNAME`, `DEMO_PASSWORD` | Single demo login, placeholder for real accounts |

Frontend:

| Key | Meaning |
| --- | --- |
| `REACT_APP_GOOGLE_CLIENT_ID` | Google OAuth client ID for the landing page |

## Roadmap

1. Real authentication: verify the Google ID token on the backend and replace
   the demo account; add user/project storage.
2. Chapter ingestion API and upload UI.
3. Scene segmentation and character-bible extraction from chapter text.
4. Illustration generation conditioned on the character bible, with a review
   and regenerate loop for consistency.
5. EPUB assembly and download.

## License

MIT — see [LICENSE](LICENSE).
