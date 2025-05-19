# Koinx Crypto Stats Platform

A robust cryptocurrency statistics and analysis platform built with Go. This project consists of two main components: an API server and a worker server, working together to provide real-time cryptocurrency data and analysis.

## 🚀 Features

- Real-time cryptocurrency statistics
- Price deviation analysis
- MongoDB integration for data persistence
- Redis pub/sub for real-time updates
- RESTful API endpoints
- Docker support for easy deployment
- Background worker for data processing

## 🏗️ Architecture

The platform consists of two main components:

### 1. API Server
- Handles HTTP requests
- Provides RESTful endpoints
- Manages data retrieval and caching
- Built with Go

### 2. Worker Server
- Processes background tasks
- Handles data aggregation
- Manages real-time updates
- Built with Go

## 📋 Prerequisites

- Go 1.21 or higher
- MongoDB
- Redis
- Docker and Docker Compose

## 🔧 Setup

### Environment Variables

Create a `.env` file in the root directory with:

```env
MONGODB_URI=mongodb://localhost:27017
MONGO_DB_NAME=koinx
PORT=8080
REDIS_HOST=redis:6379
```

### Running Locally

1. Clone the repository:
```bash
git clone https://github.com/yourusername/koinx.git
cd koinx
```

2. Install dependencies:
```bash
go mod download
```

3. Run the services:
```bash
# Run API Server
go run api-server/cmd/api/main.go

# Run Worker Server
go run worker-server/cmd/worker/main.go
```

### Running with Docker

```bash
docker-compose up --build
```

## 📡 API Endpoints

### Get Latest Crypto Stats
**GET** `/stats?coin=bitcoin|ethereum|matic-network`

**Sample Response:**
```json
{
  "price": 40000,
  "marketCap": 800000000,
  "24hChange": 3.4
}
```

### Get Price Deviation
**GET** `/deviation?coin=bitcoin|ethereum|matic-network`

**Sample Response:**
```json
{
  "deviation": 4082.48
}
```

## 📁 Project Structure

```
koinx/
├── api-server/           # API Server Component
│   ├── cmd/
│   │   └── api/         # Application entry point
│   ├── internal/
│   │   ├── config/      # Configuration
│   │   ├── models/      # Data models
│   │   └── services/    # Business logic
│   ├── Dockerfile
│   └── go.mod
│
├── worker-server/        # Worker Server Component
│   ├── cmd/
│   │   └── worker/      # Worker entry point
│   ├── internal/
│   │   ├── config/      # Configuration
│   │   ├── models/      # Data models
│   │   └── services/    # Business logic
│   ├── Dockerfile
│   └── go.mod
│
├── docker-compose.yml    # Docker compose configuration
└── README.md            # Project documentation
```

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 👩‍💻 Author

Ramya Singh 