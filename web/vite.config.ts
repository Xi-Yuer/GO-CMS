import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import viteCompression from 'vite-plugin-compression';

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
          tsWorker: [`${prefix}/language/typescript/ts.worker`],
          editorWorker: [`${prefix}/editor/editor.worker`],
        },
        chunkFileNames: 'js/[name]-[hash].js', // 引入文件名的名称
        entryFileNames: 'js/[name]-[hash].js', // 包的入口文件名称
        assetFileNames: '[ext]/[name]-[hash].[ext]', // 资源文件像 字体，图片等
      },
      plugins: [viteCompression()],
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
        target: 'http://127.124.28.77:8081',
        changeOrigin: true,
      },
    },
  },
});
