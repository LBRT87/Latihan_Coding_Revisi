import { useState } from "react";
import type { User } from "../dto/User";
import type { NavContent } from "../dto/IceCream";
import Navigation from "../components/Navigation";
import Landing from "./Landing";
import History from "./History";
import Shop from "./Shop";
import Cart from "./Cart";

interface DashboardProps {
    user: User,
    onLogout: () => void,
}

export default function Dashboard({ user, onLogout }: DashboardProps) {
    const [activeTab, setActiveTab] = useState<'Landing' | 'Profile' | 'Shop' | 'Cart' | 'History'>('Profile');
    const navContent = {
        nav1: 'Shop',
        nav2: 'Cart',
        nav3: 'History'
    } as NavContent 
    return (
        <div className="w-full bg-slate-800 p-6 rounded-2xl shadow-xl space-y-4 max-h-max">
            <Navigation user={user} onLogout={onLogout} activeTab={activeTab} onChangeTab={(tab) => setActiveTab(tab as 'Landing' | 'Profile' | 'Shop' | 'Cart' | 'History')} navigation={navContent} />
            
            {activeTab === 'Landing' && (<Landing carousel={{img1: "hal",img2: "hal",img3: "hal",img4: "hal",img5: "hal"}}/>)}

            {activeTab === 'Profile' && ( 
                <div className="max-w-2xl mx-auto bg-slate-800 p-6 rounded-xl border border-slate-700">
                <div className="flex gap-4 items-center mb-6 border-b border-slate-700 pb-4">
                    <div className="w-16 h-16 bg-slate-700 rounded-full flex-shrink-0"></div>
                    <div>
                    <h2 className="text-white text-lg font-bold">{user.email}</h2>
                    <p className="text-sm text-slate-400">{user.email}</p>
                    </div>
                    <div className="ml-auto flex gap-4 text-center">
                    </div>
                </div>

                <form className="space-y-4">
                    <div className="grid grid-cols-2 gap-4">
                    <input type="text" placeholder="Username" className="p-2 bg-slate-900 rounded text-white border border-slate-700" />
                    <input type="email" placeholder="Email" className="p-2 bg-slate-900 rounded text-white border border-slate-700" />
                    </div>
                    <textarea placeholder="Bio" className="w-full p-2 bg-slate-900 rounded text-white border border-slate-700"></textarea>

                    <div className="grid grid-cols-3 gap-2">
                    <input type="password" placeholder="Old Password" className="p-2 bg-slate-900 rounded text-white border border-slate-700" />
                    <input type="password" placeholder="New Password" className="p-2 bg-slate-900 rounded text-white border border-slate-700" />
                    <input type="password" placeholder="Confirm" className="p-2 bg-slate-900 rounded text-white border border-slate-700" />
                    </div>

                    <button className="px-4 py-2 bg-emerald-600 text-white rounded">Save Changes</button>
                </form>
                </div>
            )}

            {activeTab === 'Shop' && (<Shop iceCream={[]}/> )}

            {activeTab === 'Cart' && (<Cart/> )}

            {activeTab === 'History' && (<History transaction={{actions: "a", customer: "wilson", date:"1", status:"te", total:123 }}/>)}
        </div>
    );
}