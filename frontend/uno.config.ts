import { defineConfig } from "unocss";
import presetMini from "@unocss/preset-mini";
import { transformerDirectives } from "unocss";

export default defineConfig({
  presets: [presetMini()],
  transformers: [transformerDirectives()],
});
