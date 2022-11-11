import { Button, Tooltip } from '@mantine/core';
import { useGoClipboard } from '@src/hooks/use-go-clipboard/useGoClipboard';

export function ButtonCopy({ password, label }: { password: string; label: string }) {
    const clipboard = useGoClipboard();
    return (
        <Tooltip
            label={label}
            offset={5}
            position="top"
            radius="xl"
            transition="slide-up"
            transitionDuration={100}
            opened={clipboard.copied}
        >
            <Button
                variant={'subtle'}
                color={clipboard.copied ? 'green' : 'gray'}
                // rightIcon={
                //   clipboard.copied ? (
                //     <IconCheck size={20} stroke={1.5} />
                //   ) : null
                //   // <IconCopy size={20} stroke={1.5} />
                // }
                radius="xl"
                style={{ fontWeight: 'normal' }}
                styles={{
                    root: { paddingRight: 14 },
                    rightIcon: { marginLeft: 22 },
                }}
                onClick={() => clipboard.copy(password)}
            >
                {password}
            </Button>
        </Tooltip>
    );
}
