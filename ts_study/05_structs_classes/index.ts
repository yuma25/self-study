console.log("--- TypeScript Classes ---");

/**
 * 1. クラスの定義 (公式: Handbook > Classes)
 * TSには本格的なクラスがあり、アクセス修飾子（public/private）が使える
 */
class User {
    // コンストラクタでのプロパティ宣言（省略記法）
    constructor(protected id: number, public name: string) {}

    greet(): void {
        console.log(`Hi, I'm ${this.name} (ID: ${this.id})`);
    }

    updateName(newName: string): void {
        this.name = newName;
    }
}

/**
 * 2. 継承 (Inheritance)
 * extends を使って親クラスの機能を継承できる
 */
class Admin extends User {
    constructor(id: number, name: string, public level: number) {
        super(id, name); // 親クラスのコンストラクタを呼び出す
    }

    // メソッドのオーバーライド
    greet(): void {
        console.log(`[ADMIN] Hi, I'm ${this.name} (Level: ${this.level})`);
    }
}

// --- 実行 ---
const u = new User(1, "Alice");
u.greet();

u.updateName("Bob");
u.greet();

const a = new Admin(99, "Superuser", 10);
a.greet();
a.updateName("Admin Bob");
a.greet();
console.log("Admin Level:", a.level);

/**
 * 実行コマンド:
 * npx tsx ts_study/05_structs_classes/index.ts
 */
