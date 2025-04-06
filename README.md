# Kurosawa

[![Go Reference](https://pkg.go.dev/badge/github.com/sweetfish/kurosawa-go.svg)](https://pkg.go.dev/github.com/sweetfish/kurosawa-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/sweetfish/kurosawa-go)](https://goreportcard.com/report/github.com/sweetfish/kurosawa-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

KurosawaはGoで書かれた画面録画・マルチメディア編集ライブラリです。FFmpegをバックエンドとして使用し、シンプルで使いやすいAPIを提供します。

## プロジェクト構造

```
kurosawa-go/
├── cmd/                   # コマンドラインツール
├── examples/             # 使用例
│   └── simple/          # 基本的な使用例
├── internal/            # 内部パッケージ
│   ├── config/         # グローバル設定
│   ├── errors/         # エラー処理
│   ├── ffmpeg/         # FFmpeg実行
│   ├── filter/         # フィルター処理
│   ├── logging/        # ロギング
│   ├── metrics/        # メトリクス収集
│   └── middleware/     # ミドルウェア
└── pkg/                # 公開パッケージ
    ├── editor/         # 動画編集
    │   ├── clip/      # クリップ管理
    │   ├── effect/    # エフェクト
    │   ├── film/      # 動画構成
    │   ├── media/     # メディア処理
    │   ├── pipeline/  # 処理パイプライン
    │   └── validator/ # 入力検証
    ├── progress/      # 進捗管理
    ├── recorder/      # 画面録画
    └── resource/      # リソース管理
```

## 技術仕様

### コアコンポーネント

1. FFmpegコマンド実行 (`internal/ffmpeg`)
- コマンドビルダーパターンによるFFmpeg操作
- コンテキストベースのプロセス管理
- フィルターグラフの動的生成
- 入出力ストリーム制御

2. メディア処理パイプライン (`pkg/editor/pipeline`)
- ステージベースの処理フロー
- フレーム単位の処理
- 非同期処理のサポート
- エラーハンドリング

3. エフェクトシステム (`pkg/editor/effect`)
- プラグイン可能なエフェクトインターフェース
- 複数エフェクトの合成
- パラメータのバリデーション
- FFmpegフィルター変換

4. マルチプラットフォーム録画 (`pkg/recorder`)
- OS別の最適な録画方式
  - Windows: gdigrab
  - macOS: avfoundation
  - Linux: x11grab
- 設定可能なパラメータ
  - フレームレート
  - 録画領域
  - 品質設定

### データモデル

1. クリップ (`pkg/editor/clip`)
```go
type Clip struct {
    source   MediaSource    // メディアソース
    start    time.Duration  // 開始時間
    duration time.Duration  // 長さ
    position Position       // 位置
    effects  []Effect      // エフェクト
}
```

2. フィルム (`pkg/editor/film`)
```go
type Film struct {
    clips    []*Clip       // クリップ配列
    width    int          // 幅
    height   int          // 高さ
    duration time.Duration // 全体の長さ
}
```

3. メディアフレーム (`pkg/editor/media`)
```go
type Frame struct {
    Data      []byte         // フレームデータ
    Timestamp time.Duration  // タイムスタンプ
    Width     int           // 幅
    Height    int           // 高さ
    Format    string        // フォーマット
}
```

### エラー処理戦略

1. エラー型 (`internal/errors`)
```go
type KurosawaError struct {
    Code    ErrorCode  // エラーコード
    Message string    // エラーメッセージ
    Err     error    // 元のエラー
}
```

2. エラーコード体系
- ErrCodeInternal: 内部エラー
- ErrCodeInvalidInput: 入力エラー
- ErrCodeFFmpeg: FFmpeg関連エラー
- ErrCodeIO: I/Oエラー

### リソース管理

1. リソースマネージャ (`pkg/resource`)
- I/Oリソースの追跡
- コンテキストベースのライフサイクル管理
- 確実なクリーンアップ

2. メトリクス収集 (`internal/metrics`)
- 処理時間の計測
- エラー頻度の追跡
- FFmpegコマンドの監視

### 進捗管理システム

1. 進捗レポート (`pkg/progress`)
```go
type Progress struct {
    Percent    float64        // 進捗率
    Stage      string         // 現在のステージ
    Elapsed    time.Duration  // 経過時間
    Remaining  time.Duration  // 残り時間
    Error      error         // エラー情報
}
```

2. レポートオプション
- 更新間隔の設定
- チャネルベースの通知
- カスタムレポーター

### バリデーション

1. ファイルバリデーション (`pkg/editor/validator`)
- パスの正規化
- 存在確認
- ファイルタイプチェック

2. 時間範囲バリデーション
- 開始・終了時間の検証
- 範囲の整合性チェック

## 拡張ポイント

1. エフェクトプラグイン
- 新しいエフェクトの追加
- カスタムフィルターの実装
- パラメータの定義

2. メディアソース
- 新しいメディアタイプの追加
- カスタムデコーダーの実装
- メタデータの拡張

3. プログレスレポーター
- UIフレームワークとの統合
- カスタム進捗表示
- ログ出力の拡張

## 開発ガイドライン

### コードスタイル

1. インターフェース設計
- 小さく焦点を絞ったインターフェース
- コンポーザブルなコンポーネント
- 明確な責任範囲

2. エラー処理
- 詳細なエラーコンテキスト
- エラーのラッピング
- 回復可能なエラーの区別

3. 並行処理
- コンテキストベースのキャンセル
- リソースリークの防止
- 適切なゴルーチン管理

### テスト戦略

1. ユニットテスト
- モック可能なインターフェース
- テーブル駆動テスト
- エラーケースのカバレッジ

2. 統合テスト
- FFmpeg連携のテスト
- OS別の動作確認
- リソース管理の検証

### パフォーマンス最適化

1. メモリ管理
- フレームバッファのプーリング
- 適切なバッチサイズ
- GCプレッシャーの軽減

2. 並列処理
- パイプラインの並列化
- 効率的なチャネル使用
- コンテキスト伝播

## 将来の拡張計画

### 短期目標
1. オーディオ処理の強化
2. ハードウェアアクセラレーション
3. WebAssembly対応

### 中期目標
1. ストリーミング出力
2. リアルタイムプレビュー
3. プラグインシステム

### 長期目標
1. 分散処理対応
2. AIベースの自動編集
3. クラウドサービス統合

## トラブルシューティング

### よくある問題
1. FFmpegが見つからない
   - PATHの確認
   - FFmpegのインストール確認

2. 録画が開始されない
   - OS別の権限確認
   - デバイスアクセス権の確認

3. メモリ使用量が高い
   - フィルター設定の最適化
   - 適切なバッファサイズの設定

## ライセンスと注意事項

- MITライセンス
- FFmpegライセンスの遵守必須
- 商用利用時の注意点

## コントリビューションガイド

1. Issue作成
2. ブランチ作成
3. テスト追加
4. PRの作成

各OSでの動作確認とテストの実行を推奨します。

## サポートするOS

- Windows
- macOS
- Linux (X11)

各OSで適切な画面キャプチャ方式を自動的に選択します。
