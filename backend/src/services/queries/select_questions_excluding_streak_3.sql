WITH LatestAnswers AS (
    SELECT 
        question_id,
        streak_count,
        created_at,
        ROW_NUMBER() OVER (
            PARTITION BY question_id 
            ORDER BY created_at DESC
        ) AS rn
        -- 各 question_id ごとに最新の回答を取得するための行番号を付与  rn=1が最新となる
    FROM 
        answers
    WHERE 
        emp_id = ?
        -- 特定の従業員IDに基づいて回答をフィルタリング
    LIMIT ?  -- 最大105レコードを抽出
)
SELECT 
    question_id
FROM 
    LatestAnswers
WHERE 
    rn = 1
    -- 最新の回答（rn = 1）のみを対象とする
    AND (
        (streak_count = 2 AND created_at <= NOW() - INTERVAL '14 days') OR
        (streak_count = 1 AND created_at <= NOW() - INTERVAL '7 days') OR
        (streak_count = 0 AND created_at <= NOW() - INTERVAL '3 days')
    )
    -- streak_count が 2 かつ created_at が14日経過、1かつ7日経過、0かつ3日経過のレコードに絞る
ORDER BY 
    streak_count DESC,
    question_id ASC
    -- streak_count の降順、question_id の昇順で並べ替え
LIMIT ?;  -- 取得する質問の数を制限  現状1日5問
