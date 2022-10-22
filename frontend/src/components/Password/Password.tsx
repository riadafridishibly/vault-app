import React, { useEffect, useState } from 'react'
import { useClipboard } from '../../hooks/use-clipboard/useClipboard'
import styles from './Password.module.css'

type PasswordProps = {
    password: string,
    setPassword?: (e: React.ChangeEvent) => {}
}


function Password({ password, setPassword }: PasswordProps) {
    const [inputType, setInputType] = useState('password')
    const { copy, error, copied } = useClipboard()

    const toggleInputType = (e: React.MouseEvent) => {
        e.preventDefault()
        if (inputType === 'text') {
            setInputType('password')
        } else {
            setInputType('text')
        }
    }

    const handleCopy = (e: React.MouseEvent) => {
        e.preventDefault()
        copy(password)
    }

    return (
        <div className={styles.inputContainer}>
            <input type={inputType} value={password} className={styles.input} placeholder="Something..." onChange={setPassword} />
            <button onClick={toggleInputType} className={styles.button}>{
                inputType === 'password' ? 'SHOW' : 'HIDE'
            }</button>
            <button onClick={handleCopy} className={styles.button}> {copied ? 'COPIED' : 'COPY'} </button>
        </div>
    )
}

export default Password