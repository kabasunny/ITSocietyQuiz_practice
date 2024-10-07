SELECT 
    ROW_NUMBER() OVER (ORDER BY sub.p DESC, sub.c DESC) AS Rank,
    sub.emp_id,
    sub.username,
    sub.current_q_id,
    sub.total_questions,
    sub.correct_answers,
    ROUND(sub.c, 1) AS c, -- 小数第一位までに丸める
    ROUND(sub.p, 1) AS p  -- 小数第一位までに丸める
FROM (
    SELECT
        u.emp_id,
        u.username,
        u.current_q_id,
        u.total_questions,
        u.correct_answers,
        CASE 
            WHEN u.total_questions = 0 THEN 0 
            ELSE (u.correct_answers * 100.0 / u.total_questions)
        END AS c,
        CASE 
            WHEN u.total_questions = 0 THEN 0
            ELSE (u.current_q_id * (u.correct_answers * 1.0 / u.total_questions))
        END AS p
    FROM
        users u
    WHERE
        u.emp_id NOT IN (
            SELECT ur.emp_id
            FROM users_roles ur
            WHERE ur.role_id = 1
        )
) AS sub
ORDER BY
    sub.p DESC,
    sub.c DESC;
