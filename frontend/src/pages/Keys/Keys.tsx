import { Button, Center, Modal, Portal, Tabs } from '@mantine/core';
import React, { useEffect } from 'react';
import DataTable from '../../components/DataTable/DataTable';
import { keysDummy } from '../../shared/data/dummy-data.data';
import SSHKeyGen from './SSHKeyGen';

function Keys() {
    const [show, setShow] = React.useState(true)

    return (
        <div>
            <Modal
                title="Add New Key"
                withCloseButton={true}
                opened={show}
                onClose={() => setShow(false)}
            >
                <Tabs defaultValue="ssh_key">
                    <Tabs.List>
                        <Tabs.Tab value="ssh_key">SSH Key</Tabs.Tab>
                        <Tabs.Tab value="age_key">AGE Key</Tabs.Tab>
                    </Tabs.List>
                    <Tabs.Panel value="ssh_key" pt="xs">
                        <SSHKeyGen />
                    </Tabs.Panel>

                    <Tabs.Panel value="age_key" pt="xs">
                        <Center>
                            <Button>Generate AGE Key</Button>
                        </Center>
                    </Tabs.Panel>
                </Tabs>
            </Modal>
            <DataTable tableData={keysDummy} />
        </div>
    )
}

export default Keys;
