console.log("--- TypeScript Data Structures ---");

// 1. 配列 (Arrays) - 標準で可変長
// 公式: Handbook > Everyday Types > Arrays
const list: string[] = ["Go", "TS", "Python"];
list.push("Rust"); // 末尾に追加
console.log("Array:", list);

// 配列の操作 (TS/JSは強力な組み込みメソッドが多い)
const upperList = list.map(item => item.toUpperCase());
console.log("Uppercase:", upperList);

// 2. オブジェクト (Objects as Maps)
// 簡易的なキー・バリューとして多用される
const scores: { [key: string]: number } = {
    "Alice": 90,
    "Bob": 85
};
console.log("Object Scores:", scores);

// 存在確認 (in 演算子や undefined チェック)
if ("Charlie" in scores) {
    console.log("Charlie exists");
} else {
    console.log("Charlie not found");
}

// 3. Map オブジェクト (より厳格な Map)
// 公式: MDN > Map (TS Handbook では標準組み込み型として扱われる)
const mapScores = new Map<string, number>();
mapScores.set("Alice", 95);
console.log("Map Get Alice:", mapScores.get("Alice"));
console.log("Map Has Bob?:", mapScores.has("Bob"));

// 削除
delete scores["Alice"]; // オブジェクトの場合
mapScores.delete("Alice"); // Mapの場合
console.log("After delete scores:", scores);

/**
 * 実行コマンド:
 * npx tsx ts_study/04_data_structures/index.ts
 * (または npx ts-node --transpile-only ts_study/04_data_structures/index.ts)
 */
