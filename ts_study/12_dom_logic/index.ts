console.log("--- TypeScript: Data to DOM Rendering Logic ---");

interface User {
    id: number;
    name: string;
    role: "Admin" | "User";
    isActive: boolean;
}

const users: User[] = [
    { id: 1, name: "Alice", role: "Admin", isActive: true },
    { id: 2, name: "Bob", role: "User", isActive: true },
    { id: 3, name: "Charlie", role: "Admin", isActive: false },
];

/**
 * 1. filter と map を使った DOM 構築のシミュレーション
 * 実際の Web 開発では、React や Vue の裏側でこのような処理が行われています。
 */
const renderUserList = (users: User[]): string => {
    // アクティブな Admin だけをリスト形式の HTML に変換する
    const htmlItems = users
        .filter(u => u.isActive && u.role === "Admin") // 1. 絞り込み
        .map(u => `  <li id="user-${u.id}">${u.name} (ADMIN)</li>`) // 2. HTML要素へ変換
        .join("\n"); // 3. 結合

    return `<ul>\n${htmlItems}\n</ul>`;
};

/**
 * 2. reduce を使った集計データの DOM 表示
 */
const renderStats = (users: User[]): string => {
    const totalActive = users.reduce((count, u) => u.isActive ? count + 1 : count, 0);
    return `<div class="stats">Total Active Users: ${totalActive}</div>`;
};

// 実行結果の出力
console.log("User List HTML:");
console.log(renderUserList(users));

console.log("\nStats HTML:");
console.log(renderStats(users));

/**
 * 補足:
 * TS/JS の reduce は非常に強力で、配列を「文字列」や「数値」だけでなく、
 * 「1つの大きなオブジェクト（DOMツリーのシミュレーション）」に変換することも可能です。
 * 
 * 実行コマンド:
 * npx tsx ts_study/12_dom_logic/index.ts
 */
