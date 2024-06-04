import "./wasm_exec.js";
//import { Go } from "https://deno.land/x/godeno@v0.7.0/mod.ts";

type Animal = {
  Name: string;
  Hoge: number;
}

const go = new (globalThis as any).Go();

const wasmCode = await Deno.readFile("main.wasm");
const wasmModule = await WebAssembly.compile(wasmCode);
const wasmInstance = await WebAssembly.instantiate(wasmModule, go.importObject);

go.run(wasmInstance);

function receive(array: any) {
  //console.log("Received array from Go:", array);
}

//const array = ["aaa", "bbb", "ccc"];
const array: Array<Animal> = [{name: "aaa", hoge: 1}, {name: "bbb", hoge: 2}]
const length = array.length;
const params = [receive, length, ...array]
const result = (globalThis as any).find(...params);
console.log(result);

