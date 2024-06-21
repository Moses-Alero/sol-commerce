
import {createRoot} from "react-dom/client"
import { ProductDisplay } from "./components/product";


export function renderProductPurchase(Price: string, Description:string, Title: string){
    const productPurchase = document.getElementById("product-purchase")
    if (!productPurchase) {
        throw new Error('Could not find element with id balance');
    }
    
    const productPurchaseReactRoot = createRoot(productPurchase)

    productPurchaseReactRoot.render(ProductDisplay({Price, Description, Title}))
}


