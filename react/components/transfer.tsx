window.Buffer = window.Buffer || require("buffer").Buffer;
import { useWallet, useConnection} from "@solana/wallet-adapter-react";
import { WalletMultiButton, useWalletModal } from "@solana/wallet-adapter-react-ui";
import { LAMPORTS_PER_SOL, SystemProgram, PublicKey, Transaction } from "@solana/web3.js";
import { useCallback, useEffect } from "react";

export const Transfer = (props:{Price: string}): any => {
    const { publicKey, sendTransaction, wallet, connect, connected, connecting, autoConnect } = useWallet();
  

    const handleConnect = useCallback(async () => {
        if (!connected && !connecting) {
            try {
                await connect();
            } catch (error) {
                console.error('Failed to connect wallet:', error);
            }
        }
    }, [connect, connected, connecting]);

    useEffect(() => {
        // Auto-connect on mount if a wallet is available
        if (wallet && autoConnect) {
            handleConnect();
        }
    }, [wallet, autoConnect, handleConnect]);
    
    const number = parseFloat(props.Price)/135.28
    const amount_to_transfer = number.toFixed(3)

    console.log(amount_to_transfer)
    const { connection } = useConnection();
    const sendSol = (event) => {

    event.preventDefault();
    handleConnect()
    const transaction = new Transaction();
    const recipientPubKey = new PublicKey("CxT4fH6rTMSzN7dDUnoSMEreA3fG9UwRGtUx4FyJNeRV");
    
    const sendSolInstruction = SystemProgram.transfer({
        fromPubkey: publicKey,
        toPubkey: recipientPubKey,
        lamports: LAMPORTS_PER_SOL * parseFloat(amount_to_transfer)
    });


    transaction.add(sendSolInstruction);
    sendTransaction(transaction, connection).then((sig) => {
        console.log(sig);
    });
};
    return(
        <>
          <WalletMultiButton/>
          <button className="product-purchase-btn" onClick={sendSol}>BUY NOW</button>
        </>
    )
}