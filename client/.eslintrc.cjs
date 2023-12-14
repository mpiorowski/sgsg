/** @type { import("eslint").Linter.FlatConfig } */
module.exports = {
    root: true,
    extends: ["eslint:recommended", "plugin:svelte/recommended"],
    parserOptions: {
        sourceType: "module",
        ecmaVersion: 2020,
        extraFileExtensions: [".svelte"],
        project: "./jsconfig.json",
    },
    env: {
        browser: true,
        es2017: true,
        node: true,
    },
    rules: {
        "no-unused-vars": "error",
        "no-undef": "error",
        "no-shadow": "error",
        "no-var": "error",
        "prefer-const": "error",
    },
    ignorePatterns: ["node_modules", "service-worker.js"],
};
