# Contributing to Gotter

We'd love your contribution! Here's how to get started.

## Reporting Issues

- **Bugs**: Open an issue with a clear title, reproduction steps, and environment (Go version, OS).
- **Feature requests**: Describe the use case and desired behavior.
- Search existing issues first to avoid duplicates.

## Submitting Changes

1. Fork and clone the repository:
   ```bash
   git clone https://github.com/<your-username>/gotter.git
   cd gotter
   ```

2. Create a feature branch:
   ```bash
   git checkout -b feature/your-feature
   ```

3. Make your changes and commit with clear messages:
   ```bash
   git commit -am 'feat: add your feature'
   ```

4. Test locally:
   ```bash
   go test ./...
   gofmt -s -w .
   go vet ./...
   ```

5. Push and open a Pull Request against `main`:
   ```bash
   git push origin feature/your-feature
   ```

## PR Guidelines

- Reference related issues (e.g., "Fixes #123").
- Keep commits focused and descriptive.
- Ensure tests pass: `go test ./...`
- Format code: `gofmt -s -w .`
- If you add dependencies, run `go mod tidy` and commit `go.sum`.

## Code Style

- Follow [Effective Go](https://golang.org/doc/effective_go) conventions.
- Run `go vet` and `gofmt` before committing.
- Prefer explicit error handling.

## Security

For security vulnerabilities, see [SECURITY.md](./SECURITY.md) — do not open a public issue.

## Questions?

Open an issue or start a discussion in the repository. Thanks for contributing!