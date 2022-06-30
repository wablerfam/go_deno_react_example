import * as esbuild from "esbuild";
import { denoPlugin } from "esbuild_deno_loader";
import { copySync } from "fs/copy.ts";
import { ensureDirSync } from "fs/mod.ts";

ensureDirSync("dist");

// Copy Public
copySync("public/", "dist", { overwrite: true });

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
