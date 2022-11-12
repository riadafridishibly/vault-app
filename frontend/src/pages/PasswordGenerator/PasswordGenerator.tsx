import { Box, Button, Center, Chip, Space, Stack, Text, Tooltip } from '@mantine/core';
import { IconCheck, IconCopy, IconReload } from '@tabler/icons';
import React, { useEffect, useState } from 'react';
import { GeneratePassword } from '@wailsjs/go/main/PasswordService';
import { useGoClipboard } from '@src/hooks/use-go-clipboard/useGoClipboard';

export function ButtonCopy({ password }: { password: string }) {
    const clipboard = useGoClipboard();
    return (
        <Tooltip
            label="Password Copied"
            offset={5}
            position="top"
            radius="xl"
            transition="slide-up"
            transitionDuration={100}
            opened={clipboard.copied}
        >
            <Button
                variant="light"
                color={clipboard.copied ? 'green' : 'grape'}
                rightIcon={
                    clipboard.copied ? (
                        <IconCheck size={20} stroke={1.5} />
                    ) : (
                        <IconCopy size={20} stroke={1.5} />
                    )
                }
                radius="xl"
                size="xl"
                styles={{
                    root: { paddingRight: 14, height: 48 },
                    rightIcon: { marginLeft: 22 },
                }}
                onClick={() => clipboard.copy(password)}
            >
                {password}
            </Button>
        </Tooltip>
    );
}

function PasswordGenerator() {
    const [value, setValue] = useState(['uppercase', 'lowercase', 'numbers']);
    const [pass, setPass] = useState('');

    const generateNew = () => {
        GeneratePassword(value)
            .then((pass) => setPass(pass))
            .catch((err) => console.error(err));
    };

    useEffect(() => {
        generateNew();
    }, [JSON.stringify(value)]);

    return (
        <Box sx={{ width: '100%', height: '100%' }}>
            <Center style={{ width: '100%', height: '100%' }}>
                <Stack>
                    <Text
                        component="span"
                        align="center"
                        variant="gradient"
                        gradient={{ from: 'red', to: 'white', deg: 45 }}
                        size={40}
                        weight={700}
                        style={{ fontFamily: 'Greycliff CF, sans-serif' }}
                    >
                        Password Generator
                    </Text>
                    <Space h="xl" />
                    <Space h="sm" />
                    <Center>
                        <ButtonCopy password={pass} />
                        <Space w="sm" />
                        <Button
                            variant="light"
                            radius="xl"
                            size="sm"
                            color="yellow"
                            styles={{
                                root: { height: 48 },
                            }}
                            onClick={(e: React.MouseEvent) => {
                                e.preventDefault();
                                generateNew();
                            }}
                        >
                            <IconReload />
                        </Button>
                    </Center>
                    <Space h="xl" />
                    <Center>
                        <Chip.Group value={value} onChange={setValue} multiple>
                            <Chip value="lowercase">Lowercase</Chip>
                            <Chip value="uppercase">Uppercase</Chip>
                            <Chip value="special">Special Chars</Chip>
                            <Chip value="numbers">Numbers</Chip>
                        </Chip.Group>
                    </Center>
                </Stack>
            </Center>
        </Box>
    );
}

export default PasswordGenerator;
