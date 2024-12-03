package data // テスト用データ

import (
	"backend/src/models"

	"github.com/lib/pq"
)

var QuestionsListE = []models.Questions{
	{Question: "次のうち、画像認識において一般的に使用される畳み込みニューラルネットワークの構造はどれか？",
		Options:    pq.StringArray{"ResNet", "LSTM", "GRU", "トランスフォーマー"},
		Supplement: "ResNetは画像認識に一般的に使用される畳み込みニューラルネットワークの構造です。",
		Difficulty: 1},

	{Question: "次のうち、画像データの前処理としてよく使用される正規化手法はどれか？",
		Options:    pq.StringArray{"Batch Normalization", "標準化", "欠損値補完", "交差検証"},
		Supplement: "Batch Normalizationは画像データの前処理としてよく使用されます。",
		Difficulty: 1},

	{Question: "次のうち、ディープラーニングにおけるデータ拡張手法の一つはどれか？",
		Options:    pq.StringArray{"Random Flip", "正則化", "クロスエントロピー", "主成分分析"},
		Supplement: "Random Flipはデータ拡張手法の一つです。",
		Difficulty: 2},

	{Question: "次のうち、自己組織化マップ（SOM）の主な用途は何か？",
		Options:    pq.StringArray{"次元削減と視覚化", "分類", "クラスタリング", "予測"},
		Supplement: "自己組織化マップは次元削減と視覚化に使用されます。",
		Difficulty: 2},

	{Question: "ニューラルネットワークのトレーニングにおけるパラメータの初期化戦略はどれか？",
		Options:    pq.StringArray{"Xavier法", "SGD", "RMSprop", "Adam"},
		Supplement: "Xavier法はパラメータの初期化戦略です。",
		Difficulty: 3},

	{Question: "混合ガウス分布のパラメータ推定に使われるアルゴリズムはどれか？",
		Options:    pq.StringArray{"EMアルゴリズム", "最急降下法", "SGD", "モメンタム"},
		Supplement: "EMアルゴリズムは混合ガウス分布のパラメータ推定に使用されます。",
		Difficulty: 1},

	{Question: "LSTMの一部として使われるゲートの一つはどれか？",
		Options:    pq.StringArray{"忘却ゲート", "入力ゲート", "出力ゲート", "重み更新ゲート"},
		Supplement: "忘却ゲートはLSTMの一部として使用されます。",
		Difficulty: 3},

	{Question: "サポートベクターマシンにおけるカーネル関数として適切なものはどれか？",
		Options:    pq.StringArray{"RBFカーネル", "線形カーネル", "多項式カーネル", "シグモイドカーネル"},
		Supplement: "RBFカーネルはSVMでよく使用されるカーネル関数です。",
		Difficulty: 3},

	{Question: "サポートベクターマシンで使われるカーネル関数でないものはどれか？",
		Options:    pq.StringArray{"フーリエカーネル", "線形カーネル", "RBFカーネル", "ポリカーネル"},
		Supplement: "フーリエカーネルは一般的ではありません。",
		Difficulty: 1},

	{Question: "BERTが主に活用される分野はどれか？",
		Options:    pq.StringArray{"自然言語処理", "画像認識", "音声合成", "強化学習"},
		Supplement: "BERTは自然言語処理で活用されます。",
		Difficulty: 1},

	{Question: "WaveNetが生成するものはどれか？",
		Options:    pq.StringArray{"音声波形", "画像データ", "テキストデータ", "数値予測"},
		Supplement: "WaveNetは音声波形を生成します。",
		Difficulty: 1},

	{Question: "地域的なデータの正規化に適した手法はどれか？",
		Options:    pq.StringArray{"Batch Normalization", "Layer Normalization", "Instance Normalization", "Group Normalization"},
		Supplement: "Instance Normalizationは地域的な正規化に適しています。",
		Difficulty: 1},
}

// "次のうち、画像認識において一般的に使用される畳み込みニューラルネットワークの構造はどれか？","ResNet","LSTM","GRU","トランスフォーマー","ResNetは画像認識に一般的に使用される畳み込みニューラルネットワークの構造です。",1
// "次のうち、画像データの前処理としてよく使用される正規化手法はどれか？","Batch Normalization","標準化","欠損値補完","交差検証","Batch Normalizationは画像データの前処理としてよく使用されます。",1
// "次のうち、ディープラーニングにおけるデータ拡張手法の一つはどれか？","Random Flip","正則化","クロスエントロピー","主成分分析","Random Flipはデータ拡張手法の一つです。",2
// "次のうち、自己組織化マップ（SOM）の主な用途は何か？","次元削減と視覚化","分類","クラスタリング","予測","自己組織化マップは次元削減と視覚化に使用されます。",2
// "ニューラルネットワークのトレーニングにおけるパラメータの初期化戦略はどれか？","Xavier法","SGD","RMSprop","Adam","Xavier法はパラメータの初期化戦略です。",3
// "混合ガウス分布のパラメータ推定に使われるアルゴリズムはどれか？","EMアルゴリズム","最急降下法","SGD","モメンタム","EMアルゴリズムは混合ガウス分布のパラメータ推定に使用されます。",1
// "LSTMの一部として使われるゲートの一つはどれか？","忘却ゲート","入力ゲート","出力ゲート","重み更新ゲート","忘却ゲートはLSTMの一部として使用されます。",3
// "サポートベクターマシンにおけるカーネル関数として適切なものはどれか？","RBFカーネル","線形カーネル","多項式カーネル","シグモイドカーネル","RBFカーネルはSVMでよく使用されるカーネル関数です。",3
// "サポートベクターマシンで使われるカーネル関数でないものはどれか？","フーリエカーネル","線形カーネル","RBFカーネル","ポリカーネル","フーリエカーネルは一般的ではありません。",1
// "BERTが主に活用される分野はどれか？","自然言語処理","画像認識","音声合成","強化学習","BERTは自然言語処理で活用されます。",1
// "WaveNetが生成するものはどれか？","音声波形","画像データ","テキストデータ","数値予測","WaveNetは音声波形を生成します。",1
// "地域的なデータの正規化に適した手法はどれか？","Batch Normalization","Layer Normalization","Instance Normalization","Group Normalization","Instance Normalizationは地域的な正規化に適しています。",1
