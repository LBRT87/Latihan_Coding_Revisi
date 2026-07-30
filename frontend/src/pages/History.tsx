import HistoryTransaction from "../components/HistoryTransaction";
import type { Transaction } from "../dto/Transaction";

interface HistoryProps{
    transaction: Transaction
}

export default function History({transaction} : HistoryProps){
    return(
        <section className="py-24 bg-slate-700">
        <div className="w-full max-w-7xl mx-auto px-4 md:px-8">
            <div className="main-data p-8 sm:p-14 bg-slate-800 rounded-3xl">
                <h2 className="text-center font-manrope font-semibold text-4xl text-black mb-16">Order History</h2>
                <div className="grid grid-cols-8 pb-9">
                    <div className="col-span-8 lg:col-span-4">
                        <p className="font-medium text-lg leading-8 text-white">Product </p>
                    </div>
                    <div className="col-span-1 max-lg:hidden">
                        <p className="font-medium text-lg leading-8 text-white text-center">price </p>
                    </div>
                    <div className="col-span-1 max-lg:hidden flex items-center justify-center">
                        <p className="font-medium text-lg leading-8 text-white">Qty </p>
                    </div>
                    <div className="col-span-2 max-lg:hidden">
                        <p className="font-medium text-lg leading-8 text-white">Delivery Expected by </p>
                    </div>
                </div>
            </div>

            <HistoryTransaction transaction={transaction}/>
        </div>
        </section>
    );
}