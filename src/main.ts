import "./wasm_exec.js";

type Animal = {
  Name: string;
  Hoge: number;
}

const go = new (globalThis as any).Go();

const wasmCode = await Deno.readFile("main.wasm");
const wasmModule = await WebAssembly.compile(wasmCode);
const wasmInstance = await WebAssembly.instantiate(wasmModule, go.importObject);

go.run(wasmInstance);

const array: Array<Animal> = [{Name: "aaa", Hoge: 1}, {Name: "bbb", Hoge: 2}]
const length = array.length;
const params = [length, ...array]
const result = (globalThis as any).structure(...params);
console.log(result);

