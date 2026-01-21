# 🔒 TLS security validation

Basic go api using chi as router and Vue.js for the frontend. The general purpouse is to program a go script to check the TLS security of a given domain.

## 🛠️ Development Setup

```bash
cd api
go run main.go
```

The API will be accessible at `http://localhost:3333`.

## 📖 Routes

- `GET /`: Serves the main HTML page.
- `GET /validate?host={domain}`: Validates the TLS security of the specified
