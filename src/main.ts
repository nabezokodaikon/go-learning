import "./wasm_exec.js";
import { DduItem } from "./dduItem.ts";

function createDduItem(): DduItem {
  return {
    word: "test word",
    display: "test display",
    action: null,
    data: null,
    //highlights: ItemHighlight[],
    highlights: [
        {
          name: "test highlight name1",
          hl_group: "test highlight hl_group1",
          col: 444,
          width: 555,
        },
      ],
    status: {
        size: 111,
        time: 222,
      },
    kind: "test kind",
    level: 333,
    treePath: "test treePath",
    isExpanded: true,
    isTree: false,

    matcherKey: "test matcherKey",
    __sourceIndex: 666,
    __sourceName: "test sourceName",
    __level: 777,
    __expanded: true,
    __columnTexts: {
      888: "test columnTexts",
    },
    __groupedPath: "test groupedPath",
  };
}

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
const dduItem1 = createDduItem();
const dduItem2 = createDduItem();
dduItem2.word = "test dduItem2.word";
const array = [dduItem1, dduItem2];
const length = array.length;
const input = "hoge";
const params = [input, length, ...array]
const result = (globalThis as any).find(...params);

console.log(result);

