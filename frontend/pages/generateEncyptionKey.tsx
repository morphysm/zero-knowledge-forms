import type { NextPage } from 'next';
import Head from 'next/head';
import Headline from '../components/headline/Headline';
import Form from '../components/form/Form';
import Box from '@mui/material/Box';
import Tab from '@mui/material/Tab';
import TabContext from '@mui/lab/TabContext';
import TabList from '@mui/lab/TabList';
import TabPanel from '@mui/lab/TabPanel';
import FormSender from '../components/formSender/FormSender';
import { useEffect, useState } from 'react';
import Stack from '@mui/material/Stack';
import FormBuilder from '../components/formBuilder/FormBuilder';
import { useWeb3React } from '@web3-react/core';
import { Button, Typography } from '@mui/material';
import WalletConnector from '../components/walletConnector/WalletConnector';

const owner = process.env.NEXT_PUBLIC_OWNER;

const GenerateEnryptionKey: NextPage = () => {
  const [enryptionKey, setEnryptionKey] = useState<string>('');
  const { account, connector } = useWeb3React();

  useEffect(() => {
    setEnryptionKey('');
  }, [account]);

  const handleGeneratClick = async () => {
    const provider = await connector?.getProvider();

    try {
      const _enryptionKey = await provider.request({
        method: 'eth_getEncryptionPublicKey',
        params: [account],
      });
      setEnryptionKey(_enryptionKey);
    } catch (error: any) {
      if (error && error.code && error.code === 4001) {
        // EIP-1193 userRejectedRequest error
        console.log("We can't encrypt anything without the key.");
      } else {
        console.error(error);
      }
    }
  };

  return (
    <div className='App'>
      <Box m={2}>
        <Stack spacing={2}>
          {account ? (
            <Typography variant='subtitle1' component='div'>
              Generate encyption key for account: {account}
            </Typography>
          ) : (
            <Typography variant='subtitle1' component='div'>
              Please connect your wallet
            </Typography>
          )}
          {enryptionKey && (
            <Typography variant='subtitle1' component='div'>
              Encyption key: {enryptionKey}
            </Typography>
          )}
          <Stack direction='row' spacing={2}>
            <Button
              variant='outlined'
              onClick={handleGeneratClick}
              fullWidth={false}
            >
              Generate
            </Button>
            <WalletConnector />
          </Stack>
        </Stack>
      </Box>
    </div>
  );
};

export default GenerateEnryptionKey;
