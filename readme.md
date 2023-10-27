Certainly! Let's include the `config.conf` file and its usage in the `README.md`.

---

# ehe-web Project

This project serves as the web interface for the `ehe-web` application, with a frontend based on Webpack and TypeScript and a backend written in Go. This document provides an overview of the development environment setup, configuration management, and basic commands to aid in the development process.

## Getting Started

### Prerequisites

1. **Node.js**: Ensure you have the latest stable version installed. You can download it from [here](https://nodejs.org/).

2. **Go**: The server side of this project is written in Go. Ensure you have Go installed. You can download and install it from [here](https://golang.org/dl/).

### Installing Dependencies

**Frontend**:

Navigate to the project's root directory and run the following command:

```bash
npm install
```

This will install all the necessary frontend dependencies listed in the `package.json` file.

**Backend (Go)**:

Ensure you have set up the GOPATH properly and have organized your code according to Go's workspace instructions. If you have dependencies that aren't part of the standard library, use Go's package manager:

For Go modules (recommended):

```bash
go mod tidy
```

For older versions using `dep`:

```bash
dep ensure
```

### Configuration Management

The project utilizes a `config.conf` file to manage environment-specific settings. The file has the following format:

```
APP_ENV=development
```

In the HTML templates, this configuration is used to conditionally load resources based on the environment. For instance, in a development environment, the frontend assets are served from the Webpack dev server, while in production, they're served from the `static/dist` directory.

Example from the template:

```html
{{ if eq .Env "development" }}
    <script src="http://localhost:3000/bundle.js"></script>
{{ else }}
    <script src="/static/dist/bundle.js"></script>
{{ end }}
```

Ensure that you set the appropriate value for `APP_ENV` in `config.conf` depending on your development or deployment scenario.

## Development Workflow

### Running the Project in Development Mode

To start both the frontend development server and the Go backend server concurrently, execute:

```bash
npm run dev
```

The frontend will be accessible at [http://localhost:3000](http://localhost:3000), and the backend API will be available at [http://localhost:8080](http://localhost:8080).

### Building for Production

**Frontend**:

To create an optimized production build, use the following command:

```bash
npm run build
```

The built files will be located in the `static/dist` directory.

**Backend (Go)**:

Compile the Go server with:

```bash
go build main.go
```

This will generate an executable for your server.

## Additional Notes

- The frontend assets are located in the `static` directory, with the main application entry point being `static/js/script.ts`.
- If you encounter type-related errors, especially with external libraries, consider installing the relevant type definitions using `npm`, for instance: `npm i --save-dev @types/library-name`.
- The Go backend's main file is `main.go`. Ensure any environment variables or configurations are set appropriately before running the server in different environments.

---

This version of the README provides a comprehensive guide that covers both the frontend and backend sides of your project, as well as the configuration management aspect.