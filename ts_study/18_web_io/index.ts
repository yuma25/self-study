import * as http from 'http';
import * as fs from 'fs';

/**
 * TypeScript (Node.js): Streams and Pipes
 * Go の io.Copy に相当するのが .pipe() です。
 */

const LOG_FILE = 'server_ts.log';

// 1. ログファイルストリームの準備
const logStream = fs.createWriteStream(LOG_FILE, { flags: 'a' });

const server = http.createServer((req: http.IncomingMessage, res: http.ServerResponse) => {
    const timestamp = new Date().toISOString();
    const method = req.method ?? 'UNKNOWN';
    const url = req.url ?? 'UNKNOWN';
    const logMessage = `[${timestamp}] Request: ${method} ${url}\n`;

    // 2. 画面(Stdout)とファイルの両方に書き込み
    process.stdout.write(logMessage);
    logStream.write(logMessage);

    if (url === '/log') {
        // 3. ファイルストリーミング
        // createReadStream でファイルを読み込み、res (書き込み先) へ流す
        const readStream = fs.createReadStream(LOG_FILE);
        
        // ストリームの読み込みエラー対策
        readStream.on('error', (err) => {
            console.error('Streaming error:', err);
            if (!res.headersSent) {
                res.writeHead(500, { 'Content-Type': 'text/plain' });
                res.end('Internal Server Error');
            }
        });

        res.writeHead(200, { 'Content-Type': 'text/plain' });
        
        // Go の io.Copy(res, readStream) と同じ役割
        readStream.pipe(res);
        return;
    }

    res.writeHead(200, { 'Content-Type': 'text/plain' });
    res.end('Welcome to Node.js Stream Server!\n');
});

const PORT = 8081;
server.listen(PORT, () => {
    console.log(`TS Server starting at :${PORT}...`);
    console.log(`- http://localhost:${PORT}/`);
    console.log(`- http://localhost:${PORT}/log`);
});

/**
 * 実行コマンド:
 * npx tsx ts_study/18_web_io/index.ts
 */
