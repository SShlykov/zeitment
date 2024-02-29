import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import {fileURLToPath, URL} from 'node:url';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL("./src", import.meta.url)),
      '@store': fileURLToPath(new URL("./src/store", import.meta.url)),
      '@router': fileURLToPath(new URL('./src/router/index.js', import.meta.url)),
      '@apiServices': fileURLToPath(new URL('./src/services/apiServices', import.meta.url)),
      '@helpers': fileURLToPath(new URL('./src/helpers', import.meta.url)),
      '@cmp': fileURLToPath(new URL('/src/components', import.meta.url)),
      '@atoms': fileURLToPath(new URL('./src/components/1_atoms', import.meta.url)),
      '@molecules': fileURLToPath(new URL('./src/components/2_molecules', import.meta.url)),
      '@organisms': fileURLToPath(new URL('./src/components/3_organisms', import.meta.url)),
      '@frames': fileURLToPath(new URL('./src/components/4_frames', import.meta.url)),
      '@pages': fileURLToPath(new URL('./src/components/5_pages', import.meta.url)),
    }
  },
  test: {
    globals:true,
    environment: 'happy-dom',
  }
})

