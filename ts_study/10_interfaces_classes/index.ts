/**
 * 1. インターフェースの定義 (契約)
 * Go と同様に、どのようなメソッドを持つべきかを定義する。
 */
interface Logger {
    info(message: string): void;
    error(message: string): void;
}

/**
 * 2. クラスによる実装
 * implements を使うことで、インターフェースの契約に従っていることを明示する。
 */
class ConsoleLogger implements Logger {
    info(message: string): void {
        console.log(`[INFO] ${message}`);
    }
    error(message: string): void {
        console.error(`[ERROR] ${message}`);
    }
}

/**
 * 3. 依存性の注入 (Dependency Injection)
 * 具体的なクラス（ConsoleLogger）ではなく、
 * 抽象（Logger インターフェース）に依存するように設計する。
 * これにより、後で Logger を簡単に差し替えられる（疎結合）。
 */
class UserService {
    // 外部から Logger を受け取る（注入される）
    constructor(private logger: Logger) {}

    createUser(name: string) {
        // ビジネスロジック
        this.logger.info(`Creating user: ${name}`);
        // 成功したと仮定
        this.logger.info(`User ${name} created successfully.`);
    }
}

/**
 * 4. 構造的部分型 (Structural Typing) の再確認
 * implements を書いていないオブジェクトでも、
 * Logger インターフェースの形を満たしていれば受け入れられる。
 */
const mockLogger = {
    info: (msg: string) => console.log(`MOCK INFO: ${msg}`),
    error: (msg: string) => console.log(`MOCK ERROR: ${msg}`)
};

// 実行
console.log("--- TS Advanced: Interface & DI ---");

const logger = new ConsoleLogger();
const service = new UserService(logger); // 本物のロガーを注入
service.createUser("Alice");

console.log("\n--- Using Mock Logger (Structural Typing) ---");
const mockService = new UserService(mockLogger); // モックを注入
mockService.createUser("Bob");

/**
 * 実行コマンド:
 * npx tsx ts_study/10_interfaces_classes/index.ts
 */
