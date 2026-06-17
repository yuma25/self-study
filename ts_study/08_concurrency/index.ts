console.log("--- TypeScript Async (Promises & Async/Await) ---");

/**
 * 重い処理をシミュレートする関数
 * @returns Promise
 */
async function task(name: string): Promise<string> {
    console.log(`Task ${name} started...`);
    // 2秒待機するPromise
    await new Promise(resolve => setTimeout(resolve, 2000));
    return `Task ${name} finished`;
}

async function main() {
    const start = Date.now();

    /**
     * 1. 非同期処理の開始
     * ここで await を付けないと、バックグラウンドで処理が進む
     */
    const promiseA = task("A");
    const promiseB = task("B");

    console.log("Waiting for tasks to finish...");

    /**
     * 2. 結果の待機
     * Promise.all を使うと並列に待機できる
     */
    const [res1, res2] = await Promise.all([promiseA, promiseB]);

    console.log("Result 1:", res1);
    console.log("Result 2:", res2);

    console.log(`Total time: ${(Date.now() - start) / 1000}s`);
}

main();

/**
 * 実行コマンド:
 * npx tsx ts_study/08_concurrency/index.ts
 */
