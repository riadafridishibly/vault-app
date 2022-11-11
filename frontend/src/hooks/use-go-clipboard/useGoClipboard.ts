import { useState } from 'react';
import { CopyToClipboard, PasteFromClipboard } from '@wailsjs/go/main/App';

export function useGoClipboard({ timeout = 2000 } = {}) {
    const [error, setError] = useState<Error | null>(null);
    const [copied, setCopied] = useState(false);
    const [copyTimeout, setCopyTimeout] = useState<NodeJS.Timeout>();

    const handleCopyResult = (value: boolean) => {
        clearTimeout(copyTimeout);
        setCopyTimeout(setTimeout(() => setCopied(false), timeout));
        setCopied(value);
    };

    const copy = (valueToCopy: any) => {
        CopyToClipboard(valueToCopy)
            .then(() => handleCopyResult(true))
            .catch((err) => setError(err));
    };

    const paste = async () => {
        try {
            return await PasteFromClipboard();
        } catch (ex) {
            console.log(ex);
        }
    };

    const reset = () => {
        setCopied(false);
        setError(null);
        clearTimeout(copyTimeout);
    };

    return { copy, paste, reset, error, copied };
}
