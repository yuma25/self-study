import * as http from 'http';

/**
 * =============================================================================
 * TypeScript (Node.js) 標準モジュール定義集: http
 * Go の net/http パッケージに相当する機能を提供します。
 * =============================================================================
 */

// --- 1. サーバーの作成 ---

// 使い方: const server = http.createServer((req, res) => { ... }); server.listen(8080); と書き，HTTPサーバーを起動する．
export function MyCreateServer(): http.Server {
    return http.createServer((req: http.IncomingMessage, res: http.ServerResponse) => {
        // req: IncomingMessage のプロパティ
        const { method, url, headers } = req;
        console.log(`Method: ${method}, URL: ${url}`);

        // res: ServerResponse のメソッド
        res.setHeader('Content-Type', 'text/plain');
        res.writeHead(200);
        res.end('Hello World');
    });
}

// --- 2. 接続管理 (Agent) ---

// 使い方: const agent = new http.Agent({ keepAlive: true }); http.get({ agent, ... }); と書き，コネクションの再利用（Keep-Alive）を制御する．
export function MyAgent(): http.Agent {
    return new http.Agent({ keepAlive: true, maxSockets: 10 });
}

// --- 3. 重要なクラスと型 ---

// 使い方: req.method, req.url と書き，リクエスト情報（メソッド，パス，ヘッダー等）を参照する．
export type MyRequest = http.IncomingMessage;

// 使い方: res.writeHead(200), res.write('data'), res.end() と書き，レスポンスを構築して送信する．
export type MyResponse = http.ServerResponse;

// 使い方: http.get('http://example.com', (res) => { ... }); と書き，外部のURLに対してGETリクエストを送る．
export function MyGet(url: string, callback: (res: http.IncomingMessage) => void): http.ClientRequest {
    return http.get(url, callback);
}

// 使い方: const req = http.request({ method: 'POST', host: '...' }, (res) => { ... }); req.write(data); req.end(); 
// と書き，詳細な設定でHTTPリクエストを送信する．
export function MyRequest(options: http.RequestOptions, callback?: (res: http.IncomingMessage) => void): http.ClientRequest {
    return http.request(options, callback);
}

// 使い方: req.setTimeout(5000, () => req.destroy()); と書き，リクエストにタイムアウトを設定する．
export function MySetTimeout(req: http.ClientRequest, timeout: number, callback?: () => void): void {
    req.setTimeout(timeout, callback);
}

console.log('Loaded: http spec.');
