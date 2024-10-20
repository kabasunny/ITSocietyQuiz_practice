import axios from 'axios';
import { AdminsUser } from '../../../types';

const getHeaders = (jwt: string | null) => ({
  Authorization: `Bearer ${jwt}`,
  'Content-Type': 'application/json',
});

export const fetchUsers = (jwt: string | null, setUsers: React.Dispatch<React.SetStateAction<AdminsUser[]>>) => {
  axios
    .get(`${process.env.REACT_APP_API_URL}/admins/userdata/userslist`, {
      headers: getHeaders(jwt),
    })
    .then((response) => {
      setUsers(response.data.userlist);
      // console.log(response.data.userlist);
    })
    .catch((error) => {
      console.error('Error fetching users:', error);
    });
};

export const addUser = (
  data: AdminsUser,
  jwt: string | null,
  users: AdminsUser[],
  setUsers: React.Dispatch<React.SetStateAction<AdminsUser[]>>,
  setShowAddForm: React.Dispatch<React.SetStateAction<boolean>>
) => {
  axios
    .post(`${process.env.REACT_APP_API_URL}/admins/userdata/addusers`, data, {
      headers: getHeaders(jwt),
    })
    .then((response) => {
      setUsers([...users, response.data]);
      alert('ユーザー情報が正常に更新されました。');
      setShowAddForm(false);
    })
    .catch((error) => {
      console.error('Error adding user:', error);
      console.error('Error details:', error.response ? error.response.data : error.message);
    });
};

export const updateUser = (
  data: AdminsUser,
  editingUser: AdminsUser | null,
  jwt: string | null,
  users: AdminsUser[],
  setUsers: React.Dispatch<React.SetStateAction<AdminsUser[]>>,
  setEditingUser: React.Dispatch<React.SetStateAction<AdminsUser | null>>
) => {
  if (editingUser) {
    data.dbId = editingUser.dbId;
  }
  axios
    .put(`${process.env.REACT_APP_API_URL}/admins/userdata/updateusers/${data.dbId}`, data, {
      headers: getHeaders(jwt),
    })
    .then((response) => {
      if (response.status === 200) {
        const updatedUser = response.data;
        setUsers(
          users.map((user) => (user.dbId === data.dbId ? { ...user, ...updatedUser } : user))
        );
        alert('ユーザー情報が正常に更新されました。');
      } else {
        console.error('Unexpected response status:', response.status);
      }
    })
    .catch((error) => {
      console.error('Error updating user:', error);
    });
  setEditingUser(null);
};

export const deleteUser = (
  dbId: number,
  jwt: string | null,
  users: AdminsUser[],
  setUsers: React.Dispatch<React.SetStateAction<AdminsUser[]>>
) => {
  const confirmDelete = window.confirm('本当にこのユーザーを削除しますか？');
  if (confirmDelete) {
    axios
      .delete(`${process.env.REACT_APP_API_URL}/admins/userdata/deleteuser/${dbId}`, {
        headers: getHeaders(jwt),
      })
      .then(() => {
        setUsers(users.filter((user) => user.dbId !== dbId));
      })
      .catch((error) => {
        console.error('Error deleting user:', error);
      });
  }
};
