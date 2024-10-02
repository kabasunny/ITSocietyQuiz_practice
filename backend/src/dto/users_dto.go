package dto

// フロントへの管理者用ユーザー一覧データ返却用
type AdmUserData struct {
	ID         uint    `json:"dbId"` // GORMのIDを使用する
	EmpID      string  `json:"empId"`
	Username   *string `json:"name"`
	Email      string  `json:"email"`
	Password_1 string  `json:"password_1"` // 変更の場合、旧パスワード。新規ユーザーの場合、新しいパスワード
	Password_2 string  `json:"password_2"` // 変更の場合、新規パスワード。新規ユーザーの場合、パスワードが一致しているか
	RoleID     uint    `json:"roleId"`
	RoleName   string  `json:"roleName"`
	CreatedAt  string  `json:"createdAt"` // ユーザーテーブルの作成日
	UpdatedAt  string  `json:"updatedAt"` // ユーザーテーブルの更新日
}
