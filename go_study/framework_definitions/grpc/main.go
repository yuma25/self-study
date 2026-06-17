package main

/*
   =============================================================================
   Go フレームワーク 定義集: gRPC
   Google 製の高性能な RPC フレームワーク．Protocol Buffers を使用する．
   =============================================================================
*/

// --- 1. サーバーの実装 ---

// 使い方: s := grpc.NewServer() と書き，gRPC サーバーインスタンスを作成する．
type MyGRPCServer struct { /* 非公開 */ }

func (s *MyGRPCServer) Serve(l any) error { return nil } // 使い方: s.Serve(lis) で起動．

// 使い方: pb.RegisterUserServiceServer(s, &server{}) と書き，生成されたコードに自作の実装を登録する．

// --- 2. ストリーミングとインターセプター ---

// 使い方: grpc.UnaryInterceptor(MyInterceptor) と書き，全てのメソッドに共通のログや認証処理を挟む．
type MyInterceptor func(any, any, any, any) (any, error)

// 使い方: rpc GetChat(stream ChatReq) returns (stream ChatRes); と書き，双方向ストリーミングを実現する．
type MyStream interface {
	Send(any) error
	Recv() (any, error)
}

// 使い方: md := metadata.Pairs("key", "val"); ctx = metadata.NewOutgoingContext(ctx, md) と書き，HTTPヘッダーのようなメタデータを送受信する．

// --- 3. 運用・保守 ---

// 使い方: healthpb.RegisterHealthServer(s, health.NewServer()) と書き，Kubernetes 等からの生存確認（Liveness/Readiness Probe）に対応させる．
// 使い方: reflection.Register(s) と書き，gRPCurl 等の外部ツールからサービス定義を動的に取得できるようにする．
// 使い方: ctx, cancel := context.WithTimeout(ctx, time.Second); defer cancel() と書き，クライアント側でリクエストの制限時間（Deadline）を設ける．

// --- 4. クライアントの実装 ---

// 使い方: conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure()) と書き，サーバーに接続する．
type MyClientConn struct { /* 非公開 */ }

// 使い方: client := pb.NewUserServiceClient(conn) と書き，クライアントを作成する．

// --- 3. プロトコル定義 (proto) ---
/*
   service UserService {
       rpc GetUser (UserRequest) returns (UserResponse);
   }
   message UserRequest { string id = 1; }
   などの定義を .proto ファイルに記述し，protoc で Go コードを生成して使用する．
*/

func main() {
	println("Loaded: gRPC definition.")
}
