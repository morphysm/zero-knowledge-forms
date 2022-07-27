import React, { useContext, useEffect } from 'react';
import TextField from '@mui/material/TextField';
import Stack from '@mui/material/Stack';
import Box from '@mui/material/Box';
import { FormContext } from '../../context/FormProvider';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import IconButton from '@mui/material/IconButton';
import DeleteIcon from '@mui/icons-material/Delete';

const FormBuilder: React.FC = () => {
  const { questions, addQuestion, setQuestion, removeQuestion } =
    useContext(FormContext);

  const handleAddQuestionClick = () => {
    addQuestion();
  };

  const handleQuestionValueChange = (value: string, index: number) => {
    setQuestion(value, index);
  };

  const handleRemoveClick = (index: number) => {
    removeQuestion(index);
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
          <Button variant='contained' onClick={handleAddQuestionClick}>
            Add Question
          </Button>
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
                  <TextField
                    id='standard-basic'
                    variant='standard'
                    defaultValue={question}
                    onChange={(e) =>
                      handleQuestionValueChange(e.target.value, i)
                    }
                  />
                  <Stack
                    alignItems='center'
                    spacing={2}
                    direction='row'
                    justifyContent='space-between'
                    mt={2}
                  >
                    <Typography variant='subtitle2' component='div'>
                      Answer Text
                    </Typography>
                    <IconButton
                      size='large'
                      onClick={() => handleRemoveClick(i)}
                    >
                      <DeleteIcon fontSize='small' />
                    </IconButton>
                  </Stack>
                </Stack>
              </Box>
            );
          })}
        </Stack>
      </Box>
    </Box>
  );
};

export default FormBuilder;
