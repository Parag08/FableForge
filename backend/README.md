# FableForge backend

Go HTTP service for FableForge. See the [repository README](../README.md) for
the project overview and roadmap. API contract: [`swagger.yml`](swagger.yml).

## Endpoints

- `GET /healthz` — liveness check, returns `{"status":"ok"}`.
- `POST /login` — JSON `{"username","password"}`; checks the demo account from
  env and sets an HttpOnly `token` cookie (HS256 JWT, lifetime `TOKEN_TTL`).

## Configuration

All settings come from environment variables; there are no in-code defaults.
The server exits at startup listing every missing or invalid key
(`internal/config`). See [`.env.example`](.env.example).

| Key | Meaning |
| --- | --- |
| `PORT` | Listen port |
| `JWT_SECRET` | HMAC secret for JWTs, at least 32 characters |
| `TOKEN_TTL` | Token lifetime as a Go duration (e.g. `15m`) |
| `ALLOWED_ORIGIN` | Frontend origin allowed by CORS (e.g. `http://localhost:3000`) |
| `DEMO_USERNAME` / `DEMO_PASSWORD` | The single demo login (placeholder until real user storage exists) |

## Run and test

```sh
cp .env.example .env              # then fill in JWT_SECRET and DEMO_PASSWORD
set -a; . ./.env; set +a          # export the variables into the shell
go run .

go vet ./... && go test ./...
```
