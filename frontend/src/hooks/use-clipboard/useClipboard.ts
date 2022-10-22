import { useState } from "react";
import { CopyToClipboard } from "../../../wailsjs/go/main/App";

export function useClipboard({ timeout = 2000 } = {}) {
    const [error, setError] = useState<Error | null>(null);
    const [copied, setCopied] = useState(false);
    const [copyTimeout, setCopyTimeout] = useState<number | undefined>(undefined);

    const handleCopyResult = (value: boolean) => {
        clearTimeout(copyTimeout);
        setCopyTimeout(setTimeout(() => setCopied(false), timeout));
        setCopied(value);
    };

    const copy = (valueToCopy: any) => {
        CopyToClipboard(valueToCopy)
            .then(() => handleCopyResult(true))
            .catch(err => setError(err))
    };

    const reset = () => {
        setCopied(false);
        setError(null);
        clearTimeout(copyTimeout);
    };

    return { copy, reset, error, copied };
}