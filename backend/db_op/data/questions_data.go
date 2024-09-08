package data // テスト用データ

import (
	"backend/models"

	"github.com/lib/pq"
)

var QuestionsList = []models.Questions{
	{Question: "悟りの第二段階の状態は？",
		Options:    pq.StringArray{"一来果", "予流果", "不還果", "阿羅漢果"}, // pq.StringArrayは、PostgreSQLのtext[]型（文字列の配列）を扱う
		Supplement: "欲界の煩悩を断じ終えた位のこと",
		Difficulty: 1},

	{Question: "四聖諦（ししょうたい）はどれか？",
		Options:    pq.StringArray{"苦諦", "楽諦", "幸諦", "怒諦"},
		Supplement: "四聖諦（ししょうたい）は、仏教における4つの基本的な真理で、苦諦（くたい）、集諦（じったい）、滅諦（めったい）、道諦（どうたい）",
		Difficulty: 1},

	{Question: "テーラワーダ仏教の「テーラワーダ」とは何を意味しますか？",
		Options:    pq.StringArray{"長老の教え", "仏陀の教え", "慈悲の教え", "瞑想の教え"},
		Supplement: "テーラワーダ」はパーリ語で「長老の教え」を意味します。",
		Difficulty: 1},

	{Question: "初期仏教の教えの目的は何ですか？",
		Options:    pq.StringArray{"苦しみをなくす", "儀式儀礼を行う", "信じる者は救われる", "現世利益を得る"},
		Supplement: "初期仏教の教えは「苦しみをなくす」ことを目的としています。",
		Difficulty: 1},

	{Question: "四聖諦の中で、仏教が特に重要視しているものは何ですか？",
		Options:    pq.StringArray{"苦（Dukkha）", "集（Samudaya）", "滅（Nirodha）", "道（Magga）"},
		Supplement: "仏教では、苦というものの認識を非常に重要視しています。",
		Difficulty: 1},

	{Question: "ヴィパッサナー瞑想法の目的は何ですか？",
		Options:    pq.StringArray{"現実を正しく見る", "心を空にする", "超能力を得る", "幸福を追求する"},
		Supplement: "ヴィパッサナー瞑想法は、自分の現実を鏡に写したように見ることができる最良の方法です。",
		Difficulty: 1},

	{Question: "心の働きの一部として説明されているものはどれですか？",
		Options:    pq.StringArray{"全て", "記憶", "感情", "計算力"},
		Supplement: "心の働きには、記憶、感情、計算力などが含まれます。",
		Difficulty: 1},

	{Question: "仏教における「無常」とは何を意味しますか？",
		Options:    pq.StringArray{"常に変化する", "永遠不滅", "変化しない", "固定されたもの"},
		Supplement: "無常とは、すべてのものが常に変化し続けることを意味します。",
		Difficulty: 1},

	{Question: "宗教が死後の世界について語る理由として説明されているものはどれですか？",
		Options:    pq.StringArray{"人間は死後の世界を信じたくないから", "宗教の教義を広めるため", "死後の世界を体験した人がいるから", "全て"},
		Supplement: "宗教が死後の世界について語る理由として、人間は死後の世界を信じたくないからと説明されています。",
		Difficulty: 1},

	{Question: "お釈迦さまが主に説いたことは何ですか？",
		Options:    pq.StringArray{"生き方について", "死後の世界について", "輪廻転生のシステムについて", "全て"},
		Supplement: "お釈迦さまは主に、我々がどう生きるべきかということについて説いていました。",
		Difficulty: 1},

	{Question: "四聖諦の中で、苦しみの原因を示すものはどれですか？",
		Options:    pq.StringArray{"集", "苦", "滅", "道"},
		Supplement: "四聖諦の中で、苦しみの原因を示すのは集です。",
		Difficulty: 1},

	{Question: "四聖諦の中で、苦しみの終わりを示すものはどれですか？",
		Options:    pq.StringArray{"滅", "苦", "集", "道"},
		Supplement: "四聖諦の中で、苦しみの終わりを示すのは滅です。",
		Difficulty: 1},

	{Question: "四苦八苦の一つである愛別離苦とは何を指しますか？",
		Options:    pq.StringArray{"愛しい人との別れ", "苦手な人との付き合い", "欲しいものが手に入らない", "生きること"},
		Supplement: "四苦八苦の一つである愛別離苦は愛しい人との別れです。",
		Difficulty: 2},

	{Question: "四苦八苦の一つである怨憎会苦とは何を指しますか？",
		Options:    pq.StringArray{"苦手な人との付き合い", "愛しい人との別れ", "欲しいものが手に入らない", "生きること"},
		Supplement: "四苦八苦の一つである怨憎会苦は苦手な人との付き合いです。",
		Difficulty: 2},

	{Question: "八正道の第一である正見とは何を意味しますか？",
		Options:    pq.StringArray{"正しく四聖諦を見きわめること", "正しい言葉づかい", "正しい行い", "正しい努力"},
		Supplement: "正見は、正しく四聖諦を見きわめることを意味します。",
		Difficulty: 1},

	{Question: "八正道の中で、正しい言葉づかいを示すものはどれですか？",
		Options:    pq.StringArray{"正語", "正見", "正業", "正命"},
		Supplement: "正語は、正しい言葉づかいを示します。",
		Difficulty: 1},

	{Question: "八正道の中で、正しい行いを示すものはどれですか？",
		Options:    pq.StringArray{"正業", "正見", "正語", "正命"},
		Supplement: "正業は、正しい行いを示します。",
		Difficulty: 1},

	{Question: "八正道の中で、正しい努力を示すものはどれですか？",
		Options:    pq.StringArray{"正精進", "正見", "正語", "正業"},
		Supplement: "正精進は、正しい努力を示します。",
		Difficulty: 1},

	{Question: "八正道の中で、正しい職業を示すものはどれですか？",
		Options:    pq.StringArray{"正命", "正見", "正語", "正業"},
		Supplement: "正命は、正しい職業を示します。",
		Difficulty: 1},

	{Question: "八正道の中で、正しい精神統一を示すものはどれですか？",
		Options:    pq.StringArray{"正定", "正見", "正語", "正業"},
		Supplement: "正定は、正しい精神統一を示します。",
		Difficulty: 1},
}
