/**
 * Lab62 - Unique ID generator
 *
 * @license MIT
 * @author Harman Kang <h@h13g.com>
 *
 */
/**
 * Delegate
 */
interface Lab62Delegate {
	b62char: string;
	b62string: string;
}
/**
 * Generator
 */
interface Lab62Generator {
	(length: number): string;
}
/**
 * Lab62
 */
class Lab62 {
	// Access delegate interface
	private delegate: Lab62Delegate;
	// Access generator interface
	private generator: Lab62Generator;
	// Class constructor
	constructor() {
		// Init delegate
		this.delegate = {
			b62char: "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
			b62string: ""
		}
		// Init generator
		this.generator = function (length: number): string {
			// Run loop specified times
			for (var i = 0; i < length; i++) {
				// Get random integer
				let index = Math.floor(Math.random() * (62));
				// Build ID one character at a time
				this.delegate.b62string += this.delegate.b62char[index];
			}
			return this.delegate.b62string;
		}
	}
	// Make ID
	make(length: number): string {
		return this.generator(length);
	}
}
