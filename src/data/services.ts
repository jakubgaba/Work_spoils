import { dev } from '$app/environment';
import type {Test} from './entities.ts';

import mock_test from './mockTest.json';

const api = {
    loadTestPath: '/api/test2'
};

export const loadTest = async (): Promise<Test[]> => {
return dev ? mock_test : mock_test;
};