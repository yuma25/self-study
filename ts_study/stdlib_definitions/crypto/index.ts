import * as crypto from 'crypto';

/**
 * =============================================================================
 * TypeScript (Node.js) 標準モジュール定義集: crypto
 * Go の crypto パッケージ群に相当する機能を提供します。
 * =============================================================================
 */

// --- 1. ハッシュ計算 (SHA-256等) ---

// 使い方: const hash = crypto.createHash('sha256').update(data).digest('hex'); と書き，ハッシュ値を計算する．
export function MyHash(data: string): string {
    return crypto.createHash('sha256').update(data).digest('hex');
}

// --- 2. メッセージ認証符号 (HMAC) ---

// 使い方: const hmac = crypto.createHmac('sha256', key).update(data).digest('hex'); と書き，改ざん検知用の署名を作る．
export function MyHmac(key: string, data: string): string {
    return crypto.createHmac('sha256', key).update(data).digest('hex');
}

// --- 3. 安全な乱数生成 ---

// 使い方: const buf = crypto.randomBytes(32); と書き，暗号論的に安全な乱数を生成する．
export function MyRandomBytes(size: number): Buffer {
    return crypto.randomBytes(size);
}

// 使い方: const uuid = crypto.randomUUID(); と書き，一意な識別子（UUID v4）を生成する．
export function MyRandomUUID(): string {
    return crypto.randomUUID();
}

// --- 4. 鍵派生 (パスワードハッシュ等) ---

// 使い方: crypto.pbkdf2(password, salt, 100000, 64, 'sha512', (err, key) => { ... }); と書き，パスワードから安全な鍵を生成する．
export function MyPbkdf2(password: string, salt: string, iterations: number, keylen: number, digest: string, callback: (err: Error | null, derivedKey: Buffer) => void): void {
    crypto.pbkdf2(password, salt, iterations, keylen, digest, callback);
}

// 使い方: crypto.scrypt(password, salt, 64, (err, key) => { ... }); と書き，メモリ消費量を調整可能な安全なハッシュを生成する．
export function MyScrypt(password: string, salt: string, keylen: number, callback: (err: Error | null, derivedKey: Buffer) => void): void {
    crypto.scrypt(password, salt, keylen, callback);
}

// --- 5. 電子署名 (公開鍵暗号) ---

// 使い方: const sign = crypto.createSign('SHA256'); sign.update(data); const signature = sign.sign(privateKey); と書き，データに署名する．
export function MySign(data: string, privateKey: string): Buffer {
    const sign = crypto.createSign('SHA256');
    sign.update(data);
    return sign.sign(privateKey);
}

// 使い方: const verify = crypto.createVerify('SHA256'); verify.update(data); const isValid = verify.verify(publicKey, signature); と書き，署名を検証する．
export function MyVerify(data: string, publicKey: string, signature: Buffer): boolean {
    const verify = crypto.createVerify('SHA256');
    verify.update(data);
    return verify.verify(publicKey, signature);
}

// --- 6. 公開鍵による暗号化・復号 ---

// 使い方: const encrypted = crypto.publicEncrypt(publicKey, Buffer.from(text)); と書き，公開鍵で暗号化する．
export function MyPublicEncrypt(publicKey: string, buffer: Buffer): Buffer {
    return crypto.publicEncrypt(publicKey, buffer);
}

// 使い方: const decrypted = crypto.privateDecrypt(privateKey, encrypted); と書き，秘密鍵で復号する．
export function MyPrivateDecrypt(privateKey: string, buffer: Buffer): Buffer {
    return crypto.privateDecrypt(privateKey, buffer);
}

// 使い方: crypto.generateKeyPair('rsa', { modulusLength: 2048 }, (err, pub, priv) => { ... }); と書き，公開鍵・秘密鍵ペアを生成する．
export function MyGenerateKeyPair(type: 'rsa' | 'ed25519', options: any, callback: (err: Error | null, publicKey: string, privateKey: string) => void): void {
    crypto.generateKeyPair(type as any, options, callback);
}

// --- 7. 暗号化・復号 (AES-GCM) ---

// 使い方: const cipher = crypto.createCipheriv('aes-256-gcm', key, iv); と書き，暗号化オブジェクトを作る．
export function MyEncrypt(text: string, key: Buffer, iv: Buffer): Buffer {
    const cipher = crypto.createCipheriv('aes-256-gcm', key, iv);
    return Buffer.concat([cipher.update(text, 'utf8'), cipher.final()]);
}

// 使い方: const decipher = crypto.createDecipheriv('aes-256-gcm', key, iv); decipher.setAuthTag(tag); と書き，復号化オブジェクトを作る．
export function MyDecrypt(ciphertext: Buffer, key: Buffer, iv: Buffer, authTag: Buffer): string {
    const decipher = crypto.createDecipheriv('aes-256-gcm', key, iv);
    decipher.setAuthTag(authTag);
    return decipher.update(ciphertext) + decipher.final('utf8');
}

// --- 5. 定数時間比較 (タイミング攻撃対策) ---

// 使い方: if (crypto.timingSafeEqual(a, b)) { ... } と書き，実行時間の差から情報を漏洩させずに比較する．
export function MyTimingSafeEqual(a: Buffer, b: Buffer): boolean {
    return crypto.timingSafeEqual(a, b);
}

console.log('Loaded: crypto spec.');
