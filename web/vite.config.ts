import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

const prefix = `monaco-editor/esm/vs`;

export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      '@': '/src',
    },
  },
  build: {
    rollupOptions: {
      output: {
        manualChunks: {
          jsonWorker: [`${prefix}/language/json/json.worker`],
          cssWorker: [`${prefix}/language/css/css.worker`],
          htmlWorker: [`${prefix}/language/html/html.worker`],
          tsWorker: [`${prefix}/language/typescript/ts.worker`],
          editorWorker: [`${prefix}/editor/editor.worker`],
        },
      },
    },
    target: ['esnext'],
    terserOptions: {
      enclose: false,
      compress: true,
      sourceMap: false,
    },
  },
  server: {
    proxy: {
      '/cms': {
        target: 'http://localhost:8081',
        changeOrigin: true,
      },
    },
  },
  preview: {
    proxy: {
      '/cms': {
        target: 'http://localhost:8081',
        changeOrigin: true,
      },
    },
  },
});
