import React, { useEffect, useRef, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { IsUserExist, Register, Login } from '@wailsjs/go/main/UserService';
function Auth() {
    const navigate = useNavigate();
    const [hasUser, setHasUser] = useState(false);
    const inputRef = useRef<HTMLInputElement | null>(null);
    const handleSubmit = async (e: React.MouseEvent<HTMLFormElement>) => {
        e.preventDefault();
        const password = inputRef?.current?.value;
        if (!password) {
            // Do nothing; Show error
            return;
        }

        if (hasUser) {
            Login(password)
                .then((ok) => navigate('/app'))
                .catch((err) => console.error(err));
            return;
        }

        // FIXME: to send only error or not to send only error?
        // should we send pair of keys?
        await Register(password);

        navigate('/app');
    };

    useEffect(() => {
        IsUserExist()
            .then((ok) => {
                if (ok) {
                    setHasUser(true);
                }
                return ok;
            })
            .catch((err) => console.error(err));
    }, []);
    return (
        <form onSubmit={handleSubmit}>
            <label>
                {hasUser ? 'Login' : 'Register'}
                <input ref={inputRef} />
            </label>
            <button>Log In</button>
        </form>
    );
}

export default Auth;
