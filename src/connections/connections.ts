import { writable } from "svelte/store";
import type { test } from "../data/entities";

export const tests = writable<test[]>([]);