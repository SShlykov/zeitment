
/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
    'node_modules/flowbite-vue/**/*.{js,jsx,ts,tsx,vue}',
    'node_modules/flowbite/**/*.{js,jsx,ts,tsx}'
  ],
  darkMode: "class", // or 'media' or 'class'
  plugins: [require("@tailwindcss/forms"), require('flowbite/plugin')],
  theme: {
    screens: {
      'sm': '420px',
      'md': '768px',
      'lg': '1024px',
      'xl': '1366px',
      '2xl': '1536px', //local
      '3xl': '2048px', //local
    },
  }
}
