// 1. 変数宣言 (明示的)
let message: string = "Hello TypeScript!";

// 2. 型推論
let version = 5.4;

// 3. 定数
const language: string = "TypeScript";

// 4. 未初期化 (TypeScript/JSにはゼロ値はなく、undefinedになる)
let count: number | undefined;
let isActive: boolean | undefined;

console.log(`--- ${language} Study ---`);
console.log("Message:", message);
console.log("Version:", version);
console.log("Default Count:", count);   // undefined が表示される
console.log("Default Active:", isActive); // undefined が表示される
