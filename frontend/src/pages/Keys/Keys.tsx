import { Switch, Table } from '@mantine/core';
import { ActivateKey, DeactivateKey, ReadAllKeys } from '@wailsjs/go/main/DatabaseService';
import { ent } from '@wailsjs/go/models';
import React, { useCallback, useEffect } from 'react';
import NewKey from './NewKey';

function formatDate(date: string) {
    return new Date(date).toLocaleDateString('en-US', {
        year: 'numeric',
        month: 'short',
        day: 'numeric',
    });
}

function trunkString(s: string | undefined) {
    if (s === undefined) {
        return s;
    }

    if (s.length < 40) {
        return s;
    }

    return s.substring(0, 15) + '...' + s.substring(s.length - 5);
}

function KeyState({
    id,
    currentState,
    refetch,
}: {
    id: number | undefined;
    currentState: boolean | undefined;
    refetch: () => void;
}) {
    const [checked, setChecked] = React.useState(currentState);
    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        if (!id) {
            return;
        }

        // TODO(riad): Organize this method
        setChecked(e.currentTarget.checked);
        if (e.currentTarget.checked) {
            ActivateKey(id)
                .then((key) => console.log(key))
                .catch((err) => console.error(err));
        } else {
            // Activated
            DeactivateKey(id)
                .then((key) => console.log(key))
                .catch((err) => console.error(err));
        }
    };

    return <Switch checked={checked} onChange={handleChange} />;
}

function Keys() {
    const [keys, setKeys] = React.useState<ent.Key[]>();
    const [recalculate, setRecalculate] = React.useState(0);
    const handleRefetch = useCallback(() => {
        setRecalculate((curr) => curr + 1);
    }, []);
    useEffect(() => {
        ReadAllKeys()
            .then((keys) => setKeys(keys))
            .catch((err) => console.error(err));
    }, [recalculate]);

    if (!keys) {
        return <div>Couldn't load keys</div>;
    }

    const rows = keys.map((key) => (
        <tr key={key.id}>
            <td>{trunkString(key.public_key)}</td>
            <td>{trunkString(key.private_key)}</td>
            <td>{formatDate(key.create_time)}</td>
            <td>
                <KeyState id={key.id} currentState={key.is_active} refetch={handleRefetch} />
            </td>
        </tr>
    ));

    return (
        <>
            <NewKey refetch={handleRefetch} />
            <Table>
                <thead>
                    <tr>
                        <th>Public Key</th>
                        <th>Private Key</th>
                        <th>Created Time</th>
                        <td>Active</td>
                    </tr>
                </thead>
                <tbody>{rows}</tbody>
            </Table>
        </>
    );
}

export default Keys;
