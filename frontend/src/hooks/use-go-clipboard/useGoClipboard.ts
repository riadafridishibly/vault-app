import { useState } from 'react';
import { Copy, Paste } from '@wailsjs/go/main/ClipboardService';

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
        Copy(valueToCopy)
            .then(() => handleCopyResult(true))
            .catch((err) => setError(err));
    };

    const paste = async () => {
        try {
            return await Paste();
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
