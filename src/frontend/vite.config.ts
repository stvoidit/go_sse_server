import { defineConfig, loadEnv, ConfigEnv, UserConfigExport } from "vite";
import { fileURLToPath, URL } from "node:url";
import vue from "@vitejs/plugin-vue";


// https://vitejs.dev/config/
export default ({ mode }: ConfigEnv): UserConfigExport => {
  console.log(import.meta);
  const env = loadEnv(mode, ".");
  return defineConfig({
    plugins: [vue()],
    server: {
      https: false,
      proxy: {
        "/api": {
          target: env.VITE_PROXY_TARGET || "http://localhost:8080",
          changeOrigin: true
        }
      },
      host: env.VITE_HOST || "0.0.0.0",
      port: parseInt(env.VITE_PORT) || 3000
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
          replacement: fileURLToPath(new URL("./src", import.meta.url))
        },
        {
          find: "~bootstrap",
          replacement: fileURLToPath(new URL("node_modules/bootstrap", import.meta.url))
        }
      ]
    }
  });
};
