import React, { useState } from 'react';
import { useGoClipboard } from '@src/hooks/use-go-clipboard/useGoClipboard';
import styles from './Password.module.css';

import { Button } from '@mantine/core';

type PasswordProps = {
    password: string;
    setPassword?: (e: React.ChangeEvent) => {};
};

function Password({ password, setPassword }: PasswordProps) {
    const [inputType, setInputType] = useState('password');
    const { copy, error, copied } = useGoClipboard();

    const toggleInputType = (e: React.MouseEvent) => {
        e.preventDefault();
        if (inputType === 'text') {
            setInputType('password');
        } else {
            setInputType('text');
        }
    };

    const handleCopy = (e: React.MouseEvent) => {
        e.preventDefault();
        copy(password);
    };

    return (
        <div className={styles.inputContainer}>
            <input
                type={inputType}
                value={password}
                className={styles.input}
                onChange={setPassword}
            />
            <Button
                color="cyan"
                sx={{ height: '100%', marginLeft: '10px', width: '8rem' }}
                onClick={toggleInputType}
            >
                {inputType === 'password' ? 'SHOW' : 'HIDE'}
            </Button>
            <Button
                color={copied ? 'green' : 'blue'}
                sx={{ height: '100%', marginLeft: '10px', width: '8rem' }}
                onClick={handleCopy}
            >
                {copied ? 'COPIED' : 'COPY'}
            </Button>
        </div>
    );
}

export default Password;
