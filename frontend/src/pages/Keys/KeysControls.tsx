import { Button, Group } from '@mantine/core'
import { IconPlus } from '@tabler/icons';
import { atom, useSetRecoilState } from 'recoil';

export const showNewKeyAddModal = atom<boolean>({
    key: 'showNewKeyAddModal',
    default: false,
})

function AddNewKey() {
    const setShow = useSetRecoilState(showNewKeyAddModal)
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

function KeysControls() {
    return (
        <Group>
            <AddNewKey />
        </Group>
    )
}

export default KeysControls