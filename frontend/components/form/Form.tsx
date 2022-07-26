import React, { useState } from 'react';
import TextField from '@mui/material/TextField';
import Stack from '@mui/material/Stack';
import Box from '@mui/material/Box';
import FormControl from '@mui/material/FormControl';
import InputLabel from '@mui/material/InputLabel';
import Select from '@mui/material/Select';
import MenuItem from '@mui/material/MenuItem';

enum Severity {
  Info = 1,
  Low,
  Medium,
  High,
  Critical,
}

const Form: React.FC = () => {
  return (
    <Box display='flex' justifyContent='center'>
      <Box
        sx={{
          width: '100%',
          maxWidth: '770px',
        }}
      >
        <Stack spacing={2} direction='column'>
          <Box
            sx={{
              p: 2,
              border: '1px solid grey',
              borderRadius: '5px',
              background: 'white',
            }}
          >
            <Stack spacing={0} direction='column'>
              <h4 style={{ marginTop: '0' }}>Summary</h4>
              <TextField id='standard-basic' variant='standard' />
            </Stack>
          </Box>
          <Box
            sx={{
              p: 2,
              border: '1px solid grey',
              borderRadius: '5px',
              background: 'white',
            }}
          >
            <Stack spacing={0} direction='column'>
              <h4 style={{ marginTop: '0' }}>CVSS Score</h4>
              <FormControl fullWidth>
                <InputLabel id='demo-simple-select-label'>
                  CVSS Score
                </InputLabel>
                <Select
                  labelId='demo-simple-select-label'
                  id='demo-simple-select'
                  // value={age}
                  label='Age'
                  // onChange={handleChange}
                >
                  <MenuItem value={Severity.Info}>Info</MenuItem>
                  <MenuItem value={Severity.Low}>Low</MenuItem>
                  <MenuItem value={Severity.Medium}>Medium</MenuItem>
                  <MenuItem value={Severity.High}>High</MenuItem>
                  <MenuItem value={Severity.Critical}>Critical</MenuItem>
                </Select>
              </FormControl>
            </Stack>
          </Box>
          <Box
            sx={{
              p: 2,
              border: '1px solid grey',
              borderRadius: '5px',
              background: 'white',
            }}
          >
            <Stack spacing={0} direction='column'>
              <h4 style={{ marginTop: '0' }}>Description</h4>
              <TextField
                id='standard-basic'
                variant='standard'
                multiline
                rows={4}
                maxRows={4}
              />
            </Stack>
          </Box>
        </Stack>
      </Box>
    </Box>
  );
};

export default Form;
