import { defineConfig } from 'vite'

export default defineConfig({
  build: {
    lib: {
      entry: 'src/main.js',
      name: 'RapiDoc',
      fileName: 'rapidoc',
      formats: ['iife']
    },
    outDir: '../static',
    rollupOptions: {
      output: {
        inlineDynamicImports: true
      }
    }
  }
})