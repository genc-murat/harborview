import animate from "tailwindcss-animate"

/** @type {import('tailwindcss').Config} */
export default {
  darkMode: ["class"],
  content: [
    './pages/**/*.{js,jsx}',
    './components/**/*.{js,jsx}',
    './app/**/*.{js,jsx}',
    './src/**/*.{js,jsx}',
  ],
  theme: {
    extend: {
      colors: {
        border: "rgb(226, 232, 240)",
        background: "rgb(248, 250, 252)",
        primary: {
          DEFAULT: "rgb(59, 130, 246)",
          hover: "rgb(37, 99, 235)",
          light: "rgb(219, 234, 254)"
        },
        card: {
          DEFAULT: "rgb(255, 255, 255)",
          foreground: "rgb(15, 23, 42)"
        },
        muted: {
          DEFAULT: "rgb(241, 245, 249)",
          foreground: "rgb(100, 116, 139)"
        },
        accent: {
          DEFAULT: "rgb(239, 246, 255)",
          foreground: "rgb(59, 130, 246)"
        }
      },
      borderRadius: {
        lg: "0.5rem",
        md: "calc(0.5rem - 2px)",
        sm: "calc(0.5rem - 4px)",
      },
    },
    container: {
      center: true,
      padding: "2rem",
      screens: {
        "2xl": "1400px",
      },
    },
  },
  plugins: [animate],
}