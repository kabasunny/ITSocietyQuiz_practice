SELECT 
    ROW_NUMBER() OVER (ORDER BY sub.p DESC, sub.c DESC) AS Rank,
    sub.emp_id,
    sub.username,
    sub.current_q_id,
    sub.c,
    sub.p
FROM (
    SELECT
        u.emp_id,
        u.username,
        u.current_q_id,
        CASE 
            WHEN u.total_questions = 0 THEN 0 
            ELSE (u.correct_answers * 100.0 / CAST(u.total_questions AS FLOAT))
        END AS c, -- correctAnswerRateperでは、何故がうまくいかない
        CASE 
            WHEN u.total_questions = 0 THEN 0
            ELSE (u.current_q_id * (u.correct_answers * 1.0 / CAST(u.total_questions AS FLOAT)))
        END AS p -- performanceIndicatorでは、何故がうまくいかない
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
