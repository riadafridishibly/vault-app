export function isNullOrUndefined(value: any): boolean {
    return value === null || value === undefined;
}

export function formatDate(date: string) {
    return new Date(date).toLocaleDateString('en-US', {
        year: 'numeric',
        month: 'short',
        day: 'numeric',
    });
}

export function formatPubPrivKey(s: string | undefined) {
    if (s === undefined) {
        return s;
    }

    if (s.length < 40) {
        return s;
    }

    return s.substring(0, 15) + '...' + s.substring(s.length - 5);
}
