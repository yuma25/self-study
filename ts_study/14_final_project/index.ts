import { TaskManager } from "./taskmanager/manager";

async function main() {
    console.log("--- TypeScript: Task Manager Final Project ---");

    const m = new TaskManager();

    // 1. 追加
    m.add("Buy milk");
    m.add("Learn TS Utility Types");
    m.add("Write final project");

    // 2. 完了処理 (例外処理)
    try {
        m.complete(2);
        console.log("Task 2 completed.");
    } catch (e) {
        if (e instanceof Error) {
            console.error("Error:", e.message);
        }
    }

    // 存在しないID
    try {
        m.complete(99);
    } catch (e) {
        if (e instanceof Error) {
            console.log("Expected Error:", e.message);
        }
    }

    // 3. 一覧表示
    console.log("\nCurrent Tasks:");
    m.list();
}

main();

/**
 * 実行コマンド:
 * npx tsx ts_study/14_final_project/index.ts
 */
