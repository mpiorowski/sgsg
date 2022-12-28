import { writable, type Writable } from 'svelte/store';
import { Config } from '../config';
import type { SafeParseReturnType } from 'zod';

type UseApiResponse<D, V> = {
    fetching: Writable<boolean>;
    error: Writable<unknown>;
    formError: Writable<Record<string, string>>;
    request: (options: Options) => Promise<void>;
    response: Writable<D | null>;
    validate: (values: SafeParseReturnType<V, V>) => boolean;
}
/**
 * Use api hook
 * @returns {UseApiResponse} UseApiResponse
 */
export const useApi = <D, V = Partial<D>>(): UseApiResponse<D, V> => {
    const fetching = writable(false);
    const error = writable<unknown>(null);
    const formError = writable<Record<string, string>>({});
    const response = writable<D | null>(null);

    const request = async (options: Options): Promise<void> => {
        fetching.set(true);
        error.set(null);
        try {
            const data = await apiRequest<D>(options);
            response.set(data);
        } catch (err) {
            error.set(err);
        } finally {
            fetching.set(false);
        }
    }

    const validate = (values: SafeParseReturnType<V, V>) => {
        formError.set({});
        if (values.success) {
            return true;
        }
        console.error(values.error);
        const err: Record<string, string> = {};
        values.error.errors.forEach((val) => {
            if (val.path[0]) {
                err[val.path[0]] = val.message;
            }
        });
        formError.set(err);
        return false;
    }

    return {
        fetching,
        error,
        formError,
        request,
        response,
        validate,
    };
}

type Options = {
    url: string;
    method: 'GET' | 'POST' | 'PUT' | 'PATCH' | 'DELETE';
    body?: string | FormData;
};

/**
 * @description - Api request with json or form data
 * @param {Options} Options
 * @returns {Promise<T>} Promise
 */
export const apiRequest = async <T>({ url, method, body }: Options): Promise<T> => {
    const headers = new Headers();

    const response = await fetch(`${Config.VITE_API_URL}${url}`, {
        method: method,
        body: body,
        credentials: 'include',
        headers,
    });
    if (response.status === 401) {
        throw new Error('Unauthorized');
    }
    if (response.status === 204) {
        return {} as T;
    }
    let data: T;
    if (response.headers.get('Content-Type')?.includes('application/json')) {
        data = await response.json() as T;
    } else {
        data = await response.text() as T;
    }
    if (response.ok) {
        return data;
    }
    throw data;
};
