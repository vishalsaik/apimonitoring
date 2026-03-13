# Connector Theory — api-monitoring

## What is a Connector?

A connector opens and verifies a connection to an external service (DB, queue),
then returns a usable client object. All connection logic lives here — nothing
else in the codebase touches connection setup.

## The Universal Pattern

```
config (address + credentials)
    → connector (open + verify)
        → client/pool/connection object
            → repository (runs actual queries)
```

---

## MongoDB

**Package:** `go.mongodb.org/mongo-driver/mongo`

**What you need:**
- `*mongo.Client` — manages the connection pool internally
- `*mongo.Database` — scoped to one DB, what repositories use

**How:**
```
options.Client().ApplyURI(url)   → clientOptions
mongo.Connect(ctx, clientOptions) → client
client.Ping(ctx, nil)            → verify (Dial does NOT auto-verify)
client.Database(name)            → database
```

**Why Ping:** `mongo.Connect()` is lazy — it does NOT open a real connection.
Ping forces the actual connection attempt.

**Returns:** `*MongoDB{ Client, Database }`

---

## PostgreSQL

**Package:** `github.com/jackc/pgx/v5/pgxpool`

**What you need:**
- `*pgxpool.Pool` — manages N reusable connections (not just one)

**How:**
```
pgxpool.ParseConfig(connString)       → poolConfig
poolConfig.MaxConns = 10              → tune BEFORE creating pool
pgxpool.NewWithConfig(ctx, poolConfig) → pool
pool.Ping(ctx)                        → verify
```

**Connection string format:**
```
postgres://user:password@host:port/dbname
```

**Why pool:** Without pooling, every goroutine opens its own connection.
A busy service would exhaust Postgres's connection limit (default 100).

**Why ParseConfig + NewWithConfig (not just New):**
`pgxpool.New()` creates pool immediately with no way to set MaxConns first.
ParseConfig → modify → NewWithConfig is the correct order.

**Returns:** `*Postgres{ Pool }`

---

## RabbitMQ

**Package:** `github.com/rabbitmq/amqp091-go` (imported as `amqp`)

**What you need:**
- `*amqp.Connection` — one TCP socket to RabbitMQ (keep for lifecycle/shutdown)
- `*amqp.Channel` — virtual connection inside the TCP socket (publish/consume on this)

**How:**
```
amqp.Dial(url)    → connection (self-verifying, fails fast if unreachable)
conn.Channel()    → channel
```

**Why no Ping:** `amqp.Dial()` is blocking and fails immediately if RabbitMQ
is unreachable. Unlike MongoDB, a successful Dial means you are connected.

**Connection vs Channel:**
```
Connection (1 per service) — TCP socket, expensive to create
    └── Channel (many per connection) — lightweight, used for actual work
```
Publishers and consumers each get their own channel from the same connection.

**Returns:** `*RabbitMQ{ Connection, Channel }`

---

## Side-by-Side Summary

| | MongoDB | PostgreSQL | RabbitMQ |
|---|---|---|---|
| Driver | mongo-driver | pgx/v5 | amqp091-go |
| Main object returned | Client + Database | Pool | Connection + Channel |
| Needs explicit Ping? | Yes | Yes | No |
| Handles pooling | Internally (driver) | Explicitly (pgxpool) | N/A (channels) |
| URL source | `MONGO_URL` env | built from 5 fields | `RABBITMQ_URL` env |
| Key gotcha | Connect() is lazy | Set MaxConns before pool creation | Channel ≠ Connection |

---

## Dependency Flow in main.go

```
config.NewConfig()
    ↓
logger.NewLogger(cfg)
    ↓
mongodb.NewMongoDBConfig(cfg, log)    → *MongoDB
postgres.NewPostgres(cfg, log)        → *Postgres
rabbitmq.NewRabbitMQ(cfg, log)        → *RabbitMQ
    ↓
ingest/consumer/processor services receive these via constructor injection
```

Each service only receives the connectors it actually needs.
main.go is the only place that knows the full dependency graph.
