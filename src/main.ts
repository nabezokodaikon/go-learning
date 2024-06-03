import "./wasm_exec.js";
import { Go } from "https://deno.land/x/godeno@v0.7.0/mod.ts";

const go = new (globalThis as any).Go();

const wasmCode = await Deno.readFile("main.wasm");
const wasmModule = await WebAssembly.compile(wasmCode);
const wasmInstance = await WebAssembly.instantiate(wasmModule, go.importObject);

go.run(wasmInstance);

const result = (globalThis as any).add(["hello", "world"]);
console.log(result);

