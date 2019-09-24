function toCamelCase(name) {
	return name.replace(/-(.)/g, (match, letter) =>letter.toUpperCase());
}