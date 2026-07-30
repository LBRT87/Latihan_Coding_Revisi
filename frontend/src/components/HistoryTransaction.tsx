import type { Transaction } from "../dto/Transaction";

interface HistoryTransactionProps{
    transaction: Transaction;
}

export default function HistoryTransaction({ transaction }: HistoryTransactionProps){
    return(
        <div className="box p-8 rounded-3xl bg-slate-600 grid grid-cols-8 mb-7 cursor-pointer transition-all duration-500 hover:bg-indigo-50 max-lg:max-w-xl max-lg:mx-auto">
            <div className="col-span-8 sm:col-span-4 lg:col-span-2 flex items-center justify-center ">
                <p className="font-semibold text-xl leading-8 text-black">{transaction.actions}</p>
            </div>
            
            <div
                className="col-span-8 sm:col-span-4 lg:col-span-3 flex h-full justify-center pl-4 flex-col max-lg:items-center">
                <h5 className="font-manrope font-semibold text-2xl leading-9 text-white mb-1 whitespace-nowrap">
                    {transaction.customer}</h5>
                <p className="font-normal text-base leading-7 text-white max-md:text-center">White</p>
            </div>

            <div className="col-span-8 sm:col-span-4 lg:col-span-1 flex items-center justify-center">
                    <p className="font-semibold text-xl leading-8 text-white">{transaction.total}</p>
            </div>
            <div className="col-span-8 sm:col-span-4 lg:col-span-1 flex items-center justify-center ">
                <p className="font-semibold text-xl leading-8 text-white text-center">{transaction.status}</p>
            </div>
            <div className="col-span-8 sm:col-span-4 lg:col-span-2 flex items-center justify-center ">
                <p className="font-semibold text-xl leading-8 text-white">{transaction.date.toString()}</p>
            </div>
        </div>
    );
}