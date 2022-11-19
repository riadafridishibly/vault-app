import { Button, Group, Modal, Text } from '@mantine/core';
import { passwordManagerItems } from '@src/pages/PasswordManager/Table';
import { DeletePasswordItem, ReadAllPasswordItems } from '@wailsjs/go/main/DatabaseService';
import React from 'react';
import { atom, useRecoilState, useSetRecoilState } from 'recoil';
import { modalButtonMargin, modalButtonSize } from '../constants/modal.constants';

export const openDeletePasswordItemModal = atom<{ opened: boolean; id: number }>({
    key: 'openDeletePasswordItemModal',
    default: {
        opened: false,
        id: -1,
    },
});

const DeletePasswordItemModal = () => {
    const [openModal, setOpenModal] = useRecoilState(openDeletePasswordItemModal);
    const setPasswordItems = useSetRecoilState(passwordManagerItems);

    const handleDeletePasswordItem = async () => {
        try {
            await DeletePasswordItem(openModal.id);
        } catch (err) {
            console.log(err);
        } finally {
            setPasswordItems(await ReadAllPasswordItems());
            setOpenModal({ opened: false, id: -1 });
        }
    };

    return (
        <Modal
            centered={true}
            size="md"
            opened={openModal.opened}
            onClose={() => setOpenModal({ opened: false, id: -1 })}
        >
            <Text align="center" color={'yellow'}>
                Are you sure you want to delete the password?
            </Text>
            <Group position="center">
                <Button
                    color={'red'}
                    size={modalButtonSize}
                    m={modalButtonMargin}
                    onClick={handleDeletePasswordItem}
                >
                    Delete
                </Button>
                <Button
                    size={modalButtonSize}
                    m={modalButtonMargin}
                    onClick={() => setOpenModal({ opened: false, id: -1 })}
                >
                    Cancel
                </Button>
            </Group>
        </Modal>
    );
};

export default DeletePasswordItemModal;
