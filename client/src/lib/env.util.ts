export const checkEnv = (env: unknown) => {
    if (typeof env === 'string' && env.length > 0) {
        return env;
    }
    throw new Error('Environment variable is not set');
}
