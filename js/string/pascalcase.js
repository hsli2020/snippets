/*!
 * pascalcase <https://github.com/jonschlinkert/pascalcase>
 *
 * Copyright (c) 2015-present, Jon ("Schlink") Schlinkert.
 * Licensed under the MIT License.
 */

const titlecase = input => input[0].toLocaleUpperCase() + input.slice(1);

module.exports = value => {
  if (value === null || value === void 0) return '';
  if (typeof value.toString !== 'function') return '';

  let input = value.toString().trim();
  if (input === '') return '';
  if (input.length === 1) return input.toLocaleUpperCase();

  let match = input.match(/[a-zA-Z0-9]+/g);
  if (match) {
    return match.map(m => titlecase(m)).join('');
  }

  return input;
};