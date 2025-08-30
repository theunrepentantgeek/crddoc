# crddoc - CRD Documentation Generator

Always reference these instructions first and fallback to search or bash commands only when you encounter unexpected information that does not match the info here.

## Working Effectively

### Bootstrap and Build
- Install dependencies: `./.devcontainer/install-dependencies.sh --skip-installed` -- takes 35 seconds. NEVER CANCEL. Set timeout to 60+ seconds.
- Export PATH: `export PATH="$(pwd)/tools:$PATH"`
- Build: `go build -o build/crddoc` or `task build` -- first build takes 16 seconds, cached builds take <1 second. NEVER CANCEL. Set timeout to 60+ seconds for first build.
- Run tests: `go test ./...` or `task unit-test` -- first test run takes 13 seconds, cached runs take <1 second. NEVER CANCEL. Set timeout to 30+ seconds for first run.
- Run linting: `task lint` -- takes 1.5 minutes first time (comprehensive with 75 linters), cached runs take <1 second. NEVER CANCEL. Set timeout to 120+ seconds for first run.
- Full CI: `task ci` -- combines build, test, and lint. Takes 1.5 minutes first time, <1 second cached. NEVER CANCEL. Set timeout to 120+ seconds for first run.

### Requirements
- Go version 1.24.0 or later (check with `go version`)
- The project requires numerous tools that are installed automatically to the local `tools/` directory

### Tool Usage
The crddoc binary generates Markdown documentation for Kubernetes Custom Resource Definitions (CRDs) by parsing Go source code.

Basic commands:
- `./build/crddoc --help` -- show all available commands
- `./build/crddoc export templates --folder ./templates` -- export built-in templates for customization
- `./build/crddoc export configuration --output crddoc.yaml` -- export default configuration
- `./build/crddoc document crds --output docs.md ./package` -- generate documentation for CRDs in ./package

## Validation

### Manual Testing Requirements
ALWAYS manually validate any changes by running through these complete scenarios:

1. **Template Export Workflow**:
   - Run: `./build/crddoc export templates --folder /tmp/test-templates`
   - Verify: Check that 7 template files are created (class-diagram.tmpl, crd.tmpl, embed.tmpl, properties.tmpl, property.tmpl, usage.tmpl, values.tmpl)

2. **Configuration Export Workflow**:
   - Run: `./build/crddoc export configuration --output /tmp/test-config.yaml`
   - Verify: Check that YAML configuration file is created

3. **Documentation Generation Workflow**:
   - Run: `./build/crddoc document crds --output /tmp/test-docs.md ./internal/config`
   - Verify: Check that Markdown documentation is generated with proper structure including tables and class diagrams

### Build and Test Validation
- ALWAYS run `task ci` before finalizing changes -- this runs all build, test, and lint checks
- ALWAYS run `task tidy` to format code before committing
- The CI process uses a comprehensive linting setup with 75 different linters via golangci-lint

## Common Tasks

### Repository Structure
```
ls -la /
.devcontainer/          # Development environment setup and dependency installation
.github/workflows/      # CI/CD pipelines (pr-validation.yml, codeql.yml)
cmd/                    # CLI command implementations
internal/config/        # Configuration types and handling
internal/generator/     # Documentation generation engine
internal/model/         # CRD model types
internal/packageloader/ # Go package parsing
docs/                   # Generated documentation examples
samples/                # Example outputs
templates/              # Built-in Go templates
main.go                 # Main entry point
Taskfile.yml           # Build automation
go.mod                  # Go module definition
```

### Key Files and Their Purpose
- `main.go` -- Entry point, creates logger and executes CLI commands
- `cmd/document-crds.go` -- Main command for generating CRD documentation
- `cmd/export-templates.go` -- Command to export built-in templates
- `cmd/export-configuration.go` -- Command to export default configuration
- `internal/generator/` -- Core documentation generation logic
- `internal/packageloader/` -- Parses Go packages to extract CRD types
- `internal/model/` -- Type definitions for the documentation model
- `templates/` -- Go templates used for rendering documentation
- `Taskfile.yml` -- Defines all available build and development tasks

### Available Task Commands
Run `task --list` to see all available commands:
- `task build` -- Build the crddoc binary
- `task test` -- Run tests and linting
- `task unit-test` -- Run only unit tests
- `task lint` -- Run golangci-lint with comprehensive rule set
- `task ci` -- Run all continuous integration checks
- `task docs` -- Generate internal documentation
- `task tidy` -- Format code and tidy dependencies

### Timing Expectations
- **NEVER CANCEL**: Dependency installation takes 35 seconds
- **NEVER CANCEL**: First build takes 16 seconds, subsequent builds <1 second
- **NEVER CANCEL**: First test run takes 13 seconds, subsequent runs <1 second  
- **NEVER CANCEL**: First lint run takes 1.5 minutes, subsequent runs <1 second
- **NEVER CANCEL**: Documentation generation is very fast (~0.2 seconds)

## Development Workflow

### Making Changes
1. ALWAYS run dependency installation first: `./.devcontainer/install-dependencies.sh --skip-installed`
2. Set PATH: `export PATH="$(pwd)/tools:$PATH"`
3. Make your code changes
4. Build and test: `task build && task unit-test`
5. Run linting: `task lint`
6. Test functionality manually using the validation scenarios above
7. Format code: `task tidy`
8. Final validation: `task ci`

### Common Development Patterns
- The project uses the Cobra CLI framework for command structure
- Configuration is handled via YAML files in the `internal/config` package
- Templates use standard Go templating with Sprig functions
- All external tools are installed locally to `tools/` directory to avoid polluting the system

### Testing Philosophy
- Unit tests exist for core internal packages
- The project emphasizes comprehensive linting (75 linters) over extensive unit test coverage
- Manual validation through end-to-end scenarios is critical
- Mutation testing is available via `task gremlins`

## Troubleshooting

### Common Issues
- If builds fail, ensure Go 1.24.0+ is installed
- If tools are missing, run `./.devcontainer/install-dependencies.sh --skip-installed`
- If PATH issues occur, ensure `export PATH="$(pwd)/tools:$PATH"` is set
- The oh-my-posh installation may fail with "cannot stat 'oh-my-posh.json'" but this doesn't affect core functionality
- If linting takes longer than expected, this is normal - the first run analyzes the entire codebase with 75 linters
- If `go test` shows "no test files" for some packages, this is expected - not all packages have tests

### Build Environment
- The project can be built with just `go build` and tested with `go test ./...`
- Alternatively, use the full development environment via `.devcontainer/install-dependencies.sh`
- **Devcontainer option**: Use Visual Studio Code with devcontainers for automatic tool installation
- All build artifacts go to the `build/` directory
- The `tools/` directory contains locally installed development tools

### CI/CD
- PR validation runs in GitHub Actions using devcontainers
- The workflow builds a custom docker image and runs `task ci`
- All checks must pass for PRs to be merged