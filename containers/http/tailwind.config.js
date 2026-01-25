/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./public/views/**/*.html", "./public/css/**/*.css"],
  theme: {
    extend: {
      fontFamily: {
        sans: ["Inter", "system-ui", "sans-serif"],
      },
      colors: {
        primary: {
          50: "#f1f6fe",
          100: "#e3ebfb",
          200: "#c0d5f7",
          300: "#88b3f1",
          400: "#498ce7",
          500: "#226ed5",
          600: "#1453b5",
          700: "#114293",
          800: "#123a7a",
          900: "#153365",
          950: "#071022",
        },
      },
    },
  },
  plugins: [require("tailwind-scrollbar")({ notcompatible: true })],
};
