import type { NavContent } from "../dto/IceCream";
import type { User } from "../dto/User";

interface NavigationProps{
    user: User;
    onLogout: ()=>void;
    activeTab?: string;
    onChangeTab?: (tab: string) => void;
    navigation: NavContent;
}

export default function Navigation({ user, onLogout, navigation, onChangeTab }: NavigationProps){
    return (
        <nav className="bg-slate-900 border-b border-slate-800 p-4">
            <div className="max-w-7xl mx-auto flex items-center justify-between">

                <div className="flex items-center gap-6">
                <h1 className="text-emerald-500 font-bold text-xl"><button onClick={()=>onChangeTab("Landing")}>ESkrim</button></h1>
                <div className="hidden md:flex gap-4 text-sm text-slate-300">
                    <button onClick={()=>onChangeTab(navigation.nav1)} className="hover:text-emerald-400">{navigation.nav1}</button>
                    <button onClick={()=>onChangeTab(navigation.nav2)} className="hover:text-emerald-400">{navigation.nav2}</button>
                    <button onClick={()=>onChangeTab(navigation.nav3)} className="hover:text-emerald-400">{navigation.nav3}</button>
                </div>
                </div>

                <div className="flex items-center gap-4 text-sm">
                <div className="border-l border-slate-700 pl-4 flex items-center gap-2">
                    <div className="w-8 h-8 bg-slate-700 rounded-full text-center leading-8 text-white"><button onClick={()=>onChangeTab("Profile")}>{user.email[0]}</button></div>
                    <button onClick={onLogout} className="text-red-400 hover:underline">Logout</button>
                </div>
                </div>

            </div>
        </nav>
    );
}