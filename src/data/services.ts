//[x]: Mock data works
//[x]: SQLITE3 data working



import { dev } from '$app/environment';
import type {Test} from './entities.ts';

import mock_test from './mockTest.json';

const api = {
    loadTestPath: '/api/test2',
    loadFromRunningDB: 'http://localhost:5173/api/test2'
};

export const loadTest = async (): Promise<Test[]> => {
    console.log(import.meta.env.DEV)
return dev ? mock_test : (await fetch(api.loadFromRunningDB)).json();
};