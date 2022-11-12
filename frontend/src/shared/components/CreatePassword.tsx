import { Button, Group, Modal, TextInput } from '@mantine/core';
import { useForm } from '@mantine/form';
import { IconClipboard } from '@tabler/icons';
import React from 'react';
import { useRecoilState } from 'recoil';
import { useGoClipboard } from '@src/hooks/use-go-clipboard/useGoClipboard';
import { showNewPasswordCreateModal } from '@src/pages/PasswordManager/PasswordManagerControls';
import {
    modalButtonColor,
    modalButtonMargin,
    modalButtonSize,
    modalTextInputMargin,
} from '../constants/modal.constants';

// TODO: model may change later on
class UserPass {
    website: string = '';
    username: string = '';
    password: string = '';
}

function CreatePassword() {
    const { error, paste } = useGoClipboard();
    const [opened, setOpened] = useRecoilState(showNewPasswordCreateModal);
    const userPassForm = useForm<UserPass>({
        initialValues: new UserPass(),
    });
    const handleModalOpen = () => {
        userPassForm.reset();
        setOpened(true);
    };
    const handlePasswordSave = userPassForm.onSubmit((values) => {
        //TODO: implement actual password saving mechanism
        console.log(values);
        setOpened(false);
    });

    const handlePaste = async (e: React.MouseEvent, propName: string) => {
        e.preventDefault();
        userPassForm.setFieldValue(propName, await paste());
    };

    const pasteButton = (propName: string) => (
        <Button
            bg={modalButtonColor}
            size="sm"
            p="5px"
            onClick={(e: React.MouseEvent) => handlePaste(e, propName)}
        >
            <IconClipboard />
        </Button>
    );

    return (
        <>
            <Modal
                centered={true}
                opened={opened}
                onClose={() => setOpened(false)}
                title="Add New Password"
            >
                <form onSubmit={handlePasswordSave}>
                    <TextInput
                        m={modalTextInputMargin}
                        placeholder="Website"
                        label="Website"
                        withAsterisk
                        required
                        {...userPassForm.getInputProps('website')}
                        rightSection={pasteButton('website')}
                    />
                    <TextInput
                        m={modalTextInputMargin}
                        placeholder="Username"
                        label="Username"
                        withAsterisk
                        required
                        {...userPassForm.getInputProps('username')}
                        rightSection={pasteButton('username')}
                    />
                    <TextInput
                        m={modalTextInputMargin}
                        placeholder="Password"
                        label="Password"
                        withAsterisk
                        required
                        {...userPassForm.getInputProps('password')}
                        rightSection={pasteButton('password')}
                    />
                    <Group position="center">
                        <Button
                            bg={modalButtonColor}
                            m={modalButtonMargin}
                            size={modalButtonSize}
                            type="submit"
                        >
                            Save
                        </Button>
                        <Button
                            bg={modalButtonColor}
                            m={modalButtonMargin}
                            size={modalButtonSize}
                            type="reset"
                            onClick={() => setOpened(false)}
                        >
                            Cancel
                        </Button>
                    </Group>
                </form>
            </Modal>
        </>
    );
}

export default CreatePassword;
