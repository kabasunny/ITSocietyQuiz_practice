package data

import (
	"backend/models"

	"github.com/lib/pq"
)

var QuizDataList = []models.QuizData{
	{Question: "悟りの第二段階の状態は？",
		Options:    pq.StringArray{"予流果", "一来果", "不還果", "阿羅漢果"}, // pq.StringArrayは、PostgreSQLのtext[]型（文字列の配列）を扱う
		Correct:    "一来果",
		Supplement: "欲界の煩悩を断じ終えた位のこと"},
	{Question: "四聖諦（ししょうたい）はどれか？", Options: pq.StringArray{"苦諦", "楽諦", "幸諦", "怒諦"}, Correct: "苦諦", Supplement: "四聖諦（ししょうたい）は、仏教における4つの基本的な真理で、苦諦（くたい）、集諦（じったい）、滅諦（めったい）、道諦（どうたい）"},
	{Question: "テーラワーダ仏教の「テーラワーダ」とは何を意味しますか？", Options: pq.StringArray{"長老の教え", "仏陀の教え", "慈悲の教え", "瞑想の教え"}, Correct: "長老の教え", Supplement: "テーラワーダ」はパーリ語で「長老の教え」を意味します。"},
	{Question: "初期仏教の教えの目的は何ですか？", Options: pq.StringArray{"信じる者は救われる", "儀式儀礼を行う", "苦しみをなくす", "現世利益を得る"}, Correct: "苦しみをなくす", Supplement: "初期仏教の教えは「苦しみをなくす」ことを目的としています。"},
	{Question: "四聖諦の中で、仏教が特に重要視しているものは何ですか？", Options: pq.StringArray{"苦（Dukkha）", "集（Samudaya）", "滅（Nirodha）", "道（Magga）"}, Correct: "苦（Dukkha）", Supplement: "仏教では、苦というものの認識を非常に重要視しています。"},
	{Question: "ヴィパッサナー瞑想法の目的は何ですか？", Options: pq.StringArray{"心を空にする", "現実を正しく見る", "超能力を得る", "幸福を追求する"}, Correct: "現実を正しく見る", Supplement: "ヴィパッサナー瞑想法は、自分の現実を鏡に写したように見ることができる最良の方法です。"},
	{Question: "心の働きの一部として説明されているものはどれですか？", Options: pq.StringArray{"記憶", "感情", "計算力", "全て"}, Correct: "全て", Supplement: "心の働きには、記憶、感情、計算力などが含まれます。"},
	{Question: "仏教における「無常」とは何を意味しますか？", Options: pq.StringArray{"永遠不滅", "変化しない", "常に変化する", "固定されたもの"}, Correct: "常に変化する", Supplement: "無常とは、すべてのものが常に変化し続けることを意味します。"},
	{Question: "宗教が死後の世界について語る理由として説明されているものはどれですか？", Options: pq.StringArray{"人間は死後の世界を信じたくないから", "宗教の教義を広めるため", "死後の世界を体験した人がいるから", "全て"}, Correct: "人間は死後の世界を信じたくないから", Supplement: "宗教が死後の世界について語る理由として、人間は死後の世界を信じたくないからと説明されています。"},
	{Question: "お釈迦さまが主に説いたことは何ですか？", Options: pq.StringArray{"死後の世界について", "輪廻転生のシステムについて", "生き方について", "全て"}, Correct: "生き方について", Supplement: "お釈迦さまは主に、我々がどう生きるべきかということについて説いていました。"},
}
