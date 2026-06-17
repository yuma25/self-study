import * as os from 'os';

/**
 * =============================================================================
 * TypeScript (Node.js) 標準モジュール定義集: os
 * オペレーティングシステムに関連するユーティリティ関数と定数を提供します。
 * Go の os パッケージの一部や runtime パッケージに相当します。
 * =============================================================================
 */

// --- 1. システム基本情報 ---

// 使い方: const p = os.platform(); と書き，'linux', 'darwin', 'win32' 等のOSプラットフォーム名を取得する．
export function MyPlatform(): string { return os.platform(); }

// 使い方: const a = os.arch(); と書き，'x64', 'arm64' 等のCPUアーキテクチャを取得する．
export function MyArch(): string { return os.arch(); }

// 使い方: const v = os.version(); と書き，OSのカーネルバージョン文字列を取得する．
export function MyVersion(): string { return os.version(); }

// 使い方: const t = os.type(); と書き，OSの名前（'Linux', 'Darwin', 'Windows_NT'）を取得する．
export function MyType(): string { return os.type(); }

// --- 2. ハードウェア情報 ---

// 使い方: const cpus = os.cpus(); と書き，各論理CPUコアのモデルや速度，実行時間の詳細を取得する．
export function MyCpus(): os.CpuInfo[] { return os.cpus(); }

// 使い方: const total = os.totalmem(); と書き，システムの全物理メモリ（バイト単位）を取得する．
export function MyTotalMem(): number { return os.totalmem(); }

// 使い方: const free = os.freemem(); と書き，システムの空き物理メモリ（バイト単位）を取得する．
export function MyFreeMem(): number { return os.freemem(); }

// 使い方: const load = os.loadavg(); と書き，[1分, 5分, 15分] の平均負荷（ロードアベレージ）を取得する（Windowsでは常に[0, 0, 0]）．
export function MyLoadAvg(): number[] { return os.loadavg(); }

// --- 3. ネットワーク・ユーザー情報 ---

// 使い方: const nets = os.networkInterfaces(); と書き，ネットワークインターフェースごとのIPアドレス等の情報を取得する．
export function MyNetworkInterfaces(): NodeJS.Dict<os.NetworkInterfaceInfo[]> { return os.networkInterfaces(); }

// 使い方: const user = os.userInfo(); と書き，現在の実行ユーザーのUID, GID, ホームディレクトリ等の情報を取得する．
export function MyUserInfo(): os.UserInfo<string> { return os.userInfo(); }

// 使い方: const host = os.hostname(); と書き，システムのホスト名を取得する．
export function MyHostname(): string { return os.hostname(); }

// --- 4. パス・ディレクトリ情報 ---

// 使い方: const home = os.homedir(); と書き，現在のユーザーのホームディレクトリを取得する．
export function MyHomedir(): string { return os.homedir(); }

// 使い方: const tmp = os.tmpdir(); と書き，OSのデフォルトの一時ディレクトリ（/tmp等）を取得する．
export function MyTmpdir(): string { return os.tmpdir(); }

// --- 5. 実行状態 ---

// 使い方: const up = os.uptime(); と書き，システムが起動してからの秒数を取得する．
export function MyUptime(): number { return os.uptime(); }

// 使い方: const eol = os.EOL; と書き，OS固有の改行コード（Linuxは\n, Windowsは\r\n）を取得する．
export const MyEOL = os.EOL;

// 使い方: const priority = os.getPriority(pid); と書き，プロセスの実行優先度を取得する．
export function MyGetPriority(pid: number = 0): number { return os.getPriority(pid); }

// 使い方: os.setPriority(pid, priority); と書き，プロセスの実行優先度（-20〜19）を設定する．
export function MySetPriority(pid: number, priority: number): void { os.setPriority(pid, priority); }

// 使い方: const devNull = os.devNull; と書き，プラットフォーム固有のヌルデバイスパス（/dev/null や \\.\nul）を取得する．
export const MyDevNull = os.devNull;

// --- 6. 定数 (シグナル・エラー・定数) ---

export const MyOSConstants = {
    signals: os.constants.signals, // SIGINT, SIGTERM 等
    errno: os.constants.errno,     // EACCES, EADDRINUSE 等
    priority: os.constants.priority, // PRIORITY_LOW, PRIORITY_HIGHEST 等
};

console.log('Loaded: os spec.');
