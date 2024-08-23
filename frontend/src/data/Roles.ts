// Roles.ts
export interface Role {
    id: number;
    role_name: string;
  }
  
  export const roles: Role[] = [
    {
      id: 1,
      role_name: "管理者",
    },
    {
      id: 2,
      role_name: "一般ユーザー",
    },
  ];
  