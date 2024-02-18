import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import {fileURLToPath, URL} from 'node:url';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@': '/src',
      '@store': '/src/store',
      '@cmp': '/src/components',
      '@atoms': fileURLToPath(new URL('./src/components/1_atoms', import.meta.url)),
      '@molecules': fileURLToPath(new URL('./src/components/2_molecules', import.meta.url)),
      '@organisms': fileURLToPath(new URL('./src/components/3_organisms', import.meta.url)),
      '@pages': fileURLToPath(new URL('./src/components/4_pages', import.meta.url)),
    }
  },
})
