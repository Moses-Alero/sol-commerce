import { FC } from "react"
import { WalletContextProvider } from './wallet';
import { Transfer } from "./transfer"
// import { WalletButton } from "./balance";
import { WalletMultiButton } from "@solana/wallet-adapter-react-ui";

export const Product = (props:{Price: string, Description: string, Title: string})=>{
    const number = Number(props.Price)/135.28
    return (
        <>
            <section className="product-desc">
                <h1>{props.Title}</h1>
                <span className="product-price">{number.toFixed(5)} SOL</span>
                <p>${props.Price}</p> 
                <p>{props.Description}</p>
                <Transfer Price={props.Price}/>
            </section>
            <WalletMultiButton/>
        </>
    ) 
}

export const ProductDisplay = (props:{Price: string, Description: string, Title: string})=>{
  return (
    <>
    <WalletContextProvider {...props}>
        <Product {...props}/>
    </WalletContextProvider>
    </>
  )
}