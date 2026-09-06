import { defineConfig } from "vite";

// The build output is checked in and embedded into the Go binary, so it must
// be reproducible and self-contained. During development `npm run dev` proxies
// the API to a locally running quizserver.
export default defineConfig({
  build: {
    outDir: "dist",
    emptyOutDir: true,
  },
  server: {
    proxy: {
      "/api": "http://localhost:8080",
    },
  },
});
