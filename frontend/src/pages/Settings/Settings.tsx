import { Card, Center, Group, Text, useMantineColorScheme } from '@mantine/core'
import ColorSchemeSwitch from '@src/shared/components/ColorschemeSwitch'

function Settings() {
    return (
        <Center>
            <Card
                shadow="sm"
            >
                <Group>
                    <Text weight={400} size="lg">
                        Toggle Colorscheme
                    </Text>
                    <ColorSchemeSwitch />
                </Group>
            </Card>
        </Center>
    )
}

export default Settings