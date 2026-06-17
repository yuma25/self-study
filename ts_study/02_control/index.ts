console.log("--- TypeScript Control Structures ---");

// 1. If and Type Guard (公式: Handbook > Narrowing)
function printValue(val: string | number) {
    if (typeof val === "string") {
        // ここでは val は string として補完される
        console.log("String length:", val.length);
    } else {
        // ここでは val は number として補完される
        console.log("Number (fixed):", val.toFixed(2));
    }
}
printValue("Hello");
printValue(100);

// 2. While loop (TS/JSには標準的に存在する)
let count = 0;
while (count < 3) {
    console.log("Count (while):", count);
    count++;
}

// 3. For...of (公式: Handbook > Iterators and Generators)
const nums: number[] = [10, 20, 30];
for (const value of nums) {
    console.log("Value:", value);
}

// 4. Switch (TS/JSは明示的な break が必要)
const fruit = "apple";
switch (fruit) {
    case "apple":
        console.log("It's an apple");
        break; // これがないと次の case も実行されてしまう
    case "banana":
        console.log("It's a banana");
        break;
    default:
        console.log("Unknown fruit");
}
