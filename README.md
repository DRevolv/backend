# DRevolv Backend

A Go-based backend service for the DRevolv project.

## 🚀 Quick Start

### Prerequisites

- Go 1.24.1 or later
- Git

### Local Development

```bash
# Clone the repository
git clone <repository-url>
cd DRevolv/backend

# Download dependencies
go mod download

# Run the application
go run main.go

# Build the application
go build -o bin/app main.go

# Run the binary
./bin/app
```

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -race -coverprofile=coverage.out -covermode=atomic ./...

# View coverage report
go tool cover -html=coverage.out
```

## 🔍 Code Quality

### Linting

We use `golangci-lint` for code linting. Install it locally:

```bash
# Install golangci-lint
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run linter
golangci-lint run
```

### Formatting

```bash
# Format code
go fmt ./...

# Format imports
goimports -w .
```

## 🔄 CI/CD Pipeline

Our GitHub Actions pipeline automatically runs on:
- Push to `dev`, `test`, or `prod` branches
- Pull requests to `dev`, `test`, or `prod` branches
- Changes to backend files

### Branch Strategy

- **`dev`** - Development environment (auto-deploys on push)
- **`test`** - Testing/Staging environment (auto-deploys on push, includes security scans)
- **`prod`** - Production environment (auto-deploys on push, includes security scans and additional validation)

### Pipeline Jobs

1. **Test Job** (runs on all branches)
   - Sets up Go 1.24.1
   - Downloads and verifies dependencies
   - Runs all tests with race detection
   - Generates coverage reports
   - Uploads coverage to Codecov

2. **Lint Job** (runs on all branches)
   - Runs golangci-lint with comprehensive rules
   - Checks code formatting and style
   - Identifies potential issues and security vulnerabilities

3. **Build Job** (runs on all branches)
   - Builds the application for multiple platforms with environment info:
     - Linux (amd64)
     - Windows (amd64)
     - macOS (amd64 and arm64)
   - Embeds version and environment information
   - Creates environment-specific build artifacts
   - Uploads binaries for download (90-day retention for prod, 30-day for others)

4. **Security Job** (runs on `test` and `prod` branches only)
   - Runs Gosec security scanner
   - Uploads security analysis results
   - Integrates with GitHub Security tab

5. **Deployment Jobs** (environment-specific)
   - **Deploy-Dev**: Deploys to development environment (from `dev` branch)
   - **Deploy-Test**: Deploys to testing environment (from `test` branch, requires security scan)
   - **Deploy-Prod**: Deploys to production environment (from `prod` branch, includes pre/post-deployment checks)

### Build Artifacts

After successful builds, you can download environment-specific, platform-specific binaries from the GitHub Actions artifacts:

**Artifact Names:** `backend-binaries-{environment}` (where environment is `dev`, `test`, or `prod`)

**Platform Binaries:**
- `app-linux-amd64` - Linux binary
- `app-windows-amd64.exe` - Windows binary
- `app-darwin-amd64` - macOS Intel binary
- `app-darwin-arm64` - macOS Apple Silicon binary

**Build Information:**
Each binary includes embedded version and environment information that can be viewed when running the application.

## 📁 Project Structure

```
backend/
├── main.go              # Application entry point
├── main_test.go         # Basic tests
├── go.mod              # Go module definition
├── .golangci.yml       # Linting configuration
└── README.md           # This file
```

## 🛠️ Development Guidelines

### Adding New Features

1. Create feature branch from `dev`
2. Write tests for your code
3. Ensure all tests pass locally
4. Run linter and fix any issues
5. Create pull request to `dev`
6. CI pipeline will automatically validate your changes

### Environment Promotion

Follow this promotion path for code changes:
1. **Development**: Merge to `dev` branch → Auto-deploy to development environment
2. **Testing**: Merge `dev` to `test` branch → Auto-deploy to testing environment (includes security scans)
3. **Production**: Merge `test` to `prod` branch → Auto-deploy to production environment (includes security scans and additional validation)

### Code Standards

- Follow Go best practices and idioms
- Write comprehensive tests for new functionality
- Add comments for exported functions and types
- Use meaningful variable and function names
- Handle errors appropriately

### Commit Messages

Use conventional commit format:
- `feat:` for new features
- `fix:` for bug fixes
- `docs:` for documentation changes
- `test:` for test-related changes
- `refactor:` for code refactoring
- `ci:` for CI/CD changes

## 📊 Monitoring

The pipeline provides several monitoring capabilities:
- Test coverage reports via Codecov
- Security vulnerability scanning
- Code quality metrics
- Build success/failure notifications

## 🚀 Deployment

The pipeline includes automatic deployment to different environments based on the branch:

### Environment Configuration

Configure the following GitHub repository settings:

1. **Environments**: Create three environments in your repository settings:
   - `development` (for `dev` branch)
   - `testing` (for `test` branch) 
   - `production` (for `prod` branch)

2. **Environment Protection Rules**:
   - **Development**: No protection rules (auto-deploy)
   - **Testing**: Optional - require reviews for manual approval
   - **Production**: Recommended - require reviews and wait timer

### Deployment Process

- **Development**: Automatically deploys when code is pushed to `dev` branch
- **Testing**: Automatically deploys when code is pushed to `test` branch (after security scans)
- **Production**: Automatically deploys when code is pushed to `prod` branch (after security scans and additional validation)

### Customizing Deployments

Update the deployment steps in `.github/workflows/backend.yml` for each environment:

```yaml
# Example for deploy-prod job
- name: Deploy to Production
  run: |
    # Add your production deployment steps here
    # Example: kubectl apply, docker push, etc.
```

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'feat: add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the LICENSE file for details.