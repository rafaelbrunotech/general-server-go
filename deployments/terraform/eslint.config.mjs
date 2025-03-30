import pluginJs from '@eslint/js';
import tsParser from '@typescript-eslint/parser';
import importPlugin from 'eslint-plugin-import';
import importHelpersPlugin from 'eslint-plugin-import-helpers';
// import jsdoc from 'eslint-plugin-jsdoc';
import perfectionist from 'eslint-plugin-perfectionist';
import pluginSecurity from 'eslint-plugin-security';
import globals from 'globals';
import tsEslint from 'typescript-eslint';

/** @type {import('eslint').Linter.Config[]} */
export default [
  {
    ignores: ['**/node_modules/**', '**/.gen/**', 'eslint.config.mjs', '**/*.js'],
  },
  pluginJs.configs.recommended,
  ...tsEslint.configs.recommended,
  importPlugin.flatConfigs.recommended,
  // jsdoc.configs['flat/recommended'],
  perfectionist.configs['recommended-natural'],
  pluginSecurity.configs.recommended,
  {
    files: ['**/*.ts'],
    rules: {
      '@typescript-eslint/no-empty-object-type': 'off',
      '@typescript-eslint/no-explicit-any': 'off',
      'import-helpers/order-imports': [
        'warn',
        {
          alphabetize: {
            ignoreCase: true,
            order: 'asc',
          },
          groups: [
            'module',
            [
              '/^aws/',
              '/@aws/',
              '/^lib/',
              '/@lib/',
              '/^constructors/',
              '/@constructors/',
              '/^utils/',
              '/@utils/',
              '/^services/',
              '/@services/',
              '/opt/',
            ],
            'absolute',
            ['parent', 'sibling', 'index'],
          ],

          newlinesBetween: 'always',
        },
      ],
    },
  },
  {
    languageOptions: {
      ecmaVersion: 'latest',
      globals: globals.browser,
      parser: tsParser,
      sourceType: 'commonjs',
    },
    plugins: {
      'import-helpers': importHelpersPlugin,
    },
    settings: {
      'import/parsers': {
        '@typescript-eslint/parser': ['.ts'],
      },
      'import/resolver': {
        node: {
          extensions: ['.js', '.ts'],
        },
      },
    },
  },
];
