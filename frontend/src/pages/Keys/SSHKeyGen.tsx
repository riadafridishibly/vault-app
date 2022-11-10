import {
    TextInput,
    PasswordInput,
    Paper,
    Container,
    Button,
    NativeSelect,
} from '@mantine/core';
import React from 'react';


function SSHKeyGen() {
    const defaultSelected = 'ed25519'
    const [name, setName] = React.useState(`id_${defaultSelected}`)
    const [selected, setSelected] = React.useState(defaultSelected)

    // TODO(riad): respect user edits
    const handleSelected = React.useCallback((value: string) => {
        setSelected(value)
        setName(`id_${value}`)
    }, [])

    return (
        <Container>
            <Paper withBorder shadow="md" p={30} mt={0} radius="md">
                <NativeSelect
                    defaultChecked
                    defaultValue={defaultSelected}
                    data={['dsa', 'ecdsa', 'ecdsa-sk', 'ed25519', 'ed25519-sk', 'rsa']}
                    label="Key Type"
                    onChange={(event) => handleSelected(event.target.value)}
                />
                <TextInput
                    label="Name"
                    value={name}
                    onChange={(e) => setName(e.target.value)} />
                <TextInput label="Path" defaultValue="~/.ssh/" readOnly />
                <PasswordInput label="Password" placeholder="Optional password" mt="md" />
                <Button fullWidth mt="xl">
                    Generate
                </Button>
            </Paper>
        </Container>
    );

}

export default SSHKeyGen