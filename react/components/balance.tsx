import {WalletContextProvider} from "./wallet";
import { WalletMultiButton } from "@solana/wallet-adapter-react-ui";

export const WalletButton = () => {
  return (
    <WalletContextProvider>
      <WalletMultiButton />
    </WalletContextProvider>
  );
}