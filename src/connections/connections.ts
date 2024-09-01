import { writable } from "svelte/store";
import type { Test } from "../data/entities";

export const tests = writable<Test[]>([]);