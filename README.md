# 🔒 TLS security validation

Objective: Program a go script to check the TLS security of a given domain using sslabs APIv2.

Result:
- CMD script to run the validation with pulling results until the scan is complete if needed.
- API with a simple HTML frontend to input a domain and see the TLS security results cached from the API (does not support new scans).

## 🧑‍💻 CMD Script

To run the tool locally in the command line, use:

```bash
# go run script.go <host> [new]
# Example to force a new scan:
go run script.go www.ssllabs.com new
# Example to use cached results if available:
go run script.go www.ssllabs.com
```

## 🚀 API Demo

To see a live demo of the application, visit: [https://801r8x12a2.execute-api.us-east-1.amazonaws.com/prod/](https://801r8x12a2.execute-api.us-east-1.amazonaws.com/prod/). The API demo does not supports new scans, only validation of previous scans domain.

## 🛠️ API Development Setup

To run the tool locally, follow these steps:

```bash
cd api/
go run main.go
```

The API will be accessible at `http://localhost:3333`.

## 📖 API Routes

- `GET /`: Serves the main HTML page.
- `GET /validate?host={domain}`: Validates the TLS security of the specified
