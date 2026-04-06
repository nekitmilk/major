## Конфиг №1: `config_bad_1.json`

Содержит три проблемы:
- `debug mode enabled` (LOW)
- `password in plain text` (HIGH)
- `weak hash algorithm (MD5)` (HIGH)

```json
{
  "version": "1.0",
  "debug": true,
  "log": {
    "level": "debug",
    "output": "stdout"
  },
  "database": {
    "host": "localhost",
    "user": "admin",
    "password": "superSecret123",
    "name": "appdb"
  },
  "storage": {
    "digest_algorithm": "MD5"
  },
  "server": {
    "bind": "127.0.0.1",
    "port": 8080
  }
}
```

**Тело запроса для api/v1/check:**
```json
{
  "config": "{\"version\":\"1.0\",\"debug\":true,\"log\":{\"level\":\"debug\",\"output\":\"stdout\"},\"database\":{\"host\":\"localhost\",\"user\":\"admin\",\"password\":\"superSecret123\",\"name\":\"appdb\"},\"storage\":{\"digest_algorithm\":\"MD5\"},\"server\":{\"bind\":\"127.0.0.1\",\"port\":8080}}"
}
```

## Конфиг №2: `config_bad_2.yml`

Содержит две проблемы:
- `binding to 0.0.0.0 without restrictions` (MEDIUM)
- `TLS disabled` (HIGH)

```yml
version: "2.0"
debug: false
server:
  host: 0.0.0.0
  port: 3000
tls:
  enabled: true
  insecure_skip_verify: true
database:
  password: ENV_VAR_PASSWORD
```

**Тело запроса для api/v1/check:**
```json
{
  "config": "version: \"2.0\"\ndebug: false\nserver:\n  host: 0.0.0.0\n  port: 3000\ntls:\n  enabled: true\n  insecure_skip_verify: true\ndatabase:\n  password: ENV_VAR_PASSWORD"
}
```

## Чистый конфиг

`config_good.json`:

```json
{
  "version": "3.0",
  "debug": false,
  "log": {
    "level": "info"
  },
  "server": {
    "host": "127.0.0.1",
    "port": 8443
  },
  "tls": {
    "enabled": true,
    "insecure_skip_verify": false
  },
  "database": {
    "password": "${DB_PASSWORD}"
  },
  "storage": {
    "digest_algorithm": "SHA-256"
  }
}
```
**Тело запроса для api/v1/check:**
```json
{
"config": "{\"version\":\"3.0\",\"debug\":false,\"log\":{\"level\":\"info\"},\"server\":{\"host\":\"127.0.0.1\",\"port\":8443},\"tls\":{\"enabled\":true,\"insecure_skip_verify\":false},\"database\":{\"password\":\"${DB_PASSWORD}\"},\"storage\":{\"digest_algorithm\":\"SHA-256\"}}"
}
```