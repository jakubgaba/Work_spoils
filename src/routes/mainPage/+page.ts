import type { PageLoad } from "../$types";
import { tests } from "../../connections/connections";
import { loadTest } from '../../data/services.js';

export const load = (async () => {
    const testData = await loadTest(); 
    tests.set(testData); 
}) satisfies PageLoad;
