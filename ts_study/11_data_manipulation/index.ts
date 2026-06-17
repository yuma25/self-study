console.log("--- TypeScript: Declarative Data Manipulation ---");

interface Product {
    name: string;
    price: number;
    available: boolean;
}

const products: Product[] = [
    { name: "Laptop", price: 1000, available: true },
    { name: "Mouse", price: 20, available: true },
    { name: "Keyboard", price: 50, available: false },
    { name: "Monitor", price: 200, available: true },
];

/**
 * シナリオ: 「在庫あり」の商品の価格を「10%引き」にし、その「合計金額」を出す
 * 
 * TypeScript (JavaScript) ではメソッドをチェーンさせる
 * 宣言的（何をするか）を重視するスタイル。
 * 公式: MDN - Array.prototype.filter / map / reduce
 */
const totalPrice = products
    .filter(p => p.available)           // Filter: 条件に合うものだけ残す
    .map(p => p.price * 0.9)            // Map: 値を変換する
    .reduce((sum, price) => sum + price, 0); // Reduce: 一つの値に集約する

console.log(`Total Price (Discounted): ${totalPrice.toFixed(2)}`);

/**
 * 補足: なぜこの書き方が好まれるのか？
 * 1. 非破壊的: 元の products 配列を書き換えず、新しい結果を作る（不変性）。
 * 2. 読みやすさ: 各工程が名前（filter, map...）で分かれているため、意図が明確。
 * 
 * 実行コマンド:
 * npx tsx ts_study/11_data_manipulation/index.ts
 */
