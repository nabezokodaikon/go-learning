import "./wasm_exec.js";
import { DduItem } from "./dduItem.ts";

type Animal = {
  Name: string;
  Hoge: number;
}

const go = new (globalThis as any).Go();

const wasmCode = await Deno.readFile("main.wasm");
const wasmModule = await WebAssembly.compile(wasmCode);
const wasmInstance = await WebAssembly.instantiate(wasmModule, go.importObject);

go.run(wasmInstance);

// Structure
//const array: Array<Animal> = [{Name: "aaa", Hoge: 1}, {Name: "bbb", Hoge: 2}]
//const length = array.length;
//const params = [length, ...array]
//const result: Array<Animal> = (globalThis as any).structure(...params);

// Record
//const record: Record<number, string> = { 222: "foo" };
////console.log(record[222]);
//const params = [record];
//const result: Record<number, string> = (globalThis as any).record(...params);

// DduItem
const dduItem: DduItem = null;


console.log(result);

