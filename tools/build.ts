import { copySync } from "fs/copy.ts";
import * as esbuild from "esbuild";
import { denoPlugin } from "esbuild_deno_loader";

// Copy index.html
copySync("public/index.html", "dist/index.html", { overwrite: true });

// ESBuild
const importMapURL = new URL("../import_map.json", import.meta.url);

await esbuild.build({
  plugins: [denoPlugin({ importMapURL })],
  entryPoints: ["js/index.tsx"],
  outfile: "dist/bundle.js",
  bundle: true,
  format: "esm",
});

esbuild.stop();
