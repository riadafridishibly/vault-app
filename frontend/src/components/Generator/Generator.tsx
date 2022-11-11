import { useEffect, useState } from 'react';
import { GenerateNewPassword } from '@wailsjs/go/main/App';
import InputGroup from '../InputGroup/InputGroup';
import Password from '../Password/Password';
import styles from './Generator.module.css';

function Generator() {
    const [seed, setSeed] = useState('');
    const [password, setPassword] = useState('');
    const [generatedPassword, setGeneratedPassword] = useState('');

    useEffect(() => {
        GenerateNewPassword(seed, password)
            .then((value) => {
                setGeneratedPassword(value);
            })
            .catch((err) => console.error(err));
    }, [seed, password]);

    return (
        <div className={styles.container}>
            <div className={styles.inputWrapper}>
                <InputGroup name="Seed" value={seed} onChange={(e) => setSeed(e.target.value)} />
                <InputGroup
                    name="Password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                />
            </div>
            <h1> Generated Password </h1>
            <Password password={generatedPassword} />
        </div>
    );
}

export default Generator;
