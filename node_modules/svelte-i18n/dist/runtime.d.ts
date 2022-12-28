import * as svelte_store from 'svelte/store';
import { Readable } from 'svelte/store';
import IntlMessageFormat, { FormatXMLElementFn, Formats } from 'intl-messageformat';

declare const $locale: {
    set: (newLocale: string | null | undefined) => void | Promise<void>;
    update(this: void, updater: svelte_store.Updater<string | null | undefined>): void;
    subscribe(this: void, run: svelte_store.Subscriber<string | null | undefined>, invalidate?: ((value?: string | null | undefined) => void) | undefined): svelte_store.Unsubscriber;
};

interface LocaleDictionary {
    [key: string]: LocaleDictionary | string | Array<string | LocaleDictionary> | null;
}
type LocalesDictionary = {
    [key: string]: LocaleDictionary;
};
type InterpolationValues = Record<string, string | number | boolean | Date | FormatXMLElementFn<unknown> | null | undefined> | undefined;
interface MessageObject {
    id: string;
    locale?: string;
    format?: string;
    default?: string;
    values?: InterpolationValues;
}
type MessageFormatter = (id: string | MessageObject, options?: Omit<MessageObject, 'id'>) => string;
type TimeFormatter = (d: Date | number, options?: IntlFormatterOptions<Intl.DateTimeFormatOptions>) => string;
type DateFormatter = (d: Date | number, options?: IntlFormatterOptions<Intl.DateTimeFormatOptions>) => string;
type NumberFormatter = (d: number, options?: IntlFormatterOptions<Intl.NumberFormatOptions>) => string;
type IntlFormatterOptions<T> = T & {
    format?: string;
    locale?: string;
};
interface MemoizedIntlFormatterOptional<T, U> {
    (options?: IntlFormatterOptions<U>): T;
}
interface MessagesLoader {
    (): Promise<any>;
}
type MissingKeyHandlerInput = {
    locale: string;
    id: string;
    defaultValue: string | undefined;
};
type MissingKeyHandlerOutput = string | void | undefined;
type MissingKeyHandler = (input: MissingKeyHandlerInput) => MissingKeyHandlerOutput;
interface ConfigureOptions {
    /** The global fallback locale * */
    fallbackLocale: string;
    /** The app initial locale * */
    initialLocale?: string | null;
    /** Custom time/date/number formats * */
    formats: Formats;
    /** Loading delay interval * */
    loadingDelay: number;
    /**
     * @deprecated Use `handleMissingMessage` instead.
     * */
    warnOnMissingMessages?: boolean;
    /**
     * Optional method that is executed whenever a message is missing.
     * It may return a string to use as the fallback.
     */
    handleMissingMessage?: MissingKeyHandler;
    /**
     * Whether to treat HTML/XML tags as string literal instead of parsing them as tag token.
     * When this is false we only allow simple tags without any attributes
     * */
    ignoreTag: boolean;
}
type ConfigureOptionsInit = Pick<ConfigureOptions, 'fallbackLocale'> & Partial<Omit<ConfigureOptions, 'fallbackLocale'>>;

declare function init(opts: ConfigureOptionsInit): void | Promise<void>;

declare function registerLocaleLoader(locale: string, loader: MessagesLoader): void;

declare const getLocaleFromHostname: (hostname: RegExp) => string | null;
declare const getLocaleFromPathname: (pathname: RegExp) => string | null;
declare const getLocaleFromNavigator: () => string | null;
declare const getLocaleFromQueryString: (search: string) => string | null | undefined;
declare const getLocaleFromHash: (hash: string) => string | null | undefined;

declare const $dictionary: svelte_store.Writable<LocalesDictionary>;
declare function addMessages(locale: string, ...partials: LocaleDictionary[]): void;
declare const $locales: svelte_store.Readable<string[]>;

declare const $isLoading: svelte_store.Writable<boolean>;

declare const $format: svelte_store.Readable<MessageFormatter>;
declare const $formatTime: svelte_store.Readable<TimeFormatter>;
declare const $formatDate: svelte_store.Readable<DateFormatter>;
declare const $formatNumber: svelte_store.Readable<NumberFormatter>;
declare const $getJSON: svelte_store.Readable<(id: string, locale?: string | undefined) => unknown>;

type MemoizedNumberFormatterFactoryOptional = MemoizedIntlFormatterOptional<Intl.NumberFormat, Intl.NumberFormatOptions>;
type MemoizedDateTimeFormatterFactoryOptional = MemoizedIntlFormatterOptional<Intl.DateTimeFormat, Intl.DateTimeFormatOptions>;
declare const getNumberFormatter: MemoizedNumberFormatterFactoryOptional;
declare const getDateFormatter: MemoizedDateTimeFormatterFactoryOptional;
declare const getTimeFormatter: MemoizedDateTimeFormatterFactoryOptional;
declare const getMessageFormatter: (message: string, locale?: string) => IntlMessageFormat;

type UnwrapStore<T> = T extends Readable<infer U> ? U : T;
/**
 * Unwraps a function from a store and make it function calleable easily outside of a Svelte component.
 *
 * It works by creating a subscription to the store and getting local reference to the store value.
 * Then when the returned function is called, it will execute the function by using the local reference.
 *
 * The returned function has a 'freeze' method that will stop listening to the store.
 *
 * @example
 * // some-js-file.js
 * import { format } from 'svelte-i18n';
 *
 * const $format = unwrapFunctionStore(format);
 *
 * console.log($format('hello', { name: 'John' }));
 *
 */
declare function unwrapFunctionStore<S extends Readable<(...args: any[]) => any>, Fn extends UnwrapStore<S>>(store: S): Fn & {
    /**
     * Stops listening to the store.
     */
    freeze: () => void;
};

declare function defineMessages(i: Record<string, MessageObject>): Record<string, MessageObject>;
declare function waitLocale(locale?: string): Promise<void>;

export { $format as _, addMessages, $formatDate as date, defineMessages, $dictionary as dictionary, $format as format, getDateFormatter, getLocaleFromHash, getLocaleFromHostname, getLocaleFromNavigator, getLocaleFromPathname, getLocaleFromQueryString, getMessageFormatter, getNumberFormatter, getTimeFormatter, init, $isLoading as isLoading, $getJSON as json, $locale as locale, $locales as locales, $formatNumber as number, registerLocaleLoader as register, $format as t, $formatTime as time, unwrapFunctionStore, waitLocale };
