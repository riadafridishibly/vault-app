import { Button, Group, Loader, LoadingOverlay, Modal, Text } from '@mantine/core';
import { CheckPasswordBreach } from '@wailsjs/go/main/PasswordService';
import { useCallback, useEffect, useState } from 'react';
import { atom, useRecoilState } from 'recoil';
import { modalButtonColor, modalButtonMargin, modalButtonSize } from '../constants/modal.constants';

export const showPasswordBreachModal = atom<{
    opened: boolean;
    password: string;
}>({
    key: 'showPasswordBreachModal',
    default: {
        opened: false,
        password: '',
    },
});

const PwnCheckModal = () => {
    const [passwordBreachModalState, setPasswordBreachModalState] =
        useRecoilState(showPasswordBreachModal);
    const [color, setColor] = useState<'green' | 'red'>('green');
    const [loading, setLoading] = useState(false);
    const [msg, setMsg] = useState('No password breach found');

    const apiCallFn = useCallback(async () => {
        let count = 0;
        try {
            setLoading(true);
            count = await CheckPasswordBreach(passwordBreachModalState.password);
        } catch (ex) {
            console.error(ex);
        } finally {
            setLoading(false);
            setColor(count > 0 ? 'red' : 'green');
            if (count > 0) {
                setMsg(`Password breach found in ${count} instances!`);
            }
        }
    }, [passwordBreachModalState.password]);

    useEffect(() => {
        apiCallFn();
    }, [passwordBreachModalState.opened]);

    return (
        <>
            <Modal
                centered={true}
                opened={passwordBreachModalState.opened}
                onClose={() =>
                    setPasswordBreachModalState({
                        opened: false,
                        password: '',
                    })
                }
            >
                <div>
                    <LoadingOverlay
                        loader={<Loader size="lg" variant="dots" />}
                        overlayOpacity={1}
                        visible={loading}
                    />
                    <Text color={color} align="center">
                        {msg}
                    </Text>
                    <Group position="center">
                        <Button
                            size={modalButtonSize}
                            m={modalButtonMargin}
                            bg={modalButtonColor}
                            onClick={() =>
                                setPasswordBreachModalState({
                                    opened: false,
                                    password: '',
                                })
                            }
                        >
                            OK
                        </Button>
                    </Group>
                </div>
            </Modal>
        </>
    );
};

export default PwnCheckModal;
