package main

import (
	"context"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"time"
)

/*
   =============================================================================
   Go 標準ライブラリ 定義集: net/http パッケージ (実戦コード例配置調整版)
   ※ 型衝突を避けるため，デモ用の型名には My 接頭辞を付けている．
   ※ 解説は「である調」とし，句読点は「．」「，」を使用している．
   ※ 構造体のフィールド以外は，コードの上部に使い方を記載している．
   =============================================================================
*/

// --- 1. インターフェース (規格) ---

// 使い方: 自作構造体にこのメソッドを作り，http.Handle("/", myHandler) と登録する．
type MyHandler interface {
	ServeHTTP(w MyResponseWriter, r *MyRequest)
}

// 使い方:
// Header: w.Header().Set("Content-Type", "application/json") と書き，レスポンスの種類を教える．
// Write: w.Write([]byte(`{"ok": true}`)) と書き，生のデータを直接クライアントに流し込む．
// WriteHeader: w.WriteHeader(http.StatusNotFound) と書き，404等のステータスを確定させる．
type MyResponseWriter interface {
	Header() http.Header
	Write([]byte) (int, error)
	WriteHeader(statusCode int)
}

// 使い方: f, ok := w.(http.Flusher); if ok { f.Flush() } と書き，バッファを強制送信してストリーミングを実現する．
type MyFlusher interface {
	Flush()
}

// 使い方: conn, _, _ := w.(http.Hijacker).Hijack() と書き，WebSocket通信等へ切り替える．
type MyHijacker interface {
	Hijack() (net.Conn, *context.Context, error)
}

// 使い方: w.(http.Pusher).Push("/style.css", nil) と書き，HTMLより先にCSSを送る．
type MyPusher interface {
	Push(target string, opts *http.PushOptions) error
}

// 使い方: client.Transport = myRoundTripper と書き，通信に独自の認証やログを挟む．
type MyRoundTripper interface {
	RoundTrip(req *MyRequest) (resp *MyResponse, err error)
}

// 使い方:
// SetCookies: 通常は直接呼ばず，http.NewCookieJar() の戻り値をClientに渡して自動管理させる．
// Cookies: 同上．リクエスト送信時に，保存されたクッキーが自動で引き出される．
type MyCookieJar interface {
	SetCookies(u *url.URL, cookies []*http.Cookie)
	Cookies(u *url.URL) []*http.Cookie
}

// --- 2. 構造体 (実体) とそのメソッド ---

type MyRequest struct {
	Method           string              // 使い方: if r.Method == http.MethodPost { ... } と書き，操作の種類で処理を分ける．
	URL              *url.URL            // 使い方: id := r.URL.Query().Get("id") と書き，?id=123 のようなクエリを取得する．
	Proto            string              // 使い方: if r.Proto == "HTTP/2.0" { ... } と書き，プロトコルのバージョンを判定する．
	Header           http.Header         // 使い方: token := r.Header.Get("Authorization") と書き，認証情報を読み取る．
	Body             io.ReadCloser       // 使い方: data, _ := io.ReadAll(r.Body); defer r.Body.Close() と書き，送信されたデータを安全に読む．
	ContentLength    int64               // 使い方: if r.ContentLength > 1024 { ... } と書き，巨大なファイルのアップロードを弾く．
	TransferEncoding []string            // 使い方: for _, enc := range r.TransferEncoding { ... } と書き，チャンク転送等の状態を確認する．
	Host             string              // 使い方: if r.Host == "api.example.com" { ... } と書き，ドメインで処理を分ける．
	Form             url.Values          // 使い方: r.ParseForm(); name := r.Form.Get("name") と書き，送信されたフォーム値を取得する．
	PostForm         url.Values          // 使い方: 上記と同様だが，URLパラメータを無視して純粋なPOSTボディだけを見たい時に使う．
	MultipartForm    *multipart.Form     // 使い方: r.ParseMultipartForm(32<<20); file := r.MultipartForm.File["img"] と書き，ファイルを受け取る．
	RemoteAddr       string              // 使い方: ip := strings.Split(r.RemoteAddr, ":")[0] と書き，アクセス元のIPをログに残す．
	RequestURI       string              // 使い方: proxyReq.URL, _ = url.Parse(r.RequestURI) と書き，プロキシサーバーで生のパスを転送する．
	TLS              *http.Header        // 使い方: if r.TLS != nil { ... } と書き，通信がHTTPS(暗号化)されているか確認する．
}

// 使い方: cookie, err := r.Cookie("session") と書き，セッションIDを得る．
func (r *MyRequest) Cookie(name string) (*http.Cookie, error) { return nil, nil }

// 使い方: for _, c := range r.Cookies() { fmt.Println(c.Name) } と一覧を見る．
func (r *MyRequest) Cookies() []*http.Cookie { return nil }

