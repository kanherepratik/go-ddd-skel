# Go DDD Scaffolding Tool

A command-line tool for generating Go projects following Domain-Driven Design principles.

## Installation

```bash
# Install the tool
go install github.com/kanherepratik/go-ddd-skel@latest
```

## Usage

### Initialize a New Project

```bash
go-ddd-skel init my-project
```

### Generate Domain Components

```bash
go-ddd-skel gen domain User
```

### Generate Use Cases

```bash
go-ddd-skel gen usecase CreateUser
```

### Generate Handlers

```bash
go-ddd-skel gen handler UserHandler --type http
```

### Generate Tests

```bash
go-ddd-skel gen tests User
```

### Generate Documentation

```bash
go-ddd-skel gen docs --type markdown
```

### Visualize Architecture

```bash
go-ddd-skel graph arch
```

### Manage Plugins

```bash
go-ddd-skel plugin install ./my-plugin.so
go-ddd-skel plugin list
go-ddd-skel plugin remove my-plugin
```

### Developer Experience Features

```bash
go-ddd-skel setup lint
go-ddd-skel setup air
go-ddd-skel setup telemetry
```

### Monorepo Support

```bash
go-ddd-skel setup monorepo
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
