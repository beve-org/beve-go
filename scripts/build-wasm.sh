#!/usr/bin/env bash
# BEVE-Go WASM Builder
# Builds BEVE library for WebAssembly using TinyGo
#
# Usage:
#   ./scripts/build-wasm.sh [target]
#
# Targets:
#   wasm      - Standard WebAssembly (default)
#   wasi      - WASI (WebAssembly System Interface)
#   all       - Build all targets

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
BUILD_DIR="${ROOT_DIR}/build/wasm"
VERSION="${BEVE_VERSION:-dev}"

echo "🚀 BEVE-Go WASM Builder"
echo "📦 Version: ${VERSION}"
echo ""

# Check TinyGo installation
if ! command -v tinygo &> /dev/null; then
    echo "❌ TinyGo not found. Please install it first:"
    echo "   brew install tinygo  (macOS)"
    echo "   or visit: https://tinygo.org/getting-started/install/"
    exit 1
fi

TINYGO_VERSION=$(tinygo version | awk '{print $3}')
echo "✅ TinyGo ${TINYGO_VERSION} found"
echo ""

mkdir -p "${BUILD_DIR}"

TARGET="${1:-wasm}"

build_wasm() {
    local target="$1"
    local output="$2"
    local extra_flags="${3:-}"
    
    echo "🔨 Building ${target}..."
    
    cd "${ROOT_DIR}"
    
    # TinyGo WASM build with aggressive optimizations
    tinygo build \
        -target="${target}" \
        -o="${output}" \
        -opt=z \
        -no-debug \
        -panic=trap \
        -gc=leaking \
        ${extra_flags} \
        ./wasm/main.go
    
    if [ -f "${output}" ]; then
        local size=$(ls -lh "${output}" | awk '{print $5}')
        echo "✅ Built: ${output} (${size})"
        
        # Compress with gzip for comparison
        gzip -9 -k -f "${output}"
        local gzip_size=$(ls -lh "${output}.gz" | awk '{print $5}')
        echo "   📦 Compressed: ${output}.gz (${gzip_size})"
    else
        echo "❌ Build failed for ${target}"
        return 1
    fi
    
    echo ""
}

build_wasm_js() {
    echo "📄 Copying JavaScript glue code..."
    
    # Copy TinyGo's wasm_exec.js
    cp "$(tinygo env TINYGOROOT)/targets/wasm_exec.js" "${BUILD_DIR}/"
    echo "✅ wasm_exec.js copied"
    echo ""
}

case "${TARGET}" in
    wasm)
        build_wasm "wasm" "${BUILD_DIR}/beve.wasm"
        build_wasm_js
        ;;
    wasi)
        build_wasm "wasi" "${BUILD_DIR}/beve.wasi.wasm"
        ;;
    all)
        build_wasm "wasm" "${BUILD_DIR}/beve.wasm"
        build_wasm "wasi" "${BUILD_DIR}/beve.wasi.wasm"
        build_wasm_js
        ;;
    *)
        echo "❌ Unknown target: ${TARGET}"
        echo "Usage: $0 [wasm|wasi|all]"
        exit 1
        ;;
esac

echo "🎉 WASM build complete!"
echo "📁 Output directory: ${BUILD_DIR}"
echo ""
echo "To use in browser:"
echo "  1. Serve the files: python3 -m http.server 8080"
echo "  2. Open: http://localhost:8080/build/wasm/"
echo ""
echo "To test with Node.js:"
echo "  node --experimental-wasm-bigint build/wasm/test.js"
