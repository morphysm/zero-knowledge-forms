import React from 'react';

interface FormProps {
  answers: string[];
  questions: string[];
  setAnswer: (answer: string, index: number) => void;
  setQuestion: (questions: string, index: number) => void;
  addQuestion: () => void;
  removeQuestion: (index: number) => void;
}

type FormContextProps = {
  children?: React.ReactNode;
};

export const FormContext = React.createContext({} as FormProps);

const FormProvider: React.FC<FormContextProps> = ({
  children,
}: FormContextProps) => {
  const [answers, setAnswers] = React.useState<string[]>(['', '']);
  const [questions, setQuestions] = React.useState<string[]>([
    'What is you favourite meal?',
    'What is your favourite programming language?',
  ]);

  const setAnswer = (answer: string, index: number) => {
    setAnswers((prev) => {
      prev[index] = answer;
      return prev;
    });
  };

  const setQuestion = (question: string, index: number) => {
    setQuestions((prev) => {
      prev[index] = question;
      return prev;
    });
  };

  const addQuestion = () => {
    setQuestions((prev) => {
      return [...prev, ''];
    });
    setAnswers((prev) => {
      return [...prev, ''];
    });
  };

  const removeQuestion = (index: number) => {
    setQuestions((prev) => {
      console.log([
        ...prev.slice(0, index),
        ...prev.slice(index + 1, prev.length),
      ]);
      return [...prev.slice(0, index), ...prev.slice(index + 1, prev.length)];
    });
    setAnswers((prev) => {
      return [...prev.slice(0, index), ...prev.slice(index + 1, prev.length)];
    });
  };

  return (
    <FormContext.Provider
      value={{
        answers,
        setAnswer,
        questions,
        setQuestion,
        addQuestion,
        removeQuestion,
      }}
    >
      {children}
    </FormContext.Provider>
  );
};

export default FormProvider;
