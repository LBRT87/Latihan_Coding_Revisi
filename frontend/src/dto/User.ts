export interface User{
	email: string;
	password: string;
    role: 'customer' | 'admin';
}

export interface LoginResponse{
	access_token: string;
	refresh_token: string;
	user: User;
}

export interface RegisterResponse{
	message: string;
}

