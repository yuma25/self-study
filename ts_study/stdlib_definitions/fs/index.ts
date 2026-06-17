import * as fs from 'fs';
import * as fsp from 'fs/promises';

/**
 * =============================================================================
 * TypeScript (Node.js) 標準モジュール定義集: fs (FileSystem)
 * Go の os, io, path/filepath に相当する機能を提供します。
 * =============================================================================
 */

// --- 1. 非同期 (Promiseベース - 推奨) ---

// 使い方: const data = await fs.promises.readFile('file.txt', 'utf-8'); と書き，ファイルを読み込む．
// Go の os.ReadFile に相当．
export async function MyReadFile(path: string): Promise<string> {
    return await fsp.readFile(path, 'utf-8');
}

// 使い方: await fs.promises.writeFile('file.txt', 'hello'); と書き，ファイルに書き込む（上書き）．
// Go の os.WriteFile に相当．
export async function MyWriteFile(path: string, data: string): Promise<void> {
    await fsp.writeFile(path, data);
}

// 使い方: await fs.promises.appendFile('file.txt', 'more data'); と書き，追記する．
export async function MyAppendFile(path: string, data: string): Promise<void> {
    await fsp.appendFile(path, data);
}

// 使い方: await fs.promises.mkdir('dir/subdir', { recursive: true }); と書き，ディレクトリを作成する．
// Go の os.MkdirAll に相当．
export async function MyMkdir(path: string): Promise<void> {
    await fsp.mkdir(path, { recursive: true });
}

// 使い方: const stats = await fs.promises.stat('file.txt'); if (stats.isFile()) { ... } と書き，ファイル情報を取得する．
// Go の os.Stat に相当．
export async function MyStat(path: string): Promise<fs.Stats> {
    return await fsp.stat(path);
}

// 使い方: await fs.promises.unlink('file.txt'); と書き，ファイルを削除する．
// Go の os.Remove に相当．
export async function MyRemoveFile(path: string): Promise<void> {
    await fsp.unlink(path);
}

// 使い方: const files = await fs.promises.readdir('dir'); と書き，ディレクトリ内のファイル一覧を取得する．
export async function MyReadDir(path: string): Promise<string[]> {
    return await fsp.readdir(path);
}

// 使い方: await fs.promises.rename('old.txt', 'new.txt'); と書き，ファイル名の変更や移動を行う．
export async function MyRename(oldPath: string, newPath: string): Promise<void> {
    await fsp.rename(oldPath, newPath);
}

// 使い方: await fs.promises.access('file.txt', fs.constants.R_OK); と書き，ファイルの読み書き権限を確認する．
export async function MyAccess(path: string, mode?: number): Promise<void> {
    await fsp.access(path, mode);
}

// 使い方: await fs.promises.copyFile('src.txt', 'dest.txt'); と書き，ファイルをコピーする．
export async function MyCopyFile(src: string, dest: string): Promise<void> {
    await fsp.copyFile(src, dest);
}

// 使い方: await fs.promises.chmod('file.txt', 0o755); と書き，ファイルの権限（パーミッション）を変更する．
export async function MyChmod(path: string, mode: number): Promise<void> {
    await fsp.chmod(path, mode);
}

// 使い方: await fs.promises.chown('file.txt', uid, gid); と書き，ファイルの所有者（UID/GID）を変更する．
export async function MyChown(path: string, uid: number, gid: number): Promise<void> {
    await fsp.chown(path, uid, gid);
}

// 使い方: await fs.promises.utimes('file.txt', atime, mtime); と書き，ファイルのアクセス時刻と修正時刻を変更する．
export async function MyUtimes(path: string, atime: number | Date, mtime: number | Date): Promise<void> {
    await fsp.utimes(path, atime, mtime);
}

// 使い方: const realPath = await fs.promises.realpath('symlink.txt'); と書き，シンボリックリンクのリンク先の実パスを取得する．
export async function MyRealpath(path: string): Promise<string> {
    return await fsp.realpath(path);
}

// 使い方: await fs.promises.symlink('target.txt', 'link.txt'); と書き，シンボリックリンクを作成する．
export async function MySymlink(target: string, path: string): Promise<void> {
    await fsp.symlink(target, path);
}

// 使い方: const target = await fs.promises.readlink('link.txt'); と書き，リンクの内容を読み取る．
export async function MyReadlink(path: string): Promise<string> {
    return await fsp.readlink(path);
}

// 使い方: await fs.promises.mkdtemp(path.join(os.tmpdir(), 'foo-')); と書き，一意なテンポラリディレクトリを作成する．
export async function MyMkdtemp(prefix: string): Promise<string> {
    return await fsp.mkdtemp(prefix);
}

// 使い方: const watcher = fs.watch('dir', (eventType, filename) => { ... }); と書き，ファイルやディレクトリの変更を監視する．
export function MyWatch(path: string, listener: (event: string, filename: string | null) => void): fs.FSWatcher {
    return fs.watch(path, listener);
}

// --- 2. ストリーム (大容量データ用) ---

// 使い方: const rs = fs.createReadStream('large.txt'); rs.on('data', (chunk) => { ... }); と書き，逐次読み込む．
// Go の os.Open と io.Reader の組み合わせに相当．
export function MyReadStream(path: string): fs.ReadStream {
    return fs.createReadStream(path);
}

// 使い方: const ws = fs.createWriteStream('out.txt'); ws.write('chunk'); と書き，逐次書き込む．
// Go の os.Create と io.Writer の組み合わせに相当．
export function MyWriteStream(path: string): fs.WriteStream {
    return fs.createWriteStream(path);
}

// --- 3. 同期 (Sync - 特殊な場合のみ) ---

// 使い方: const data = fs.readFileSync('file.txt', 'utf-8'); と書き，処理を止めて読み込む．
// CLIツール等の初期化時以外は非推奨．
export function MyReadFileSync(path: string): string {
    return fs.readFileSync(path, 'utf-8');
}

// 使い方: const exists = fs.existsSync('file.txt'); と書き，存在確認を行う．
// 非同期版がない（推奨されない）珍しいメソッド．
export function MyExistsSync(path: string): boolean {
    return fs.existsSync(path);
}

console.log('Loaded: fs spec.');
