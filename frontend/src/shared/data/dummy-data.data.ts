import { CellData } from '../models/cell-data.model';
import { TableData } from './../models/table-data.model';


export const passwordDummy: TableData = {
    headers: ['Website', 'Username', 'Password', 'Tags'],
    rows: [
        {
            website: new CellData('https://google.com'),
            username: new CellData('testUser'),
            password: new CellData('something', true),
            tags: new CellData(['1', '2', '3'].join(','), false, true),
        },
        {
            website: new CellData('https://google.com'),
            username: new CellData('testUszxzczer'),
            password: new CellData('someasdaxzthing'),
            tags: new CellData(['1', '2', '213'].join(','), false, true),
        },
        {
            website: new CellData('https://google.com'),
            username: new CellData('testU1231ser'),
            password: new CellData('something', true),
            tags: new CellData(['1zxc', '2', '3'].join(','), false, true),
        },
        {
            website: new CellData('https://google.com'),
            username: new CellData('testUser'),
            password: new CellData('something', true),
            tags: new CellData(['1', '2', '3'].join(','), false, true),
        },
        {
            website: new CellData('https://google.com'),
            username: new CellData('testUser'),
            password: new CellData('someaszxthing', true),
            tags: new CellData(['1', '2', '3'].join(','), false, true),
        },
        {
            website: new CellData('https://google.com'),
            username: new CellData('tes123tUser'),
            password: new CellData('something', true),
            tags: new CellData(['1', '2', '3'].join(','), false, true),
        },
        {
            website: new CellData('https://google.com'),
            username: new CellData('te65456stUser'),
            password: new CellData('something', true),
            tags: new CellData(['1', '2', '3'].join(','), false, true),
        },
    ],
};
