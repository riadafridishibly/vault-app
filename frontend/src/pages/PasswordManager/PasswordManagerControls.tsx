import { Button, Group } from '@mantine/core';
import { passwordItemFormMode } from '@src/shared/components/PasswordItemModal';
import { IconPlus } from '@tabler/icons';
import React from 'react';
import { atom, useSetRecoilState } from 'recoil';

export const showNewPasswordCreateModal = atom<boolean>({
    key: 'showCreateModal',
    default: false,
});

function AddNewPassword() {
    const setShow = useSetRecoilState(showNewPasswordCreateModal);
    const setPasswordItemMode = useSetRecoilState(passwordItemFormMode);
    return (
        <Button
            onClick={(e: React.MouseEvent) => {
                e.preventDefault();
                setPasswordItemMode('create');
                setShow(true);
            }}
            leftIcon={<IconPlus />}
        >
            Add
        </Button>
    );
}

function PasswordManagerControls() {
    return (
        <Group>
            <AddNewPassword />
        </Group>
    );
}

export default PasswordManagerControls;
