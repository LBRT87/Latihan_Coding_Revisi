import { useState } from "react";
import type { RegisterResponse } from "../dto/User";
import { req } from "../api/Api";

interface RegisterProps{
	onNavigate: (pages: 'login' | 'register') => void;
}

export default function Register({ onNavigate }: RegisterProps) {
	const [, setError] = useState('');

	const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
		e.preventDefault();
		setError('');
		const formData = new FormData(e.currentTarget);
        const fields = {
            email: formData.get('email') as string,
            password: formData.get('password') as string,
            role: "customer"
        }
        try{
            req<RegisterResponse>('/auth/register', fields);
            onNavigate('login');
        } catch (err) {
            setError(err instanceof Error ? err.message : "Error");
        }
 	};

    return(
        <div className="max-w-md mx-auto bg-slate-800 p-6 rounded-xl border border-slate-700">
            <h2 className="text-xl font-bold text-white mb-4 text-center">Create Account</h2>

            <form className="space-y-4 flex flex-col" onSubmit={handleSubmit}>
                <div className="grid grid-cols-2 gap-4"/>

                <input type="email" name="email" placeholder="Email Address" className="p-2 bg-slate-900 border border-slate-700 rounded text-white" />
                <input type="password" name="password" placeholder="Password" className="p-2 bg-slate-900 border border-slate-700 rounded text-white" />
                <input type="password" name="password" placeholder="Confirm Password" className="p-2 bg-slate-900 border border-slate-700 rounded text-white" />

                <div className="p-2 border border-slate-700 rounded bg-slate-900">
                <label className="text-sm text-slate-400 block mb-1">Profile Picture</label>
                <input type="file" accept="image/*" className="text-slate-300 text-sm" />
                </div>

                <button className="py-2 bg-emerald-600 text-white font-medium rounded">
                    Register
                </button>

                <div className="grid grid-cols-2 gap-4"/>
                
                <button onClick={() => {onNavigate('register')}} className="w-full py-2 bg-white text-slate-900 font-medium rounded flex items-center justify-center gap-2">
                    <span>Already have an account? login here.</span>
                </button>
            </form>
        </div>
    );
}