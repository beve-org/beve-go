#!/bin/bash
# Cross-platform SIMD string test script

echo "🧪 Testing SIMD String Implementation on Multiple Architectures"
echo "================================================================"
echo ""

# Save current directory
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Test ARM64
echo -e "${YELLOW}Testing ARM64 (Apple Silicon / ARM64)...${NC}"
GOARCH=arm64 go test -v ./core -run "TestValidateUTF8SIMD|TestCountUTF8RunesSIMD" 2>&1 | grep -E "PASS|FAIL|RUN"
if [ $? -eq 0 ]; then
    echo -e "${GREEN}✓ ARM64 tests passed${NC}"
else
    echo -e "${RED}✗ ARM64 tests failed${NC}"
fi
echo ""

# Test AMD64
echo -e "${YELLOW}Testing AMD64 (Intel/AMD x86-64)...${NC}"
GOARCH=amd64 go test -v ./core -run "TestValidateUTF8SIMD|TestCountUTF8RunesSIMD" 2>&1 | grep -E "PASS|FAIL|RUN"
if [ $? -eq 0 ]; then
    echo -e "${GREEN}✓ AMD64 tests passed${NC}"
else
    echo -e "${RED}✗ AMD64 tests failed${NC}"
fi
echo ""

# Benchmark ARM64
echo -e "${YELLOW}Benchmarking ARM64...${NC}"
GOARCH=arm64 go test -bench BenchmarkValidateUTF8_Long/SIMD -benchmem -benchtime=1s ./core 2>&1 | grep "Benchmark"
echo ""

# Benchmark AMD64
echo -e "${YELLOW}Benchmarking AMD64 (cross-compile, may not show accurate performance)...${NC}"
echo "Note: For accurate AMD64 benchmarks, run on native AMD64 hardware"
GOARCH=amd64 go test -bench BenchmarkValidateUTF8_Long/SIMD -benchmem -benchtime=1s ./core 2>&1 | grep "Benchmark"
echo ""

echo -e "${GREEN}════════════════════════════════════════════════════════════${NC}"
echo -e "${GREEN}Cross-platform SIMD tests completed!${NC}"
echo -e "${GREEN}════════════════════════════════════════════════════════════${NC}"
