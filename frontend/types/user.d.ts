export interface IUserApiResponse {
  data: {
    user: IUser;
  };
  success: boolean;
}

export interface IUser {
  _id: string;
  fullName: string;
  email: string;
  password: string;
  role: Role;
  permissions: null | string[];
  profile: Profile;
}

interface Role {
  _id: string;
  name: string;
  description: string;
  createdAt: string;
  updatedAt: string;
}

interface Profile {
  _id: string;
  authId: string;
  fullName: string;
  bio: string;
  avatar: string;
  username: string;
  role: string;
  status: string;
  socketId: string;
  createdAt: string;
  updatedAt: string;
}
