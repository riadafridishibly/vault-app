import { Button, Group } from '@mantine/core';
import React from 'react'
import { atom, useSetRecoilState } from 'recoil';
import { OpenDialog } from '../../../wailsjs/go/main/App';


export const selectedFiles = atom<string[]>({
    key: 'selectedFiles',
    default: []
})


function OpenFiles() {
    const setFiles = useSetRecoilState(selectedFiles)
    const openFileDialog = () => {
        OpenDialog()
            .then((files) => {
                if (files) {
                    setFiles((curr) => [...curr, ...files])
                }
            })
            .catch((err) => console.error(err));
    };
    return (
        <Button
            onClick={(e: React.MouseEvent) => {
                e.preventDefault();
                openFileDialog();
            }}
        >
            Open Files
        </Button>
    );
}

function EncryptFilesControls() {
    return (
        <Group>
            <OpenFiles />
        </Group>
    )
}

export default EncryptFilesControls