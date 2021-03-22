const sveltePreprocess = require("svelte-preprocess");
const esbuildSvelte = require("esbuild-svelte");
const esbuild = require("esbuild");

esbuild.build({
  entryPoints: ["src/main.js"],
  bundle: true,
  minify: true,
  outdir: "./public",
  plugins: [esbuildSvelte({
    preprocess: sveltePreprocess()
  })],
}).catch(() => process.exit(1));