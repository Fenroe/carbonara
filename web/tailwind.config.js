/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
  theme: {
    extend: {
      colors: {
        "carb-black": "#181D27",
        "carb-green-dark": "#254D32",
        "carb-green-mid": "#3A7D44",
        "carb-green-light": "#69B578",
        "carb-yellow": "#D0DB97",
        "carb-white": "#F2F4F8",
      },
    },
  },
  plugins: [],
};
