export async function signReport(
  report: string,
  fromAddress: string,
  providerRequest: (request: {
    method: string;
    params?: Array<any>;
    from?: string;
  }) => Promise<any>
): Promise<string> {
  const msgParams = buildMsgParams(report, fromAddress);

  const result = await providerRequest({
    method: 'eth_signTypedData_v4',
    params: [fromAddress, msgParams],
    from: fromAddress,
  });

  console.log('TYPED SIGNED:' + JSON.stringify(result));

  return JSON.stringify(result);
}

//TODO change chaineId
function buildMsgParams(report: string, fromAddress: string) {
  return JSON.stringify({
    domain: {
      chainId: 5,
      name: 'Ethereum Private Message over Waku',
      version: '1',
    },
    message: {
      report,
      fromAddress,
    },
    // Refers to the keys of the *types* object below.
    primaryType: 'Report',
    types: {
      EIP712Domain: [
        { name: 'name', type: 'string' },
        { name: 'version', type: 'string' },
        { name: 'chainId', type: 'uint256' },
      ],
      Report: [
        { name: 'report', type: 'string' },
        { name: 'fromAddress', type: 'string' },
      ],
    },
  });
}
