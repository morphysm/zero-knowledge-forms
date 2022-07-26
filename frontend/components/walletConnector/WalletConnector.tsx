import React, { useState } from 'react';
import Button from '@mui/material/Button';
import styles from './Button.module.css';
import Modal from '@mui/material/Modal';
import Box from '@mui/material/Box';
import Stack from '@mui/material/Stack';

import { useWeb3React } from '@web3-react/core';
import { AbstractConnector } from '@web3-react/abstract-connector';
import { InjectedConnector } from '@web3-react/injected-connector';

const Injected = new InjectedConnector({
  supportedChainIds: [5],
});

const WalletConnector: React.FC = () => {
  const [open, setOpen] = useState<boolean>(false);
  const { active, activate, deactivate, connector, account } = useWeb3React();

  const handleToggleOpenClick = () => {
    setOpen(!open);
  };

  const handleActivationClick = async (connector: AbstractConnector) => {
    await activate(connector, function (error) {
      // TODO handle error
      console.log(error);
    });

    setOpen(false);
  };

  // TODO add connection info & move to component
  return (
    <div>
      {!active ? (
        <Button variant='contained' onClick={handleToggleOpenClick}>
          Connect Wallet
        </Button>
      ) : (
        <Button variant='contained' onClick={deactivate}>
          Disconnect
        </Button>
      )}
      <Modal
        open={open}
        onClose={handleToggleOpenClick}
        aria-labelledby='modal-modal-title'
        aria-describedby='modal-modal-description'
      >
        <Box>
          <Stack spacing={2} direction='column'>
            <Button
              variant='contained'
              onClick={() => {
                handleActivationClick(Injected);
              }}
            >
              Metamask
            </Button>
          </Stack>
        </Box>
      </Modal>
    </div>
  );
};

export default WalletConnector;
