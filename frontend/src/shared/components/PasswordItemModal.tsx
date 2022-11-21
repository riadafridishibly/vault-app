import { ActionIcon, Modal, MultiSelect, Textarea } from '@mantine/core';
import React, { useEffect, useState } from 'react';
import { useForm } from '@mantine/form';
import { IconLock, IconAt, IconClipboardCopy, IconEyeOff, IconEye, IconCopy } from '@tabler/icons';
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
import { atom, useRecoilState, useRecoilValue, useSetRecoilState } from 'recoil';
import { showNewPasswordCreateModal } from '@src/pages/PasswordManager/PasswordManagerControls';
import {
    CreatePasswordItem,
    ReadAllPasswordItems,
    ReadSinglePasswordItems,
    UpdatePasswordItem,
} from '@wailsjs/go/main/DatabaseService';
import { passwordManagerItems } from '@src/pages/PasswordManager/Table';
import { isNullOrUndefined } from '../utils/utils';

export interface CreatePasswordItemProps {
    noShadow?: boolean;
    noPadding?: boolean;
    style?: React.CSSProperties;
}

export const defaultPasswordItem = new ent.PasswordItem({
    description: '',
    site_name: '',
    site_url: '',
    username: '',
    password: '',
    tags: [],
});

export const passwordItemFormMode = atom<'create' | 'edit'>({
    key: 'passwordItemFormMode',
    default: 'create',
});

export const passwordItemId = atom<ent.PasswordItem>({
    key: 'passwordItemId',
    default: defaultPasswordItem,
    effects: [
        ({ setSelf, onSet }) => {
            onSet(async (value) => {
                try {
                    if (isNullOrUndefined(value.id) === false && (value.id as number) > 0) {
                        const item = await ReadSinglePasswordItems(value.id as number);
                        setSelf(item);
                    }
                } catch (ex) {
                    console.log(ex);
                }
            });
        },
    ],
});

export function PasswordItemForm({ noShadow, noPadding, style }: CreatePasswordItemProps) {
    const passwordItemMode = useRecoilValue(passwordItemFormMode);
    const setPasswordItems = useSetRecoilState(passwordManagerItems);
    const [openModal, setOpenModal] = useRecoilState(showNewPasswordCreateModal);
    const [editPasswordItem, setEditPasswordItem] = useRecoilState(passwordItemId);
    const [inputType, setInputType] = useState<'text' | 'password'>('password');
    const { paste, copy } = useGoClipboard();
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState<string | null>(null);
    const theme = useMantineTheme();

    const form = useForm<ent.PasswordItem>({
        initialValues: editPasswordItem,
    });
    //  kinda hax need to figure out a better solution later
    useEffect(() => {
        form.setValues(editPasswordItem);
    }, [editPasswordItem]);

    const toggleInputType = () => {
        setInputType((curr) => (curr === 'text' ? 'password' : 'text'));
    };

    const handlePaste = async (e: React.MouseEvent, propName: string) => {
        e.preventDefault();
        form.setFieldValue(propName, await paste());
    };

    const handleCopy = async (e: React.MouseEvent, propName: string) => {
        e.preventDefault();
        const dynamicKey = propName as keyof typeof form.values;
        const value = form.values[dynamicKey];
        if (value) {
            copy(value);
        }
    };

    const handleCopyPaste = async (e: React.MouseEvent, propName: string) => {
        if (passwordItemMode === 'create') {
            handlePaste(e, propName);
        } else {
            handleCopy(e, propName);
        }
    };

    const copyPasteButton = (propName: string) => (
        <ActionIcon
            variant="default"
            sx={{ opacity: 0.4 }}
            // size="sm"
            onClick={(e: React.MouseEvent) => handleCopyPaste(e, propName)}
            tabIndex={-1}
        >
            {passwordItemMode === 'create' ? (
                <IconClipboardCopy size={18} stroke={1.5} />
            ) : (
                <IconCopy size={18} stroke={1.5} />
            )}
        </ActionIcon>
    );

    const handleSubmit = async (values: ent.PasswordItem) => {
        setLoading(true);
        try {
            let addedPasswordItem =
                (values.id as number) > 0
                    ? await UpdatePasswordItem(values)
                    : await CreatePasswordItem(values);

            setPasswordItems(await ReadAllPasswordItems());
        } catch (err) {
            console.log(err);
        } finally {
            setLoading(false);
            form.reset();
            setOpenModal(false);
        }
    };

    return (
        <Modal
            title={passwordItemMode === 'create' ? 'Add New Item' : 'Edit Item'}
            size="lg"
            centered={true}
            opened={openModal}
            onClose={() => {
                setEditPasswordItem(defaultPasswordItem);
                setOpenModal(false);
            }}
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
                            rightSection={copyPasteButton('site_name')}
                        />
                        <TextInput
                            placeholder="https://accounts.google.com/v3/signin/identifier"
                            label="Address"
                            {...form.getInputProps('site_url')}
                            rightSection={copyPasteButton('site_url')}
                        />
                    </Group>
                    {/* TODO: Use constants instead of 'md','sm', etc. This should provide better tracking and easier updatability */}
                    <TextInput
                        mt="md"
                        placeholder="Your email"
                        label="Email"
                        icon={<IconAt size={16} stroke={1.5} />}
                        {...form.getInputProps('username')}
                        rightSection={copyPasteButton('username')}
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
                                {copyPasteButton('password')}
                            </Group>
                        }
                    />

                    <Textarea
                        mt="md"
                        label="Description"
                        placeholder="Add description"
                        {...form.getInputProps('description')}
                    ></Textarea>

                    <MultiSelect
                        mt="md"
                        label="Tags"
                        value={form.values.tags}
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
