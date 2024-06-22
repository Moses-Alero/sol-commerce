import { FC, ReactNode, useEffect, useMemo, useState } from "react";
import {
  ConnectionProvider,
  WalletProvider,
  useWallet,
  useConnection,
} from "@solana/wallet-adapter-react";
import * as web3 from "@solana/web3.js";
// import { WalletMultiButton } from '@solana/wallet-adapter-react-ui'
import * as walletAdapterWallets from "@solana/wallet-adapter-wallets";
import { WalletModalProvider, WalletMultiButton } from "@solana/wallet-adapter-react-ui";
require("@solana/wallet-adapter-react-ui/styles.css");

const WalletStuff = ()=>{
  const [balance, setBalance] = useState(0);
  const {publicKey} = useWallet() 
  const {connection} = useConnection()
  useEffect(() => {
    if (!connection || !publicKey) { return }

    // Ensure the balance updates after the transaction completes
    connection.onAccountChange(
        publicKey, 
        (updatedAccountInfo) => {
            setBalance(updatedAccountInfo.lamports / web3.LAMPORTS_PER_SOL)
        }, 
        'confirmed'
    )
   
    connection.getAccountInfo(publicKey).then(info => {
        setBalance(info.lamports);

        console.log("rererer")
    })
    
}, [connection, publicKey])
  return (
    <></>
  )
}

export const WalletContextProvider: FC<{ children: ReactNode }> = ({ children }) => {
  const endpoint = web3.clusterApiUrl("devnet");
//   const wallets = [new walletAdapterWallets.PhantomWalletAdapter()]
  const wallets = useMemo(() => [], []);
  return (
   <div>
     <ConnectionProvider endpoint={endpoint}>
      <WalletProvider wallets={wallets} autoConnect={true}>
        <WalletModalProvider>
          {children}
        </WalletModalProvider>
      </WalletProvider>
    </ConnectionProvider>
   </div>
  );
};