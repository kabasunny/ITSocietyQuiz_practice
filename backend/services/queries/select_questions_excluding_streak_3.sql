WITH LatestAnswers AS (
    SELECT 
        question_id,
        streak_count,
        ROW_NUMBER() OVER (PARTITION BY question_id ORDER BY created_at DESC) AS rn
        -- 各 question_id ごとに最新の回答を取得するための行番号を付与  rn=1が最新となる
    FROM 
        answers
    WHERE 
        emp_id = ?
        -- 特定の従業員IDに基づいて回答をフィルタリング
    LIMIT ?
    -- 取得する行数の上限を設定  現状105レコード
)
SELECT 
    question_id
FROM 
    LatestAnswers
WHERE 
    rn = 1 AND streak_count IN (0, 1, 2)
    -- 最新の回答（rn = 1）かつ 連続正解数streak_count が 0, 1, 2 の質問を選択  連続正解数=3は除外する
ORDER BY 
    streak_count DESC,
    question_id ASC
    -- streak_count の降順、question_id の昇順で並べ替え
LIMIT ?;
    -- 取得する質問の数を制限  現状1日5問
