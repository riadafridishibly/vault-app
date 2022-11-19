import { Avatar, Table, Group, Text, ActionIcon, Menu, ScrollArea, Popover } from '@mantine/core';
import {
    IconPencil,
    IconTrash,
    IconDots,
    IconClipboardCopy,
    IconShare,
    IconAlertOctagon,
} from '@tabler/icons';
import { atom, useRecoilValue, useSetRecoilState } from 'recoil';
import { ent } from '@wailsjs/go/models';

import { useGoClipboard } from '@src/hooks/use-go-clipboard/useGoClipboard';
import { showNotification } from '@mantine/notifications';
import { showPasswordBreachModal } from '@src/shared/components/PwnCheckModal';
import { ReadAllPasswordItems } from '@wailsjs/go/main/DatabaseService';
import { passwordItemId } from '@src/shared/components/PasswordItemModal';
import { showNewPasswordCreateModal } from './PasswordManagerControls';
import { isNullOrUndefined } from '@src/shared/utils/utils';
import { openDeletePasswordItemModal } from '@src/shared/components/DeletePasswordItemModal';

function PasswordPopover({ password }: { password: string }) {
    const len = password.length;
    return (
        <Popover width="auto" position="bottom" withArrow shadow="md">
            <Popover.Target>
                <Text>{'*'.repeat(len)}</Text>
            </Popover.Target>
            <Popover.Dropdown>
                <Text size="sm">{password}</Text>
            </Popover.Dropdown>
        </Popover>
    );
}

export const passwordManagerItems = atom<ent.PasswordItem[]>({
    key: 'passwordManagerItems',
    default: [],
    effects: [
        ({ setSelf, onSet }) => {
            ReadAllPasswordItems()
                .then((data) => setSelf(data))
                .catch((err) => console.log(err));
        },
    ],
});

export function PasswordManagerTable() {
    const setEditPasswordItemId = useSetRecoilState(passwordItemId);
    const setOpenModal = useSetRecoilState(showNewPasswordCreateModal);
    const setOpenDeletePasswordItemModal = useSetRecoilState(openDeletePasswordItemModal);
    const setPasswordBreachModalState = useSetRecoilState(showPasswordBreachModal);

    const data = useRecoilValue(passwordManagerItems);
    const { copy } = useGoClipboard();
    const rows = data.map((item) => (
        <tr key={item.id}>
            <td>
                <Group spacing="sm">
                    <Avatar size={40} src={item.avatar} radius={40} />
                    <div>
                        <Text size="sm" weight={500}>
                            {item.site_name}
                        </Text>
                        <Text
                            lineClamp={1}
                            sx={{ maxWidth: '17ch', textOverflow: 'ellipsis', overflow: 'hidden' }}
                            size="xs"
                            color="dimmed"
                        >
                            {item.site_url}
                        </Text>
                    </div>
                </Group>
            </td>
            <td>
                <Text size="sm">{item.username}</Text>
                <Text size="xs" color="dimmed">
                    Email
                </Text>
            </td>
            <td>
                {isNullOrUndefined(item.password) === false && (
                    <PasswordPopover password={item.password as string} />
                )}
                <Text size="xs" color="dimmed">
                    Password
                </Text>
            </td>
            <td>
                <Group spacing={0} position="right">
                    <ActionIcon>
                        <IconPencil
                            size={16}
                            stroke={1.5}
                            onClick={() => {
                                setEditPasswordItemId(item);
                                setOpenModal(true);
                            }}
                        />
                    </ActionIcon>
                    <Menu transition="pop" withArrow position="bottom-end">
                        <Menu.Target>
                            <ActionIcon>
                                <IconDots size={16} stroke={1.5} />
                            </ActionIcon>
                        </Menu.Target>
                        <Menu.Dropdown>
                            <Menu.Item
                                onClick={() => {}}
                                icon={<IconShare size={16} stroke={1.5} />}
                            >
                                Share
                            </Menu.Item>
                            <Menu.Item
                                onClick={() => {
                                    copy(item.site_url);
                                    showNotification({
                                        title: 'Site URL Copied',
                                        message: `Site URL copied to clipboard`,
                                    });
                                }}
                                icon={<IconClipboardCopy size={16} stroke={1.5} />}
                            >
                                Copy URL
                            </Menu.Item>
                            <Menu.Item
                                onClick={() => {
                                    copy(item.username);
                                    showNotification({
                                        title: 'Email Copied',
                                        message: `${item.username} copied to clipboard`,
                                    });
                                }}
                                icon={<IconClipboardCopy size={16} stroke={1.5} />}
                            >
                                Copy Email
                            </Menu.Item>
                            <Menu.Item
                                onClick={() => {
                                    copy(item.password);
                                    showNotification({
                                        title: 'Password Copied',
                                        message: `Password copied to clipboard`,
                                    });
                                }}
                                icon={<IconClipboardCopy size={16} stroke={1.5} />}
                            >
                                Copy Password
                            </Menu.Item>
                            <Menu.Item
                                icon={<IconAlertOctagon size={16} stroke={1.5} />}
                                onClick={() =>
                                    setPasswordBreachModalState({
                                        opened: true,
                                        password: item?.password ?? '',
                                        count: -1,
                                    })
                                }
                            >
                                Check Breach
                            </Menu.Item>
                            <Menu.Item
                                icon={<IconTrash size={16} stroke={1.5} />}
                                color="red"
                                onClick={() => {
                                    console.log('opening delete modal');
                                    setOpenDeletePasswordItemModal({
                                        opened: true,
                                        id: item.id as number,
                                    });
                                }}
                            >
                                Delete
                            </Menu.Item>
                        </Menu.Dropdown>
                    </Menu>
                </Group>
            </td>
        </tr>
    ));

    return (
        <ScrollArea sx={{ height: '100%' }}>
            <Table sx={{ width: '100%' }} verticalSpacing="md">
                <tbody>{rows}</tbody>
            </Table>
        </ScrollArea>
    );
}
