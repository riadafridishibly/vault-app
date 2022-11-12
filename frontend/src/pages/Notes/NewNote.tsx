import { useCallback, useState } from 'react';
import { RichTextEditor } from '@mantine/rte';
import { GetAddress } from '@wailsjs/go/main/AssetServerHack';

const initialValue =
    '<p>Your initial <b>html value</b> or an empty string to init editor without value</p>';

export default function NewNote() {
    const [value, onChange] = useState(initialValue);
    const handleChange = useCallback((value: string, delta: any, sources: any) => {
        console.log(value);
        onChange(value);
    }, []);

    const handleImageUpload = useCallback((img: File): Promise<string> => {
        return new Promise(async (resolve, reject) => {
            console.log(img);
            const baseURL = await GetAddress();
            const formData = new FormData();
            formData.append('image', img);
            const reqURL = new URL('/upload', baseURL).href;
            fetch(reqURL, {
                method: 'POST',
                body: formData,
            })
                .then((res) => res.json())
                .then((res) => {
                    console.log(res);
                    return resolve(`${baseURL}/${res.file}`);
                })
                .catch((err) => onChange(JSON.stringify(err)));
        });
    }, []);

    return (
        <RichTextEditor
            // sticky={true}
            // stickyOffset={"60px"}
            value={value}
            onImageUpload={handleImageUpload}
            onChange={handleChange}
            id="rte"
        />
    );
}
