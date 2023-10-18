export type Safe<T> =
    | {
          error: false;
          data: T;
      }
    | {
          error: true;
          msg: string;
          fields?: {
              field: string;
              tag: string;
          }[];
      };

export declare function safe<T>(promise: Promise<T>): Promise<Safe<T>>;
export declare function safe<T>(fn: () => T): Safe<T>;

export declare function grpcSafe<T>(
    res: (value: Safe<T>) => void,
): (err: ServiceError | null, data: T | undefined) => void;
