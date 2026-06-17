console.log("--- TypeScript Interfaces ---");

/**
 * 1. インターフェースの定義 (公式: Handbook > Interfaces)
 */
interface Shape {
    area(): number;
}

/**
 * 2. 実装 (implements を使うのが一般的だが...)
 */
class Circle implements Shape {
    constructor(public radius: number) {}
    area(): number {
        return Math.PI * this.radius ** 2;
    }
}

/**
 * 3. 構造的部分型 (Structural Subtyping)
 * TSは「implements」を書いていなくても、形が合っていれば代入できる
 * (公式: Handbook > Type Compatibility)
 */
const square = {
    side: 10,
    area() {
        return this.side * this.side;
    }
};

function printArea(s: Shape) {
    console.log(`Area: ${s.area().toFixed(2)}`);
}

// 実行
const c = new Circle(5);
printArea(c);      // もちろんOK
printArea(square); // implementsを書いていないオブジェクトでもOK！(ここがTSの特徴)

/**
 * 4. type alias との比較 (公式: Handbook > Interfaces vs Types)
 */
type Point = {
    x: number;
    y: number;
};

const p: Point = { x: 10, y: 20 };
console.log("Point:", p);

/**
 * 実行コマンド:
 * npx tsx ts_study/06_interfaces/index.ts
 */
