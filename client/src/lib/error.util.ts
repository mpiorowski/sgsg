import { toast, ToastType } from '@mpiorowski/svelte-init';
import { Config } from '../config';
import { t } from "svelte-i18n"

type GrpcError = {
    code: string;
    error: string;
}

const isGrpcError = (error: unknown): error is GrpcError => {
    if (error && typeof error === 'object' && "code" in error && 'error' in error) {
        return true;
    }
    return false;
};

export const handleError = (error: unknown) => {
    if (Config.VITE_NODE_ENV === 'development') {
        console.error(error);
    }
    if (isGrpcError(error)) {
        t.subscribe((value) => {
            toast(value(error.error), ToastType.ERROR)
        });
    } else {
        t.subscribe((value) => {
            toast(value('errors.somethingWentWrong'), ToastType.ERROR);
        });
    }
};
