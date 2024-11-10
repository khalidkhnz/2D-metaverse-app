export interface IRegister {
  fullName: string;
  email: string;
  password: string;
  roleId: string;
  permissions: string[];
}

export interface ILogin {
  email: string;
  password: string;
}