// 使い方: if strings.Contains(r.UserAgent(), "Mobile") { ... } と判定する．
func (r *MyRequest) UserAgent() string { return "" }

// 使い方: req = req.WithContext(ctx) と書き，タイムアウト設定を紐付ける．
func (r *MyRequest) WithContext(ctx context.Context) *MyRequest { return nil }

type MyResponse struct {
	Status     string        // 使い方: fmt.Println("Result:", resp.Status) と書き，"200 OK" をログに出す．
	StatusCode int           // 使い方: if resp.StatusCode >= 400 { return fmt.Errorf("error") } とエラーハンドリングする．
	Header     http.Header   // 使い方: contentType := resp.Header.Get("Content-Type") と書き，返ってきたデータの種類を知る．
	Body       io.ReadCloser // 使い方: io.Copy(os.Stdout, resp.Body); defer resp.Body.Close() と書き，結果を画面に垂れ流す．
}

type MyClient struct {
	Transport MyRoundTripper // 使い方: client := &http.Client{Transport: myTrans} と書き，プロキシ通信等を実現する．
	Jar       MyCookieJar    // 使い方: client := &http.Client{Jar: jar} と書き，ブラウザのようにログイン状態を保持させる．
	Timeout   time.Duration  // 使い方: client := &http.Client{Timeout: 5 * time.Second} と書き，無限に応答を待つのを防ぐ．
}

// 使い方: resp, err := client.Do(req) と書き，詳細設定したリクエストを放つ．
func (c *MyClient) Do(req *MyRequest) (*MyResponse, error) { return nil, nil }

type MyServer struct {
	Addr              string        // 使い方: srv := &http.Server{Addr: ":8080"} と書き，ポート8080で待ち受ける設定をする．
	Handler           MyHandler     // 使い方: srv.Handler = mux と書き，自作のルーターにすべての処理を任せる．
	ReadTimeout       time.Duration // 使い方: srv.ReadTimeout = 10 * time.Second と書き，遅いクライアントから防御する．
}

// 使い方: log.Fatal(srv.ListenAndServe()) と書き，エラーで落ちるまでサーバーを回し続ける．
func (srv *MyServer) ListenAndServe() error { return nil }

// --- 3. パッケージレベルの関数 (操作) ---

// 使い方: http.Handle("/images/", http.StripPrefix("/images/", fileServer)) と書く．
func Handle(pattern string, handler http.Handler) {}

// 使い方: http.HandleFunc("/ping", func(w, r) { w.Write([]byte("pong")) }) と書く．
func HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {}

// 使い方: http.ListenAndServe(":80", nil) と書き，手っ取り早く起動する．
func ListenAndServe(addr string, handler http.Handler) error{ return nil }

// 使い方: http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil) と書く．
func ListenAndServeTLS(addr, certFile, keyFile string, handler http.Handler) error { return nil }

// 使い方: mux := http.NewServeMux(); mux.HandleFunc(...) と独立したルーターを作る．
func NewServeMux() *http.ServeMux { return nil }

// 使い方: http.ServeFile(w, r, "index.html") と書き，静的ファイルを返す．
func ServeFile(w http.ResponseWriter, r *http.Request, name string) {}

// 使い方: 動画のシーク再生（途中から再生）等に対応させるために使う．
func ServeContent(w http.ResponseWriter, req *http.Request, name string, modtime time.Time, content io.ReadSeeker) {}

// 使い方: resp, err := http.Get("https://api.github.com/") と書き，データを取る．
func Get(url string) (resp *http.Response, err error) { return nil, nil }

// 使い方: http.Post(url, "application/json", strings.NewReader(`{"id":1}`)) と書く．
func Post(url, contentType string, body io.Reader) (resp *http.Response, err error) { return nil, nil }

// 使い方: http.SetCookie(w, &http.Cookie{Name: "session", Value: "xyz"}) と書く．
func SetCookie(w http.ResponseWriter, cookie *http.Cookie) {}

// 使い方: http.Redirect(w, r, "/login", http.StatusSeeOther) と書き，画面遷移させる．
func Redirect(w http.ResponseWriter, r *http.Request, url string, code int) {}

// 使い方: http.Error(w, "Bad Request", http.StatusBadRequest) と書き，処理を弾く．
func Error(w http.ResponseWriter, error string, code int) {}

// --- 4. 特殊なアダプター型 ---

// 使い方: 関数を Handler インターフェースに適合させるための型である．
type MyHandlerFunc func(http.ResponseWriter, *http.Request)

// 使い方: 関数自身を呼び出すことで ServeHTTP メソッドを実装したことになる．
func (f MyHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) { f(w, r) }

func main() {
	println("Updated: net/http spec with top-level comments for funcs/interfaces and inline for structs.")
}
