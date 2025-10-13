/**
 * BEVE Axios Interceptor - Usage Examples
 */

import axios from 'axios';
import { setupBeveInterceptor, createBeveAxios, setupBeveWithFallback } from './axios-beve';
// Import your BEVE TypeScript library
import { encode, decode } from '@your-org/beve-ts'; // Replace with your actual library

// ============================================================================
// Example 1: Manual Setup (per-request control)
// ============================================================================

const api = axios.create({
  baseURL: 'https://api.example.com',
});

setupBeveInterceptor(api, {
  encode,
  decode,
  enableByDefault: false, // Only use BEVE when explicitly requested
  debug: true,
});

// Use BEVE for this request
await api.post('/users', 
  { name: 'Alice', email: 'alice@example.com' },
  { useBeve: true } // Enable BEVE for this request
);

// Use JSON for this request (default)
await api.get('/users/123');

// ============================================================================
// Example 2: Global BEVE (all requests use BEVE)
// ============================================================================

const beveApi = axios.create({
  baseURL: 'https://api.example.com',
});

setupBeveInterceptor(beveApi, {
  encode,
  decode,
  enableByDefault: true, // Use BEVE for all requests
  fallbackToJson: true,  // Fallback to JSON on errors
  debug: false,
});

// All requests automatically use BEVE
await beveApi.post('/users', userData);
await beveApi.get('/posts');

// Opt-out of BEVE for specific request
await beveApi.get('/legacy-endpoint', { useBeve: false });

// ============================================================================
// Example 3: Pre-configured Instance
// ============================================================================

const beveClient = createBeveAxios(
  {
    baseURL: 'https://api.example.com',
    timeout: 5000,
  },
  {
    encode,
    decode,
    enableByDefault: true,
    debug: process.env.NODE_ENV === 'development',
  }
);

// Ready to use!
const response = await beveClient.get('/api/data');
console.log(response.data); // Automatically decoded

// ============================================================================
// Example 4: Auto-detect with Fallback
// ============================================================================

const smartApi = axios.create({
  baseURL: 'https://api.example.com',
});

// Automatically detect WASM support and setup
const beveEnabled = await setupBeveWithFallback(smartApi, {
  encode,
  decode,
});

console.log(`Using ${beveEnabled ? 'BEVE' : 'JSON'}`);

// ============================================================================
// Example 5: React Integration
// ============================================================================

import { createContext, useContext, useEffect, useState } from 'react';

const ApiContext = createContext<typeof api | null>(null);

export function ApiProvider({ children }: { children: React.ReactNode }) {
  const [api] = useState(() => {
    const instance = axios.create({
      baseURL: process.env.REACT_APP_API_URL,
    });

    setupBeveInterceptor(instance, {
      encode,
      decode,
      enableByDefault: true,
      debug: process.env.NODE_ENV === 'development',
    });

    return instance;
  });

  return <ApiContext.Provider value={api}>{children}</ApiContext.Provider>;
}

export function useApi() {
  const api = useContext(ApiContext);
  if (!api) throw new Error('useApi must be used within ApiProvider');
  return api;
}

// Usage in components
function UserList() {
  const api = useApi();
  const [users, setUsers] = useState([]);

  useEffect(() => {
    api.get('/users').then(res => setUsers(res.data));
  }, [api]);

  return <div>{/* render users */}</div>;
}

// ============================================================================
// Example 6: Vue 3 Integration
// ============================================================================

import { inject, provide, InjectionKey } from 'vue';
import type { AxiosInstance } from 'axios';

const ApiKey: InjectionKey<AxiosInstance> = Symbol('api');

export function provideApi() {
  const api = axios.create({
    baseURL: import.meta.env.VITE_API_URL,
  });

  setupBeveInterceptor(api, {
    encode,
    decode,
    enableByDefault: true,
  });

  provide(ApiKey, api);
}

export function useApi() {
  const api = inject(ApiKey);
  if (!api) throw new Error('useApi must be used with provideApi');
  return api;
}

// In App.vue setup()
// provideApi();

// In component
// const api = useApi();
// const data = await api.get('/data');

// ============================================================================
// Example 7: Performance Monitoring
// ============================================================================

const monitoredApi = axios.create({
  baseURL: 'https://api.example.com',
});

let totalBeveBytes = 0;
let totalJsonBytes = 0;
let requestCount = 0;

setupBeveInterceptor(monitoredApi, {
  encode: (data) => {
    const encoded = encode(data);
    const jsonSize = JSON.stringify(data).length;
    totalBeveBytes += encoded.length;
    totalJsonBytes += jsonSize;
    requestCount++;
    
    console.log(`Saved: ${jsonSize - encoded.length} bytes (${
      ((1 - encoded.length / jsonSize) * 100).toFixed(1)
    }%)`);
    
    return encoded;
  },
  decode,
  enableByDefault: true,
  debug: true,
});

// Report savings periodically
setInterval(() => {
  if (requestCount > 0) {
    const saved = totalJsonBytes - totalBeveBytes;
    const percentage = (saved / totalJsonBytes * 100).toFixed(1);
    console.log(`📊 BEVE Stats:
      Requests: ${requestCount}
      JSON size: ${(totalJsonBytes / 1024).toFixed(2)} KB
      BEVE size: ${(totalBeveBytes / 1024).toFixed(2)} KB
      Saved: ${(saved / 1024).toFixed(2)} KB (${percentage}%)
    `);
  }
}, 60000); // Every minute

// ============================================================================
// Example 8: Conditional BEVE (based on payload size)
// ============================================================================

const smartBeveApi = axios.create({
  baseURL: 'https://api.example.com',
});

// Only use BEVE for large payloads (>1KB)
const MIN_SIZE_FOR_BEVE = 1024;

smartBeveApi.interceptors.request.use((config) => {
  if (config.data) {
    const jsonSize = JSON.stringify(config.data).length;
    config.useBeve = jsonSize > MIN_SIZE_FOR_BEVE;
    
    if (config.useBeve) {
      console.log(`Using BEVE for large payload (${jsonSize} bytes)`);
    }
  }
  return config;
});

setupBeveInterceptor(smartBeveApi, {
  encode,
  decode,
  enableByDefault: false, // Controlled by the interceptor above
});

// ============================================================================
// Example 9: TypeScript Typing
// ============================================================================

interface User {
  id: number;
  name: string;
  email: string;
}

interface ApiResponse<T> {
  data: T;
  status: number;
}

// Fully typed BEVE API
const typedApi = createBeveAxios(
  { baseURL: 'https://api.example.com' },
  { encode, decode, enableByDefault: true }
);

// Type-safe request
const user = await typedApi.get<User>('/users/123');
console.log(user.data.name); // ✅ Type-safe

const users = await typedApi.post<User[]>('/users', {
  name: 'Bob',
  email: 'bob@example.com',
});

// ============================================================================
// Example 10: Error Handling
// ============================================================================

const robustApi = createBeveAxios(
  { baseURL: 'https://api.example.com' },
  {
    encode,
    decode,
    enableByDefault: true,
    fallbackToJson: true, // Important!
    debug: true,
  }
);

try {
  const response = await robustApi.get('/data');
  console.log('Data:', response.data);
} catch (error) {
  if (axios.isAxiosError(error)) {
    console.error('API Error:', error.response?.data);
  } else {
    console.error('BEVE Error:', error);
  }
}
