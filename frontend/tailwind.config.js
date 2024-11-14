/** @type {import('tailwindcss').Config} */
import scrollbarPlugin from 'tailwind-scrollbar';

export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        primary: '#0073c6', //blue
        secondary: '#00274d', //dark blue
        accent: '#00c77b', //green 
        background: '#f8f8f8', //light grey
        textPrimary: '#333', //dark grey
      },
    },
  },
  plugins: [
    scrollbarPlugin ({nocompatible: true})
  ],
}


