import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import { useWeb3React } from '@web3-react/core';
import { utils, Waku, WakuMessage } from 'js-waku';
import protobuf from 'protobufjs';
import * as sigUtil from 'eth-sig-util';
import React, { useContext, useEffect, useState } from 'react';
import { signReport } from '../../crypto';
import { FormContext } from '../../context/FormProvider';

const ContentTopic = `/relay-reactjs-chat/1/chat/proto`;
// publicEncryptionKey holds the public key of the vulnerability receiver.
const publicEncryptionKey = process.env.NEXT_PUBLIC_PUBLIC_ENCRYPTION_KEY;

const FormMessage = new protobuf.Type('FormMessage')
  .add(new protobuf.Field('timestamp', 1, 'uint64'))
  .add(new protobuf.Field('payload', 2, 'bytes'))
  .add(new protobuf.Field('ethAddress', 3, 'bytes'))
  .add(new protobuf.Field('signature', 4, 'bytes'));

const FormSender: React.FC = () => {
  const [waku, setWaku] = useState<Waku | undefined>(undefined);
  const [wakuStatus, setWakuStatus] = useState('None');
  // Using a counter just for the messages to be different
  const { active, connector, account } = useWeb3React();
  const { questions, answers } = useContext(FormContext);

  useEffect(() => {
    if (!!waku) return;
    if (wakuStatus !== 'None') return;

    setWakuStatus('Starting');

    Waku.create({ bootstrap: { default: true } }).then((waku) => {
      setWaku(waku);
      setWakuStatus('Connecting');
      waku.waitForRemotePeer().then(() => {
        setWakuStatus('Ready');
      });
    });
  }, [waku, wakuStatus]);

  const sendMessageOnClick = async () => {
    // Check Waku is started and connected first.
    if (
      waku == undefined ||
      wakuStatus !== 'Ready' ||
      !active ||
      !account ||
      !connector ||
      !publicEncryptionKey
    )
      return;

    // TODO think about a better way to encode questions and answers
    const msg = JSON.stringify({ questions, answers });
    const provider = await connector?.getProvider();

    sendMessage(
      msg,
      waku,
      new Date(),
      account,
      publicEncryptionKey,
      provider.request
    ).then(() => console.log('Message sent'));
  };

  return (
    <Box display='flex' justifyContent='center'>
      <Box
        sx={{
          p: 2,
          border: '1px solid grey',
          borderRadius: '5px',
          width: '100%',
          maxWidth: '770px',
        }}
      >
        <p>
          Waku Network Connection State: <b>{wakuStatus}</b>
        </p>
        {!active && <p>Please connect your wallet.</p>}
        <Button
          onClick={sendMessageOnClick}
          variant='contained'
          disabled={
            wakuStatus !== 'Ready' ||
            !active ||
            !connector ||
            !publicEncryptionKey
          }
        >
          Send Form
        </Button>
      </Box>
    </Box>
  );
};

async function sendMessage(
  message: string,
  waku: Waku,
  timestamp: Date,
  account: string,
  publicEncryptionKey: string,
  providerRequest: (request: {
    method: string;
    params?: Array<any>;
    from?: string;
  }) => Promise<any>
) {
  const time = timestamp.getTime();

  console.log('Sending Message: ' + message);

  // Encode to protobuf
  const protoMsg = FormMessage.create({
    timestamp: time,
    payload: new TextEncoder().encode(message),
    ethAddress: new TextEncoder().encode(account),
    signature: new TextEncoder().encode(
      await signReport(message, account, providerRequest)
    ),
  });
  const payload = FormMessage.encode(protoMsg).finish();

  const encObj = sigUtil.encrypt(
    Buffer.from(publicEncryptionKey, 'base64').toString('base64'),
    { data: utils.bytesToHex(payload) },
    'x25519-xsalsa20-poly1305'
  );

  const encryptedPayload = Buffer.from(JSON.stringify(encObj), 'utf8');
  console.log(JSON.stringify(encObj));

  // Wrap in a Waku Message
  return WakuMessage.fromBytes(encryptedPayload, ContentTopic).then(
    (wakuMessage) =>
      // Send over Waku Relay
      waku.relay.send(wakuMessage)
  );
}

export default FormSender;
