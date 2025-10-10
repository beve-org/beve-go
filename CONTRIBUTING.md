# Contributing to BEVE

Thank you for your interest in contributing to BEVE! 🎉

## 🚀 Getting Started

### Prerequisites
- Go 1.18 or higher
- Git

### Setup
```bash
git clone https://github.com/beve-org/beve-go.git
cd beve-go
go mod download
go test ./...
```

## 📝 How to Contribute

### Reporting Bugs
1. Check if the bug already exists in [Issues](https://github.com/beve-org/beve-go/issues)
2. If not, create a new issue with:
   - Clear description of the bug
   - Steps to reproduce
   - Expected vs actual behavior
   - Go version and OS
   - Code sample if possible

### Suggesting Features
1. Check [existing feature requests](https://github.com/beve-org/beve-go/issues?q=is%3Aissue+is%3Aopen+label%3Aenhancement)
2. Open a new issue with:
   - Clear use case description
   - Why this feature would be valuable
   - Proposed API (if applicable)

### Submitting Pull Requests

#### Before You Start
1. Fork the repository
2. Create a feature branch: `git checkout -b feature/amazing-feature`
3. Make sure tests pass: `go test ./...`

#### Development Guidelines

**Code Style:**
- Follow [Effective Go](https://golang.org/doc/effective_go.html)
- Run `gofmt -s -w .` before committing
- Use meaningful variable names
- Add comments for exported functions

**Testing:**
- Write tests for new features
- Maintain or improve code coverage
- Run benchmarks for performance-sensitive changes
- Test edge cases and error conditions

**Commit Messages:**
```
<type>(<scope>): <subject>

<body>

Examples:
feat(encoder): add typed array support for int8 slices
fix(decoder): handle empty map keys correctly
docs(readme): update installation instructions
test(core): add edge case tests for struct encoding
perf(buffer): optimize buffer pre-allocation
```

#### Pull Request Process

1. **Update Documentation**
   - Update README.md if adding features
   - Add/update GoDoc comments
   - Update CHANGELOG.md

2. **Run Full Test Suite**
   ```bash
   go test -v -race -cover ./...
   go test -bench=. -benchmem ./...
   ```

3. **Submit PR**
   - Write clear PR description
   - Reference related issues
   - Include benchmark results for performance changes
   - Wait for CI to pass

4. **Code Review**
   - Respond to feedback promptly
   - Make requested changes
   - Keep discussion focused and professional

## 🧪 Testing

### Running Tests
```bash
# Unit tests
go test ./...

# With coverage
go test -cover ./...

# With race detector
go test -race ./...

# Benchmarks
go test -bench=. -benchmem ./...

# Specific benchmark
go test -bench=BenchmarkMarshal -benchtime=10000x ./...
```

### Writing Tests
```go
func TestFeatureName(t *testing.T) {
    // Arrange
    input := ...
    expected := ...

    // Act
    result, err := FunctionUnderTest(input)

    // Assert
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if result != expected {
        t.Errorf("got %v, want %v", result, expected)
    }
}
```

## 📊 Benchmarking

When making performance changes:

1. **Baseline**: Run benchmarks before changes
   ```bash
   go test -bench=. -benchmem > old.txt
   ```

2. **Compare**: Run after changes
   ```bash
   go test -bench=. -benchmem > new.txt
   benchstat old.txt new.txt
   ```

3. **Report**: Include results in PR description

## 🔍 Code Review Checklist

Before submitting:
- [ ] Tests pass (`go test ./...`)
- [ ] Code formatted (`gofmt -s -w .`)
- [ ] Documentation updated
- [ ] Benchmarks run (for performance changes)
- [ ] No new linter warnings
- [ ] Commit messages follow convention
- [ ] CHANGELOG.md updated

## 🐛 Debugging Tips

### Profiling
```bash
# CPU profiling
go test -cpuprofile=cpu.prof -bench=BenchmarkMarshal
go tool pprof cpu.prof

# Memory profiling
go test -memprofile=mem.prof -bench=BenchmarkMarshal
go tool pprof mem.prof
```

### Debugging Tests
```bash
# Verbose output
go test -v ./...

# Run specific test
go test -run=TestSpecificFunction ./...

# With debugger (delve)
dlv test -- -test.run TestSpecificFunction
```

## 📚 Resources

- [BEVE Binary Format Spec](BINARY_FORMAT.md)
- [Performance Guide](PERFORMANCE.md)
- [Architecture Documentation](ARCHITECTURE.md)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)

## 💬 Communication

- GitHub Issues: Bug reports and feature requests
- GitHub Discussions: General questions and ideas
- Pull Requests: Code contributions

## 📄 License

By contributing, you agree that your contributions will be licensed under the MIT License.

## 🙏 Thank You!

Your contributions make BEVE better for everyone. We appreciate your time and effort! ❤️
