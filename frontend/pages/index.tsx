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
import { useState } from 'react';
import Stack from '@mui/material/Stack';
import FormBuilder from '../components/formBuilder/FormBuilder';

const Home: NextPage = () => {
  const [tab, setTab] = useState<string>('1');

  const handleChange = (newValue: string) => {
    setTab(newValue);
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
      <Headline />
      <TabContext value={tab}>
        <Box
          sx={{ borderBottom: 1, borderColor: 'divider', background: 'white' }}
          display='flex'
          justifyContent='center'
        >
          <TabList
            onChange={(_, value) => handleChange(value)}
            aria-label='lab API tabs example'
          >
            <Tab label='Respond' value='1' />
            <Tab label='Edit' value='2' />
          </TabList>
        </Box>
        <TabPanel value='1'>
          <Stack spacing={2} direction='column'>
            <Form />
            <FormSender />
          </Stack>
        </TabPanel>
        <TabPanel value='2'>
          <FormBuilder />
        </TabPanel>
      </TabContext>
    </div>
  );
};

export default Home;
