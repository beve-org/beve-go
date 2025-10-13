/**
 * Axios BEVE Interceptor
 * Automatically encodes requests and decodes responses using BEVE format
 * 
 * Usage:
 *   import { setupBeveInterceptor } from './axios-beve';
 *   setupBeveInterceptor(axios);
 */

import type { AxiosInstance, InternalAxiosRequestConfig, AxiosResponse } from 'axios';

export interface BeveConfig {
  /**
   * BEVE encoder/decoder (from your TypeScript library)
   */
  encode: (data: any) => Uint8Array;
  decode: (data: Uint8Array) => any;
  
  /**
   * Enable BEVE for all requests (default: false)
   * If false, only requests with `useBeve: true` in config will use BEVE
   */
  enableByDefault?: boolean;
  
  /**
   * Fallback to JSON on BEVE decode errors (default: true)
   */
  fallbackToJson?: boolean;
  
  /**
   * Log BEVE usage (default: false)
   */
  debug?: boolean;
}

declare module 'axios' {
  export interface AxiosRequestConfig {
    useBeve?: boolean;
  }
}

/**
 * Setup BEVE interceptor for Axios instance
 */
export function setupBeveInterceptor(
  axios: AxiosInstance,
  config: BeveConfig
): void {
  const {
    encode,
    decode,
    enableByDefault = false,
    fallbackToJson = true,
    debug = false,
  } = config;

  // Request interceptor: Encode request body with BEVE
  axios.interceptors.request.use(
    (requestConfig: InternalAxiosRequestConfig) => {
      const shouldUseBeve = requestConfig.useBeve ?? enableByDefault;

      if (!shouldUseBeve || !requestConfig.data) {
        return requestConfig;
      }

      try {
        // Encode data with BEVE
        const encoded = encode(requestConfig.data);
        
        if (debug) {
          console.log('[BEVE Request]', {
            url: requestConfig.url,
            originalSize: JSON.stringify(requestConfig.data).length,
            beveSize: encoded.length,
            compression: (1 - encoded.length / JSON.stringify(requestConfig.data).length) * 100,
          });
        }

        // Replace data with BEVE bytes
        requestConfig.data = encoded;
        
        // Set BEVE content type
        requestConfig.headers['Content-Type'] = 'application/beve';
        
        // Request BEVE response
        requestConfig.headers['Accept'] = 'application/beve, application/json;q=0.9';
        
        // Important: Tell Axios to handle binary data
        requestConfig.responseType = 'arraybuffer';

      } catch (error) {
        console.error('[BEVE] Request encoding failed:', error);
        // Keep original data and use JSON
        requestConfig.headers['Content-Type'] = 'application/json';
      }

      return requestConfig;
    },
    (error) => Promise.reject(error)
  );

  // Response interceptor: Decode BEVE response
  axios.interceptors.response.use(
    (response: AxiosResponse) => {
      const contentType = response.headers['content-type'] || '';
      
      if (!contentType.includes('application/beve')) {
        // Not a BEVE response, return as-is
        return response;
      }

      try {
        // Decode BEVE data
        const beveData = new Uint8Array(response.data);
        const decoded = decode(beveData);
        
        if (debug) {
          console.log('[BEVE Response]', {
            url: response.config.url,
            beveSize: beveData.length,
            decodedSize: JSON.stringify(decoded).length,
            expansion: (JSON.stringify(decoded).length / beveData.length).toFixed(2) + 'x',
          });
        }

        // Replace data with decoded value
        response.data = decoded;

      } catch (error) {
        console.error('[BEVE] Response decoding failed:', error);
        
        if (fallbackToJson && response.data) {
          try {
            // Try to parse as JSON fallback
            const text = new TextDecoder().decode(response.data);
            response.data = JSON.parse(text);
            if (debug) {
              console.warn('[BEVE] Fallback to JSON parsing');
            }
          } catch (jsonError) {
            console.error('[BEVE] JSON fallback also failed:', jsonError);
            throw error;
          }
        } else {
          throw error;
        }
      }

      return response;
    },
    (error) => Promise.reject(error)
  );

  if (debug) {
    console.log('[BEVE] Interceptor installed', {
      enableByDefault,
      fallbackToJson,
    });
  }
}

/**
 * Create a pre-configured Axios instance with BEVE support
 */
export function createBeveAxios(
  baseConfig: any,
  beveConfig: BeveConfig
): AxiosInstance {
  const axios = require('axios').create(baseConfig);
  setupBeveInterceptor(axios, beveConfig);
  return axios;
}

/**
 * Helper: Check if current browser supports BEVE (WASM)
 */
export function isBEVESupported(): boolean {
  try {
    return typeof WebAssembly !== 'undefined' &&
           typeof WebAssembly.instantiate === 'function';
  } catch {
    return false;
  }
}

/**
 * Auto-detect and setup BEVE with fallback
 */
export async function setupBeveWithFallback(
  axios: AxiosInstance,
  beveConfig: Omit<BeveConfig, 'enableByDefault'>
): Promise<boolean> {
  if (!isBEVESupported()) {
    console.warn('[BEVE] WebAssembly not supported, using JSON');
    return false;
  }

  try {
    // Setup BEVE interceptor
    setupBeveInterceptor(axios, {
      ...beveConfig,
      enableByDefault: true,
      debug: true,
    });
    
    console.log('[BEVE] Enabled for all requests');
    return true;
  } catch (error) {
    console.error('[BEVE] Setup failed:', error);
    return false;
  }
}
