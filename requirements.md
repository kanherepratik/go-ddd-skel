# 🧾 Product Requirements Document (PRD) – `go-ddd-skel`

## 1. 🎯 Project Overview

`go-ddd-skel` is a command-line scaffolding tool designed to generate Go backend projects following **Domain-Driven Design (DDD)**. It accelerates project bootstrapping by providing ready-to-use folder structures, templates, boilerplate code, and tooling integrations.

---

## 2. 🔥 Goals

- Provide an opinionated DDD project structure for Go.
- Generate domains, use cases, handlers, and infrastructure components.
- Support generation of test stubs and mocks.
- Enable auto-documentation, architecture visualization, and API generation.
- Allow plugin-based extensibility and monorepo support.
- Improve developer experience with linting, live reload, and telemetry integration.

---

## 3. 🧱 Key Features

### 🎬 Initialization

- `init`: Set up a new Go project with DDD structure.

### 📦 Domain & Application Layer Generation

- `gen domain <name>`: Generate domain entity, repository interface, and value object (optional).
- `gen usecase <name>`: Generate service logic and interfaces.
- `gen handler <name> --type [http|grpc]`: Generate handler and route binding.

### 🧪 Test Generation

- `gen tests <name>`: Generate test stubs.
- `--with-mocks`: Integrate mocks using `mockery`.

### 📊 Architecture & Documentation

- `graph arch`: Visualize DDD layer dependency graph using `go-callvis`.
- `gen docs --type [openapi|md]`: Generate documentation from annotations.

### 🧩 Plugins & Extensibility

- `plugin install <path>`: Support `.so` plugins.
- `plugin list/remove`: Plugin lifecycle management.

### 📚 Developer Experience

- `setup lint`: Integrate `golangci-lint`.
- `setup air`: Add `air.toml` for live reload.
- `setup telemetry`: Integrate OpenTelemetry, Prometheus, and Grafana support.
- `setup monorepo`: Generate multi-service folder structure.

---

## 4. 🔧 Project Structure (Generated)

```bash
my-ddd-app/
├── cmd/
│   ├── crons/
│   ├── grpc/
│   └── http/
├── internal/
│   ├── adapters/
│   │   ├── external/
│   │   ├── persistence/
│   │   └── ports/
│   ├── config/
│   ├── core/
│   ├── interfaces/
│   └── usecase/
├── migrations/
├── pkg/
├── scripts/
├── sql/
├── static/
└── Dockerfile, *.yml, etc.
```

| Command                      | Description                                 |
| ---------------------------- | ------------------------------------------- |
| `init`                       | Initialize base project                     |
| `gen domain <name>`          | Generate domain entity and repo interface   |
| `gen usecase <name>`         | Generate usecase logic                      |
| `gen handler <name>`         | Generate handler and route                  |
| `gen tests <name>`           | Generate test boilerplate                   |
| `setup lint`                 | Add golangci-lint configuration             |
| `setup air`                  | Add air live reload config                  |
| `setup telemetry`            | Add OpenTelemetry boilerplate               |
| `setup monorepo`             | Create monorepo-compatible folder structure |
| `graph arch`                 | Visualize architecture graph                |
| `gen docs`                   | Generate markdown/OpenAPI spec              |
| `plugin install/remove/list` | Manage CLI plugins                          |

6. 📌 Non-Goals

- Will not generate frontend code.
- Will not handle CI/CD pipeline generation.
- No automatic DB migrations beyond folder scaffolding.

7. 🗓️ Milestones
   Phase Features ETA
   Phase 1 – Core init, gen domain/usecase/handler Week 1
   Phase 2 – Testing gen tests, mock generation Week 2
   Phase 3 – Docs arch graph, docs gen Week 3
   Phase 4 – Plugins plugin system Week 4
   Phase 5 – DX telemetry, lint, live reload, monorepo Week 5

8. 👥 Target Users
   Backend Go developers

Engineering teams following Clean/DDD architecture

Tech leads and architects building scalable Go systems

9. 🧪 Tooling Dependencies
   Go 1.20+

cobra for CLI

text/template for code generation

Optional: go-callvis, mockery, golangci-lint, air

10. 🔮 Future Enhancements
    GitHub/GitLab template integration

Automatic Swagger doc generation from comments

Interactive CLI mode (e.g., with survey)

GraphQL support
