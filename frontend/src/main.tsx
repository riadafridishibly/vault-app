import React from 'react';
import { createRoot } from 'react-dom/client';
import './style.css';
import { HashRouter } from 'react-router-dom';
import App from './App';
import { RecoilRoot } from 'recoil';

const container = document.getElementById('root');

const root = createRoot(container!);

root.render(
    <React.StrictMode>
        <RecoilRoot>
            <App />
        </RecoilRoot>
    </React.StrictMode>,
);
