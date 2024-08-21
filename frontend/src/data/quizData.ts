interface QuizData {
    question: string;
    options: string[];
    correct: string;
    supplement: string;
  }
  
  export const quizData: QuizData[] = [
    {
      question: "悟りの第二段階の状態は？",
      options: ["予流果", "一来果", "不還果", "阿羅漢果"],
      correct: "一来果",
      supplement: "欲界の煩悩を断じ終えた位のこと",
    },
    {
      question: "四聖諦（ししょうたい）はどれか？",
      options: ["苦諦", "楽諦", "幸諦", "怒諦"],
      correct: "苦諦",
      supplement: "四聖諦（ししょうたい）は、仏教における4つの基本的な真理で、苦諦（くたい）、集諦（じったい）、滅諦（めったい）、道諦（どうたい）",
    },
    {
      question: "テーラワーダ仏教の「テーラワーダ」とは何を意味しますか？",
      options: ["長老の教え", "仏陀の教え", "慈悲の教え", "瞑想の教え"],
      correct: "長老の教え",
      supplement: "テーラワーダ」はパーリ語で「長老の教え」を意味します。",
    },
    {
      question: "初期仏教の教えの目的は何ですか？",
      options: ["信じる者は救われる", "儀式儀礼を行う", "苦しみをなくす", "現世利益を得る"],
      correct: "苦しみをなくす",
      supplement: "初期仏教の教えは「苦しみをなくす」ことを目的としています。",
    },
    {
      question: "四聖諦の中で、仏教が特に重要視しているものは何ですか？",
      options: ["苦（Dukkha）", "集（Samudaya）", "滅（Nirodha）", "道（Magga）"],
      correct: "苦（Dukkha）",
      supplement: "仏教では、苦というものの認識を非常に重要視しています。",
    },
    {
      question: "ヴィパッサナー瞑想法の目的は何ですか？",
      options: ["心を空にする", "現実を正しく見る", "超能力を得る", "幸福を追求する"],
      correct: "現実を正しく見る",
      supplement: "ヴィパッサナー瞑想法は、自分の現実を鏡に写したように見ることができる最良の方法です。",
    },
    {
      question: "心の働きの一部として説明されているものはどれですか？",
      options: ["記憶", "感情", "計算力", "全て"],
      correct: "全て",
      supplement: "心の働きには、記憶、感情、計算力などが含まれます。",
    },
    {
      question: "仏教における「無常」とは何を意味しますか？",
      options: ["永遠不滅", "変化しない", "常に変化する", "固定されたもの"],
      correct: "常に変化する",
      supplement: "無常とは、すべてのものが常に変化し続けることを意味します。",
    }
  ];
  