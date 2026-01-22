# 🔒 TLS security validation

Basic go api using chi as router and Vue.js for the frontend. The general purpouse is to program a go script to check the TLS security of a given domain.

## 🧑‍💻 CMD Script

To run the tool locally in the command line, use:

```bash
# go run script.go <host> <startNew|optional>
go run script.go www.ssllabs.com new
```

## 🚀 API Demo

To see a live demo of the application, visit: [https://801r8x12a2.execute-api.us-east-1.amazonaws.com/prod/](https://801r8x12a2.execute-api.us-east-1.amazonaws.com/prod/). The API demo does not supports new scans, only validation of previous scans domain.

## 🛠️ API Development Setup

To run the fool locally, follow these steps:

```bash
cd api/
go run main.go
```

The API will be accessible at `http://localhost:3333`.

## 📖 API Routes

- `GET /`: Serves the main HTML page.
- `GET /validate?host={domain}`: Validates the TLS security of the specified
