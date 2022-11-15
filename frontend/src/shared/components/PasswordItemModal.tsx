import { ActionIcon, Modal, MultiSelect } from '@mantine/core';
import React, { useState } from 'react';
import { useForm } from '@mantine/form';
import { IconLock, IconAt, IconClipboardCopy, IconEyeOff, IconEye } from '@tabler/icons';
import { ent } from '@wailsjs/go/models';
import {
    TextInput,
    Group,
    Button,
    Paper,
    Text,
    LoadingOverlay,
    useMantineTheme,
} from '@mantine/core';
import { useGoClipboard } from '@src/hooks/use-go-clipboard/useGoClipboard';
import { useRecoilState, useSetRecoilState } from 'recoil';
import { showNewPasswordCreateModal } from '@src/pages/PasswordManager/PasswordManagerControls';
import { CreatePasswordItem } from '@wailsjs/go/main/DatabaseService';
import { passwordManagerItems } from '@src/pages/PasswordManager/Table';

export interface CreatePasswordItemProps {
    noShadow?: boolean;
    noPadding?: boolean;
    style?: React.CSSProperties;
}

export function PasswordItemForm({ noShadow, noPadding, style }: CreatePasswordItemProps) {
    const setPasswordItems = useSetRecoilState(passwordManagerItems);
    const [openModal, setOpenModal] = useRecoilState(showNewPasswordCreateModal);
    const [inputType, setInputType] = useState<'text' | 'password'>('password');
    const { paste } = useGoClipboard();
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState<string | null>(null);
    const theme = useMantineTheme();

    const form = useForm<ent.PasswordItem>({
        initialValues: new ent.PasswordItem({
            description: '',
            site_name: '',
            site_url: '',
            username: '',
            password: '',
            tags: [],
        }),
    });

    const toggleInputType = () => {
        setInputType((curr) => (curr === 'text' ? 'password' : 'text'));
    };

    const handlePaste = async (e: React.MouseEvent, propName: string) => {
        e.preventDefault();
        form.setFieldValue(propName, await paste());
    };

    const pasteButton = (propName: string) => (
        <ActionIcon
            variant="default"
            sx={{ opacity: 0.4 }}
            // size="sm"
            onClick={(e: React.MouseEvent) => handlePaste(e, propName)}
        >
            <IconClipboardCopy size={18} stroke={1.5} />
        </ActionIcon>
    );

    const handleSubmit = (values: ent.PasswordItem) => {
        setLoading(true);
        CreatePasswordItem(values)
            .then((resp) => {
                setPasswordItems((current) => [...current, resp]);
                console.log(resp);
                setLoading(false);
                form.reset();
                setOpenModal(false);
            })
            .catch((err) => console.error(err));
        console.log(values);
    };

    return (
        <Modal
            title="Add New Item"
            size="md"
            opened={openModal}
            onClose={() => setOpenModal((curr) => !curr)}
        >
            <Paper
                p={noPadding ? 0 : 'lg'}
                shadow={noShadow ? undefined : 'sm'}
                style={{
                    position: 'relative',
                    backgroundColor:
                        theme.colorScheme === 'dark' ? theme.colors.dark[7] : theme.white,
                    ...style,
                }}
            >
                <form onSubmit={form.onSubmit(handleSubmit)}>
                    <LoadingOverlay visible={loading} />
                    <Group grow>
                        <TextInput
                            data-autofocus
                            placeholder="Google"
                            label="Website"
                            {...form.getInputProps('site_name')}
                            rightSection={pasteButton('site_name')}
                        />
                        <TextInput
                            placeholder="https://accounts.google.com/v3/signin/identifier"
                            label="Address"
                            {...form.getInputProps('site_url')}
                            rightSection={pasteButton('site_url')}
                        />
                    </Group>
                    {/* TODO: Use constants instead of 'md','sm', etc. This should provide better tracking and easier updatability */}
                    <TextInput
                        mt="md"
                        placeholder="Your email"
                        label="Email"
                        icon={<IconAt size={16} stroke={1.5} />}
                        {...form.getInputProps('username')}
                        rightSection={pasteButton('username')}
                    />

                    <TextInput
                        mt="md"
                        placeholder="Password"
                        label="Password"
                        type={inputType}
                        icon={<IconLock size={18} stroke={1.5} />}
                        {...form.getInputProps('password')}
                        rightSectionWidth={75}
                        rightSection={
                            <Group spacing="xs">
                                <LockUnlock
                                    inputType={inputType}
                                    toggleInputType={toggleInputType}
                                />
                                {pasteButton('password')}
                            </Group>
                        }
                    />

                    <MultiSelect
                        mt="md"
                        label="Tags"
                        data={form.getInputProps('tags').value}
                        placeholder="Select items"
                        searchable
                        creatable
                        getCreateLabel={(query) => `+ Create ${query}`}
                        onCreate={(query) => {
                            form.setFieldValue('tags', [
                                ...form.getInputProps('tags').value,
                                query,
                            ]);
                            return query;
                        }}
                    />

                    {error && (
                        <Text color="red" size="sm" mt="sm">
                            {error}
                        </Text>
                    )}

                    <Group position="apart" mt="xl">
                        <Button color="blue" type="submit">
                            Save
                        </Button>
                    </Group>
                </form>
            </Paper>
        </Modal>
    );
}

function LockUnlock({
    inputType,
    toggleInputType,
}: {
    inputType: string;
    toggleInputType: () => void;
}) {
    return (
        <ActionIcon
            variant="default"
            sx={{ opacity: 0.4 }}
            onClick={() => {
                toggleInputType();
            }}
        >
            {inputType === 'password' ? (
                <IconEyeOff size={18} stroke={1.5} />
            ) : (
                <IconEye size={18} stroke={1.5} />
            )}
        </ActionIcon>
    );
}
