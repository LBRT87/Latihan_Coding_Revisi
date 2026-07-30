import type { IceCream } from "../dto/IceCream";

interface ShopCardProps{
    iceCream: IceCream;
}

export default function ShopCard({iceCream}: ShopCardProps){
    return(
        <div className="bg-slate-800 rounded-lg overflow-hidden border border-slate-700 w-64">
            <img src={iceCream.photo} alt="Course" className="w-full h-32 object-cover bg-slate-700" />
            <div className="p-4 flex flex-col gap-2">
                <span className="text-xs text-emerald-400 uppercase font-semibold">Technology</span>
                <h3 className="text-white font-medium truncate">{iceCream.name}</h3>
                <div className="flex justify-between items-center mt-2">
                <span className="text-sm text-slate-400">{iceCream.description}</span>
                <span className="text-emerald-500 font-bold">{iceCream.price}</span>
                </div>
            </div>
        </div>
    );
}