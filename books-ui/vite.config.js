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
      '@router': fileURLToPath(new URL('./src/router/index.js', import.meta.url)),
      '@helpers': fileURLToPath(new URL('./src/helpers', import.meta.url)),
      '@cmp': '/src/components',
      '@atoms': fileURLToPath(new URL('./src/components/1_atoms', import.meta.url)),
      '@molecules': fileURLToPath(new URL('./src/components/2_molecules', import.meta.url)),
      '@organisms': fileURLToPath(new URL('./src/components/3_organisms', import.meta.url)),
      '@frames': fileURLToPath(new URL('./src/components/4_frames', import.meta.url)),
      '@pages': fileURLToPath(new URL('./src/components/5_pages', import.meta.url)),
    }
  },
})


// export default defineConfig({
//   plugins: [vue()],
//   server: {
//     port: 5444,
//     host: '0.0.0.0',
//   },
//   resolve: {
//     alias: {
//       '@': fileURLToPath(new URL('./src', import.meta.url)),
//       '@cmp': fileURLToPath(new URL('./src/components', import.meta.url)),
//       '@store': fileURLToPath(new URL('./src/store', import.meta.url)),
//       '@utils': fileURLToPath(new URL('./src/utils', import.meta.url)),
//       '@atoms': fileURLToPath(new URL('./src/components/1_atoms', import.meta.url)),
//       '@molecules': fileURLToPath(new URL('./src/components/2_molecules', import.meta.url)),
//       '@frames': fileURLToPath(new URL('./src/components/3_frames', import.meta.url)),
//       '@pages': fileURLToPath(new URL('./src/components/4_pages', import.meta.url)),
//       '@router': fileURLToPath(new URL('./src/router/index.js', import.meta.url)),
//       '@helpers': fileURLToPath(new URL('./src/helpers/index.js', import.meta.url)),
//     },
//   },
// });