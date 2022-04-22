import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

/**
 * @type {import('vite').UserConfig}
 */

export default defineConfig({
  plugins: [vue()],
  define: {
    '__APP_VERSION__': JSON.stringify(process.env.npm_package_version),
},
  css: { preprocessorOptions: { scss: { charset: false } } },

  build: {
    outDir: "web-static",
    brotliSize: false,
    cssCodeSplit: true,
    rollupOptions: {
      output: {
        //manualChunks: () => "everything.js", // create one js bundle
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
