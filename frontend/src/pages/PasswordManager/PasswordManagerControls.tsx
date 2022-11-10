import { Button, Group } from '@mantine/core'
import { IconPlus } from '@tabler/icons';
import React from 'react'
import { atom, useSetRecoilState } from 'recoil';

export const showNewPasswordCreateModal = atom<boolean>({
    key: 'showCreateModal',
    default: false,
});

function AddNewPassword() {
    const setShow = useSetRecoilState(showNewPasswordCreateModal)
    return (
        <Button
            onClick={(e: React.MouseEvent) => {
                e.preventDefault();
                setShow(true)
            }}
            leftIcon={<IconPlus />}
        >
            Add
        </Button>
    )
}

function PasswordManagerControls() {
    return (
        <Group>
            <AddNewPassword />
        </Group>
    )
}

export default PasswordManagerControls