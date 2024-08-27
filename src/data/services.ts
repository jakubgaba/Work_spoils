import { dev } from '$app/environment';
import type {test} from './entities.ts';

import mock_test from './mockTest.json';

const api = {
    loadTestPath: '/api/test2'
};

export const loadTest = async (): Promise<test[]> => {
return dev ? mock_test : await (await fetch(api.loadTestPath)).json();
};