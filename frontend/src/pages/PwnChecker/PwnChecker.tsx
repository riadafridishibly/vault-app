import { Button, Group, Text, TextInput } from '@mantine/core';
import { useForm } from '@mantine/form';
import { useGoClipboard } from '@src/hooks/use-go-clipboard/useGoClipboard';
import {
    modalButtonColor,
    modalButtonMargin,
    modalButtonSize,
    modalTextInputMargin,
} from '@src/shared/constants/modal.constants';
import { IconClipboard } from '@tabler/icons';
import { CheckPasswordBreach } from '@wailsjs/go/main/App';
import React, { useState } from 'react';

export const PwnChecker = () => {
    const { error, paste } = useGoClipboard();
    const [breachFound, setBreachFound] = useState(0);
    const passForm = useForm<{
        password: string;
    }>({
        initialValues: { password: '' },
    });

    const handlePasswordBreachCheck = passForm.onSubmit(async (values) => {
        try {
            const value = await CheckPasswordBreach(values.password);
            setBreachFound(value);
        } catch (ex) {
            console.log(`exception found`, ex);
        }
    });

    const handlePaste = async (e: React.MouseEvent, propName: string) => {
        e.preventDefault();
        passForm.setFieldValue(propName, (await paste())?.trim());
    };
    const pasteButton = (propName: string) => (
        <Button
            bg={modalButtonColor}
            size="sm"
            p="5px"
            onClick={(e: React.MouseEvent) => handlePaste(e, propName)}
        >
            <IconClipboard />
        </Button>
    );
    return (
        <>
            <form onSubmit={handlePasswordBreachCheck}>
                <TextInput
                    m={modalTextInputMargin}
                    placeholder="Password"
                    label="Password"
                    withAsterisk
                    required
                    {...passForm.getInputProps('password')}
                    rightSection={pasteButton('password')}
                />
                <Group position="center">
                    <Button
                        bg={modalButtonColor}
                        m={modalButtonMargin}
                        size={modalButtonSize}
                        type="submit"
                    >
                        Check Breach
                    </Button>
                </Group>
            </form>
            <div>
                {breachFound > 0 && (
                    <Text color={'red'} align={'center'}>
                        Password breach has been found in {breachFound} instances!
                    </Text>
                )}
            </div>
        </>
    );
};
