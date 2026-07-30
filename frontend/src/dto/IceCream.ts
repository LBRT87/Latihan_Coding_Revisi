export interface IceCream{
    name: string;
    photo: string;
    price: number;
    description: string;
    flavour:string;
}

export interface NavContent{
    nav1: string,
    nav2: string,
    nav3: string,
}

export interface Carousel{
    img1: string;
    img2: string;
    img3: string;
    img4: string;
    img5: string;
}

export interface CreateIceCreamResponse {
    message : string;
}