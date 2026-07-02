# Napoleon Email

Backend microservice that receives contact-form submissions, persists them to Google Firestore, and forwards them as HTML emails via SMTP. Built with Go, Fiber, and a clean hexagonal architecture.

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/` | Health check |
| `POST` | `/api/v1/napoleon/contact` | Napoleon contact form |
| `POST` | `/api/v1/group-napoleon/contact` | Grupo Napoleon contact form |
| `POST` | `/api/v1/bank-gold/contact` | Bank Gold contact form |
| `POST` | `/api/v1/napoleon-mine/contact` | Napoleon Gold Mine contact form |

### Request Schema

**`/napoleon/contact` & `/napoleon-mine/contact`:**
```json
{
  "email": "user@example.com",
  "name": "John Doe",
  "subject": "Hello",
  "message": "I'd like to know more about..."
}
```

**`/group-napoleon/contact`:**
```json
{
  "email": "user@example.com",
  "name": "John Doe",
  "phone": "+1234567890",
  "message": "I'm interested in..."
}
```

**`/bank-gold/contact`:**
```json
{
  "email": "user@example.com",
  "name": "John Doe",
  "nickname": "JD",
  "phone": "+1234567890"
}
```

### Response

```json
// 200 OK
{ "message": "Napoleon email sent successfully" }

// 400 Bad Request
{ "error": "all fields are required" }
{ "error": "invalid email format" }

// 500 Internal Server Error
{ "error": "Error sending email" }
```

## Processing Flow

1. Request body is parsed and validated (required fields, email format, length limits).
2. An HTML email is rendered using the embedded `contacto.html` template.
3. A contact record is created in Firestore with `sending = true`.
4. The HTML email is sent via SMTP to the configured recipient.
5. The Firestore record is updated to `sending = false`.

## Environment Variables

Copy `.env.example` to `.env` and fill in the values:

| Variable | Description | Default |
|----------|-------------|---------|
| `APP_ENV` | `development` or `production` | `development` |
| `APP_VERSION` | Application version | `1.0.0` |
| `SERVER_PORT` | HTTP server port | `3000` |
| `GOOGLE_APPLICATION_CREDENTIALS` | Path to Firebase service account JSON file | *(required)* |
| `EMAIL_HOST` | SMTP host | *(required)* |
| `EMAIL_PORT` | SMTP port | *(required)* |
| `EMAIL_ADDRESS` | Sender email address | *(required)* |
| `EMAIL_PASSWORD` | SMTP password | *(required)* |
| `EMAIL_NAPOLEON_TO` | Recipient for Napoleon contacts | *(required)* |
| `EMAIL_GROUP_NAPOLEON_TO` | Recipient for Grupo Napoleon contacts | *(required)* |
| `EMAIL_NAPOLEON_MINE_TO` | Recipient for Napoleon Gold Mine contacts | *(required)* |
| `EMAIL_BANK_GOLD_TO` | Recipient for Bank Gold contacts | *(required)* |

## Firestore Collections

| Collection | Entity |
|------------|--------|
| `contacto` | Napoleon contacts |
| `contacto_group` | Grupo Napoleon contacts |
| `contacto_bank_gold` | Bank Gold contacts |
| `contacto_mine` | Napoleon Gold Mine contacts |

Each document stores: `clientEmail`, `clientName`, `createdAt`, `to`, `message` (`html`, `subject`, `text`), `sending`.

## Running

```bash
# Install dependencies
go mod tidy

# Start the server
go run cmd/main.go server
```

The server listens on `http://localhost:3000` (or `SERVER_PORT`).

## Tech Stack

| Layer | Technology |
|-------|-----------|
| HTTP | [Fiber v2](https://github.com/gofiber/fiber) |
| CLI | [Cobra](https://github.com/spf13/cobra) |
| Database | [Google Cloud Firestore](https://cloud.google.com/firestore) |
| Email | [gomail v2](https://github.com/go-gomail/gomail) |
| Logging | [Logrus](https://github.com/sirupsen/logrus) + [Lumberjack](https://github.com/natefinch/lumberjack) |
| Config | [godotenv](https://github.com/joho/godotenv) |
| Templates | `html/template` + `go:embed` (baked into binary) |

## Architecture

```
cmd/                         Entry points
├── main.go                  CLI bootstrapping
└── server/server.go         HTTP server + graceful shutdown

src/
├── config/                  App, Firebase & Logger configuration
├── routes/                  HTTP route definitions (health + API v1)
├── pkg/                     Shared utilities
│   ├── logger/              Structured logging wrappers
│   ├── parse/               Parsing helpers
│   └── send_email/          HTML template rendering + SMTP sending
└── app/
    ├── domain/              Pure interfaces & models (no deps)
    │   ├── model/           Firestore entities
    │   ├── application/     Use-case interfaces
    │   └── repository/      Repository interfaces
    ├── application/         Use-case implementations
    ├── http/                Fiber handlers + request DTOs
    ├── infrastructure/      Firestore repository implementations + DI Kernel
    └── resource/template/   Embedded HTML email template
```

## License

Private — All rights reserved.
