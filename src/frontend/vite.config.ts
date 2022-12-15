import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import path from "path";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  server: {
    https: false,
    proxy: {
      "/api": {
        target: "http://0.0.0.0:8080"
      }
    },
    host: "0.0.0.0",
    port: 3000
  },
  build: {
    target: "modules",
    outDir: "dist",
    manifest: false,
    minify: "esbuild",
    emptyOutDir: true,
    sourcemap: false,
    cssCodeSplit: false
  },
  resolve: {
    alias: [
      {
        find: "@",
        replacement: path.resolve(__dirname, "src")
      },
      {
        find: "~bootstrap",
        replacement: path.resolve(__dirname, "node_modules/bootstrap")
      }
    ]
  }
});
