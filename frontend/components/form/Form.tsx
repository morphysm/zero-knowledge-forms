import React, { useContext, useState } from 'react';
import TextField from '@mui/material/TextField';
import Stack from '@mui/material/Stack';
import Box from '@mui/material/Box';
import { FormContext } from '../../context/FormProvider';

const Form: React.FC = () => {
  const { questions, answers, setAnswer } = useContext(FormContext);

  const handleAnswerChange = (answer: string, index: number) => {
    setAnswer(answer, index);
  };

  return (
    <Box display='flex' justifyContent='center'>
      <Box
        sx={{
          width: '100%',
          maxWidth: '770px',
        }}
      >
        <Stack spacing={2} direction='column'>
          {questions.map((question, i) => {
            return (
              <Box
                sx={{
                  p: 2,
                  border: '1px solid grey',
                  borderRadius: '5px',
                  background: 'white',
                }}
                key={`form-builder-${question}-${i}`}
              >
                <Stack spacing={0} direction='column'>
                  <h4 style={{ marginTop: '0' }}>{question}</h4>
                  <TextField
                    id='standard-basic'
                    variant='standard'
                    defaultValue={answers[i]}
                    onChange={(e) => handleAnswerChange(e.target.value, i)}
                  />
                </Stack>
              </Box>
            );
          })}
        </Stack>
      </Box>
    </Box>
  );
};

export default Form;
