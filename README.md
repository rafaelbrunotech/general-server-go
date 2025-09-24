# general-server-go

A boilerplate, **enterprise-grade server application** built with Go, implementing **Domain-Driven Design (DDD)** and **Clean Architecture** principles. Demonstrates Go best practices, including proper project structure, dependency injection, and comprehensive testing.


## 🏛️ Architecture

This project follows **Clean Architecture** and **DDD** principles to ensure **separation of concerns, testability, and maintainability**.


### Layers

- **Domain Layer**: Core business logic, **entities**, **value objects**, **domain events**, and repository interfaces. Framework-independent.  
- **Application Layer**: Implements **use cases** (commands and queries), orchestrates domain objects, follows **CQRS** principles.  
- **Infrastructure Layer**: Handles **external concerns** such as databases, services, and repository implementations.  
- **Presentation Layer**: Manages **API endpoints**, request validation, and response formatting.

### DDD Concepts

- **Entities**: Rich domain models with unique identity (e.g., `User`, `AuthUser`)  
- **Value Objects**: Immutable objects defined by attributes, not identity (e.g., `Email`, `Id`, `Money`)  
- **Domain Events**: Decoupled notifications representing important domain occurrences  
- **Repositories**: Abstract interfaces for persistence and retrieval  
- **Use Cases / Application Services**: Coordinate domain logic and implement workflows  


## 🚀 Features

- JWT-based **authentication & authorization**
- **Dependency injection** with a container-based approach
- Comprehensive **unit and integration tests**
- **TypeSpec**-generated OpenAPI documentation
- Docker multi-stage builds for **dev, test, prod**


## 🏗️ Feature Under Progress
- Event-driven architecture with **domain events**
- Terraform CDK for **cloud deployment**


## 🛠️ Technology Stack

- **Language**: Go 1.25.1  
- **Web Framework**: Gin  
- **Database**: PostgreSQL (`lib/pq`)  
- **API Documentation**: TypeSpec  
- **Containerization**: Docker multi-stage builds  
- **Infrastructure**: Terraform CDK (TypeScript)


## 📁 Project Structure

Following [Go project layout standards](https://github.com/golang-standards/project-layout):


```
├── cmd/server/                 # Application entry point
├── internal/                   # Private application code
│   ├── app.go                  # Application container setup
│   ├── common/                 # Shared domain and infrastructure
│   │   ├── domain/             # Common domain models and events
│   │   └── infrastructure/     # Shared infrastructure services
│   └── packages/               # Feature modules
│       └── user/               # User package
│           ├── domain/         # Entities, value objects, repositories
│           ├── application/    # Use cases (commands & queries)
│           └── infrastructure/ # Controllers, database implementations
├── api/                        # API documentation and HTTP files
├── deployments/                # Infrastructure as Code (Terraform)
├── build/                      # Docker and deployment configurations
└── test/                       # End-to-end tests
```


## 🚀 Getting Started

### Prerequisites

- Go 1.25.1 or later
- Docker and Docker Compose
- TypeSpec CLI (for API documentation)

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/rafaelbrunotech/general-server-go.git
   cd general-server-go
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables**
   ```bash
   # Edit configs/.env.dev with your env variables
   ```

4. **Run with Docker Compose**
   ```bash
   sh ./scripts/dev.sh
   ```


## 📚 API Documentation

This project uses **TypeSpec** for API documentation generation:

### Prerequisites

- Node.js >= v22.0.0


### Generate API Documentation

1. **Install TypeSpec CLI**
   ```bash
   npm install -g @typespec/compiler
   ```

2. **Navigate to API documentation directory**
   ```bash
   cd api/api-doc
   ```

3. **Install dependencies and compile**
   ```bash
   tsp install
   tsp compile .
   ```

4. **View generated documentation**
   ```bash
   # OpenAPI specifications are generated in:
   # tsp-output/schema/
   ```

### Available API Specifications

- `openapi.AuthApi.1.0.yaml` - Authentication endpoints
- `openapi.UsersApi.1.0.yaml` - User management endpoints


## 🧪 Testing

Run the comprehensive test suite:

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...
```

## 🐳 Docker

### Multi-stage Build

The Dockerfile supports multiple targets:
- `dev`: Development environment with hot reload
- `test`: Runs test suite
- `prod`: Optimized production image (scratch-based)


## 🏗️ Infrastructure

Infrastructure is managed using **Terraform CDK** (TypeScript):

You can check the commands [here](docs/TERRAFORM.md)


## 🔧 Development

### Code Quality

- **Linting**: Pre-commit hooks for code quality
- **Testing**: Comprehensive unit and integration tests
- **Architecture**: DDD nd Clean Architecture with dependency inversion
- **Error Handling**: Structured error handling with custom domain errors

### Git Hooks

Pre-configured git hooks for code quality:
- `pre-commit`: Runs linting and basic checks
- `pre-push`: Runs tests before pushing
- `prepare-commit-msg`: Formats commit messages


## 📖 Key Design Patterns

- **Dependency Injection**: Container-based DI for loose coupling
- **Repository Pattern**: Abstract data access layer
- **CQRS**: Command Query Responsibility Segregation
- **Domain Events**: Event-driven architecture
- **Value Objects**: Immutable domain concepts


## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request


## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.


## 👨‍💻 Author

**Rafael Bruno** - [LinkedIn](https://linkedin/in/rafaelbrunotech)
