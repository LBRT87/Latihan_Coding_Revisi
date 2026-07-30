import { useState } from "react";
import type { CreateIceCreamResponse } from "../dto/IceCream";
import { req } from "../api/Api";


export default function CreateIceCream(){
    const [, setError] = useState('');
    
    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        setError('');
        const formData = new FormData(e.currentTarget);
        const fields = {
            name: formData.get('name') as string,
            flavour: formData.get('flavour') as string,
            description: formData.get('description') as string,
            price: formData.get('price') ,
        }
        try{
            req<CreateIceCreamResponse>('product/create',fields)
            console.log('tes dulu')
        } catch (err) {
            setError(err instanceof Error ? err.message : "Error");
        }
    };
    return (
        <div className="max-w-md mx-auto bg-slate-800 p-6 rounded-xl border border-slate-700">
        <h2 className="text-xl font-bold text-white mb-4 text-center">Create Ice Cream</h2>

        <form className="space-y-4 flex flex-col" onSubmit={handleSubmit}>
            <div className="grid grid-cols-2 gap-4">
            <input type="text" name="name" placeholder="Name" className="p-2 bg-slate-900 border border-slate-700 rounded text-white" />
            <input type="text" name="flavour" placeholder="Flavour" className="p-2 bg-slate-900 border border-slate-700 rounded text-white" />
            <input type="text" name="description" placeholder="Description" className="p-2 bg-slate-900 border border-slate-700 rounded text-white" />
            <input type="number" name="price" placeholder="Price(Rp)" className="p-2 bg-slate-900 border border-slate-700 rounded text-white" />
            </div>

            <div className="p-2 border border-slate-700 rounded bg-slate-900">
            <label className="text-sm text-slate-400 block mb-1">Image Ice Cream</label>
            <input type="file" accept="image/*" name="imageicecream" className="text-slate-300 text-sm" />
            </div>

            <button className="py-2 bg-emerald-600 text-white font-medium rounded">
            Create
            </button>
        </form>
        </div>
    )
}