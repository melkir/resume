/** @type {import('tailwindcss').Config}*/
const config = {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {
      spacing: {
        '1cm': '1cm'
      }
    }
  },

  plugins: []
};

module.exports = config;
