import ShopCard from "../components/ShopCard";
import type { IceCream } from "../dto/IceCream";

interface ShopProps{
    iceCream: IceCream[]
}

export default function Shop({ iceCream }: ShopProps){
    return(
        <div className="grid grid-cols-4 gap-4">
            <ShopCard iceCream={iceCream[0]}/>
            <ShopCard iceCream={iceCream[1]}/>
            <ShopCard iceCream={iceCream[2]}/>
            <ShopCard iceCream={iceCream[3]}/>
            <ShopCard iceCream={iceCream[4]}/>
        </div>
    );
}