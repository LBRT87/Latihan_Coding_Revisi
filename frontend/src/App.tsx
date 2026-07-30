import { useState } from "react";
import Login from "./pages/Login";
import Register from "./pages/Register";
import Dashboard from "./pages/Dashboard";
import type { User } from "./dto/User";
import DashboardA from "./pages/DashboardA";

export default function App() {
  const [page, setPage] = useState<'login' | 'register'  | 'dashboard' | 'dashboard2'>('dashboard');
  // const [user, setUser] = useState<User | null>(null);
  const [user, setUser] = useState<User | null>({
    email: "wilson@gmail.com",
    password: "wilson",
    role: "customer",
  });
  
  const handleLogout = () => {
    localStorage.clear();
    setUser(null);
    setPage('login');
  }

  return (
      <div className={`min-h-screen bg-slate-900 text-white p-4 flex ${
        page === 'dashboard' ? 'flex-col items-stretch' : 'items-center justify-center'
      }`}>
        { page === 'login' && (<Login onSuccessLogin={(u) => { setUser(u); setPage('dashboard'); }} onNavigate={setPage}/>) }
        { page === 'register' && (<Register onNavigate={setPage}/>) }
        { page === 'dashboard' && user.role === 'customer' && (<Dashboard user={user} onLogout={handleLogout}/>)}
        { page === 'dashboard2' && user.role === 'admin' && (<DashboardA user={user} onLogout={handleLogout}/>)}
      </div>
  );
}