package dto

// フロントへの管理者用ユーザー一覧データ返却用
type AdmUserData struct {
	ID        uint    `json:"id"` // GORMのIDを使用する
	EmpID     string  `json:"empid"`
	Username  *string `json:"name"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	Role      string  `json:"role"`
	CreatedAt string  `json:"createdAt"` // ユーザーテーブルの作成日
	UpdatedAt string  `json:"updatedAt"` // ユーザーテーブルの更新日
}
