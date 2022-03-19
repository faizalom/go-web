import { useState } from 'react';

const useInputAjax = () => {
  const [enteredValue, setEnteredValue] = useState('');
  const [isTouched, setIsTouched] = useState(false);
  const [valueIsValid, setValueIsValid] = useState(false);
  const [isLoading, setLoading] = useState(false);
  const [message, setMessage] = useState("");

  const hasError = !valueIsValid && isTouched && !isLoading;

  const valueChangeHandler = (event) => {
    setIsTouched(true);
    setEnteredValue(event.target.value);
  };

  const inputBlurHandler = (event) => {
    setIsTouched(true);
  };

  const reset = () => {
    setEnteredValue('');
    setIsTouched(false);
  };

  return {
    value: enteredValue,
    setValue: setEnteredValue,
    isValid: valueIsValid,
    setIsValid: setValueIsValid,
    isLoading,
    setLoading,
    message,
    setMessage,
    hasError,
    isTouched,
    valueChangeHandler,
    inputBlurHandler,
    reset
  };
};

export default useInputAjax;