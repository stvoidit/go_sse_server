import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import { resolve } from "path";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  server: {
    https: false,
    proxy: {
      "/api": {
        target: "http://0.0.0.0:8080",
        changeOrigin: true
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
        replacement: resolve(__dirname, "src")
      },
      {
        find: "~bootstrap",
        replacement: resolve(__dirname, "node_modules/bootstrap")
      }
    ]
  }
});
