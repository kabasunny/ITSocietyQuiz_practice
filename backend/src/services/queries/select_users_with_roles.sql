SELECT DISTINCT
    users.id,
    users.emp_id,
    users.username,
    users.email,
    roles.role_name,
    users.created_at,
    users.updated_at
FROM 
    users
JOIN 
    users_roles ON users.emp_id = users_roles.emp_id
JOIN 
    roles ON users_roles.role_id = roles.role_id
ORDER BY 
    users.id;
