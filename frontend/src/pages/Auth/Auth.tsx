import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { IsUserExist, Register, Login } from '@wailsjs/go/main/UserService';

import { Box, Progress, PasswordInput, Group, Text, Center, Flex } from '@mantine/core';
import { useInputState } from '@mantine/hooks';
import { IconCheck, IconX } from '@tabler/icons';

function PasswordRequirement({ meets, label }: { meets: boolean; label: string }) {
    return (
        <Text color={meets ? 'teal' : 'red'} mt={5} size="sm">
            <Center inline>
                {meets ? <IconCheck size={14} stroke={1.5} /> : <IconX size={14} stroke={1.5} />}
                <Box ml={7}>{label}</Box>
            </Center>
        </Text>
    );
}

const requirements = [
    { re: /[0-9]/, label: 'Includes number' },
    { re: /[a-z]/, label: 'Includes lowercase letter' },
    { re: /[A-Z]/, label: 'Includes uppercase letter' },
    { re: /[$&+,:;=?@#|'<>.^*()%!-]/, label: 'Includes special symbol' },
];

function getStrength(password: string) {
    let multiplier = password.length > 5 ? 0 : 1;

    requirements.forEach((requirement) => {
        if (!requirement.re.test(password)) {
            multiplier += 1;
        }
    });

    return Math.max(100 - (100 / (requirements.length + 1)) * multiplier, 0);
}

function Auth() {
    const navigate = useNavigate();
    const [hasUser, setHasUser] = useState(false);
    const [password, setPassword] = useInputState('');
    const [reload, setReload] = useState(0);
    const handleSubmit = async (e: React.MouseEvent<HTMLFormElement>) => {
        e.preventDefault();
        if (!password) {
            // Do nothing
            // TODO: Show error
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
        Register(password)
            .then((ok) => {
                setReload((v) => v + 1);
                setPassword('');
            })
            .catch((err) => console.error(err));
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
    }, [reload]);

    const strength = getStrength(password);
    const checks = requirements.map((requirement, index) => (
        <PasswordRequirement
            key={index}
            label={requirement.label}
            meets={requirement.re.test(password)}
        />
    ));
    const bars = Array(4)
        .fill(0)
        .map((_, index) => (
            <Progress
                styles={{ bar: { transitionDuration: '0ms' } }}
                value={
                    password.length > 0 && index === 0
                        ? 100
                        : strength >= ((index + 1) / 4) * 100
                        ? 100
                        : 0
                }
                color={strength > 80 ? 'teal' : strength > 50 ? 'yellow' : 'red'}
                key={index}
                size={4}
            />
        ));

    return (
        <div>
            <Center style={{ width: '100vw', height: '95vh' }}>
                <Flex
                    mih={50}
                    gap="md"
                    justify="center"
                    align="center"
                    direction="column"
                    wrap="nowrap"
                >
                    <Text
                        sx={() => ({
                            fontSize: 32,
                            fontWeight: 300,
                        })}
                    >
                        {hasUser ? 'Login' : 'Register'}
                    </Text>
                    <div style={{ minWidth: '300px' }}>
                        <form onSubmit={handleSubmit}>
                            <PasswordInput
                                value={password}
                                onChange={setPassword}
                                placeholder="Your password"
                                label="Master Password"
                                required
                            />
                        </form>

                        {!hasUser && (
                            <>
                                <Group spacing={5} grow mt="xs" mb="md">
                                    {bars}
                                </Group>

                                <PasswordRequirement
                                    label="Has at least 6 characters"
                                    meets={password.length > 5}
                                />
                                {checks}
                            </>
                        )}
                    </div>
                </Flex>
            </Center>
        </div>
    );
}

// function Auth() {
//     const navigate = useNavigate();
//     const [hasUser, setHasUser] = useState(false);
//     const [password, setPassword] = useInputState('');
//     const inputRef = useRef<HTMLInputElement | null>(null);
//     const handleSubmit = async (e: React.MouseEvent<HTMLFormElement>) => {
//         e.preventDefault();
//         const password = inputRef?.current?.value;
//         if (!password) {
//             // Do nothing; Show error
//             return;
//         }

//         if (hasUser) {
//             Login(password)
//                 .then((ok) => navigate('/app'))
//                 .catch((err) => console.error(err));
//             return;
//         }

//         // FIXME: to send only error or not to send only error?
//         // should we send pair of keys?
//         await Register(password);

//         navigate('/app');
//     };

//     useEffect(() => {
//         IsUserExist()
//             .then((ok) => {
//                 if (ok) {
//                     setHasUser(true);
//                 }
//                 return ok;
//             })
//             .catch((err) => console.error(err));
//     }, []);
//     return (
//         <>
//             <Center style={{ width: '100vw', height: '95vh' }}>
//                 <div style={{ minWidth: '300px' }}>
//                     <PasswordStrength
//                         isLogin={hasUser}
//                         password={password}
//                         setPassword={setPassword}
//                     />
//                 </div>
//             </Center>
//         </>
//     );
// }

export default Auth;
