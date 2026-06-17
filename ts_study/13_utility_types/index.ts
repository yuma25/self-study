console.log("--- TypeScript: Utility Types and Type Manipulation ---");

/**
 * 1. 基礎となる型
 */
interface User {
    id: number;
    name: string;
    email: string;
    role: "Admin" | "User";
    age?: number; // オプショナル
}

/**
 * 2. Partial<T>: すべてのプロパティをオプショナル（任意）にする
 * 用途: 更新処理（一部のフィールドだけ送る場合）
 */
type UpdateUserDto = Partial<User>;

const updateData: UpdateUserDto = {
    name: "New Name",
    // id や email がなくてもエラーにならない
};

/**
 * 3. Pick<T, Keys>: 特定のプロパティだけを抽出する
 * 用途: 公開プロフィール用など
 */
type PublicUser = Pick<User, "name" | "role">;

const publicInfo: PublicUser = {
    name: "Alice",
    role: "Admin",
    // id や email を入れるとコンパイルエラーになる
};

/**
 * 4. Omit<T, Keys>: 特定のプロパティを除外する
 * 用途: IDを除いた「新規作成用」の型など
 */
type CreateUserDto = Omit<User, "id">;

const newUser: CreateUserDto = {
    name: "Bob",
    email: "bob@example.com",
    role: "User",
};

/**
 * 5. Readonly<T>: すべてを読み取り専用にする
 */
type ImmutableUser = Readonly<User>;

const frozenUser: ImmutableUser = {
    id: 99,
    name: "Charlie",
    email: "charlie@example.com",
    role: "User"
};
// frozenUser.name = "New Charlie"; // コンパイルエラー！

console.log("Update Data:", updateData);
console.log("Public Info:", publicInfo);
console.log("New User:", newUser);
console.log("Frozen User:", frozenUser);

/**
 * 補足:
 * TS の Utility Types は、DRY（Don't Repeat Yourself）原則を守るために不可欠です。
 * 1つの User 型から、用途に合わせて何種類もの型を「計算」で作ることができます。
 * 
 * 実行コマンド:
 * npx tsx ts_study/13_utility_types/index.ts
 */
