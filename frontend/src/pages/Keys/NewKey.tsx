import { Button, Center, Modal, Tabs } from '@mantine/core';
import { CreateKey } from '@wailsjs/go/main/DatabaseService';
import { useRecoilState } from 'recoil';
import { showNewKeyAddModal } from './KeysControls';
import SSHKeyGen from './SSHKeyGen';

function NewKey() {
    const [show, setShow] = useRecoilState(showNewKeyAddModal);

    const generateNewAgeKey = () => {
        CreateKey()
            .then((key) => console.log(key))
            .catch((err) => console.error(err));
    };

    return (
        <Modal
            title="Add New Key"
            withCloseButton={true}
            opened={show}
            onClose={() => {
                setShow(false);
            }}
        >
            <Tabs defaultValue="age_key">
                <Tabs.List>
                    <Tabs.Tab value="ssh_key">SSH Key</Tabs.Tab>
                    <Tabs.Tab value="age_key">AGE Key</Tabs.Tab>
                </Tabs.List>
                <Tabs.Panel value="ssh_key" pt="xs">
                    <SSHKeyGen />
                </Tabs.Panel>
                <Tabs.Panel value="age_key" pt="xs">
                    <Center>
                        <Button onClick={generateNewAgeKey}>Generate AGE Key</Button>
                    </Center>
                </Tabs.Panel>
            </Tabs>
        </Modal>
    );
}

export default NewKey;
