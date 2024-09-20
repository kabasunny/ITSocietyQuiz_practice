import { Dispatch, SetStateAction } from 'react';

const goToNextQuestion = (
  setNext: Dispatch<SetStateAction<boolean>>,
) => {
  setNext(false);
};

export default goToNextQuestion;
