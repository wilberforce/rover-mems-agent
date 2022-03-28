import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

/**
 * @type {import('vite').UserConfig}
 */

export default defineConfig({
  plugins: [vue()],

  css: { preprocessorOptions: { scss: { charset: false } } },

  build: {
    outDir: "web-static",
    brotliSize: false,
    cssCodeSplit: true,
    rollupOptions: {
      output: {
        manualChunks: () => "everything.js", // create one js bundle
      }      
    },
  },
  server: {
    open: "/", // auto open browser in dev mode
    host: true, // and on dev ip,
    https: true,
    proxy: {
      '^/ws': {
        target: 'ws://localhost:8080',
        ws: true,
        preserveHeaderKeyCase: true,
      }
    }
  },
});
