import React from 'react'
import styles from './InputGroup.module.css'

type InputGroupProps = {
    name: string,
    value: string,
    onChange: React.ChangeEventHandler<HTMLInputElement>
}

function InputGroup({ name, value, onChange }: InputGroupProps) {
    return (
        <div className={styles.container}>
            <div className={styles.group}>
                <span>{name}</span>
            </div>
            <input type="text" onChange={onChange} value={value} />
        </div>
    )
}

export default InputGroup