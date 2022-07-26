import type { NextPage } from 'next';
import Head from 'next/head';
import { useState, useEffect, useCallback } from 'react';
import { Waku, WakuMessage, utils } from 'js-waku';
import protobuf from 'protobufjs';
import styles from '../styles/Home.module.css';
import WalletConnector from '../components/walletConnector/WalletConnector';
import * as sigUtil from 'eth-sig-util';
import { signReport } from '../crypto';
import { useWeb3React } from '@web3-react/core';

interface Message {
  timestamp: Date;
  payload: string;
}

const ContentTopic = `/relay-reactjs-chat/1/chat/proto`;
// encPublicKey holds the public key of the vulnerability receiver.
const encPublicKey = 'Pf8fRRhD8d+a/YlBah2tmxgXLjgDpXlS7JtwnX+r5U8=';

const SimpleChatMessage = new protobuf.Type('SimpleChatMessage')
  .add(new protobuf.Field('timestamp', 1, 'uint64'))
  .add(new protobuf.Field('payload', 2, 'bytes'))
  .add(new protobuf.Field('ethAddress', 3, 'bytes'))
  .add(new protobuf.Field('signature', 4, 'bytes'));

const Home: NextPage = () => {
  const [waku, setWaku] = useState<Waku | undefined>(undefined);
  const [wakuStatus, setWakuStatus] = useState('None');
  // Using a counter just for the messages to be different
  const [sendCounter, setSendCounter] = useState(0);
  const [messages, setMessages] = useState<Message[]>([]);
  const { active, connector, account } = useWeb3React();

  useEffect(() => {
    if (!!waku) return;
    if (wakuStatus !== 'None') return;

    setWakuStatus('Starting');

    Waku.create({ bootstrap: { default: true } }).then((waku) => {
      setWaku(waku);
      setWakuStatus('Connecting');
      waku.waitForRemotePeer().then(() => {
        console.log(waku.relay.peers);
        setWakuStatus('Ready');
      });
    });
  }, [waku, wakuStatus]);

  const processIncomingMessage = useCallback((wakuMessage: WakuMessage) => {
    if (!wakuMessage.payload) return;

    const { payload, timestamp } = SimpleChatMessage.decode(
      wakuMessage.payload
    ) as unknown as { payload: Uint8Array; timestamp: number };

    const time = new Date();
    time.setTime(timestamp);
    const message = {
      payload: new TextDecoder('utf-8').decode(payload),
      timestamp: time,
    };

    setMessages((prev) => {
      return [message].concat(prev);
    });
  }, []);

  useEffect(() => {
    if (!waku) return;

    // Pass the content topic to only process messages related to your dApp
    waku.relay.addObserver(processIncomingMessage, [ContentTopic]);

    // `cleanUp` is called when the component is unmounted, see ReactJS doc.
    return function cleanUp() {
      waku.relay.deleteObserver(processIncomingMessage, [ContentTopic]);
    };
  }, [waku, wakuStatus, processIncomingMessage]);

  const sendMessageOnClick = async () => {
    // Check Waku is started and connected first.
    if (
      waku == undefined ||
      wakuStatus !== 'Ready' ||
      !active ||
      !account ||
      !connector
    )
      return;

    const provider = await connector?.getProvider();

    sendMessage(
      `Here is message #${sendCounter}`,
      waku,
      new Date(),
      account,
      provider.request
    ).then(() => console.log('Message sent'));

    // For demonstration purposes.
    setSendCounter(sendCounter + 1);
  };

  return (
    <div className='App'>
      <Head>
        <title>Famed Submission Form</title>
        <meta
          name='description'
          content='Submit Bug Reports to the Famed Protocol'
        />
        <link rel='icon' href='/favicon.ico' />
      </Head>
      <p>{wakuStatus}</p>
      <button onClick={sendMessageOnClick} disabled={wakuStatus !== 'Ready'}>
        Send Message
      </button>
      <ul>
        {messages.map((msg) => {
          return (
            <li key={msg.timestamp.valueOf()}>
              <p>
                {msg.timestamp.toString()}: {msg.payload}
              </p>
            </li>
          );
        })}
      </ul>
      <WalletConnector />
    </div>
  );
};

async function sendMessage(
  message: string,
  waku: Waku,
  timestamp: Date,
  account: string,
  providerRequest: (request: {
    method: string;
    params?: Array<any>;
    from?: string;
  }) => Promise<any>
) {
  const time = timestamp.getTime();

  // Encode to protobuf
  const protoMsg = SimpleChatMessage.create({
    timestamp: time,
    payload: new TextEncoder().encode(message),
    ethAddress: new TextEncoder().encode(account),
    signature: new TextEncoder().encode(
      await signReport(message, account, providerRequest)
    ),
  });
  const payload = SimpleChatMessage.encode(protoMsg).finish();

  const encObj = sigUtil.encrypt(
    Buffer.from(encPublicKey, 'base64').toString('base64'),
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

export default Home;
