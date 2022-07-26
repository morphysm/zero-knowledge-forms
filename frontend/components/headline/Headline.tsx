import React from 'react';
import WalletConnector from '../walletConnector/WalletConnector';

const Headline: React.FC = () => {
  return (
    <div
      style={{
        display: 'flex',
        flexDirection: 'row',
        alignItems: 'center',
        justifyContent: 'space-between',
        padding: '20px 30px',
        background: 'white',
        marginBottom: '20px',
      }}
    >
      <h3>Famed Submission Form Prototype</h3>
      <WalletConnector />
    </div>
  );
};

export default Headline;
