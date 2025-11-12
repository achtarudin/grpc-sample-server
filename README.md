# 🚀 gRPC Sample Server

**Server gRPC modern yang dibangun dengan Go untuk demonstrasi implementasi microservices**

[![Go](https://img.shields.io/badge/Go-1.24.4+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![gRPC](https://img.shields.io/badge/gRPC-Latest-4285F4?style=flat&logo=grpc)](https://grpc.io/)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=flat&logo=docker)](https://docker.com/)
[![Protocol Buffers](https://img.shields.io/badge/Protocol_Buffers-Latest-blue?style=flat)](https://protobuf.dev/)

## 🌐 Live Demo

**URL Live Server:** https://grpc-server.cutbray.tech

> **Catatan:** Server ini sudah dalam tahap production dan siap untuk testing!

## 📋 Deskripsi

Proyek ini adalah implementasi gRPC server yang komprehensif menggunakan Go (Golang) dengan port default 9000. Server ini mendemonstrasikan best practices dalam pengembangan microservices modern dengan arsitektur yang bersih dan scalable. Cocok untuk referensi pengembangan aplikasi enterprise atau sebagai starting point untuk proyek gRPC yang kompleks.

### ✨ Fitur Utama Berdasarkan Struktur Project

- 🎯 **Clean Architecture** - Implementasi hexagonal architecture dengan layer yang terpisah
- 🔄 **4 Tipe gRPC Services** - Unary, Server Streaming, Client Streaming, dan Bidirectional Streaming
- 🌍 **Multi-language Greetings** - Mendukung sapaan dalam berbagai bahasa (Hello, Bonjour, Halo, Hola, Ciao)
- �️ **gRPC Interceptors** - Custom unary dan stream interceptors untuk middleware
- 📊 **Structured Logging** - Custom logging format dengan color output
- 🔧 **Environment Configuration** - Flexible configuration dengan .env files
- 🧪 **Protocol Buffer Validation** - Validasi otomatis menggunakan protovalidate
- � **Docker Ready** - Development dan production containerization
- ⚡ **Graceful Shutdown** - Proper shutdown handling dengan context timeout
- 🔄 **Live Reload** - Development mode dengan auto-reload menggunakan gow

## 🧪 Testing dengan Postman

⚠️ **PENTING:** Server ini adalah gRPC server murni (port 9000), bukan HTTP REST API. Untuk testing dengan Postman, Anda memerlukan:

### 🛠️ Cara Testing gRPC dengan Postman:

1. **Buka Postman** dan pilih request type **gRPC**
2. **Server URL:** `https://grpc-server.cutbray.tech`
4. **Pilih Service Method** yang tersedia

### 📝 Available gRPC Methods:

**1. SayHello (Unary)**
```protobuf
// Request
grpcurl -d '{  "name": "celebrer"}' grpc-server.cutbray.tech:443 hello.v1.HelloService/SayHello


// Response
{
  "message": "Ciao, celebrer!",
  "createdAt": "2025-11-12T12:57:37.478209782Z"
}

```

**2. SayManyHellos (Server Streaming)**
```protobuf
// Request
grpcurl -d '{  "name": "veniam"}' grpc-server.cutbray.tech:443  hello.v1.HelloService/SayManyHellos


// Response (10 messages stream)
{
  "message": "Bonjour, celebrer! - 1",
  "createdAt": "2025-11-12T12:55:27.657200532Z"
}
{
  "message": "Hello, celebrer! - 2",
  "createdAt": "2025-11-12T12:55:28.157448884Z"
}
{
  "message": "Bonjour, celebrer! - 3",
  "createdAt": "2025-11-12T12:55:28.657686300Z"
}

....
```

**3. SayHelloToEveryone (Client Streaming)**
```protobuf
// Request
grpcurl -d @ grpc-server.cutbray.tech:443 hello.v1.HelloService/SayHelloToEveryone << EOF
{  "name": "veniam"}
{  "name": "celebrer"}
{  "name": "celebrer"}
EOF

// Response
{
  "message": "Halo, veniam! Ciao, celebrer! Bonjour, celebrer! ",
  "createdAt": "2025-11-12T13:01:31.418213757Z"
}
```

**4. SayHelloContinuous (Bidirectional Streaming)**
```protobuf
// Request
grpcurl -d @ grpc-server.cutbray.tech:443 hello.v1.HelloService/SayHelloContinuous << EOF
{  "name": "perferendis"}
{  "name": "veniam"}
{  "name": "celebrer"}
EOF

{
  "message": "Ciao, perferendis!",
  "createdAt": "2025-11-12T13:02:55.724872800Z"
}
{
  "message": "Hello, veniam!",
  "createdAt": "2025-11-12T13:02:55.724900665Z"
}
{
  "message": "Hello, celebrer!",
  "createdAt": "2025-11-12T13:02:55.724906304Z"
}
```

## 🔗 Protocol Buffers

**Source Protobuf:** [https://github.com/achtarudin/grpc-sample](https://github.com/achtarudin/grpc-sample)

Repository terpisah yang berisi:
- 📄 **Proto Definitions** - Service dan message definitions untuk hello service
- 🔄 **Generated Go Code** - Protobuf generated code dengan gRPC stubs
- 🌐 **gRPC Gateway** - HTTP REST API mapping untuk gRPC services
- � **OpenAPI/Swagger** - API documentation dan schema definitions

## 🏗️ Arsitektur Project (Actual Structure)

```
grpc-sample-server/
├── cmd/
│   ├── cli/                    # CLI tools dan utilities
│   └── server/                 # Main gRPC server (port 9000)
├── internal/
│   ├── adapter/
│   │   ├── grpc_adapter/       # gRPC server implementation
│   │   │   ├── interceptors/   # Custom unary & stream interceptors
│   │   │   ├── grpc_adapter.go
│   │   │   └── hello_service_server.go  # 4 gRPC method implementations
│   │   └── logging/            # Structured logging dengan color
│   ├── port/
│   │   ├── grpc_adapter_port/  # gRPC adapter interfaces
│   │   └── hello_service_port/ # Business logic interfaces
│   ├── service/
│   │   └── hello_service/      # Business logic implementation
│   │       ├── hello_service.go    # Multi-language greetings
│   │       └── hello_service_test.go
│   └── utils/
│       ├── console/            # Console utilities
│       └── helper/             # Environment helper functions
├── docker/
│   ├── go.dev.Dockerfile      # Development container
│   └── go.prod.Dockerfile     # Production container
├── docker-compose.yml         # Development orchestration (port 7000:9000)
├── docker-compose.prod.yml    # Production orchestration
├── Makefile                   # Build automation (dev-server, prod-server, cli-server)
├── .env.example               # Environment configuration template
└── go.mod                     # Dependencies (Go 1.24.4)
```

### 🔄 Flow Arsitektur

1. **gRPC Client** → **gRPC Interceptors** → **gRPC Adapter** → **Service Port** → **Business Logic**
2. **Dependency Injection** - Interface-based design untuk testability
3. **Context Management** - Proper context handling dengan graceful shutdown
4. **Protocol Buffer Validation** - Automatic request validation
5. **Metadata Handling** - Custom headers dan metadata processing

## 🚀 Quick Start

### Prerequisites
- Go 1.24.4+
- Docker & Docker Compose
- Make
- gow (untuk development live reload)

### 🐳 Menjalankan dengan Docker

```bash
# Clone repository
git clone https://github.com/achtarudin/grpc-sample-server.git
cd grpc-sample-server

# Install tools dan dependencies
make install-tools
make install-deps

# Development mode (port 7000:9000)
docker-compose up dev

# Production mode
docker-compose -f docker-compose.prod.yml up
```

### 💻 Development Setup

```bash
# Install development tools
make install-tools    # Install gow untuk live reload

# Install dependencies
make install-deps     # Install semua Go modules

# Development server dengan live reload
make dev-server       # Jalankan server dengan auto-reload

# CLI mode
make cli-server       # Jalankan dalam mode CLI

# Production build
make build-server     # Build binary ke ./bin/
make prod-server      # Jalankan binary production
```

## 🔧 Konfigurasi

Server menggunakan environment variables dengan fallback ke default values:

```env
# Server Configuration (dari .env file)
GRPC_PORT=9000        # Default gRPC port

# Network Configuration (Docker)
# Development: localhost:7000 → container:9000
# Production: grpc-server.cutbray.tech:9000
```

**File Konfigurasi:**
- `.env` - Environment variables untuk development
- `.env.example` - Template konfigurasi
- `.env.prod` - Production configuration

## 📈 Logging & Monitoring

- **Structured Logging:** Custom format dengan color console output
- **Metadata Handling:** Automatic logging dari gRPC headers
- **Graceful Shutdown:** 5 detik timeout untuk proper cleanup
- **Context Management:** Proper context cancellation dan timeout handling
- **Error Handling:** Comprehensive error response dengan gRPC status codes

## � Testing & Development

### Available Make Commands:
```bash
make install-tools     # Install gow untuk live reload
make install-deps      # Update dependencies dari GitHub
make dev-server        # Development server dengan live reload
make cli-server        # CLI mode untuk testing
make build-server      # Build production binary
make prod-server       # Run production binary
make clean-server-bin  # Clean build artifacts
```

### gRPC Testing Tools:
- **Postman gRPC** - GUI testing untuk gRPC services
- **grpcurl** - Command line testing
- **BloomRPC/Kreya** - Alternative gRPC clients
- **evans** - Interactive gRPC client

## 🔍 Fitur Khusus

### 1. Multi-Language Greetings
Server secara random memilih sapaan dari 5 bahasa:
- **English:** "Hello"
- **French:** "Bonjour" 
- **Indonesian:** "Halo"
- **Spanish:** "Hola"
- **Italian:** "Ciao"

### 2. gRPC Streaming Support
- **Server Streaming:** 10 pesan berurutan dengan delay 500ms
- **Client Streaming:** Aggregate multiple names ke satu response
- **Bidirectional Streaming:** Real-time two-way communication

### 3. Production Features
- **Graceful Shutdown:** Proper cleanup dengan 5 detik timeout
- **Context Handling:** Cancellation dan timeout management
- **Custom Interceptors:** Logging dan middleware processing
- **Docker Multi-stage:** Optimized production images


## 👨‍💻 Author

**Achtarudin**
- 🌐 Live Server: [grpc-server.cutbray.tech](https://grpc-server.cutbray.tech)
- � Protobuf Repo: [github.com/achtarudin/grpc-sample](https://github.com/achtarudin/grpc-sample)
- 🐙 GitHub: [@achtarudin](https://github.com/achtarudin)

---

> **💼 Catatan Portofolio:** Proyek ini mendemonstrasikan kemampuan menggunakan gRPC, Go 1.24.4, Protocol Buffers, Docker containerization, dan implementasi clean architecture dengan hexagonal pattern. Menunjukkan Bagaimana Protocol Buffers atau GRPC bekerja.
