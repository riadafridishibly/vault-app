import { Switch, Table } from '@mantine/core';
import { formatDate, formatPubPrivKey } from '@src/shared/utils/utils';
import { ActivateKey, DeactivateKey, ReadAllKeys } from '@wailsjs/go/main/DatabaseService';
import { ent } from '@wailsjs/go/models';
import React, { useCallback, useEffect } from 'react';
import NewKey from './NewKey';

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
            <td>{formatPubPrivKey(key.public_key)}</td>
            <td>{formatPubPrivKey(key.private_key)}</td>
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
