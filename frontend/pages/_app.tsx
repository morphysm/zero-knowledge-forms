import '../styles/globals.css';
import type { AppProps } from 'next/app';
import {
  Web3Provider,
  ExternalProvider,
  JsonRpcFetchFunc,
} from '@ethersproject/providers';
import { Web3ReactProvider } from '@web3-react/core';
import Background from '../components/background/Background';

function getLibrary(provider: ExternalProvider | JsonRpcFetchFunc) {
  return new Web3Provider(provider);
}

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <Web3ReactProvider getLibrary={getLibrary}>
      <Background />
      <Component {...pageProps} />
    </Web3ReactProvider>
  );
}

export default MyApp;
