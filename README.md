# Go DDD Scaffolding Tool

A command-line tool for generating Go projects following Domain-Driven Design principles.

## Installation

1. Install the tool globally:

```bash
go install github.com/kanherepratik/go-ddd-skel@latest
```

2. Add the Go bin directory to your PATH (if not already added):

```bash
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshrc
source ~/.zshrc
```

## Usage

### Initialize a New Project

```bash
go-ddd-skel init my-project
```

### Generate Domain Entities

```bash
go-ddd-skel domain User
```

### Generate Use Cases

```bash
go-ddd-skel usecase CreateUser
```

### Generate Handlers

```bash
go-ddd-skel handler UserHandler
```

### Generate Tests

```bash
go-ddd-skel tests User
```

### Visualize Architecture

```bash
go-ddd-skel arch
```

### Generate Documentation

```bash
go-ddd-skel docs --type markdown
```

### Developer Experience Features

```bash
go-ddd-skel dx lint
go-ddd-skel dx air
go-ddd-skel dx telemetry
```

### Monorepo Support

```bash
go-ddd-skel monorepo
```

### Manage Plugins

```bash
go-ddd-skel plugin install ./my-plugin.so
go-ddd-skel plugin list
go-ddd-skel plugin remove my-plugin
```

## Features

- **Project Initialization**: Creates a DDD-compliant project structure
- **Code Generation**: Generates domains, use cases, handlers, and tests
- **Documentation**: Supports markdown and OpenAPI documentation
- **Architecture Visualization**: Generates dependency graphs
- **Plugin System**: Extensible through .so plugins
- **Developer Experience**: Includes linting, live reload, and telemetry
- **Monorepo Support**: Creates multi-service folder structure

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)
