/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./cmd/web/**/*.{templ,go}", "./assets/**/*.{html,js}"],
  theme: {
    container: {
      center: true,
      padding: {
        DEFAULT: "1rem",
        sm: "4rem",
        md: "3rem",
        lg: "2rem",
        xl: "2rem",
        "2xl": "1rem",
      },
    },
    extend: {},
  },
  plugins: [require("@tailwindcss/typography")],
};
